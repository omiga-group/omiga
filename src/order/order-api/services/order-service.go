package services

import (
	"context"

	"github.com/omiga-group/omiga/src/order/order-api/models"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
)

type OrderService interface {
	Submit(context.Context, models.Order) (*models.Order, error)
}

type orderService struct {
	client *repositories.Client
}

func NewOrderService(entgoClient repositories.EntgoClient) (OrderService, error) {
	client, err := entgoClient.GetClient()
	if err != nil {
		return nil, err
	}

	return &orderService{
		client: client,
	}, nil
}

func (os *orderService) Submit(
	ctx context.Context,
	order models.Order) (*models.Order, error) {

	createdOrder, err := os.client.Order.
		Create().
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &models.Order{
		ID: createdOrder.ID,
	}, nil
}
