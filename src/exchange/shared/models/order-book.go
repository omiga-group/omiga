package models

import "time"

type Currency struct {
	Name         string
	Code         string
	MaxPrecision int32
	Digital      bool
}

type Quantity struct {
	Amount int64
	Scale  int32
}

type Money struct {
	Quantity Quantity
	Currency Currency
}

type OrderBookEntry struct {
	Quantity Quantity
	Price    Money
}

type OrderBook struct {
	ExchangeId      string
	BaseCurrency    Currency
	CounterCurrency Currency
	Time            time.Time
	Bids            []OrderBookEntry
	Asks            []OrderBookEntry
}
