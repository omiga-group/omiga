package mappers

import (
	"github.com/life4/genesis/slices"
	dextradev1 "github.com/omiga-group/omiga/src/venue/dextrade-processor/dextradeclient/v1"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func DextradeSymbolsToTradingPairs(symbols []dextradev1.Symbol) []models.TradingPair {
	return slices.Map(symbols, func(symbol dextradev1.Symbol) models.TradingPair {
		return models.TradingPair{
			Symbol:  symbol.Pair,
			Base:    symbol.Base,
			Counter: symbol.Quote,
		}
	})
}
