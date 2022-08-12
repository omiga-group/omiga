package outbox

import (
	"context"
	"encoding/json"
	"time"

	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"github.com/omiga-group/omiga/src/order/shared/repositories/outbox"
	"go.uber.org/zap"
)

type OutboxPublisher interface {
	PublishWithoutTransaction(
		ctx context.Context,
		topic string,
		key string,
		headers map[string]string,
		event interface{}) error
	Publish(
		ctx context.Context,
		transaction *repositories.Tx,
		topic string,
		key string,
		headers map[string]string,
		event interface{}) error
}

type outboxPublisher struct {
	logger      *zap.SugaredLogger
	entgoClient repositories.EntgoClient
}

func NewOutboxPublisher(
	logger *zap.SugaredLogger,
	entgoClient repositories.EntgoClient) (OutboxPublisher, error) {
	return &outboxPublisher{
		logger:      logger,
		entgoClient: entgoClient,
	}, nil
}

func (op *outboxPublisher) PublishWithoutTransaction(
	ctx context.Context,
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

	if _, err := op.entgoClient.GetClient().Outbox.
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
