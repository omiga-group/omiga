package mappers

import (
	"github.com/life4/genesis/slices"
	ftxv1 "github.com/omiga-group/omiga/src/venue/ftx-processor/ftxclient/v1"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func FtxMarketToTradingPairs(markets []ftxv1.Market) []models.TradingPair {
	return slices.Map(markets, func(market ftxv1.Market) models.TradingPair {
		return models.TradingPair{
			Symbol:  market.Name,
			Base:    market.BaseCurrency,
			Counter: market.QuoteCurrency,
		}
	})
}
