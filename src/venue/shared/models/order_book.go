package models

import "time"

type OrderCurrency struct {
	Name         string
	Code         string
	MaxPrecision int32
	Digital      bool
}

type Quantity struct {
	Amount int64
	Scale  int32
}

type OrderBookEntry struct {
	Time     time.Time
	Quantity Quantity
	Price    Quantity
}

type OrderBook struct {
	ExchangeId      string
	BaseCurrency    OrderCurrency
	CounterCurrency OrderCurrency
	Bids            []OrderBookEntry
	Asks            []OrderBookEntry
}
