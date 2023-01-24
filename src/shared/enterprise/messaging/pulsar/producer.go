package pulsar

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type pulsarMessageProducer struct {
	logger       *zap.SugaredLogger
	pulsarClient PulsarClient
}

func NewPulsarMessageProducer(
	logger *zap.SugaredLogger,
	pulsarClient PulsarClient) (messaging.MessageProducer, error) {
	return &pulsarMessageProducer{
		logger:       logger,
		pulsarClient: pulsarClient,
	}, nil
}

func (pmp *pulsarMessageProducer) Close() {
}

func (pmp *pulsarMessageProducer) Produce(
	ctx context.Context,
	topic string,
	key string,
	data []byte) error {
	producer, err := pmp.pulsarClient.CreateProducer(topic)
	if err != nil {
		return err
	}

	_, err = producer.Send(ctx, &pulsar.ProducerMessage{
		Key:     key,
		Payload: data,
	})
	if err != nil {
		pmp.logger.Errorf("Failed to send message on topic: %s. Error: %v", topic, err)

		return err
	}

	return nil
}
