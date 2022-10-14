package mappers

import (
	"github.com/aopoltorzhicky/go_kraken/rest"
	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/shared/models"
)

func KrakenAssetPairsToTradingPairs(assetPairs map[string]rest.AssetPair) []models.TradingPair {
	return slices.Map(maps.Values(assetPairs), func(assetPair rest.AssetPair) models.TradingPair {
		return models.TradingPair{
			Symbol:                      assetPair.Altname,
			Base:                        assetPair.Base,
			BasePriceMinPrecision:       &assetPair.PairDecimals,
			BasePriceMaxPrecision:       &assetPair.PairDecimals,
			BaseQuantityMinPrecision:    &assetPair.LotDecimals,
			BaseQuantityMaxPrecision:    &assetPair.LotDecimals,
			Counter:                     assetPair.Quote,
			CounterPriceMinPrecision:    &assetPair.PairDecimals,
			CounterPriceMaxPrecision:    &assetPair.PairDecimals,
			CounterQuantityMinPrecision: &assetPair.LotDecimals,
			CounterQuantityMaxPrecision: &assetPair.LotDecimals,
		}
	})
}
