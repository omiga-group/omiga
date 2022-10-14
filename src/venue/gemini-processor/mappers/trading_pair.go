package mappers

import (
	"github.com/life4/genesis/slices"
	geminiv1 "github.com/omiga-group/omiga/src/shared/clients/openapi/gemini/v1"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func GeminiTradingPairsToTradingPairs(tradingPairs []geminiv1.TradingPair) []models.TradingPair {
	return slices.Map(tradingPairs, func(tradingPair geminiv1.TradingPair) models.TradingPair {
		return models.TradingPair{
			Symbol:  tradingPair.Symbol,
			Base:    tradingPair.BaseCurrency,
			Counter: tradingPair.QuoteCurrency,
		}
	})
}
