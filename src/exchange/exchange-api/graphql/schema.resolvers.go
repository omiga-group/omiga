package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/generated"
	"github.com/omiga-group/omiga/src/exchange/exchange-api/graphql/models"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	marketrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/market"
	tickerrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/ticker"
	tradingpairrepo "github.com/omiga-group/omiga/src/exchange/shared/entities/tradingpair"
	venuerepo "github.com/omiga-group/omiga/src/exchange/shared/entities/venue"
)

// Type is the resolver for the type field.
func (r *currencyResolver) Type(ctx context.Context, obj *entities.Currency) (models.CurrencyType, error) {
	return models.CurrencyType(obj.Type), nil
}

// Type is the resolver for the type field.
func (r *marketResolver) Type(ctx context.Context, obj *entities.Market) (models.MarketType, error) {
	return models.MarketType(obj.Type), nil
}

// Currency is the resolver for the currency field.
func (r *queryResolver) Currency(ctx context.Context, where *entities.CurrencyWhereInput) (*entities.Currency, error) {
	query, err := where.Filter(r.client.Currency.Query())
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

// Currencies is the resolver for the currencies field.
func (r *queryResolver) Currencies(ctx context.Context, after *entities.Cursor, first *int, before *entities.Cursor, last *int, orderBy []*entities.CurrencyOrder, where *entities.CurrencyWhereInput) (*entities.CurrencyConnection, error) {
	orderBy = slices.Reverse(orderBy)

	pageOrder := slices.Map(orderBy, func(item *entities.CurrencyOrder) entities.CurrencyPaginateOption {
		return entities.WithCurrencyOrder(item)
	})

	pageOrderAndFilter := append(pageOrder, entities.WithCurrencyFilter(where.Filter))

	return r.client.Currency.
		Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			pageOrderAndFilter...)
}

// Venue is the resolver for the exchange field.
func (r *queryResolver) Venue(ctx context.Context, where *entities.VenueWhereInput) (*entities.Venue, error) {
	query, err := where.Filter(r.client.Venue.Query())
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

// Venues is the resolver for the exchanges field.
func (r *queryResolver) Venues(ctx context.Context, after *entities.Cursor, first *int, before *entities.Cursor, last *int, orderBy []*entities.VenueOrder, where *entities.VenueWhereInput) (*entities.VenueConnection, error) {
	orderBy = slices.Reverse(orderBy)

	pageOrder := slices.Map(orderBy, func(item *entities.VenueOrder) entities.VenuePaginateOption {
		return entities.WithVenueOrder(item)
	})

	pageOrderAndFilter := append(pageOrder, entities.WithVenueFilter(where.Filter))

	return r.client.Venue.
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
func (r *tickerResolver) Market(ctx context.Context, obj *entities.Ticker) (*models.TickerMarket, error) {
	return &models.TickerMarket{
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

// Markets is the resolver for the markets field.
func (r *tradingPairResolver) Markets(ctx context.Context, obj *entities.TradingPair) ([]*entities.Market, error) {
	return obj.Market(ctx)
}

// Type is the resolver for the type field.
func (r *venueResolver) Type(ctx context.Context, obj *entities.Venue) (models.VenueType, error) {
	return models.VenueType(obj.Type), nil
}

// Links is the resolver for the links field.
func (r *venueResolver) Links(ctx context.Context, obj *entities.Venue) (*models.Links, error) {
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
func (r *venueResolver) Tickers(ctx context.Context, obj *entities.Venue) ([]*entities.Ticker, error) {
	return r.client.Ticker.
		Query().
		Where(tickerrepo.HasVenueWith(venuerepo.IDEQ(obj.ID))).
		All(ctx)
}

// TradingPairs is the resolver for the tradingPairs field.
func (r *venueResolver) TradingPairs(ctx context.Context, obj *entities.Venue) ([]*entities.TradingPair, error) {
	return r.client.TradingPair.
		Query().
		Where(tradingpairrepo.HasVenueWith(venuerepo.IDEQ(obj.ID))).
		All(ctx)
}

// Markets is the resolver for the markets field.
func (r *venueResolver) Markets(ctx context.Context, obj *entities.Venue) ([]*entities.Market, error) {
	return r.client.Market.
		Query().
		Where(marketrepo.HasVenueWith(venuerepo.IDEQ(obj.ID))).
		All(ctx)
}

// Currency returns generated.CurrencyResolver implementation.
func (r *Resolver) Currency() generated.CurrencyResolver { return &currencyResolver{r} }

// Market returns generated.MarketResolver implementation.
func (r *Resolver) Market() generated.MarketResolver { return &marketResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Ticker returns generated.TickerResolver implementation.
func (r *Resolver) Ticker() generated.TickerResolver { return &tickerResolver{r} }

// TradingPair returns generated.TradingPairResolver implementation.
func (r *Resolver) TradingPair() generated.TradingPairResolver { return &tradingPairResolver{r} }

// Venue returns generated.VenueResolver implementation.
func (r *Resolver) Venue() generated.VenueResolver { return &venueResolver{r} }

type currencyResolver struct{ *Resolver }
type marketResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tickerResolver struct{ *Resolver }
type tradingPairResolver struct{ *Resolver }
type venueResolver struct{ *Resolver }
