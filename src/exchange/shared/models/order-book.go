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
	Time     time.Time
	Quantity Quantity
	Price    Money
}

type OrderBook struct {
	ExchangeId      string
	BaseCurrency    Currency
	CounterCurrency Currency
	Bids            []OrderBookEntry
	Asks            []OrderBookEntry
}
