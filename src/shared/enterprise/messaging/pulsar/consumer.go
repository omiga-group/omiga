package pulsar

import (
	"context"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"go.uber.org/zap"
)

type pulsarMessageConsumer struct {
	logger            *zap.SugaredLogger
	pulsarConfig      PulsarConfig
	pulsarClient      pulsar.Client
	consumer          pulsar.Consumer
	osHelper          os.OsHelper
	operationTimeout  time.Duration
	connectionTimeout time.Duration
	authentication    interface{}
}

func NewPulsarMessageConsumer(
	logger *zap.SugaredLogger,
	pulsarConfig PulsarConfig,
	osHelper os.OsHelper) (messaging.MessageConsumer, error) {
	operationTimeout, err := time.ParseDuration(pulsarConfig.OperationTimeout)
	if err != nil {
		return nil, err
	}

	connectionTimeout, err := time.ParseDuration(pulsarConfig.ConnectionTimeout)
	if err != nil {
		return nil, err
	}

	var authentication interface{} = nil

	if pulsarConfig.EnableAuthenticationOAuth2 {
		privateKeyFilePath, err := osHelper.CreateTemporaryTextFile(
			pulsarConfig.AuthenticationOAuth2.PrivateKey)
		if err != nil {
			return nil, err
		}

		authentication = pulsar.NewAuthenticationOAuth2(map[string]string{
			"type":       pulsarConfig.AuthenticationOAuth2.Type,
			"issuerUrl":  pulsarConfig.AuthenticationOAuth2.IssuerUrl,
			"audience":   pulsarConfig.AuthenticationOAuth2.Audience,
			"privateKey": privateKeyFilePath,
			"clientId":   pulsarConfig.AuthenticationOAuth2.ClientId,
		})
	}

	return &pulsarMessageConsumer{
		logger:            logger,
		pulsarConfig:      pulsarConfig,
		osHelper:          osHelper,
		operationTimeout:  operationTimeout,
		connectionTimeout: connectionTimeout,
		authentication:    authentication,
	}, nil
}

func (pmc *pulsarMessageConsumer) Close() {
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

func (pmc *pulsarMessageConsumer) Consume(ctx context.Context, topic string) (
	messaging.Message,
	messaging.MessageProcessedCallback,
	messaging.MessageFailedCallback,
	error) {
	if err := pmc.connect(topic); err != nil {
		return messaging.Message{}, nil, nil, err
	}

	msg, err := pmc.consumer.Receive(ctx)
	if err != nil {
		return messaging.Message{}, nil, nil, err
	}

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

func (pmc *pulsarMessageConsumer) connect(topic string) (err error) {
	if pmc.consumer != nil {
		return nil
	}

	pmc.pulsarClient, err = pulsar.NewClient(
		pulsar.ClientOptions{
			URL:               pmc.pulsarConfig.Url,
			OperationTimeout:  pmc.operationTimeout,
			ConnectionTimeout: pmc.connectionTimeout,
			Authentication:    pmc.authentication,
		})
	if err != nil {
		pmc.Close()

		return err
	}

	pmc.consumer, err = pmc.pulsarClient.Subscribe(
		pulsar.ConsumerOptions{
			Topic:            topic,
			SubscriptionName: pmc.pulsarConfig.SubscriptionName,
			Type:             pulsar.KeyShared,
		})
	if err != nil {
		pmc.Close()

		return err
	}

	return nil
}
