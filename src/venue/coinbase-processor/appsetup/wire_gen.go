// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/venue/coinbase-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/coinbase-processor/subscribers"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewSyntheticOrderConsumer(logger *zap.SugaredLogger, pulsarClient pulsar.PulsarClient, pulsarConfig pulsar.PulsarConfig) (syntheticorderv1.Consumer, error) {
	subscriber, err := subscribers.NewSyntheticOrderSubscriber(logger)
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

func NewCoinbaseTradingPairSubscriber(ctx context.Context, logger *zap.SugaredLogger, venueConfig configuration.CoinbaseConfig, jobScheduler *gocron.Scheduler, postgresConfig postgres.PostgresConfig) (subscribers.CoinbaseTradingPairSubscriber, error) {
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
	coinbaseTradingPairSubscriber, err := subscribers.NewCoinbaseTradingPairSubscriber(ctx, logger, venueConfig, jobScheduler, tradingPairRepository)
	if err != nil {
		return nil, err
	}
	return coinbaseTradingPairSubscriber, nil
}
