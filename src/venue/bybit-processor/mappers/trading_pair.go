package mappers

import (
	"strings"

	"github.com/life4/genesis/slices"
	bybitpotv3 "github.com/omiga-group/omiga/src/venue/bybit-processor/bybitclient/spot/v3"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func BybitSymbolToTradingPairs(symbols []bybitpotv3.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol bybitpotv3.Symbol) models.TradingPair {
		basePricePrecision := 0
		if strings.Contains(symbol.BasePrecision, ".") {
			basePricePrecision = len(strings.Split(symbol.BasePrecision, ".")[1])
		}

		counterPricePrecision := 0
		if strings.Contains(symbol.QuotePrecision, ".") {
			counterPricePrecision = len(strings.Split(symbol.QuotePrecision, ".")[1])
		}

		return models.TradingPair{
			Symbol:                   symbol.Name,
			Base:                     symbol.BaseCoin,
			BasePriceMinPrecision:    &basePricePrecision,
			BasePriceMaxPrecision:    &basePricePrecision,
			Counter:                  symbol.QuoteCoin,
			CounterPriceMinPrecision: &counterPricePrecision,
			CounterPriceMaxPrecision: &counterPricePrecision,
		}
	})
}
