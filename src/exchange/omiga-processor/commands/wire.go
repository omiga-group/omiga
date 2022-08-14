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
	"github.com/omiga-group/omiga/src/exchange/omiga-processor/subscribers"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"go.uber.org/zap"
)

func NewMessageConsumer(logger *zap.SugaredLogger, pulsarSettings pulsar.PulsarSettings, topic string) (messaging.MessageConsumer, error) {
	wire.Build(pulsar.NewPulsarMessageConsumer)

	return nil, nil
}

func NewSyntheticOrderConsumer(logger *zap.SugaredLogger, messageConsumer messaging.MessageConsumer) (syntheticorderv1.Consumer, error) {
	wire.Build(syntheticorderv1.NewConsumer, subscribers.NewSyntheticOrderSubscriber)

	return nil, nil
}
