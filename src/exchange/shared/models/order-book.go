package models

type OrderBook struct {
	ExchangeId      string
	BaseCurrency    Currency
	CounterCurrency Currency
	Bids            []OrderBookEntry
	Asks            []OrderBookEntry
}

type Currency struct {
	Name         string
	Code         string
	MaxPrecision int
	Digital      bool
}

type OrderBookEntry struct {
	Quantity Money
	Price    Money
}

type Money struct {
	Amount   int
	Scale    int
	Currency Currency
}
