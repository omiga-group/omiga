package pulsar

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
)

type pulsarMessageConsumer struct {
	pulsarClient PulsarClient
}

func NewPulsarMessageConsumer(
	pulsarClient PulsarClient) (messaging.MessageConsumer, error) {
	return &pulsarMessageConsumer{
		pulsarClient: pulsarClient,
	}, nil
}

func (pmc *pulsarMessageConsumer) Close() {
	pmc.pulsarClient.Close()
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
			consumer.Ack(msg)
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
