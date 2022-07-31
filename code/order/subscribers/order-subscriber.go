package subscribers

import (
	"context"

	orderv1 "github.com/omiga-group/omiga/code/shared/events/events/omiga/order/v1"
)

type OrderSubscriber struct {
}

func NewOrderSubscriber() (orderv1.Subscriber, error) {
	return OrderSubscriber{}, nil
}

func (o OrderSubscriber) Handle(ctx context.Context, event orderv1.DomainEvent) error {
	return nil
}
