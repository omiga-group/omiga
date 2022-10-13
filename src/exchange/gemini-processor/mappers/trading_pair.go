package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	geminiv1 "github.com/omiga-group/omiga/src/shared/clients/openapi/gemini/v1"
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
