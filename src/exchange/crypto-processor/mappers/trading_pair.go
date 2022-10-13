package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	cryptov2 "github.com/omiga-group/omiga/src/shared/clients/openapi/crypto/v2"
)

func CryptoInstrumentsToTradingPairs(instruments []cryptov2.Instrument) []models.TradingPair {
	return slices.Map(instruments, func(instrument cryptov2.Instrument) models.TradingPair {
		return models.TradingPair{
			Symbol:                      instrument.InstrumentName,
			Base:                        instrument.BaseCurrency,
			BasePriceMinPrecision:       &instrument.PriceDecimals,
			BasePriceMaxPrecision:       &instrument.PriceDecimals,
			BaseQuantityMinPrecision:    &instrument.QuantityDecimals,
			BaseQuantityMaxPrecision:    &instrument.QuantityDecimals,
			Counter:                     instrument.QuoteCurrency,
			CounterPriceMinPrecision:    &instrument.PriceDecimals,
			CounterPriceMaxPrecision:    &instrument.PriceDecimals,
			CounterQuantityMinPrecision: &instrument.QuantityDecimals,
			CounterQuantityMaxPrecision: &instrument.QuantityDecimals,
		}
	})
}
