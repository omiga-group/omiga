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
	"github.com/omiga-group/omiga/src/order/order-api/graphql"
	"github.com/omiga-group/omiga/src/order/order-api/http"
	"github.com/omiga-group/omiga/src/order/order-api/publishers"
	orderrepositories "github.com/omiga-group/omiga/src/order/order-api/repositories"
	"github.com/omiga-group/omiga/src/order/order-api/services"
	"github.com/omiga-group/omiga/src/order/shared/outbox"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	enterpriseOutbox "github.com/omiga-group/omiga/src/shared/enterprise/outbox"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

func NewCronService(
	logger *zap.SugaredLogger) (cron.CronService, error) {
	wire.Build(
		cron.NewCronService,
		time.NewTimeHelper)

	return nil, nil
}

func NewEntgoClient(
	logger *zap.SugaredLogger,
	postgresConfig postgres.PostgresConfig) (repositories.EntgoClient, error) {
	wire.Build(
		postgres.NewPostgres,
		repositories.NewEntgoClient)

	return nil, nil
}

func NewOutboxBackgroundService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pulsarConfig pulsar.PulsarConfig,
	outboxConfig enterpriseOutbox.OutboxConfig,
	entgoClient repositories.EntgoClient,
	cronService cron.CronService) (outbox.OutboxBackgroundService, error) {
	wire.Build(
		os.NewOsHelper,
		pulsar.NewPulsarClient,
		pulsar.NewPulsarMessageProducer,
		outbox.NewOutboxBackgroundService)

	return nil, nil
}

func NewHttpServer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	entgoClient repositories.EntgoClient,
	orderOutboxBackgroundService outbox.OutboxBackgroundService) (http.HttpServer, error) {
	wire.Build(
		http.NewHttpServer,
		graphql.NewGraphQLServer,
		services.NewOrderService,
		publishers.NewOrderPublisher,
		outbox.NewOutboxPublisher,
		orderrepositories.NewOrderRepository)

	return nil, nil
}
