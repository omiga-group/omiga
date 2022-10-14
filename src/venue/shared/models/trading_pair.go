package models

type TradingPair struct {
	Symbol string

	Base                     string
	BasePriceMinPrecision    *int
	BasePriceMaxPrecision    *int
	BaseQuantityMinPrecision *int
	BaseQuantityMaxPrecision *int

	Counter                     string
	CounterPriceMinPrecision    *int
	CounterPriceMaxPrecision    *int
	CounterQuantityMinPrecision *int
	CounterQuantityMaxPrecision *int
}
