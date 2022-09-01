const { GoGenerator } = require("@asyncapi/modelina");
const { File } = require("@asyncapi/generator-react-sdk");

export default async function ({ asyncapi, params }) {
  const generator = new GoGenerator();
  const models = await generator.generate(asyncapi);

  const payloadContent = `
// Code generated by go-omiga-template, DO NOT EDIT.

package ${params.packageName}

import (
	"context"
	"encoding/json"

	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type Producer interface {
	Produce(ctx context.Context, key string, event ${models[0].modelName}) error
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

func (c *producer) Produce(ctx context.Context, key string, event ${models[0].modelName}) error {
	data, err := json.Marshal(event)
	if err != nil {
		c.logger.Errorf(
			"Failed to serialize ${models[0].modelName} message to json. Error: %v",
			err)

		return err
	}

	if err := c.messageProducer.Produce(ctx, key, data); err != nil {
		return err
	}

	return nil
}
`;

  return <File name="producer_eventgen.go">{payloadContent}</File>;
}
