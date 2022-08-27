// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package commands

import (
	"context"
	"github.com/omiga-group/omiga/src/exchange/omiga-processor/simulators"
	"github.com/omiga-group/omiga/src/exchange/omiga-processor/subscribers"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewMessageConsumer(logger *zap.SugaredLogger, pulsarConfig pulsar.PulsarConfig, topic string) (messaging.MessageConsumer, error) {
	messageConsumer, err := pulsar.NewPulsarMessageConsumer(logger, pulsarConfig, topic)
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

func NewOrderBookSimulator(ctx context.Context, logger *zap.SugaredLogger, appConfig configuration.AppConfig, pulsarConfig pulsar.PulsarConfig, topic string, orderBookSimulatorSettings simulators.OrderBookSimulatorSettings) (simulators.OrderBookSimulator, error) {
	timeHelper, err := time.NewTimeHelper()
	if err != nil {
		return nil, err
	}
	cronService, err := cron.NewCronService(logger, timeHelper)
	if err != nil {
		return nil, err
	}
	messageProducer, err := pulsar.NewPulsarMessageProducer(logger, pulsarConfig, topic)
	if err != nil {
		return nil, err
	}
	producer := orderbookv1.NewProducer(logger, messageProducer)
	orderBookPublisher, err := publishers.NewOrderBookPublisher(logger, appConfig, producer)
	if err != nil {
		return nil, err
	}
	orderBookSimulator, err := simulators.NewOrderBookSimulator(ctx, logger, cronService, orderBookPublisher, orderBookSimulatorSettings)
	if err != nil {
		return nil, err
	}
	return orderBookSimulator, nil
}
