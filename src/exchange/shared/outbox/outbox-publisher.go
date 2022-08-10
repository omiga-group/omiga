package outbox

import (
	"context"
	"encoding/json"
	"time"

	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/outbox"
	"go.uber.org/zap"
)

type OutboxPublisher[EventType interface{}] interface {
	Publish(
		ctx context.Context,
		transaction *repositories.Tx,
		topic string,
		key string,
		headers map[string]string,
		event EventType) error
}

type outboxPublisher struct {
	logger *zap.SugaredLogger
}

func NewOutboxPublisher(
	logger *zap.SugaredLogger) (OutboxPublisher[interface{}], error) {
	return &outboxPublisher{
		logger: logger,
	}, nil
}

func (op *outboxPublisher) Publish(
	ctx context.Context,
	transaction *repositories.Tx,
	topic string,
	key string,
	headers map[string]string,
	event interface{}) error {
	payload, err := json.Marshal(event)
	if err != nil {
		op.logger.Errorf(
			"Failed to serialize event to json. Error: %v",
			err)

		return err
	}

	if _, err := transaction.Outbox.
		Create().
		SetTimestamp(time.Now()).
		SetTopic(topic).
		SetKey(key).
		SetPayload(payload).
		SetHeaders(headers).
		SetRetryCount(0).
		SetStatus(outbox.StatusPending).
		SetNillableLastRetry(nil).
		Save(ctx); err != nil {
		op.logger.Errorf(
			"Failed to save outbox item. Error: %v",
			err)

		return err
	}

	return nil
}
