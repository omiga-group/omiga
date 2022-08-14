package simulators

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/omiga-processor/publishers"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type OrderBookSimulator interface {
}

type orderBookSimulator struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	orderBookPublisher publishers.OrderBookPublisher
}

func NewOrderBookSimulator(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	orderBookPublisher publishers.OrderBookPublisher) (OrderBookSimulator, error) {
	instance := &orderBookSimulator{
		ctx:                ctx,
		logger:             logger,
		orderBookPublisher: orderBookPublisher,
	}

	if _, err := cronService.GetCron().AddJob("0/5 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (obs *orderBookSimulator) Run() {
}
