// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package commands

import (
	"context"
	"github.com/omiga-group/omiga/src/exchange/coingeko/configuration"
	"github.com/omiga-group/omiga/src/exchange/coingeko/subscribers"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewTimeHelper() (time.TimeHelper, error) {
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	return timeHelper, nil
}

func NewCoingekoSubscriber(ctx context.Context, logger *zap.SugaredLogger, coingekoConfig configuration.CoingekoConfig, postgresConfig postgres.PostgresConfig) (subscribers.CoingekoSubscriber, error) {
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	cronService, err := cron.NewCronService(logger, timeHelper)
	if err != nil {
		return nil, err
	}
	database, err := postgres.NewPostgres(logger, postgresConfig)
	if err != nil {
		return nil, err
	}
	entgoClient, err := repositories.NewEntgoClient(logger, database)
	if err != nil {
		return nil, err
	}
	coingekoSubscriber, err := subscribers.NewCoingekoSubscriber(ctx, logger, cronService, coingekoConfig, entgoClient, timeHelper)
	if err != nil {
		return nil, err
	}
	return coingekoSubscriber, nil
}
