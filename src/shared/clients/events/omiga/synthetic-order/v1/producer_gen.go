
// Code generated by go-omiga-template, DO NOT EDIT.

package syntheticorderv1

import (
	"context"
	"encoding/json"

	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type Producer interface {
	Produce(ctx context.Context, key string, event SyntheticOrderEvent) error
}

type producer struct {
	logger          *zap.SugaredLogger
	messageProducer messaging.MessageProducer
}

func NewProducer(
	logger *zap.SugaredLogger,
	messageProducer messaging.MessageProducer) Producer {
	return &producer{
		logger:          logger,
		messageProducer: messageProducer,
	}
}

func (c *producer) Produce(ctx context.Context, key string, event SyntheticOrderEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		c.logger.Errorf(
			"Failed to serialize SyntheticOrderEvent message to json. Error: %v",
			err)

		return err
	}

	if err := c.messageProducer.Produce(ctx, key, data); err != nil {
		return err
	}

	return nil
}