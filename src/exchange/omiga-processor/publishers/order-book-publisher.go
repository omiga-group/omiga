package publishers

import (
	"context"
	"time"

	"github.com/google/uuid"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"go.uber.org/zap"
)

type OrderBookPublisher interface {
	Publish(ctx context.Context) error
}

type OrderBookPublisherSettings struct {
	Name string
}

type orderBookPublisher struct {
	logger                     *zap.SugaredLogger
	appSettings                configuration.AppSettings
	orderBookProducer          orderbookv1.Producer
	orderBookPublisherSettings OrderBookPublisherSettings
}

func NewOrderBookPublisher(
	logger *zap.SugaredLogger,
	appSettings configuration.AppSettings,
	orderBookPublisherSettings OrderBookPublisherSettings,
	orderBookProducer orderbookv1.Producer) (OrderBookPublisher, error) {
	return &orderBookPublisher{
		logger:                     logger,
		orderBookPublisherSettings: orderBookPublisherSettings,
		appSettings:                appSettings,
		orderBookProducer:          orderBookProducer,
	}, nil
}

func (obp *orderBookPublisher) Publish(ctx context.Context) error {
	orderBookEvent := orderbookv1.OrderBookEvent{
		Metadata: orderbookv1.Metadata{
			Id:     orderbookv1.ID(uuid.New()),
			Time:   time.Now(),
			Source: obp.appSettings.Source + "::" + obp.orderBookPublisherSettings.Name,
			Type:   orderbookv1.TypeOrderBookUpdated,
		},
	}

	return obp.orderBookProducer.Produce(
		ctx,
		"",
		orderBookEvent)
}
