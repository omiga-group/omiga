package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/generated"
	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/models"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/exchange"
)

// Links is the resolver for the links field.
func (r *exchangeResolver) Links(ctx context.Context, obj *repositories.Exchange) (*models.Links, error) {
	links := models.Links{}

	if link, ok := obj.Links["website"]; ok {
		links.Website = &link
	}

	if link, ok := obj.Links["facebook"]; ok {
		links.Facebook = &link
	}

	if link, ok := obj.Links["reddit"]; ok {
		links.Reddit = &link
	}

	if link, ok := obj.Links["twitter"]; ok {
		links.Twitter = &link
	}

	if link, ok := obj.Links["slack"]; ok {
		links.Slack = &link
	}

	if link, ok := obj.Links["telegram"]; ok {
		links.Telegram = &link
	}

	return &links, nil
}

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
func (r *queryResolver) Exchanges(ctx context.Context, after *repositories.Cursor, first *int, before *repositories.Cursor, last *int, orderBy []*repositories.ExchangeOrder, where *repositories.ExchangeWhereInput) (*repositories.ExchangeConnection, error) {
	orderBy = slices.Reverse(orderBy)

	pageOrder := slices.Map(orderBy, func(item *repositories.ExchangeOrder) repositories.ExchangePaginateOption {
		return repositories.WithExchangeOrder(item)
	})

	pageOrderAndFilter := append(pageOrder, repositories.WithExchangeFilter(where.Filter))

	return r.client.Exchange.
		Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			pageOrderAndFilter...)
}

// Exchange returns generated.ExchangeResolver implementation.
func (r *Resolver) Exchange() generated.ExchangeResolver { return &exchangeResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type exchangeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
