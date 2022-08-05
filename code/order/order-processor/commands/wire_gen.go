// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package commands

import (
	"github.com/omiga-group/omiga/code/order/order-processor/subscribers"
	"github.com/omiga-group/omiga/code/shared/events/events/omiga/order/v1"
	"github.com/omiga-group/omiga/code/shared/messaging/pulsar"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewOrderConsumer(logger *zap.SugaredLogger, pulsarSettings pulsar.PulsarSettings) (orderv1.Consumer, error) {
	subscriber, err := subscribers.NewOrderSubscriber(logger)
	if err != nil {
		return nil, err
	}
	messageConsumerService, err := pulsar.NewPulsarMessageConsumerService(logger, pulsarSettings)
	if err != nil {
		return nil, err
	}
	consumer := orderv1.NewConsumer(logger, subscriber, messageConsumerService)
	return consumer, nil
}
