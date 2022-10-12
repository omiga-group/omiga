package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/preichenberger/go-coinbasepro/v2"
)

func BinanceSymbolsToTradingPairs(symbols []coinbasepro.Product) []models.TradingPair {
	return slices.Map(symbols, func(symbol coinbasepro.Product) models.TradingPair {
		return models.TradingPair{
			Symbol:  symbol.DisplayName,
			Base:    symbol.BaseCurrency,
			Counter: symbol.QuoteCurrency,
		}
	})
}
