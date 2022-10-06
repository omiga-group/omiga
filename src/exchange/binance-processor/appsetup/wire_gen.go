// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	configuration2 "github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/subscribers"
	configuration3 "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/services"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewCronService(logger *zap.SugaredLogger) (cron.CronService, error) {
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	cronService, err := cron.NewCronService(logger, timeHelper)
	if err != nil {
		return nil, err
	}
	return cronService, nil
}

func NewTimeHelper() (time.TimeHelper, error) {
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	return timeHelper, nil
}

func NewSyntheticOrderConsumer(logger *zap.SugaredLogger, pulsarConfig pulsar.PulsarConfig) (syntheticorderv1.Consumer, error) {
	subscriber, err := subscribers.NewSyntheticOrderSubscriber(logger)
	if err != nil {
		return nil, err
	}
	osHelper, err := os.NewOsHelper()
	if err != nil {
		return nil, err
	}
	pulsarClient, err := pulsar.NewPulsarClient(logger, pulsarConfig, osHelper)
	if err != nil {
		return nil, err
	}
	messageConsumer, err := pulsar.NewPulsarMessageConsumer(pulsarClient)
	if err != nil {
		return nil, err
	}
	consumer := syntheticorderv1.NewConsumer(logger, subscriber, messageConsumer)
	return consumer, nil
}

func NewBinanceOrderBookSubscriber(ctx context.Context, logger *zap.SugaredLogger, appConfig configuration.AppConfig, binanceConfig configuration2.BinanceConfig, pairConfig configuration2.PairConfig, pulsarConfig pulsar.PulsarConfig, postgresConfig postgres.PostgresConfig, topic string) (subscribers.BinanceOrderBookSubscriber, error) {
	osHelper, err := os.NewOsHelper()
	if err != nil {
		return nil, err
	}
	pulsarClient, err := pulsar.NewPulsarClient(logger, pulsarConfig, osHelper)
	if err != nil {
		return nil, err
	}
	messageProducer, err := pulsar.NewPulsarMessageProducer(logger, pulsarClient)
	if err != nil {
		return nil, err
	}
	producer := orderbookv1.NewProducer(logger, messageProducer)
	orderBookPublisher, err := publishers.NewOrderBookPublisher(logger, appConfig, producer)
	if err != nil {
		return nil, err
	}
	database, err := postgres.NewPostgres(logger, postgresConfig)
	if err != nil {
		return nil, err
	}
	entgoClient, err := entities.NewEntgoClient(logger, database)
	if err != nil {
		return nil, err
	}
	coinHelper, err := services.NewCoinHelper(entgoClient)
	if err != nil {
		return nil, err
	}
	binanceOrderBookSubscriber, err := subscribers.NewBinanceOrderBookSubscriber(ctx, logger, pairConfig, orderBookPublisher, coinHelper)
	if err != nil {
		return nil, err
	}
	return binanceOrderBookSubscriber, nil
}

func NewBinanceTradingPairsSubscriber(ctx context.Context, logger *zap.SugaredLogger, binanceConfig configuration2.BinanceConfig, exchangeConfig configuration3.ExchangeConfig, cronService cron.CronService, postgresConfig postgres.PostgresConfig) (subscribers.BinanceTradingPairsSubscriber, error) {
	database, err := postgres.NewPostgres(logger, postgresConfig)
	if err != nil {
		return nil, err
	}
	entgoClient, err := entities.NewEntgoClient(logger, database)
	if err != nil {
		return nil, err
	}
	tradingPairRepository, err := repositories.NewTradingPairRepository(logger, entgoClient)
	if err != nil {
		return nil, err
	}
	binanceTradingPairsSubscriber, err := subscribers.NewBinanceTradingPairsSubscriber(ctx, logger, binanceConfig, exchangeConfig, cronService, tradingPairRepository)
	if err != nil {
		return nil, err
	}
	return binanceTradingPairsSubscriber, nil
}
