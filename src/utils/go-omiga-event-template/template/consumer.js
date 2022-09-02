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

type Consumer interface {
	StartAsync(ctx context.Context) error
}

type consumer struct {
	logger                 *zap.SugaredLogger
	subscriber             Subscriber
	messageConsumer messaging.MessageConsumer
}

func NewConsumer(
	logger *zap.SugaredLogger,
	subscriber Subscriber,
	messageConsumer messaging.MessageConsumer) Consumer {
	return &consumer{
		logger:                 logger,
		subscriber:             subscriber,
		messageConsumer: messageConsumer,
	}
}

func (c *consumer) StartAsync(ctx context.Context) error {
	go func() {
		for {
			if ctx.Err() == context.Canceled {
				return
			}

			message, messageProcessedCallback, messageFailedCallback, err := c.messageConsumer.Consume(ctx)
			if err != nil && err != context.Canceled {
				c.logger.Errorf("Failed to consume message. Error: %v", err)
				return
			}

			event := ${models[0].modelName}{}
			if err := json.Unmarshal(message.Payload, &event); err != nil {
				c.logger.Errorf("Failed to de-serialize ${models[0].modelName} message. Error: %v", err)

				messageFailedCallback()

				continue
			}

			if err := c.subscriber.Handle(ctx, event); err != nil {
				c.logger.Errorf("Failed to handle ${models[0].modelName} message. Error: %v", err)

				messageFailedCallback()

				continue
			}

			messageProcessedCallback()
		}
	}()

	return nil
}
`;

  return <File name="consumer_eventgen.go">{payloadContent}</File>;
}