package mappers

import (
	"github.com/aopoltorzhicky/go_kraken/rest"
	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
)

func KrakenAssetPairsToTradingPairs(assetPairs map[string]rest.AssetPair) []models.TradingPair {
	return slices.Map(maps.Values(assetPairs), func(assetPair rest.AssetPair) models.TradingPair {
		return models.TradingPair{
			Symbol:           assetPair.Altname,
			Base:             assetPair.Base,
			BasePrecision:    &assetPair.PairDecimals,
			Counter:          assetPair.Quote,
			CounterPrecision: &assetPair.PairDecimals,
		}
	})
}
