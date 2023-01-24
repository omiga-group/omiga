package outbox

import (
	"context"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/order/shared/entities"
	outboxmodel "github.com/omiga-group/omiga/src/order/shared/entities/outbox"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/outbox"
	"go.uber.org/zap"
)

type OutboxBackgroundService interface {
	RunAsync()
	Close()
}

type outboxBackgroundService struct {
	ctx             context.Context
	logger          *zap.SugaredLogger
	outboxConfig    outbox.OutboxConfig
	messageProducer messaging.MessageProducer
	entgoClient     entities.EntgoClient
	globalMutex     sync.Mutex
}

func NewOutboxBackgroundService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	outboxConfig outbox.OutboxConfig,
	messageProducer messaging.MessageProducer,
	entgoClient entities.EntgoClient,
	jobScheduler *gocron.Scheduler) (OutboxBackgroundService, error) {

	instance := &outboxBackgroundService{
		ctx:             ctx,
		logger:          logger,
		outboxConfig:    outboxConfig,
		messageProducer: messageProducer,
		entgoClient:     entgoClient,
		globalMutex:     sync.Mutex{},
	}

	if _, err := jobScheduler.Every(1).Seconds().Do(func() {
		instance.Run()
	}); err != nil {
		return nil, err
	}

	return instance, nil
}

func (obs *outboxBackgroundService) RunAsync() {
	go obs.Run()
}

func (obs *outboxBackgroundService) Close() {
	obs.messageProducer.Close()
}

func (obs *outboxBackgroundService) Run() {
	obs.globalMutex.Lock()
	defer obs.globalMutex.Unlock()

	client := obs.entgoClient.GetClient()

	records, err := client.Outbox.Query().
		Where(
			outboxmodel.And(
				outboxmodel.StatusEQ(outboxmodel.StatusPending),
				outboxmodel.Or(
					outboxmodel.LastRetryLTE(time.Now().Add(-1*obs.outboxConfig.RetryDelay)),
					outboxmodel.LastRetryIsNil()),
			),
		).
		All(obs.ctx)
	if err != nil {
		obs.logger.Errorf("Failed to fetch outbox items. Error: %v", err)

		return
	}

	numOfRecords := len(records)
	if numOfRecords == 0 {
		return
	}

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	failedRecords := make([]failedRecord, 0)
	succeededRecords := make([]*entities.Outbox, 0)

	for _, record := range records {
		wg.Add(1)

		go func(record *entities.Outbox) {
			defer wg.Done()

			if err := obs.messageProducer.Connect(record.Topic); err != nil {
				mu.Lock()
				defer mu.Unlock()

				failedRecords = append(failedRecords, failedRecord{
					record: record,
					err:    err,
				})

				return
			}

			if err := obs.messageProducer.Produce(
				obs.ctx,
				record.Key,
				record.Payload); err != nil {
				mu.Lock()
				defer mu.Unlock()

				failedRecords = append(failedRecords, failedRecord{
					record: record,
					err:    err,
				})

				return
			}

			succeededRecords = append(succeededRecords, record)
		}(record)
	}

	wg.Wait()

	now := time.Now()

	for _, failedRecord := range failedRecords {
		record := failedRecord.record

		if record.RetryCount == obs.outboxConfig.MaxRetryCount {
			_, err := client.Outbox.
				UpdateOne(record).
				SetStatus(outboxmodel.StatusFailed).
				SetLastRetry(now).
				SetProcessingErrors(append(record.ProcessingErrors, failedRecord.err.Error())).
				Save(obs.ctx)
			if err != nil {
				obs.logger.Errorf("Failed to update failed outbox item for topic %s. Error: %v", record.Topic, err)
			}
		} else {
			_, err := client.Outbox.
				UpdateOne(record).
				AddRetryCount(1).
				SetLastRetry(now).
				SetProcessingErrors(append(record.ProcessingErrors, failedRecord.err.Error())).
				Save(obs.ctx)
			if err != nil {
				obs.logger.Errorf("Failed to update failed outbox item for topic %s. Error: %v", record.Topic, err)
			}
		}
	}

	for _, record := range succeededRecords {
		_, err := client.Outbox.
			UpdateOne(record).
			SetStatus(outboxmodel.StatusSucceeded).
			AddRetryCount(1).
			SetLastRetry(now).
			Save(obs.ctx)
		if err != nil {
			obs.logger.Errorf("Failed to update succeeded outbox item for topic %s. Error: %v", record.Topic, err)
		}
	}
}

type failedRecord struct {
	record *entities.Outbox
	err    error
}
