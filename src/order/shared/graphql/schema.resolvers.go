package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/order/shared"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input shared.CreateOrderInput) (*shared.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, input shared.UpdateOrderInput) (*shared.OrderPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Order is the resolver for the order field.
func (r *queryResolver) Order(ctx context.Context, id *int) (*repositories.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context, after *repositories.Cursor, first *int, before *repositories.Cursor, last *int, where *repositories.OrderWhereInput) (*repositories.OrderConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns shared.MutationResolver implementation.
func (r *Resolver) Mutation() shared.MutationResolver { return &mutationResolver{r} }

// Query returns shared.QueryResolver implementation.
func (r *Resolver) Query() shared.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
