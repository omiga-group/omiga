package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	bitmartspotv1 "github.com/omiga-group/omiga/src/shared/clients/openapi/bitmart/spot/v1"
)

func BitMartSymbolsToTradingPairs(symbols []bitmartspotv1.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol bitmartspotv1.Symbol) models.TradingPair {
		return models.TradingPair{
			Symbol:                   symbol.Symbol,
			Base:                     symbol.BaseCurrency,
			BasePriceMinPrecision:    &symbol.PriceMinPrecision,
			BasePriceMaxPrecision:    &symbol.PriceMaxPrecision,
			Counter:                  symbol.QuoteCurrency,
			CounterPriceMinPrecision: &symbol.PriceMinPrecision,
			CounterPriceMaxPrecision: &symbol.PriceMaxPrecision,
		}
	})
}
