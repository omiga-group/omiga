// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	configuration2 "github.com/omiga-group/omiga/src/venue/ftx-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/subscribers"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewSyntheticOrderConsumer(logger *zap.SugaredLogger, pulsarClient pulsar.PulsarClient, pulsarConfig pulsar.PulsarConfig) (syntheticorderv1.Consumer, error) {
	subscriber, err := subscribers.NewSyntheticOrderSubscriber(logger)
	if err != nil {
		return nil, err
	}
	messageConsumer, err := pulsar.NewPulsarMessageConsumer(logger, pulsarClient, pulsarConfig)
	if err != nil {
		return nil, err
	}
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	consumer := syntheticorderv1.NewConsumer(logger, subscriber, messageConsumer, timeHelper)
	return consumer, nil
}

func NewFtxOrderBookSubscriber(ctx context.Context, logger *zap.SugaredLogger, pulsarClient pulsar.PulsarClient, appConfig configuration.AppConfig, venueConfig configuration2.FtxConfig, pulsarConfig pulsar.PulsarConfig, topic string) (subscribers.FtxOrderBookSubscriber, error) {
	messageProducer, err := pulsar.NewPulsarMessageProducer(logger, pulsarClient, pulsarConfig)
	if err != nil {
		return nil, err
	}
	producer := orderbookv1.NewProducer(logger, messageProducer)
	orderBookPublisher, err := publishers.NewOrderBookPublisher(logger, appConfig, producer)
	if err != nil {
		return nil, err
	}
	ftxOrderBookSubscriber, err := subscribers.NewFtxOrderBookSubscriber(ctx, logger, venueConfig, orderBookPublisher)
	if err != nil {
		return nil, err
	}
	return ftxOrderBookSubscriber, nil
}

func NewFtxTradingPairSubscriber(ctx context.Context, logger *zap.SugaredLogger, venueConfig configuration2.FtxConfig, jobScheduler *gocron.Scheduler, postgresConfig postgres.PostgresConfig) (subscribers.FtxTradingPairSubscriber, error) {
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
	ftxTradingPairSubscriber, err := subscribers.NewFtxTradingPairSubscriber(ctx, logger, venueConfig, jobScheduler, tradingPairRepository)
	if err != nil {
		return nil, err
	}
	return ftxTradingPairSubscriber, nil
}
