// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/client"
	configuration2 "github.com/omiga-group/omiga/src/exchange/gemini-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/subscribers"
	"github.com/omiga-group/omiga/src/exchange/shared/publishers"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
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

func NewGeminiOrderBookSubscriber(ctx context.Context, logger *zap.SugaredLogger, appConfig configuration.AppConfig, geminiConfig configuration2.GeminiConfig, pulsarConfig pulsar.PulsarConfig, topic string) (subscribers.GeminiOrderBookSubscriber, error) {
	apiClient := client.NewGeminiApiClient(geminiConfig)
	messageProducer, err := pulsar.NewPulsarMessageProducer(logger, pulsarConfig, topic)
	if err != nil {
		return nil, err
	}
	producer := orderbookv1.NewProducer(logger, messageProducer)
	orderBookPublisher, err := publishers.NewOrderBookPublisher(logger, appConfig, producer)
	if err != nil {
		return nil, err
	}
	geminiOrderBookSubscriber, err := subscribers.NewGeminiOrderBookSubscriber(ctx, logger, apiClient, geminiConfig, orderBookPublisher)
	if err != nil {
		return nil, err
	}
	return geminiOrderBookSubscriber, nil
}