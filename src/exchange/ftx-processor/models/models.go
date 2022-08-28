package models

import (
	"time"
)

type PriceLevel struct {
	Price    float64
	Quantity float64
}

type OrderBookEntry struct {
	Symbol string
	Time   time.Time
	Bid    *PriceLevel
	Ask    *PriceLevel
}
