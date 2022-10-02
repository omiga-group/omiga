// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package appsetup

import (
	"context"

	"github.com/google/wire"
	"github.com/omiga-group/omiga/src/exchange/coingecko-processor/configuration"
	coingeckorepositories "github.com/omiga-group/omiga/src/exchange/coingecko-processor/repositories"
	"github.com/omiga-group/omiga/src/exchange/coingecko-processor/subscribers"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

func NewCronService(
	logger *zap.SugaredLogger) (cron.CronService, error) {
	wire.Build(
		time.NewTimeHelper,
		cron.NewCronService)

	return nil, nil
}

func NewTimeHelper() (time.TimeHelper, error) {
	wire.Build(
		time.NewTimeHelper)

	return nil, nil
}

func NewCoingeckoExchangeSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingeckoConfig configuration.CoingeckoConfig,
	exchanges map[string]configuration.Exchange,
	postgresConfig postgres.PostgresConfig) (subscribers.CoingeckoExchangeSubscriber, error) {
	wire.Build(
		postgres.NewPostgres,
		repositories.NewEntgoClient,
		time.NewTimeHelper,
		subscribers.NewCoingeckoExchangeSubscriber,
		coingeckorepositories.NewExchangeRepository)

	return nil, nil
}

func NewCoingeckoCoinSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingeckoConfig configuration.CoingeckoConfig,
	exchanges map[string]configuration.Exchange,
	postgresConfig postgres.PostgresConfig) (subscribers.CoingeckoCoinSubscriber, error) {
	wire.Build(
		postgres.NewPostgres,
		repositories.NewEntgoClient,
		time.NewTimeHelper,
		subscribers.NewCoingeckoCoinSubscriber,
		coingeckorepositories.NewCoinRepository)

	return nil, nil
}
