
// Code generated by go-omiga-template, DO NOT EDIT.

package syntheticorderv1

import (
	"context"
	"encoding/json"
	"time"

	"github.com/omiga-group/omiga/src/shared/enterprise/messaging"
	enterpriseTime "github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

type Consumer interface {
	StartAsync(ctx context.Context) error
	Close()
}

type consumer struct {
	logger          *zap.SugaredLogger
	subscriber      Subscriber
	messageConsumer messaging.MessageConsumer
	timeHelper      enterpriseTime.TimeHelper
}

func NewConsumer(
	logger *zap.SugaredLogger,
	subscriber Subscriber,
	messageConsumer messaging.MessageConsumer,
	timeHelper enterpriseTime.TimeHelper) Consumer {
	return &consumer{
		logger:          logger,
		subscriber:      subscriber,
		messageConsumer: messageConsumer,
		timeHelper:      timeHelper,
	}
}

func (c *consumer) StartAsync(ctx context.Context) error {
	go func() {
		for {
			if ctx.Err() == context.Canceled {
				return
			}

			if err := c.messageConsumer.Connect(TopicName); err == nil {
				break
			} else {
				c.logger.Warnf("Can't connect to broker on topic: %s. Error: %v", TopicName, err)
			}

			c.timeHelper.SleepOrWaitForContextGetCancelled(ctx, time.Millisecond*100)
		}

		for {
			if ctx.Err() == context.Canceled {
				return
			}

			message, messageProcessedCallback, messageFailedCallback, err := c.messageConsumer.Consume(ctx)
			if err != nil && err == context.Canceled {
				return
			} else if err != nil && err != context.Canceled {
				c.logger.Errorf("Failed to consume message. Error: %v", err)

				continue
			}

			event := SyntheticOrderEvent{}
			if err := json.Unmarshal(message.Payload, &event); err != nil {
				c.logger.Errorf("Failed to de-serialize SyntheticOrderEvent message. Error: %v", err)

				messageFailedCallback()

				continue
			}

			if err := c.subscriber.Handle(ctx, event); err != nil {
				c.logger.Errorf("Failed to handle SyntheticOrderEvent message. Error: %v", err)

				messageFailedCallback()

				continue
			}

			messageProcessedCallback()
		}
	}()

	return nil
}

func (c *consumer) Close()  {
	c.messageConsumer.Close()
}
