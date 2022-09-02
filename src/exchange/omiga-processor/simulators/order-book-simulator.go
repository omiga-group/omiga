package simulators

import (
	"context"
	"time"

	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type OrderBookSimulator interface {
}

type OrderBookSimulatorConfig struct {
	ExchangeName string
}

type orderBookSimulator struct {
	ctx                      context.Context
	logger                   *zap.SugaredLogger
	orderBookPublisher       publishers.OrderBookPublisher
	orderBookSimulatorConfig OrderBookSimulatorConfig
}

func NewOrderBookSimulator(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	orderBookPublisher publishers.OrderBookPublisher,
	orderBookSimulatorConfig OrderBookSimulatorConfig) (OrderBookSimulator, error) {
	instance := &orderBookSimulator{
		ctx:                      ctx,
		logger:                   logger,
		orderBookPublisher:       orderBookPublisher,
		orderBookSimulatorConfig: orderBookSimulatorConfig,
	}

	if _, err := cronService.GetCron().AddJob("0/5 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (obs *orderBookSimulator) Run() {
	now := time.Now()
	orderBook := models.OrderBook{
		ExchangeId: obs.orderBookSimulatorConfig.ExchangeName,
		BaseCurrency: models.Currency{
			Name:         "Bitcoin",
			Code:         "BTC",
			MaxPrecision: 10,
			Digital:      true,
		},
		CounterCurrency: models.Currency{
			Name:         "Ethereum",
			Code:         "ETH",
			MaxPrecision: 20,
			Digital:      true,
		},
		Asks: []models.OrderBookEntry{
			{
				Time: now,
				Quantity: models.Quantity{
					Amount: 1,
					Scale:  1,
				},
				Price: models.Quantity{
					Amount: 1,
					Scale:  1,
				},
			},
			{
				Time: now,
				Quantity: models.Quantity{
					Amount: 1,
					Scale:  1,
				},
				Price: models.Quantity{
					Amount: 1,
					Scale:  1,
				},
			},
		},
		Bids: []models.OrderBookEntry{
			{
				Time: now,
				Quantity: models.Quantity{
					Amount: 1,
					Scale:  1,
				},
				Price: models.Quantity{
					Amount: 1,
					Scale:  1,
				},
			},
			{
				Time: now,
				Quantity: models.Quantity{
					Amount: 1,
					Scale:  1,
				},
				Price: models.Quantity{
					Amount: 1,
					Scale:  1,
				},
			},
		},
	}

	if err := obs.orderBookPublisher.Publish(
		obs.ctx,
		obs.orderBookSimulatorConfig.ExchangeName,
		orderBook); err != nil {
		obs.logger.Errorf("Failed to publish order book. Error: %v", err)
	}
}
