package mappers

import (
	"strings"

	"github.com/life4/genesis/maps"
	"github.com/omiga-group/omiga/src/venue/shared/models"
	xtv1 "github.com/omiga-group/omiga/src/venue/xt-processor/xtclient/v1"
)

func XtMarketConfigsToTradingPairs(marketConfigs map[string]xtv1.MarketConfig) []models.TradingPair {
	return maps.Values(
		maps.Map(marketConfigs, func(symbol string, marketConfig xtv1.MarketConfig) (string, models.TradingPair) {
			baseCounter := strings.Split(symbol, "_")

			return symbol, models.TradingPair{
				Symbol:                   symbol,
				Base:                     baseCounter[0],
				BasePriceMinPrecision:    &marketConfig.PricePoint,
				BasePriceMaxPrecision:    &marketConfig.PricePoint,
				Counter:                  baseCounter[1],
				CounterPriceMinPrecision: &marketConfig.PricePoint,
				CounterPriceMaxPrecision: &marketConfig.PricePoint,
			}
		}))
}
