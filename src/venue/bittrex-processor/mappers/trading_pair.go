package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/shared/models"
	"github.com/toorop/go-bittrex"
)

func BittrexMarketsToTradingPairs(markets []bittrex.MarketV3) []models.TradingPair {
	return slices.Map(markets, func(market bittrex.MarketV3) models.TradingPair {

		precision := int(market.Precision)

		return models.TradingPair{
			Symbol:                   market.Symbol,
			Base:                     market.BaseCurrencySymbol,
			BasePriceMinPrecision:    &precision,
			BasePriceMaxPrecision:    &precision,
			Counter:                  market.QuoteCurrencySymbol,
			CounterPriceMinPrecision: &precision,
			CounterPriceMaxPrecision: &precision,
		}
	})
}
