package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/order/shared"
	"github.com/omiga-group/omiga/src/order/shared/mappers"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	"github.com/omiga-group/omiga/src/order/shared/repositories/order"
)

// SubmitOrder is the resolver for the submitOrder field.
func (r *mutationResolver) SubmitOrder(ctx context.Context, input shared.SubmitOrderInput) (*shared.OrderPayload, error) {
	order := mappers.FromSubmitOrderInputToOrder(input)

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
