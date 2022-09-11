package pulsar

import (
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/lucsky/cuid"
	"github.com/omiga-group/omiga/src/shared/enterprise/os"
	"go.uber.org/zap"
)

type PulsarClient interface {
	CreateProducer(topic string) (pulsar.Producer, error)
	CreateConsumer(topic string) (pulsar.Consumer, error)
	Close()
}

type pulsarClient struct {
	logger            *zap.SugaredLogger
	pulsarConfig      PulsarConfig
	pulsarClient      pulsar.Client
	producer          pulsar.Producer
	consumer          pulsar.Consumer
	osHelper          os.OsHelper
	operationTimeout  time.Duration
	connectionTimeout time.Duration
	authentication    interface{}
}

func NewPulsarClient(
	logger *zap.SugaredLogger,
	pulsarConfig PulsarConfig,
	osHelper os.OsHelper) (PulsarClient, error) {
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

	return &pulsarClient{
		logger:            logger,
		pulsarConfig:      pulsarConfig,
		osHelper:          osHelper,
		operationTimeout:  operationTimeout,
		connectionTimeout: connectionTimeout,
		authentication:    authentication,
	}, nil
}

func (pc *pulsarClient) CreateProducer(topic string) (pulsar.Producer, error) {
	if pc.producer != nil {
		return pc.producer, nil
	}

	var err error

	if err = pc.connect(); err != nil {
		return nil, err
	}

	pc.producer, err = pc.pulsarClient.CreateProducer(
		pulsar.ProducerOptions{
			Topic: topic,
			Name:  pc.pulsarConfig.ProducerName + "-" + cuid.New(),
		})
	if err != nil {
		pc.Close()

		return nil, err
	}

	return pc.producer, nil
}

func (pc *pulsarClient) CreateConsumer(topic string) (pulsar.Consumer, error) {
	if pc.consumer != nil {
		return pc.consumer, nil
	}

	var err error

	if err = pc.connect(); err != nil {
		return nil, err
	}

	pc.consumer, err = pc.pulsarClient.Subscribe(
		pulsar.ConsumerOptions{
			Topic:            topic,
			SubscriptionName: pc.pulsarConfig.SubscriptionName,
			Type:             pulsar.KeyShared,
		})
	if err != nil {
		pc.Close()

		return nil, err
	}

	return pc.consumer, nil
}

func (pc *pulsarClient) Close() {
	if pc.producer != nil {
		if err := pc.producer.Flush(); err != nil {
			pc.logger.Errorf("Failed to flush. Error: %v", err)
		}

		pc.producer.Close()
		pc.producer = nil
	}

	if pc.consumer != nil {
		if err := pc.consumer.Unsubscribe(); err != nil {
			pc.logger.Errorf("Failed to unsubscribe. Error: %v", err)
		}

		pc.consumer.Close()
		pc.consumer = nil
	}

	if pc.pulsarClient == nil {
		pc.pulsarClient.Close()
		pc.pulsarClient = nil
	}
}

func (pc *pulsarClient) connect() (err error) {
	if pc.pulsarClient != nil {
		return nil
	}

	pc.pulsarClient, err = pulsar.NewClient(
		pulsar.ClientOptions{
			URL:               pc.pulsarConfig.Url,
			OperationTimeout:  pc.operationTimeout,
			ConnectionTimeout: pc.connectionTimeout,
			Authentication:    pc.authentication,
		})
	if err != nil {
		pc.Close()

		return err
	}

	return nil
}
