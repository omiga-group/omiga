// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Coin) CoinBase(ctx context.Context) ([]*TradingPair, error) {
	result, err := c.NamedCoinBase(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = c.QueryCoinBase().All(ctx)
	}
	return result, err
}

func (c *Coin) CoinCounter(ctx context.Context) ([]*TradingPair, error) {
	result, err := c.NamedCoinCounter(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = c.QueryCoinCounter().All(ctx)
	}
	return result, err
}

func (e *Exchange) Ticker(ctx context.Context) ([]*Ticker, error) {
	result, err := e.NamedTicker(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = e.QueryTicker().All(ctx)
	}
	return result, err
}

func (e *Exchange) TradingPair(ctx context.Context) ([]*TradingPair, error) {
	result, err := e.NamedTradingPair(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = e.QueryTradingPair().All(ctx)
	}
	return result, err
}

func (e *Exchange) Market(ctx context.Context) ([]*Market, error) {
	result, err := e.NamedMarket(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = e.QueryMarket().All(ctx)
	}
	return result, err
}

func (m *Market) Exchange(ctx context.Context) (*Exchange, error) {
	result, err := m.Edges.ExchangeOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryExchange().Only(ctx)
	}
	return result, err
}

func (m *Market) TradingPair(ctx context.Context) ([]*TradingPair, error) {
	result, err := m.NamedTradingPair(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = m.QueryTradingPair().All(ctx)
	}
	return result, err
}

func (t *Ticker) Exchange(ctx context.Context) (*Exchange, error) {
	result, err := t.Edges.ExchangeOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryExchange().Only(ctx)
	}
	return result, err
}

func (tp *TradingPair) Exchange(ctx context.Context) (*Exchange, error) {
	result, err := tp.Edges.ExchangeOrErr()
	if IsNotLoaded(err) {
		result, err = tp.QueryExchange().Only(ctx)
	}
	return result, err
}

func (tp *TradingPair) Base(ctx context.Context) (*Coin, error) {
	result, err := tp.Edges.BaseOrErr()
	if IsNotLoaded(err) {
		result, err = tp.QueryBase().Only(ctx)
	}
	return result, err
}

func (tp *TradingPair) Counter(ctx context.Context) (*Coin, error) {
	result, err := tp.Edges.CounterOrErr()
	if IsNotLoaded(err) {
		result, err = tp.QueryCounter().Only(ctx)
	}
	return result, err
}

func (tp *TradingPair) Market(ctx context.Context) ([]*Market, error) {
	result, err := tp.NamedMarket(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = tp.QueryMarket().All(ctx)
	}
	return result, err
}
