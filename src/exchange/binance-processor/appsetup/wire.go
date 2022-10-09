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
	"github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/subscribers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/services"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	enterpriseConfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
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

func NewSyntheticOrderConsumer(
	logger *zap.SugaredLogger,
	pulsarConfig pulsar.PulsarConfig) (syntheticorderv1.Consumer, error) {
	wire.Build(
		os.NewOsHelper,
		pulsar.NewPulsarClient,
		pulsar.NewPulsarMessageConsumer,
		syntheticorderv1.NewConsumer,
		subscribers.NewSyntheticOrderSubscriber)

	return nil, nil
}

func NewBinanceOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	appConfig enterpriseConfiguration.AppConfig,
	binanceConfig configuration.BinanceConfig,
	pairConfig configuration.PairConfig,
	pulsarConfig pulsar.PulsarConfig,
	postgresConfig postgres.PostgresConfig,
	topic string) (subscribers.BinanceOrderBookSubscriber, error) {
	wire.Build(
		os.NewOsHelper,
		postgres.NewPostgres,
		entities.NewEntgoClient,
		orderbookv1.NewProducer,
		pulsar.NewPulsarClient,
		pulsar.NewPulsarMessageProducer,
		publishers.NewOrderBookPublisher,
		subscribers.NewBinanceOrderBookSubscriber,
		services.NewCoinHelper)

	return nil, nil
}

func NewBinanceTradingPairsSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	binanceConfig configuration.BinanceConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	postgresConfig postgres.PostgresConfig) (subscribers.BinanceTradingPairSubscriber, error) {
	wire.Build(
		postgres.NewPostgres,
		entities.NewEntgoClient,
		repositories.NewCoinRepository,
		repositories.NewExchangeRepository,
		repositories.NewTradingPairRepository,
		subscribers.NewBinanceTradingPairsSubscriber)

	return nil, nil
}
