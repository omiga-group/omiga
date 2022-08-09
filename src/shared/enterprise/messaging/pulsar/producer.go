package pulsar

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type pulsarMessageProducer struct {
	logger         *zap.SugaredLogger
	pulsarSettings PulsarSettings
	pulsarClient   pulsar.Client
	producer       pulsar.Producer
}

func NewPulsarMessageProducer(
	logger *zap.SugaredLogger,
	pulsarSettings PulsarSettings) (messaging.MessageProducer, error) {

	return &pulsarMessageProducer{
		logger:         logger,
		pulsarSettings: pulsarSettings,
	}, nil
}

func (pmcs *pulsarMessageProducer) Produce(
	ctx context.Context,
	topic string,
	key string,
	data []byte) error {

	return nil
}
