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
package commands

import (
	"github.com/google/wire"
	"github.com/omiga-group/omiga/src/order/order-processor/services"
	"github.com/omiga-group/omiga/src/order/order-processor/subscribers"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"go.uber.org/zap"
)

func NewEntgoClient(
	logger *zap.SugaredLogger,
	postgresConfig postgres.PostgresConfig) (repositories.EntgoClient, error) {
	wire.Build(
		postgres.NewPostgres,
		repositories.NewEntgoClient)

	return nil, nil
}

func NewMessageConsumer(
	logger *zap.SugaredLogger,
	pulsarConfig pulsar.PulsarConfig,
	topic string) (messaging.MessageConsumer, error) {
	wire.Build(pulsar.NewPulsarMessageConsumer)

	return nil, nil
}

func NewOrderConsumer(
	logger *zap.SugaredLogger,
	messageConsumer messaging.MessageConsumer) (orderv1.Consumer, error) {
	wire.Build(orderv1.NewConsumer, subscribers.NewOrderSubscriber)

	return nil, nil
}

func NewOrderBookConsumer(
	logger *zap.SugaredLogger,
	messageConsumer messaging.MessageConsumer,
	entgoClient repositories.EntgoClient) (orderbookv1.Consumer, error) {
	wire.Build(orderbookv1.NewConsumer, subscribers.NewOrderBookSubscriber, services.NewOrderBookService)

	return nil, nil
}
