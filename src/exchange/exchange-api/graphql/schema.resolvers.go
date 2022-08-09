package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/omiga-group/omiga/src/exchange/shared"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
)

// Exchange is the resolver for the exchange field.
func (r *queryResolver) Exchange(ctx context.Context, id int) (*repositories.Exchange, error) {
	panic(fmt.Errorf("not implemented"))
}

// Exchanges is the resolver for the exchanges field.
func (r *queryResolver) Exchanges(ctx context.Context, after *repositories.Cursor, first *int, before *repositories.Cursor, last *int, where *repositories.ExchangeWhereInput) (*repositories.ExchangeConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns shared.QueryResolver implementation.
func (r *Resolver) Query() shared.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
