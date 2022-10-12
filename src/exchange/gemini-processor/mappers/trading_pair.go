package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	gemini "github.com/omiga-group/omiga/src/shared/clients/openapi/gemini/v1"
)

func GeminiSymbolsToTradingPairs(symbols []gemini.TradingPair) []models.TradingPair {
	return slices.Map(symbols, func(tp gemini.TradingPair) models.TradingPair {
		return models.TradingPair{
			Symbol:  tp.Symbol,
			Base:    tp.BaseCurrency,
			Counter: tp.QuoteCurrency,
		}
	})
}
