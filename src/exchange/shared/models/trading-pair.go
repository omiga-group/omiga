package models

type TradingPair struct {
	Symbol           string
	Base             string
	BasePrecision    *int
	Counter          string
	CounterPrecision *int
}
