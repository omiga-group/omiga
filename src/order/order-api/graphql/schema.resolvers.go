package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/order/shared"
	"github.com/omiga-group/omiga/src/order/shared/models"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"github.com/omiga-group/omiga/src/order/shared/repositories/order"
)

// SubmitOrder is the resolver for the submitOrder field.
func (r *mutationResolver) SubmitOrder(ctx context.Context, input shared.SubmitOrderInput) (*shared.OrderPayload, error) {
	order := models.Order{
		OrderDetails: models.OrderDetails{
			BaseCurrency: models.Currency{
				Name:         input.OrderDetails.BaseCurrency.Name,
				Code:         input.OrderDetails.BaseCurrency.Code,
				MaxPrecision: input.OrderDetails.BaseCurrency.MaxPrecision,
				Digital:      input.OrderDetails.BaseCurrency.Digital,
			},
			CounterCurrency: models.Currency{
				Name:         input.OrderDetails.CounterCurrency.Name,
				Code:         input.OrderDetails.CounterCurrency.Code,
				MaxPrecision: input.OrderDetails.CounterCurrency.MaxPrecision,
				Digital:      input.OrderDetails.CounterCurrency.Digital,
			},
			Type: models.OrderType(*input.OrderDetails.Type),
			Side: models.OrderSide(*input.OrderDetails.Side),
			Quantity: models.Money{
				Amount: input.OrderDetails.Quantity.Amount,
				Scale:  input.OrderDetails.Quantity.Scale,
				Currency: models.Currency{
					Name:         input.OrderDetails.Quantity.Currency.Name,
					Code:         input.OrderDetails.Quantity.Currency.Code,
					MaxPrecision: input.OrderDetails.Quantity.Currency.MaxPrecision,
					Digital:      input.OrderDetails.Quantity.Currency.Digital,
				},
			},
			Price: models.Money{
				Amount: input.OrderDetails.Price.Amount,
				Scale:  input.OrderDetails.Price.Scale,
				Currency: models.Currency{
					Name:         input.OrderDetails.Price.Currency.Name,
					Code:         input.OrderDetails.Price.Currency.Code,
					MaxPrecision: input.OrderDetails.Price.Currency.MaxPrecision,
					Digital:      input.OrderDetails.Price.Currency.Digital,
				},
			},
		},
	}

	submittedOrder, err := r.orderService.Submit(ctx, order)
	if err != nil {
		return nil, err
	}

	r.orderOutboxBackgroundService.RunAsync()

	return &shared.OrderPayload{
		ClientMutationID: input.ClientMutationID,
		Order: &repositories.Order{
			ID: submittedOrder.Id,
		},
	}, nil
}

// CancelOrder is the resolver for the cancelOrder field.
func (r *mutationResolver) CancelOrder(ctx context.Context, input shared.CancelOrderInput) (*shared.OrderPayload, error) {
	return nil, fmt.Errorf("not implemented")
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id int) (*repositories.Order, error) {
	query := r.client.Order.Query()
	query = query.Where(order.IDEQ(id))

	result, err := query.First(ctx)
	if _, ok := err.(*repositories.NotFoundError); ok {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, after *repositories.Cursor, first *int, before *repositories.Cursor, last *int, where *repositories.OrderWhereInput) (*repositories.OrderConnection, error) {
	return r.client.Order.
		Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			repositories.WithOrderFilter(where.Filter))
}

// Mutation returns shared.MutationResolver implementation.
func (r *Resolver) Mutation() shared.MutationResolver { return &mutationResolver{r} }

// Query returns shared.QueryResolver implementation.
func (r *Resolver) Query() shared.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
