package messaging

import (
	"context"
)

type MessageProducer interface {
	Close()
	Produce(
		ctx context.Context,
		topic string,
		key string,
		data []byte) error
}
