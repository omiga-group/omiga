// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	configuration2 "github.com/omiga-group/omiga/src/venue/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/binance-processor/subscribers"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"github.com/omiga-group/omiga/src/venue/shared/services"
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
	messageConsumer, err := pulsar.NewPulsarMessageConsumer(logger, pulsarClient)
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
	currencyHelper, err := services.NewCurrencyHelper(entgoClient)
	if err != nil {
		return nil, err
	}
	binanceOrderBookSubscriber, err := subscribers.NewBinanceOrderBookSubscriber(ctx, logger, pairConfig, orderBookPublisher, currencyHelper)
	if err != nil {
		return nil, err
	}
	return binanceOrderBookSubscriber, nil
}

func NewBinanceTradingPairSubscriber(ctx context.Context, logger *zap.SugaredLogger, venueConfig configuration2.BinanceConfig, cronService cron.CronService, postgresConfig postgres.PostgresConfig) (subscribers.BinanceTradingPairSubscriber, error) {
	database, err := postgres.NewPostgres(logger, postgresConfig)
	if err != nil {
		return nil, err
	}
	entgoClient, err := entities.NewEntgoClient(logger, database)
	if err != nil {
		return nil, err
	}
	currencyRepository, err := repositories.NewCurrencyRepository(logger, entgoClient)
	if err != nil {
		return nil, err
	}
	venueRepository, err := repositories.NewVenueRepository(logger, entgoClient)
	if err != nil {
		return nil, err
	}
	tradingPairRepository, err := repositories.NewTradingPairRepository(logger, entgoClient, currencyRepository, venueRepository)
	if err != nil {
		return nil, err
	}
	binanceTradingPairSubscriber, err := subscribers.NewBinanceTradingPairSubscriber(ctx, logger, venueConfig, cronService, tradingPairRepository)
	if err != nil {
		return nil, err
	}
	return binanceTradingPairSubscriber, nil
}
