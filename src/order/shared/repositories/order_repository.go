package repositories

import (
	"context"

	"github.com/omiga-group/omiga/src/order/shared/entities"
	"github.com/omiga-group/omiga/src/order/shared/models"
)

type OrderRepository interface {
	CreateOrder(
		ctx context.Context,
		tx *entities.Tx,
		order models.Order) (models.Order, error)
}

type orderRepository struct {
}

func NewOrderRepository() (OrderRepository, error) {
	return &orderRepository{}, nil
}

func (or *orderRepository) CreateOrder(
	ctx context.Context,
	tx *entities.Tx,
	order models.Order) (models.Order, error) {
	savedOrder, err := tx.Order.
		Create().
		SetOrderDetails(order.OrderDetails).
		SetPreferredExchanges(order.PreferredExchanges).
		Save(ctx)
	if err != nil {
		return models.Order{}, err
	}

	order.Id = savedOrder.ID

	return order, nil
}
