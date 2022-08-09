package pulsar

import (
	"context"
	"errors"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type pulsarMessageConsumer struct {
	logger         *zap.SugaredLogger
	pulsarSettings PulsarSettings
	pulsarClient   pulsar.Client
	consumer       pulsar.Consumer
}

func NewPulsarMessageConsumer(
	logger *zap.SugaredLogger,
	pulsarSettings PulsarSettings,
	topic string) (messaging.MessageConsumer, error) {
	pulsarClient, err := pulsar.NewClient(
		pulsar.ClientOptions{
			URL:               pulsarSettings.Url,
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
	if err != nil {
		return nil, err
	}

	consumer, err := pulsarClient.Subscribe(
		pulsar.ConsumerOptions{
			Topic:            topic,
			SubscriptionName: pulsarSettings.SubscriptionName,
		})
	if err != nil {
		return nil, err
	}

	return &pulsarMessageConsumer{
		logger:         logger,
		pulsarSettings: pulsarSettings,
		pulsarClient:   pulsarClient,
		consumer:       consumer,
	}, nil
}

func (pmc *pulsarMessageConsumer) Close(ctx context.Context) {
	if pmc.consumer != nil {
		if err := pmc.consumer.Unsubscribe(); err != nil {
			pmc.logger.Errorf("Failed to unsubscribe. Error: %v", err)
		}

		pmc.consumer.Close()
		pmc.consumer = nil
	}

	if pmc.pulsarClient == nil {
		pmc.pulsarClient.Close()
		pmc.pulsarClient = nil
	}
}

func (pmc *pulsarMessageConsumer) Consume(ctx context.Context) (
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

		msg, err := pmc.consumer.Receive(ctxWithTimeout)
		if err != nil {
			cancel()

			if !errors.Is(err, context.DeadlineExceeded) {
				pmc.logger.Errorf("Failed to receive message from. Error: %v", err)
			}

			continue
		}

		cancel()

		messageProcessedCallback := func() {
			if pmc.consumer != nil {
				pmc.consumer.Ack(msg)
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
}
