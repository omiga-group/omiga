package outbox

import (
	"context"

	"github.com/omiga-group/omiga/src/order/shared/repositories"
	outboxmodel "github.com/omiga-group/omiga/src/order/shared/repositories/outbox"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/outbox"
	"go.uber.org/zap"
)

type OutboxBackgroundService interface {
}

type outboxBackgroundService struct {
	ctx             context.Context
	logger          *zap.SugaredLogger
	outboxSettings  outbox.OutboxSettings
	messageProducer messaging.MessageProducer
	entgoClient     repositories.EntgoClient
}

func NewOutboxBackgroundService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	outboxSettings outbox.OutboxSettings,
	messageProducer messaging.MessageProducer,
	entgoClient repositories.EntgoClient,
	cronService cron.CronService) (OutboxBackgroundService, error) {
	instance := &outboxBackgroundService{
		ctx:             ctx,
		logger:          logger,
		outboxSettings:  outboxSettings,
		messageProducer: messageProducer,
		entgoClient:     entgoClient,
	}

	if _, err := cronService.GetCron().AddJob("0/5 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (obs *outboxBackgroundService) Run() {
	client, err := obs.entgoClient.GetClient()
	if err != nil {
		obs.logger.Errorf("Failed to get client. Error: %v", err)

		return
	}

	items, err := client.Outbox.Query().
		Where(outboxmodel.StatusEQ(outboxmodel.StatusPending)).
		Limit(1000).
		All(obs.ctx)

	if err != nil {
		obs.logger.Errorf("Failed to fetch outbox items. Error: %v", err)

		return
	}

	if len(items) == 0 {
		return
	}
}
