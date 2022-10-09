
func ToTradingPairs(assetPairs map[string]rest.AssetPair) []models.TradingPair {
	return slices.Map(maps.Values(assetPairs), func(assetPair rest.AssetPair) models.TradingPair {
		return models.TradingPair{
			Symbol:           assetPair.Altname,
			Base:             assetPair.Base,
			BasePrecision:    assetPair.PairDecimals,
			Counter:          assetPair.Quote,
			CounterPrecision: assetPair.PairDecimals,
		}
	})
}