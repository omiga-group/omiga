package messaging

import (
	"context"
)

type MessageProducer interface {
	Close(ctx context.Context)
	Produce(
		ctx context.Context,
		key string,
		data []byte) error
}
