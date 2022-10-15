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
	"go.uber.org/zap"

	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	enterpriseConfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"github.com/omiga-group/omiga/src/venue/gemini-processor/client"
	"github.com/omiga-group/omiga/src/venue/gemini-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/gemini-processor/subscribers"
	venueConfiguration "github.com/omiga-group/omiga/src/venue/shared/configuration"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
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

func NewGeminiOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	appConfig enterpriseConfiguration.AppConfig,
	geminiConfig configuration.GeminiConfig,
	pulsarConfig pulsar.PulsarConfig,
	topic string) (subscribers.GeminiOrderBookSubscriber, error) {
	wire.Build(
		os.NewOsHelper,
		orderbookv1.NewProducer,
		client.NewGeminiApiClient,
		pulsar.NewPulsarClient,
		pulsar.NewPulsarMessageProducer,
		publishers.NewOrderBookPublisher,
		subscribers.NewGeminiOrderBookSubscriber,
	)

	return nil, nil
}

func NewGeminiTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	geminiConfig configuration.GeminiConfig,
	venueConfig venueConfiguration.VenueConfig,
	cronService cron.CronService,
	postgresConfig postgres.PostgresConfig) (subscribers.GeminiTradingPairSubscriber, error) {
	wire.Build(
		postgres.NewPostgres,
		entities.NewEntgoClient,
		repositories.NewCurrencyRepository,
		repositories.NewVenueRepository,
		repositories.NewTradingPairRepository,
		subscribers.NewGeminiTradingPairSubscriber)

	return nil, nil
}
