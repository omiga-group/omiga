// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"github.com/omiga-group/omiga/src/order/order-processor/subscribers"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewTimeHelper() (time.TimeHelper, error) {
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	return timeHelper, nil
}

func NewMessageConsumer(logger *zap.SugaredLogger, pulsarConfig pulsar.PulsarConfig, topic string) (messaging.MessageConsumer, error) {
	osHelper, err := os.NewOsHelper()
	if err != nil {
		return nil, err
	}
	messageConsumer, err := pulsar.NewPulsarMessageConsumer(logger, pulsarConfig, osHelper, topic)
	if err != nil {
		return nil, err
	}
	return messageConsumer, nil
}

func NewOrderConsumer(logger *zap.SugaredLogger, messageConsumer messaging.MessageConsumer) (orderv1.Consumer, error) {
	subscriber, err := subscribers.NewOrderSubscriber(logger)
	if err != nil {
		return nil, err
	}
	consumer := orderv1.NewConsumer(logger, subscriber, messageConsumer)
	return consumer, nil
}
