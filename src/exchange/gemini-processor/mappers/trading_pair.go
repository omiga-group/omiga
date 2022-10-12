package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	gemini "github.com/omiga-group/omiga/src/shared/clients/openapi/gemini/v1"
)

func GeminiTradingPairsToTradingPairs(tradingPairs []gemini.TradingPair) []models.TradingPair {
	return slices.Map(tradingPairs, func(tradingPair gemini.TradingPair) models.TradingPair {
		return models.TradingPair{
			Symbol:  tradingPair.Symbol,
			Base:    tradingPair.BaseCurrency,
			Counter: tradingPair.QuoteCurrency,
		}
	})
}
