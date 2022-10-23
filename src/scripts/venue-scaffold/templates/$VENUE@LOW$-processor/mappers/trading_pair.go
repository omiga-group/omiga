package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func $VENUE@PAS$SymbolsToTradingPairs(symbols []$VENUE@LOW$.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol $VENUE@LOW$.Symbol) models.TradingPair {
		return models.TradingPair{
			Symbol:                   symbol.Symbol,
			Base:                     symbol.BaseAsset,
			BasePriceMinPrecision:    &symbol.BaseAssetPrecision,
			BasePriceMaxPrecision:    &symbol.BaseAssetPrecision,
			Counter:                  symbol.QuoteAsset,
			CounterPriceMinPrecision: &symbol.QuoteAssetPrecision,
			CounterPriceMaxPrecision: &symbol.QuoteAssetPrecision,
		}
	})
}
