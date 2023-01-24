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

	"github.com/go-co-op/gocron"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	enterpriseConfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"github.com/omiga-group/omiga/src/venue/gemini-processor/client"
	"github.com/omiga-group/omiga/src/venue/gemini-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/gemini-processor/subscribers"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
)

func NewSyntheticOrderConsumer(
	logger *zap.SugaredLogger,
	pulsarClient pulsar.PulsarClient,
	pulsarConfig pulsar.PulsarConfig) (syntheticorderv1.Consumer, error) {
	wire.Build(
		pulsar.NewPulsarMessageConsumer,
		syntheticorderv1.NewConsumer,
		subscribers.NewSyntheticOrderSubscriber,
		time.NewTimeHelper)

	return nil, nil
}

func NewGeminiOrderBookSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pulsarClient pulsar.PulsarClient,
	appConfig enterpriseConfiguration.AppConfig,
	venueConfig configuration.GeminiConfig,
	pulsarConfig pulsar.PulsarConfig,
	topic string) (subscribers.GeminiOrderBookSubscriber, error) {
	wire.Build(
		orderbookv1.NewProducer,
		client.NewGeminiApiClient,
		pulsar.NewPulsarMessageProducer,
		publishers.NewOrderBookPublisher,
		subscribers.NewGeminiOrderBookSubscriber)

	return nil, nil
}

func NewGeminiTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.GeminiConfig,
	jobScheduler *gocron.Scheduler,
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
