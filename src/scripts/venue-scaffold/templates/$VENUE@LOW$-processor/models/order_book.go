package models

import (
	"time"
)

type $VENUE@PAS$OrderBookEntry struct {
	Symbol string
	Time   time.Time
	Bid    *binance.Bid
	Ask    *binance.Ask
}
