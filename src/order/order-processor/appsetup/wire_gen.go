// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"github.com/omiga-group/omiga/src/order/order-processor/subscribers"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewOrderConsumer(logger *zap.SugaredLogger, pulsarClient pulsar.PulsarClient, pulsarConfig pulsar.PulsarConfig) (orderv1.Consumer, error) {
	subscriber, err := subscribers.NewOrderSubscriber(logger)
	if err != nil {
		return nil, err
	}
	messageConsumer, err := pulsar.NewPulsarMessageConsumer(logger, pulsarClient, pulsarConfig)
	if err != nil {
		return nil, err
	}
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	consumer := orderv1.NewConsumer(logger, subscriber, messageConsumer, timeHelper)
	return consumer, nil
}
