package mappers

import (
	"github.com/life4/genesis/slices"
	rainv1 "github.com/omiga-group/omiga/src/venue/rain-processor/rainclient/v1"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func RainCoinsToTradingPairs(coins []rainv1.Coin) []models.TradingPair {
	return slices.Reduce(
		coins,
		make([]models.TradingPair, 0),
		func(coin rainv1.Coin, acc []models.TradingPair) []models.TradingPair {
			otherCoins := slices.Filter(coins, func(otherCoin rainv1.Coin) bool {
				return otherCoin.Code != coin.Code
			})

			return slices.Concat(
				acc,
				slices.Map(otherCoins, func(otherCoin rainv1.Coin) models.TradingPair {
					return models.TradingPair{
						Symbol:                      coin.Code + otherCoin.Code,
						Base:                        coin.Code,
						BaseQuantityMinPrecision:    &coin.Precision,
						BaseQuantityMaxPrecision:    &coin.Precision,
						Counter:                     otherCoin.Code,
						CounterQuantityMinPrecision: &otherCoin.Precision,
						CounterQuantityMaxPrecision: &otherCoin.Precision,
					}
				}))

		})
}
