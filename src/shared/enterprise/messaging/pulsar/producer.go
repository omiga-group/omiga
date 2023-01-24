package pulsar

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/lucsky/cuid"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type pulsarMessageProducer struct {
	logger       *zap.SugaredLogger
	pulsarClient PulsarClient
	pulsarConfig PulsarConfig
	producer     pulsar.Producer
}

func NewPulsarMessageProducer(
	logger *zap.SugaredLogger,
	pulsarClient PulsarClient,
	pulsarConfig PulsarConfig) (messaging.MessageProducer, error) {
	return &pulsarMessageProducer{
		logger:       logger,
		pulsarClient: pulsarClient,
		pulsarConfig: pulsarConfig,
	}, nil
}

func (pmp *pulsarMessageProducer) Connect(topic string) error {
	if pmp.producer != nil {
		return nil
	}

	client, err := pmp.pulsarClient.CreatePulsarClient()
	if err != nil {
		return err
	}

	pmp.producer, err = client.CreateProducer(
		pulsar.ProducerOptions{
			Topic: topic,
			Name:  pmp.pulsarConfig.ProducerName + "-" + cuid.New(),
		})
	if err != nil {
		return err
	}

	return nil
}

func (pmp *pulsarMessageProducer) Produce(
	ctx context.Context,
	key string,
	data []byte) error {
	_, err := pmp.producer.Send(ctx, &pulsar.ProducerMessage{
		Key:     key,
		Payload: data,
	})
	if err != nil {
		pmp.logger.Errorf("Failed to send message on topic: %s, producer name: %s. Error: %v", pmp.producer.Topic(), pmp.producer.Name(), err)

		return err
	}

	return nil
}

func (pmp *pulsarMessageProducer) Close() {
	if pmp.producer != nil {
		if err := pmp.producer.Flush(); err != nil {
			pmp.logger.Errorf("Failed to flush messages on topic: %s, producer name: %s. Error: %v", pmp.producer.Topic(), pmp.producer.Name(), err)
		}

		pmp.producer.Close()
	}
}
