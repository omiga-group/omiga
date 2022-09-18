// Code generated by ent, DO NOT EDIT.

package repositories

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (e *Exchange) Ticker(ctx context.Context) ([]*Ticker, error) {
	result, err := e.NamedTicker(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = e.QueryTicker().All(ctx)
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
