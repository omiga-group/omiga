package mappers

import (
	"github.com/adshao/go-binance/v2"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func BinanceSymbolsToTradingPairs(symbols []binance.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol binance.Symbol) models.TradingPair {
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
