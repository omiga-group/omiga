// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package commands

import (
	"github.com/omiga-group/omiga/src/exchange/omiga-processor/subscribers"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewMessageConsumer(logger *zap.SugaredLogger, pulsarSettings pulsar.PulsarSettings, topic string) (messaging.MessageConsumer, error) {
	messageConsumer, err := pulsar.NewPulsarMessageConsumer(logger, pulsarSettings, topic)
	if err != nil {
		return nil, err
	}
	return messageConsumer, nil
}

func NewSyntheticOrderConsumer(logger *zap.SugaredLogger, messageConsumer messaging.MessageConsumer) (syntheticorderv1.Consumer, error) {
	subscriber, err := subscribers.NewSyntheticOrderSubscriber(logger)
	if err != nil {
		return nil, err
	}
	consumer := syntheticorderv1.NewConsumer(logger, subscriber, messageConsumer)
	return consumer, nil
}
