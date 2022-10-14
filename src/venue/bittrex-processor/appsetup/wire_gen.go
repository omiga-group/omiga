// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"github.com/omiga-group/omiga/src/venue/bittrex-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/bittrex-processor/subscribers"
	configuration2 "github.com/omiga-group/omiga/src/venue/shared/configuration"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
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

func NewBittrexTradingPairSubscriber(ctx context.Context, logger *zap.SugaredLogger, bittrexConfig configuration.BittrexConfig, exchangeConfig configuration2.VenueConfig, cronService cron.CronService, postgresConfig postgres.PostgresConfig) (subscribers.BittrexTradingPairSubscriber, error) {
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
	bittrexTradingPairSubscriber, err := subscribers.NewBittrexTradingPairSubscriber(ctx, logger, bittrexConfig, exchangeConfig, cronService, tradingPairRepository)
	if err != nil {
		return nil, err
	}
	return bittrexTradingPairSubscriber, nil
}
