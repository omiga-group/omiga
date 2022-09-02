package services

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/coin"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/predicate"
)

type CoinHelper interface {
	GetCoinsNames(ctx context.Context, symbols []string) (map[string]string, error)
}

type coinHelper struct {
	entgoClient repositories.EntgoClient
}

func NewCoinHelper(
	entgoClient repositories.EntgoClient) (CoinHelper, error) {
	return &coinHelper{
		entgoClient: entgoClient,
	}, nil
}

func (se *coinHelper) GetCoinsNames(
	ctx context.Context,
	symbols []string) (map[string]string, error) {
	client := se.entgoClient.GetClient()
	predicates := slices.Map(symbols, func(symbol string) predicate.Coin {
		return coin.SymbolEQ(symbol)
	})

	results, err := client.Coin.Query().Where(coin.Or(predicates...)).All(ctx)
	if err != nil {
		return nil, err
	}

	return slices.Reduce(
		results,
		make(map[string]string),
		func(coin *repositories.Coin, acc map[string]string) map[string]string {
			return acc
		}), nil
}
