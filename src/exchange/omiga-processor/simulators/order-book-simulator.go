package simulators

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type OrderBookSimulator interface {
}

type OrderBookSimulatorSettings struct {
	ExchangeName string
}

type orderBookSimulator struct {
	ctx                        context.Context
	logger                     *zap.SugaredLogger
	orderBookPublisher         publishers.OrderBookPublisher
	orderBookSimulatorSettings OrderBookSimulatorSettings
}

func NewOrderBookSimulator(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	orderBookPublisher publishers.OrderBookPublisher,
	orderBookSimulatorSettings OrderBookSimulatorSettings) (OrderBookSimulator, error) {
	instance := &orderBookSimulator{
		ctx:                        ctx,
		logger:                     logger,
		orderBookPublisher:         orderBookPublisher,
		orderBookSimulatorSettings: orderBookSimulatorSettings,
	}

	if _, err := cronService.GetCron().AddJob("0/5 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (obs *orderBookSimulator) Run() {
	orderBook := models.OrderBook{
		ExchangeId: obs.orderBookSimulatorSettings.ExchangeName,
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
				Quantity: models.Money{
					Amount: 1,
					Scale:  1,
					Currency: models.Currency{
						Name:         "Bitcoin",
						Code:         "BTC",
						MaxPrecision: 1,
						Digital:      true,
					},
				},
				Price: models.Money{
					Amount: 1,
					Scale:  1,
					Currency: models.Currency{
						Name:         "Ethereum",
						Code:         "ETH",
						MaxPrecision: 1,
						Digital:      true,
					},
				},
			},
			{
				Quantity: models.Money{
					Amount: 2,
					Scale:  2,
					Currency: models.Currency{
						Name:         "Bitcoin",
						Code:         "BTC",
						MaxPrecision: 1,
						Digital:      true,
					},
				},
				Price: models.Money{
					Amount: 2,
					Scale:  2,
					Currency: models.Currency{
						Name:         "Ethereum",
						Code:         "ETH",
						MaxPrecision: 1,
						Digital:      true,
					},
				},
			},
		},
		Bids: []models.OrderBookEntry{
			{
				Quantity: models.Money{
					Amount: 3,
					Scale:  3,
					Currency: models.Currency{
						Name:         "Bitcoin",
						Code:         "BTC",
						MaxPrecision: 1,
						Digital:      true,
					},
				},
				Price: models.Money{
					Amount: 1,
					Scale:  1,
					Currency: models.Currency{
						Name:         "Ethereum",
						Code:         "ETH",
						MaxPrecision: 1,
						Digital:      true,
					},
				},
			},
			{
				Quantity: models.Money{
					Amount: 4,
					Scale:  4,
					Currency: models.Currency{
						Name:         "Bitcoin",
						Code:         "BTC",
						MaxPrecision: 1,
						Digital:      true,
					},
				},
				Price: models.Money{
					Amount: 2,
					Scale:  2,
					Currency: models.Currency{
						Name:         "Ethereum",
						Code:         "ETH",
						MaxPrecision: 1,
						Digital:      true,
					},
				},
			},
		},
	}

	if err := obs.orderBookPublisher.Publish(
		obs.ctx,
		obs.orderBookSimulatorSettings.ExchangeName,
		orderBook); err != nil {
		obs.logger.Errorf("Failed to publish order book. Error: %v", err)
	}
}
