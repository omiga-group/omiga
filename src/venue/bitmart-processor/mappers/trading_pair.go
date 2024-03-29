package mappers

import (
	"github.com/life4/genesis/slices"
	bitmartspotv1 "github.com/omiga-group/omiga/src/venue/bitmart-processor/bitmartclient/spot/v1"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func BitmartSymbolsToTradingPairs(symbols []bitmartspotv1.Symbol) []models.TradingPair {
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
