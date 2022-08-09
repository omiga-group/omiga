package pulsar

import (
	"context"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type pulsarMessageProducer struct {
	logger         *zap.SugaredLogger
	pulsarSettings PulsarSettings
	pulsarClient   pulsar.Client
	producer       pulsar.Producer
	topic          string
}

func NewPulsarMessageProducer(
	logger *zap.SugaredLogger,
	pulsarSettings PulsarSettings,
	topic string) (messaging.MessageProducer, error) {
	pulsarClient, err := pulsar.NewClient(
		pulsar.ClientOptions{
			URL:               pulsarSettings.Url,
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
	if err != nil {
		return nil, err
	}

	producer, err := pulsarClient.CreateProducer(
		pulsar.ProducerOptions{
			Topic: topic,
			Name:  pulsarSettings.ProducerName,
		})
	if err != nil {
		return nil, err
	}

	return &pulsarMessageProducer{
		logger:         logger,
		pulsarSettings: pulsarSettings,
		pulsarClient:   pulsarClient,
		producer:       producer,
	}, nil
}

func (pmp *pulsarMessageProducer) Close(ctx context.Context) {
	if pmp.producer != nil {
		if err := pmp.producer.Flush(); err != nil {
			pmp.logger.Errorf("Failed to flush. Error: %v", err)
		}

		pmp.producer.Close()
		pmp.producer = nil
	}

	if pmp.pulsarClient == nil {
		pmp.pulsarClient.Close()
		pmp.pulsarClient = nil
	}
}

func (pmp *pulsarMessageProducer) Produce(
	ctx context.Context,
	topic string,
	key string,
	data []byte) error {
	_, err := pmp.producer.Send(ctx, &pulsar.ProducerMessage{
		Key:     key,
		Payload: data,
	})
	if err != nil {
		pmp.logger.Errorf("Failed to send message on topic: %s. Error: %v", pmp.topic, err)

		return err
	}

	return nil
}
