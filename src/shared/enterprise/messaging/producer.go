package messaging

import (
	"context"
)

type MessageProducer interface {
	Produce(
		ctx context.Context,
		topic string,
		key string,
		data []byte) error
}
