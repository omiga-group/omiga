// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"github.com/omiga-group/omiga/src/shared/enterprise/time"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/client"
	configuration2 "github.com/omiga-group/omiga/src/venue/ftx-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/subscribers"
	"github.com/omiga-group/omiga/src/venue/shared/publishers"
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

func NewSyntheticOrderConsumer(logger *zap.SugaredLogger, pulsarConfig pulsar.PulsarConfig) (syntheticorderv1.Consumer, error) {
	subscriber, err := subscribers.NewSyntheticOrderSubscriber(logger)
	if err != nil {
		return nil, err
	}
	osHelper, err := os.NewOsHelper()
	if err != nil {
		return nil, err
	}
	pulsarClient, err := pulsar.NewPulsarClient(logger, pulsarConfig, osHelper)
	if err != nil {
		return nil, err
	}
	messageConsumer, err := pulsar.NewPulsarMessageConsumer(pulsarClient)
	if err != nil {
		return nil, err
	}
	consumer := syntheticorderv1.NewConsumer(logger, subscriber, messageConsumer)
	return consumer, nil
}

func NewFtxOrderBookSubscriber(ctx context.Context, logger *zap.SugaredLogger, appConfig configuration.AppConfig, ftxConfig configuration2.FtxConfig, pulsarConfig pulsar.PulsarConfig, topic string) (subscribers.FtxOrderBookSubscriber, error) {
	apiClient := client.NewFtxApiClient(ftxConfig)
	osHelper, err := os.NewOsHelper()
	if err != nil {
		return nil, err
	}
	pulsarClient, err := pulsar.NewPulsarClient(logger, pulsarConfig, osHelper)
	if err != nil {
		return nil, err
	}
	messageProducer, err := pulsar.NewPulsarMessageProducer(logger, pulsarClient)
	if err != nil {
		return nil, err
	}
	producer := orderbookv1.NewProducer(logger, messageProducer)
	orderBookPublisher, err := publishers.NewOrderBookPublisher(logger, appConfig, producer)
	if err != nil {
		return nil, err
	}
	ftxOrderBookSubscriber, err := subscribers.NewFtxOrderBookSubscriber(ctx, logger, apiClient, ftxConfig, orderBookPublisher)
	if err != nil {
		return nil, err
	}
	return ftxOrderBookSubscriber, nil
}