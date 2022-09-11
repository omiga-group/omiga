package publishers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/omiga-group/omiga/src/exchange/shared/mappers"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"go.uber.org/zap"
)

type OrderBookPublisher interface {
	Publish(ctx context.Context, key string, orderBook models.OrderBook) error
	Close()
}

type orderBookPublisher struct {
	logger            *zap.SugaredLogger
	appConfig         configuration.AppConfig
	orderBookProducer orderbookv1.Producer
}

func NewOrderBookPublisher(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	orderBookProducer orderbookv1.Producer) (OrderBookPublisher, error) {
	return &orderBookPublisher{
		logger:            logger,
		appConfig:         appConfig,
		orderBookProducer: orderBookProducer,
	}, nil
}

func (obp *orderBookPublisher) Publish(ctx context.Context, key string, orderBook models.OrderBook) error {
	orderBookEvent := orderbookv1.OrderBookEvent{
		Metadata: orderbookv1.Metadata{
			Id:            uuid.New(),
			Source:        obp.appConfig.Source,
			Type:          orderbookv1.TypeOrderBookUpdated,
			Time:          time.Now(),
			CorrelationId: uuid.New(),
		},
		Data: mappers.FromOrderBookToEventOrderBook(orderBook),
	}

	return obp.orderBookProducer.Produce(
		ctx,
		key,
		orderBookEvent)
}

func (obp *orderBookPublisher) Close() {
	obp.orderBookProducer.Close()
}
