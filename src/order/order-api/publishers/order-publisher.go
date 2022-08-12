package publishers

import (
	"context"

	"github.com/omiga-group/omiga/src/order/order-api/models"
	"github.com/omiga-group/omiga/src/order/shared/outbox"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"go.uber.org/zap"
)

type OrderPublisher interface {
	Publish(ctx context.Context, tx *repositories.Tx, order models.Order) error
}

type orderPublisher struct {
	logger               *zap.SugaredLogger
	orderOutboxPublisher outbox.OutboxPublisher[orderv1.OrderEvent]
}

func NewOrderPublisher(
	logger *zap.SugaredLogger,
	orderOutboxPublisher outbox.OutboxPublisher[orderv1.OrderEvent]) (OrderPublisher, error) {
	return &orderPublisher{
		logger:               logger,
		orderOutboxPublisher: orderOutboxPublisher,
	}, nil
}

func (op *orderPublisher) Publish(ctx context.Context, tx *repositories.Tx, order models.Order) error {
	orderEvent := orderv1.OrderEvent{
		Data: orderv1.Data{
			AfterState: orderv1.Order{
				Id: order.OrderId,
			},
		},
	}

	return op.orderOutboxPublisher.Publish(
		ctx,
		tx,
		orderv1.TopicName,
		"",
		map[string]string{},
		orderEvent)
}
