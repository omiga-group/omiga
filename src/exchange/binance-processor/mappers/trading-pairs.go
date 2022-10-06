package mappers

import (
	"github.com/adshao/go-binance/v2"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
)

func FromBinanceSymbolsToTradingPairs(symbols []binance.Symbol) []models.TradingPairs {
	return slices.Map(symbols, func(symbol binance.Symbol) models.TradingPairs {
		return models.TradingPairs{
			Symbol:           symbol.Symbol,
			Base:             symbol.BaseAsset,
			BasePrecision:    symbol.BaseAssetPrecision,
			Counter:          symbol.QuoteAsset,
			CounterPrecision: symbol.QuoteAssetPrecision,
		}
	})
}
