package simulators

import (
	"context"

	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type OrderBookSimulator interface {
}

type orderBookSimulator struct {
	ctx               context.Context
	logger            *zap.SugaredLogger
	orderBookProducer orderbookv1.Producer
}

func NewOrderBookSimulator(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	orderBookProducer orderbookv1.Producer) (OrderBookSimulator, error) {
	instance := &orderBookSimulator{
		ctx:               ctx,
		logger:            logger,
		orderBookProducer: orderBookProducer,
	}

	if _, err := cronService.GetCron().AddJob("0/5 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (obs *orderBookSimulator) Run() {
}
