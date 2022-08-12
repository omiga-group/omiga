package outbox

import (
	"context"
	"sync"
	"time"

	"github.com/omiga-group/omiga/src/order/shared/repositories"
	outboxmodel "github.com/omiga-group/omiga/src/order/shared/repositories/outbox"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/outbox"
	"go.uber.org/zap"
)

type OutboxBackgroundService interface {
	RunAsync()
}

type outboxBackgroundService struct {
	ctx             context.Context
	logger          *zap.SugaredLogger
	outboxSettings  outbox.OutboxSettings
	topic           string
	messageProducer messaging.MessageProducer
	entgoClient     repositories.EntgoClient
	retryDelay      time.Duration
}

func NewOutboxBackgroundService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	outboxSettings outbox.OutboxSettings,
	messageProducer messaging.MessageProducer,
	topic string,
	entgoClient repositories.EntgoClient,
	cronService cron.CronService) (OutboxBackgroundService, error) {
	retryDelay, err := time.ParseDuration(outboxSettings.RetryDelay)
	if err != nil {
		return nil, err
	}

	instance := &outboxBackgroundService{
		ctx:             ctx,
		logger:          logger,
		outboxSettings:  outboxSettings,
		messageProducer: messageProducer,
		entgoClient:     entgoClient,
		retryDelay:      retryDelay,
	}

	if _, err := cronService.GetCron().AddJob("0/5 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (obs *outboxBackgroundService) RunAsync() {
	go obs.Run()
}

func (obs *outboxBackgroundService) Run() {
	client := obs.entgoClient.GetClient()

	records, err := client.Outbox.Query().
		Where(
			outboxmodel.And(
				outboxmodel.TopicEQ(obs.topic),
				outboxmodel.StatusEQ(outboxmodel.StatusPending),
				outboxmodel.Or(
					outboxmodel.LastRetryLTE(time.Now().Add(-1*obs.retryDelay)),
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
	succeededRecords := make([]*repositories.Outbox, 0)

	for _, record := range records {
		wg.Add(1)

		go func(record *repositories.Outbox) {
			defer wg.Done()

			err := obs.messageProducer.Produce(
				obs.ctx,
				record.Key,
				record.Payload)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {

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

		if record.RetryCount == obs.outboxSettings.MaxRetryCount {
			_, err := client.Outbox.
				UpdateOne(record).
				SetStatus(outboxmodel.StatusFailed).
				SetLastRetry(now).
				SetProcessingErrors(append(record.ProcessingErrors, failedRecord.err.Error())).
				Save(obs.ctx)
			if err != nil {
				obs.logger.Errorf("Failed to update failed outbox item for topic %s. Error: %v", obs.topic, err)
			}
		} else {
			_, err := client.Outbox.
				UpdateOne(record).
				AddRetryCount(1).
				SetLastRetry(now).
				SetProcessingErrors(append(record.ProcessingErrors, failedRecord.err.Error())).
				Save(obs.ctx)
			if err != nil {
				obs.logger.Errorf("Failed to update failed outbox item for topic %s. Error: %v", obs.topic, err)
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
			obs.logger.Errorf("Failed to update succeeded outbox item for topic %s. Error: %v", obs.topic, err)
		}
	}
}

type failedRecord struct {
	record *repositories.Outbox
	err    error
}
