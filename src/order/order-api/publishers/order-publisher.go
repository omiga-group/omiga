package publishers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/omiga-group/omiga/src/order/shared/models"
	"github.com/omiga-group/omiga/src/order/shared/outbox"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"go.uber.org/zap"
)

type OrderPublisher interface {
	Publish(
		ctx context.Context,
		tx *repositories.Tx,
		orderBeforeState *models.Order,
		orderAfterState models.Order) error
}

type orderPublisher struct {
	logger               *zap.SugaredLogger
	appSettings          configuration.AppSettings
	orderOutboxPublisher outbox.OutboxPublisher
}

func NewOrderPublisher(
	logger *zap.SugaredLogger,
	appSettings configuration.AppSettings,
	orderOutboxPublisher outbox.OutboxPublisher) (OrderPublisher, error) {
	return &orderPublisher{
		logger:               logger,
		orderOutboxPublisher: orderOutboxPublisher,
		appSettings:          appSettings,
	}, nil
}

func (op *orderPublisher) Publish(
	ctx context.Context,
	tx *repositories.Tx,
	orderBeforeState *models.Order,
	orderAfterState models.Order) error {

	orderEvent := orderv1.OrderEvent{
		Metadata: orderv1.Metadata{
			Id:     orderv1.ID(uuid.New()),
			Time:   time.Now(),
			Source: op.appSettings.Source,
			Type:   orderv1.TypeOrderSubmitted,
		},
		Data: orderv1.Data{},
	}

	if orderBeforeState != nil {
		orderEvent.Data.BeforeState = &orderv1.Order{
			Id: orderBeforeState.Id,
		}
	}

	orderEvent.Data.AfterState = orderv1.Order{
		Id: orderAfterState.Id,
	}

	return op.orderOutboxPublisher.Publish(
		ctx,
		tx,
		orderv1.TopicName,
		"",
		map[string]string{},
		orderEvent)
}
