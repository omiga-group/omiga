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

	"github.com/go-co-op/gocron"
	"github.com/google/wire"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"github.com/omiga-group/omiga/src/venue/xt-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/xt-processor/subscribers"
	"go.uber.org/zap"
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

func NewXtTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.XtConfig,
	jobScheduler *gocron.Scheduler,
	postgresConfig postgres.PostgresConfig) (subscribers.XtTradingPairSubscriber, error) {
	wire.Build(
		postgres.NewPostgres,
		entities.NewEntgoClient,
		repositories.NewCurrencyRepository,
		repositories.NewVenueRepository,
		repositories.NewTradingPairRepository,
		subscribers.NewXtTradingPairSubscriber)

	return nil, nil
}
