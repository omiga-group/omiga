package messaging

import (
	"context"
	"time"
)

type Message struct {
	Topic       string
	Key         string
	Payload     []byte
	Headers     map[string]string
	PublishTime time.Time
	EventTime   time.Time
}

type MessageProcessedCallback func()
type MessageFailedCallback func()

type MessageConsumer interface {
	Close(ctx context.Context)
	Consume(ctx context.Context) (
		Message,
		MessageProcessedCallback,
		MessageFailedCallback,
		error)
}
