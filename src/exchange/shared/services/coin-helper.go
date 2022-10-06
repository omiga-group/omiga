package services

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/coin"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/predicate"
)

type CoinHelper interface {
	GetCoinsNames(ctx context.Context, symbols []string) (map[string]string, error)
}

type coinHelper struct {
	entgoClient entities.EntgoClient
}

func NewCoinHelper(
	entgoClient entities.EntgoClient) (CoinHelper, error) {
	return &coinHelper{
		entgoClient: entgoClient,
	}, nil
}

func (ch *coinHelper) GetCoinsNames(
	ctx context.Context,
	symbols []string) (map[string]string, error) {
	client := ch.entgoClient.GetClient()
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
		func(coin *entities.Coin, acc map[string]string) map[string]string {
			return acc
		}), nil
}
