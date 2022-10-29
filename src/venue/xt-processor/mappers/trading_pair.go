package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/shared/models"
	xtv4 "github.com/omiga-group/omiga/src/venue/xt-processor/xtclient/v4"
)

func XtSymbolsToTradingPairs(symbols []xtv4.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol xtv4.Symbol) models.TradingPair {
		return models.TradingPair{
			Symbol:                      symbol.Symbol,
			Base:                        symbol.BaseCurrency,
			BasePriceMinPrecision:       &symbol.BaseCurrencyPrecision,
			BasePriceMaxPrecision:       &symbol.BaseCurrencyPrecision,
			BaseQuantityMinPrecision:    &symbol.QuantityPrecision,
			BaseQuantityMaxPrecision:    &symbol.QuantityPrecision,
			Counter:                     symbol.QuoteCurrency,
			CounterPriceMinPrecision:    &symbol.QuoteCurrencyPrecision,
			CounterPriceMaxPrecision:    &symbol.QuoteCurrencyPrecision,
			CounterQuantityMinPrecision: &symbol.QuantityPrecision,
			CounterQuantityMaxPrecision: &symbol.QuantityPrecision,
		}
	})
}
