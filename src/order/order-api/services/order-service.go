package services

import (
	"context"

	"github.com/omiga-group/omiga/src/order/order-api/publishers"
	orderrepositories "github.com/omiga-group/omiga/src/order/order-api/repositories"
	"github.com/omiga-group/omiga/src/order/shared/models"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"go.uber.org/zap"
)

type OrderService interface {
	Submit(ctx context.Context, order models.Order) (models.Order, error)
}

type orderService struct {
	logger          *zap.SugaredLogger
	entgoClient     repositories.EntgoClient
	orderRepository orderrepositories.OrderRepository
	orderPublisher  publishers.OrderPublisher
}

func NewOrderService(
	logger *zap.SugaredLogger,
	entgoClient repositories.EntgoClient,
	orderRepository orderrepositories.OrderRepository,
	orderPublisher publishers.OrderPublisher) (OrderService, error) {
	return &orderService{
		logger:          logger,
		entgoClient:     entgoClient,
		orderRepository: orderRepository,
		orderPublisher:  orderPublisher,
	}, nil
}

func (os *orderService) Submit(
	ctx context.Context,
	order models.Order) (models.Order, error) {
	tx, err := os.entgoClient.CreateTransaction(ctx)
	if err != nil {
		return models.Order{}, err
	}

	order, err = os.orderRepository.
		CreateOrder(
			ctx,
			tx,
			order)
	if err != nil {
		rollbackErr := os.entgoClient.RollbackTransaction(tx)
		if rollbackErr != nil {
			os.logger.Errorf("Failed to rollback transaction. Error: %v", rollbackErr)
		}

		return models.Order{}, err
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

		return models.Order{}, err
	}

	err = os.entgoClient.CommitTransaction(tx)
	if err != nil {
		rollbackErr := os.entgoClient.RollbackTransaction(tx)
		if rollbackErr != nil {
			os.logger.Errorf("Failed to rollback transaction. Error: %v", rollbackErr)
		}

		return models.Order{}, err
	}

	return order, nil
}
