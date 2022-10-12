package mappers

import (
	"strings"

	"github.com/huobirdcenter/huobi_golang/pkg/model/common"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
)

func HuobiSymbolsToTradingPairs(symbols []common.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol common.Symbol) models.TradingPair {
		return models.TradingPair{
			Symbol:           strings.ToUpper(symbol.Symbol),
			Base:             strings.ToUpper(symbol.BaseCurrency),
			BasePrecision:    &symbol.PricePrecision,
			Counter:          strings.ToUpper(symbol.QuoteCurrency),
			CounterPrecision: &symbol.PricePrecision,
		}
	})
}
