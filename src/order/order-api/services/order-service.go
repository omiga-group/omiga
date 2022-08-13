package services

import (
	"context"

	"github.com/omiga-group/omiga/src/order/order-api/publishers"
	"github.com/omiga-group/omiga/src/order/shared/models"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"go.uber.org/zap"
)

type OrderService interface {
	Submit(ctx context.Context, request models.Order) (*models.Order, error)
}

type orderService struct {
	logger         *zap.SugaredLogger
	entgoClient    repositories.EntgoClient
	orderPublisher publishers.OrderPublisher
}

func NewOrderService(
	logger *zap.SugaredLogger,
	entgoClient repositories.EntgoClient,
	orderPublisher publishers.OrderPublisher) (OrderService, error) {
	return &orderService{
		logger:         logger,
		entgoClient:    entgoClient,
		orderPublisher: orderPublisher,
	}, nil
}

func (os *orderService) Submit(
	ctx context.Context,
	request models.Order) (*models.Order, error) {
	tx, err := os.entgoClient.CreateTransaction(ctx)
	if err != nil {
		return nil, err
	}

	createdOrder, err := tx.Order.
		Create().
		Save(ctx)
	if err != nil {
		rollbackErr := os.entgoClient.RollbackTransaction(tx)
		if rollbackErr != nil {
			os.logger.Errorf("Failed to rollback transaction. Error: %v", rollbackErr)
		}

		return nil, err
	}

	order := models.Order{
		Id: createdOrder.ID,
	}

	err = os.orderPublisher.Publish(
		ctx,
		tx,
		nil,
		order)
	if err != nil {
		rollbackErr := os.entgoClient.RollbackTransaction(tx)
		if rollbackErr != nil {
			os.logger.Errorf("Failed to rollback transaction. Error: %v", rollbackErr)
		}

		return nil, err
	}

	err = os.entgoClient.CommitTransaction(tx)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
