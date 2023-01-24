package pulsar

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type pulsarMessageConsumer struct {
	logger       *zap.SugaredLogger
	pulsarClient PulsarClient
}

func NewPulsarMessageConsumer(
	logger *zap.SugaredLogger,
	pulsarClient PulsarClient) (messaging.MessageConsumer, error) {
	return &pulsarMessageConsumer{
		logger:       logger,
		pulsarClient: pulsarClient,
	}, nil
}

func (pmc *pulsarMessageConsumer) Close() {
}

func (pmc *pulsarMessageConsumer) Consume(ctx context.Context, topic string) (
	messaging.Message,
	messaging.MessageProcessedCallback,
	messaging.MessageFailedCallback,
	error) {
	consumer, err := pmc.pulsarClient.CreateConsumer(topic)
	if err != nil {
		return messaging.Message{}, nil, nil, err
	}

	msg, err := consumer.Receive(ctx)
	if err != nil {
		return messaging.Message{}, nil, nil, err
	}

	messageProcessedCallback := func() {
		if consumer != nil {
			if ackErr := consumer.Ack(msg); ackErr != nil {
				pmc.logger.Errorf("Failed to ack message. Error: %v", err)

			}
		}
	}

	messageFailedCallback := func() {
		if consumer != nil {
			consumer.Nack(msg)
		}
	}

	return messaging.Message{
			Topic:       msg.Topic(),
			Key:         msg.Key(),
			Payload:     msg.Payload(),
			Headers:     msg.Properties(),
			PublishTime: msg.PublishTime(),
			EventTime:   msg.EventTime(),
		},
		messageProcessedCallback,
		messageFailedCallback,
		nil
}
