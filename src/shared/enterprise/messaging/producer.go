package messaging

import (
	"context"
)

type MessageProducer interface {
	Connect(topic string) error
	Produce(
		ctx context.Context,
		key string,
		data []byte) error
	Close()
}
