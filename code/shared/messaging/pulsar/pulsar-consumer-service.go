package pulsar

import (
	"context"
	"errors"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/omiga-group/omiga/code/shared/messaging"
	"go.uber.org/zap"
)

type pulsarMessageConsumerService struct {
	logger         *zap.SugaredLogger
	pulsarSettings PulsarSettings
	pulsarClient   pulsar.Client
	consumer       pulsar.Consumer
}

func NewPulsarMessageConsumerService(
	logger *zap.SugaredLogger,
	pulsarSettings PulsarSettings) (messaging.MessageConsumerService, error) {

	return &pulsarMessageConsumerService{
		logger:         logger,
		pulsarSettings: pulsarSettings,
	}, nil
}

func (pmcs *pulsarMessageConsumerService) Connect(ctx context.Context, topic string) error {
	if pmcs.pulsarClient == nil {
		pulsarClient, err := pulsar.NewClient(
			pulsar.ClientOptions{
				URL:               pmcs.pulsarSettings.Url,
				OperationTimeout:  30 * time.Second,
				ConnectionTimeout: 30 * time.Second,
			})
		if err != nil {
			return err
		}

		pmcs.pulsarClient = pulsarClient
	}

	if pmcs.consumer == nil {
		consumer, err := pmcs.pulsarClient.Subscribe(
			pulsar.ConsumerOptions{
				Topic:            topic,
				SubscriptionName: pmcs.pulsarSettings.SubscriptionName,
			})
		if err != nil {
			return err
		}

		pmcs.consumer = consumer
	}

	return nil
}

func (pmcs *pulsarMessageConsumerService) Diconnect(ctx context.Context) {
	if pmcs.consumer != nil {
		if err := pmcs.consumer.Unsubscribe(); err != nil {
			pmcs.logger.Errorf("Failed to unsubscribe. Error: %v", err)
		}

		pmcs.consumer.Close()
		pmcs.consumer = nil
	}

	if pmcs.pulsarClient == nil {
		pmcs.pulsarClient.Close()
		pmcs.pulsarClient = nil
	}

	return
}

func (pmcs *pulsarMessageConsumerService) Consume(ctx context.Context) (
	messaging.Message,
	messaging.MessageProcessedCallback,
	messaging.MessageFailedCallback,
	error) {
	for {
		if ctx.Err() == context.Canceled {
			return messaging.Message{},
				nil,
				nil,
				context.Canceled
		}

		ctxWithTimeout, cancel := context.WithTimeout(ctx, 1*time.Second)

		msg, err := pmcs.consumer.Receive(ctxWithTimeout)
		if err != nil {
			cancel()

			if !errors.Is(err, context.DeadlineExceeded) {
				pmcs.logger.Errorf("Failed to receive message from. Error: %v", err)
			}

			continue
		}

		cancel()

		messageProcessedCallback := func() {
			if pmcs.consumer != nil {
				pmcs.consumer.Ack(msg)
			}
		}

		messageFailedCallback := func() {
			if pmcs.consumer != nil {
				pmcs.consumer.Nack(msg)
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
}
