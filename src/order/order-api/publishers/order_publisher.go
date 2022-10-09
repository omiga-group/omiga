package publishers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/omiga-group/omiga/src/order/shared/entities"
	"github.com/omiga-group/omiga/src/order/shared/mappers"
	"github.com/omiga-group/omiga/src/order/shared/models"
	"github.com/omiga-group/omiga/src/order/shared/outbox"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"go.uber.org/zap"
)

type OrderPublisher interface {
	Publish(
		ctx context.Context,
		tx *entities.Tx,
		orderBeforeState *models.Order,
		orderAfterState models.Order) error
}

type orderPublisher struct {
	logger               *zap.SugaredLogger
	appConfig            configuration.AppConfig
	orderOutboxPublisher outbox.OutboxPublisher
}

func NewOrderPublisher(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	orderOutboxPublisher outbox.OutboxPublisher) (OrderPublisher, error) {
	return &orderPublisher{
		logger:               logger,
		orderOutboxPublisher: orderOutboxPublisher,
		appConfig:            appConfig,
	}, nil
}

func (op *orderPublisher) Publish(
	ctx context.Context,
	tx *entities.Tx,
	orderBeforeState *models.Order,
	orderAfterState models.Order) error {

	orderEvent := orderv1.OrderEvent{
		Metadata: orderv1.Metadata{
			Id:            uuid.New(),
			Source:        op.appConfig.Source,
			Type:          orderv1.TypeOrderSubmitted,
			Time:          time.Now(),
			CorrelationId: uuid.New(),
		},
		Data: orderv1.Data{},
	}

	if orderBeforeState != nil {
		mappedBeforeState := mappers.FromOrderToEventOrder(*orderBeforeState)
		orderEvent.Data.BeforeState = &mappedBeforeState
	}

	orderEvent.Data.AfterState = mappers.FromOrderToEventOrder(orderAfterState)

	return op.orderOutboxPublisher.Publish(
		ctx,
		tx,
		orderv1.TopicName,
		"",
		map[string]string{},
		orderEvent)
}
