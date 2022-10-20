package mappers

import (
	"strings"

	"github.com/life4/genesis/slices"
	mexcpotv2 "github.com/omiga-group/omiga/src/venue/mexc-processor/mexcclient/spot/v2"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func MexcSymbolsToTradingPairs(symbols []mexcpotv2.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol mexcpotv2.Symbol) models.TradingPair {
		return models.TradingPair{
			Symbol:                      symbol.Symbol,
			Base:                        symbol.VcoinName,
			BasePriceMinPrecision:       &symbol.PriceScale,
			BasePriceMaxPrecision:       &symbol.PriceScale,
			BaseQuantityMinPrecision:    &symbol.QuantityScale,
			BaseQuantityMaxPrecision:    &symbol.QuantityScale,
			Counter:                     strings.Split(symbol.Symbol, "_")[1],
			CounterPriceMinPrecision:    &symbol.PriceScale,
			CounterPriceMaxPrecision:    &symbol.PriceScale,
			CounterQuantityMinPrecision: &symbol.QuantityScale,
			CounterQuantityMaxPrecision: &symbol.QuantityScale,
		}
	})
}
