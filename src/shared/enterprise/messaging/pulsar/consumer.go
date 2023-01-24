package pulsar

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type pulsarMessageConsumer struct {
	logger       *zap.SugaredLogger
	pulsarClient PulsarClient
	pulsarConfig PulsarConfig
	consumer     pulsar.Consumer
}

func NewPulsarMessageConsumer(
	logger *zap.SugaredLogger,
	pulsarClient PulsarClient,
	pulsarConfig PulsarConfig) (messaging.MessageConsumer, error) {
	return &pulsarMessageConsumer{
		logger:       logger,
		pulsarClient: pulsarClient,
		pulsarConfig: pulsarConfig,
	}, nil
}

func (pmc *pulsarMessageConsumer) Connect(topic string) error {
	if pmc.consumer != nil {
		return nil
	}

	client, err := pmc.pulsarClient.CreatePulsarClient()
	if err != nil {
		return err
	}

	pmc.consumer, err = client.Subscribe(
		pulsar.ConsumerOptions{
			Topic:            topic,
			SubscriptionName: pmc.pulsarConfig.SubscriptionName,
			Type:             pulsar.KeyShared,
		})
	if err != nil {
		return err
	}

	return nil
}

func (pmc *pulsarMessageConsumer) Consume(ctx context.Context) (
	messaging.Message,
	messaging.MessageProcessedCallback,
	messaging.MessageFailedCallback,
	error) {
	msg, err := pmc.consumer.Receive(ctx)
	if err != nil {
		return messaging.Message{}, nil, nil, err
	}

	messageProcessedCallback := func() {
		if pmc.consumer != nil {
			if ackErr := pmc.consumer.Ack(msg); ackErr != nil {
				pmc.logger.Errorf("Failed to ack message. Error: %v", err)

			}
		}
	}

	messageFailedCallback := func() {
		if pmc.consumer != nil {
			pmc.consumer.Nack(msg)
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

func (pmc *pulsarMessageConsumer) Close() {
	if pmc.consumer != nil {
		pmc.consumer.Close()
	}
}
