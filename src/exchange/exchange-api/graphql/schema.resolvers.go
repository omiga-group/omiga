package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/generated"
	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/models"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/ticker"
)

// Links is the resolver for the links field.
func (r *exchangeResolver) Links(ctx context.Context, obj *entities.Exchange) (*models.Links, error) {
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

// Tickers is the resolver for the tickers field.
func (r *exchangeResolver) Tickers(ctx context.Context, obj *entities.Exchange) ([]*entities.Ticker, error) {
	return r.client.Ticker.
		Query().
		Where(ticker.HasExchangeWith(exchange.IDEQ(obj.ID))).
		All(ctx)
}

// Coin is the resolver for the coin field.
func (r *queryResolver) Coin(ctx context.Context, where *entities.CoinWhereInput) (*entities.Coin, error) {
	query, err := where.Filter(r.client.Coin.Query())
	if err != nil {
		return nil, err
	}

	result, err := query.First(ctx)
	if _, ok := err.(*entities.NotFoundError); ok {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

// Coins is the resolver for the coins field.
func (r *queryResolver) Coins(ctx context.Context, after *entities.Cursor, first *int, before *entities.Cursor, last *int, orderBy []*entities.CoinOrder, where *entities.CoinWhereInput) (*entities.CoinConnection, error) {
	orderBy = slices.Reverse(orderBy)

	pageOrder := slices.Map(orderBy, func(item *entities.CoinOrder) entities.CoinPaginateOption {
		return entities.WithCoinOrder(item)
	})

	pageOrderAndFilter := append(pageOrder, entities.WithCoinFilter(where.Filter))

	return r.client.Coin.
		Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			pageOrderAndFilter...)
}

// Exchange is the resolver for the exchange field.
func (r *queryResolver) Exchange(ctx context.Context, where *entities.ExchangeWhereInput) (*entities.Exchange, error) {
	query, err := where.Filter(r.client.Exchange.Query())
	if err != nil {
		return nil, err
	}

	result, err := query.First(ctx)
	if _, ok := err.(*entities.NotFoundError); ok {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

// Exchanges is the resolver for the exchanges field.
func (r *queryResolver) Exchanges(ctx context.Context, after *entities.Cursor, first *int, before *entities.Cursor, last *int, orderBy []*entities.ExchangeOrder, where *entities.ExchangeWhereInput) (*entities.ExchangeConnection, error) {
	orderBy = slices.Reverse(orderBy)

	pageOrder := slices.Map(orderBy, func(item *entities.ExchangeOrder) entities.ExchangePaginateOption {
		return entities.WithExchangeOrder(item)
	})

	pageOrderAndFilter := append(pageOrder, entities.WithExchangeFilter(where.Filter))

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

// Market is the resolver for the market field.
func (r *tickerResolver) Market(ctx context.Context, obj *entities.Ticker) (*models.Market, error) {
	return &models.Market{
		HasTradingIncentive: obj.Market.HasTradingIncentive,
		Identifier:          obj.Market.Identifier,
		Name:                &obj.Market.Name,
	}, nil
}

// ConvertedLast is the resolver for the convertedLast field.
func (r *tickerResolver) ConvertedLast(ctx context.Context, obj *entities.Ticker) (*models.ConvertedDetails, error) {
	return &models.ConvertedDetails{
		Btc: obj.ConvertedLast.Btc,
		Eth: obj.ConvertedLast.Eth,
		Usd: obj.ConvertedLast.Usd,
	}, nil
}

// ConvertedVolume is the resolver for the convertedVolume field.
func (r *tickerResolver) ConvertedVolume(ctx context.Context, obj *entities.Ticker) (*models.ConvertedDetails, error) {
	return &models.ConvertedDetails{
		Btc: obj.ConvertedVolume.Btc,
		Eth: obj.ConvertedVolume.Eth,
		Usd: obj.ConvertedVolume.Usd,
	}, nil
}

// Exchange returns generated.ExchangeResolver implementation.
func (r *Resolver) Exchange() generated.ExchangeResolver { return &exchangeResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Ticker returns generated.TickerResolver implementation.
func (r *Resolver) Ticker() generated.TickerResolver { return &tickerResolver{r} }

type exchangeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tickerResolver struct{ *Resolver }
