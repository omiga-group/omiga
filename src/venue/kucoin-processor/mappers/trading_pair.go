package mappers

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func KuCoinSymbolModelToTradingPairs(symbols kucoin.SymbolsModel) []models.TradingPair {
	return slices.Map(symbols, func(symbol *kucoin.SymbolModel) models.TradingPair {
		return models.TradingPair{
			Symbol:  symbol.Symbol,
			Base:    symbol.BaseCurrency,
			Counter: symbol.QuoteCurrency,
		}
	})
}
