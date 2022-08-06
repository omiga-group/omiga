package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/exchange/shared"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
)

// CreateExchange is the resolver for the createExchange field.
func (r *mutationResolver) CreateExchange(ctx context.Context, input shared.CreateExchangeInput) (*shared.ExchangePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateExchange is the resolver for the updateExchange field.
func (r *mutationResolver) UpdateExchange(ctx context.Context, input shared.UpdateExchangeInput) (*shared.ExchangePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Exchange is the resolver for the exchange field.
func (r *queryResolver) Exchange(ctx context.Context, id *int) (*repositories.Exchange, error) {
	panic(fmt.Errorf("not implemented"))
}

// Exchanges is the resolver for the exchanges field.
func (r *queryResolver) Exchanges(ctx context.Context, after *repositories.Cursor, first *int, before *repositories.Cursor, last *int, where *repositories.ExchangeWhereInput) (*repositories.ExchangeConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns shared.MutationResolver implementation.
func (r *Resolver) Mutation() shared.MutationResolver { return &mutationResolver{r} }

// Query returns shared.QueryResolver implementation.
func (r *Resolver) Query() shared.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
