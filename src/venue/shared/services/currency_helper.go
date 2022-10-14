package services

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	currencyrepo "github.com/omiga-group/omiga/src/venue/shared/entities/currency"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
)

type CurrencyHelper interface {
	GetCoinsNames(ctx context.Context, symbols []string) (map[string]string, error)
}

type currencyHelper struct {
	entgoClient entities.EntgoClient
}

func NewCurrencyHelper(
	entgoClient entities.EntgoClient) (CurrencyHelper, error) {
	return &currencyHelper{
		entgoClient: entgoClient,
	}, nil
}

func (ch *currencyHelper) GetCoinsNames(
	ctx context.Context,
	symbols []string) (map[string]string, error) {
	client := ch.entgoClient.GetClient()
	predicates := slices.Map(symbols, func(symbol string) predicate.Currency {
		return currencyrepo.SymbolEQ(symbol)
	})

	results, err := client.Currency.Query().Where(currencyrepo.Or(predicates...)).All(ctx)
	if err != nil {
		return nil, err
	}

	return slices.Reduce(
		results,
		make(map[string]string),
		func(coin *entities.Currency, acc map[string]string) map[string]string {
			return acc
		}), nil
}
