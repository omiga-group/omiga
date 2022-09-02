package models

import (
	"time"

	"github.com/adshao/go-binance/v2"
)

type BinanceOrderBookEntry struct {
	Symbol string
	Time   time.Time
	Bid    *binance.Bid
	Ask    *binance.Ask
}
