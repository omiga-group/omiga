package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/generated"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/exchange"
)

// Exchange is the resolver for the exchange field.
func (r *queryResolver) Exchange(ctx context.Context, id int) (*repositories.Exchange, error) {
	query := r.client.Exchange.Query()
	query = query.Where(exchange.IDEQ(id))

	result, err := query.First(ctx)
	if _, ok := err.(*repositories.NotFoundError); ok {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

// Exchanges is the resolver for the exchanges field.
func (r *queryResolver) Exchanges(ctx context.Context, after *repositories.Cursor, first *int, before *repositories.Cursor, last *int, where *repositories.ExchangeWhereInput) (*repositories.ExchangeConnection, error) {
	return r.client.Exchange.
		Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			repositories.WithExchangeFilter(where.Filter))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
