package pulsar

import (
	"context"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"go.uber.org/zap"
)

type pulsarMessageProducer struct {
	logger       *zap.SugaredLogger
	pulsarConfig PulsarConfig
	pulsarClient pulsar.Client
	producer     pulsar.Producer
	topic        string
}

func NewPulsarMessageProducer(
	logger *zap.SugaredLogger,
	pulsarConfig PulsarConfig,
	osHelper os.OsHelper,
	topic string) (messaging.MessageProducer, error) {
	operationTimeout, err := time.ParseDuration(pulsarConfig.OperationTimeout)
	if err != nil {
		return nil, err
	}

	connectionTimeout, err := time.ParseDuration(pulsarConfig.ConnectionTimeout)
	if err != nil {
		return nil, err
	}

	var authentication interface{} = nil

	if pulsarConfig.EnableAuthenticationOAuth2 && pulsarConfig.AuthenticationOAuth2 != nil {
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

	pulsarClient, err := pulsar.NewClient(
		pulsar.ClientOptions{
			URL:               pulsarConfig.Url,
			OperationTimeout:  operationTimeout,
			ConnectionTimeout: connectionTimeout,
			Authentication:    authentication,
		})
	if err != nil {
		return nil, err
	}

	producer, err := pulsarClient.CreateProducer(
		pulsar.ProducerOptions{
			Topic: topic,
			Name:  pulsarConfig.ProducerName,
		})
	if err != nil {
		return nil, err
	}

	return &pulsarMessageProducer{
		logger:       logger,
		pulsarConfig: pulsarConfig,
		pulsarClient: pulsarClient,
		producer:     producer,
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
