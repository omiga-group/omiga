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
	"github.com/omiga-group/omiga/src/order/shared/entities"
	"github.com/omiga-group/omiga/src/order/shared/outbox"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	enterpriseOutbox "github.com/omiga-group/omiga/src/shared/enterprise/outbox"
	"go.uber.org/zap"
)

func NewEntgoClient(
	logger *zap.SugaredLogger,
	postgresConfig postgres.PostgresConfig) (entities.EntgoClient, error) {
	wire.Build(
		postgres.NewPostgres,
		entities.NewEntgoClient)

	return nil, nil
}

func NewOutboxBackgroundService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	pulsarClient pulsar.PulsarClient,
	pulsarConfig pulsar.PulsarConfig,
	outboxConfig enterpriseOutbox.OutboxConfig,
	entgoClient entities.EntgoClient,
	jobScheduler *gocron.Scheduler) (outbox.OutboxBackgroundService, error) {
	wire.Build(
		pulsar.NewPulsarMessageProducer,
		outbox.NewOutboxBackgroundService)

	return nil, nil
}
