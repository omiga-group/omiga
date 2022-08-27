// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package commands

import (
	"context"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/subscribers"
	"github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
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

func NewBinanceOrderBookSubscriber(ctx context.Context, logger *zap.SugaredLogger, binanceSettings configuration.BinanceSettings, symbol string) (subscribers.BinanceOrderBookSubscriber, error) {
	binanceOrderBookSubscriber, err := subscribers.NewBinanceOrderBookSubscriber(ctx, logger, binanceSettings, symbol)
	if err != nil {
		return nil, err
	}
	return binanceOrderBookSubscriber, nil
}
