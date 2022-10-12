package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/toorop/go-bittrex"
)

func BittrexMarketsToTradingPairs(markets []bittrex.MarketV3) []models.TradingPair {
	return slices.Map(markets, func(market bittrex.MarketV3) models.TradingPair {

		precision := int(market.Precision)

		return models.TradingPair{
			Symbol:           market.Symbol,
			Base:             market.BaseCurrencySymbol,
			BasePrecision:    &precision,
			Counter:          market.QuoteCurrencySymbol,
			CounterPrecision: &precision,
		}
	})
}
