package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/order/order-processor/services"
	"github.com/omiga-group/omiga/src/order/shared/models"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"go.uber.org/zap"
)

type orderBookSubscriber struct {
	logger           *zap.SugaredLogger
	orderBookService services.OrderBookService
}

func NewOrderBookSubscriber(
	logger *zap.SugaredLogger,
	orderBookService services.OrderBookService) (orderbookv1.Subscriber, error) {
	return &orderBookSubscriber{
		logger:           logger,
		orderBookService: orderBookService,
	}, nil
}

func (obs *orderBookSubscriber) Handle(ctx context.Context, event orderbookv1.OrderBookEvent) error {
	return obs.orderBookService.Handle(
		ctx,
		event.Data.ExchangeId,
		event.Metadata.Time,
		models.OrderBook{})
}
