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
	"github.com/omiga-group/omiga/src/exchange/coingeko/configuration"
	coingekorepositories "github.com/omiga-group/omiga/src/exchange/coingeko/repositories"
	"github.com/omiga-group/omiga/src/exchange/coingeko/subscribers"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

func NewTimeHelper() (time.TimeHelper, error) {
	wire.Build(
		time.NewTimeHelper)

	return nil, nil
}

func NewCoingekoSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	coingekoConfig configuration.CoingekoConfig,
	postgresConfig postgres.PostgresConfig) (subscribers.CoingekoSubscriber, error) {
	wire.Build(
		postgres.NewPostgres,
		repositories.NewEntgoClient,
		cron.NewCronService,
		time.NewTimeHelper,
		subscribers.NewCoingekoSubscriber,
		coingekorepositories.NewExchangeRepository)

	return nil, nil
}
