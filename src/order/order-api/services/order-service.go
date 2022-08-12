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
	return &orderService{
		client: entgoClient.GetClient(),
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
		Id:      createdOrder.ID,
		OrderID: createdOrder.OrderID,
	}, nil
}
