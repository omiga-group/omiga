package mappers

import (
	"github.com/huobirdcenter/huobi_golang/pkg/model/common"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func HuobiSymbolsToTradingPairs(symbols []common.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol common.Symbol) models.TradingPair {
		return models.TradingPair{
			Symbol:                      symbol.Symbol,
			Base:                        symbol.BaseCurrency,
			BasePriceMinPrecision:       &symbol.PricePrecision,
			BasePriceMaxPrecision:       &symbol.PricePrecision,
			BaseQuantityMinPrecision:    &symbol.AmountPrecision,
			BaseQuantityMaxPrecision:    &symbol.AmountPrecision,
			Counter:                     symbol.QuoteCurrency,
			CounterPriceMinPrecision:    &symbol.PricePrecision,
			CounterPriceMaxPrecision:    &symbol.PricePrecision,
			CounterQuantityMinPrecision: &symbol.AmountPrecision,
			CounterQuantityMaxPrecision: &symbol.AmountPrecision,
		}
	})
}
