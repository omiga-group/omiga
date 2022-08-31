package models

import (
	"fmt"
	"strings"
	"time"
)

const (
	MarketTypeFuture MarketType = "future"
	MarketTypeSpot   MarketType = "spot"
)

type MarketType string

func (mt *MarketType) UnmarshalJSON(raw []byte) error {
	rawStr := strings.TrimSuffix(strings.TrimPrefix(string(raw), "\""), "\"")
	mtSrt := strings.ToLower(rawStr)
	switch mtSrt {
	case "future":
		*mt = MarketTypeFuture
	case "spot":
		*mt = MarketTypeSpot
	default:
		return fmt.Errorf("market type `%s` is not valid or supported", mtSrt)
	}

	return nil
}

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

// There are far more fields to ftx market object but for now we only handle those that we actually
// need, ok!
type Market struct {
	Name          string     `json:"name"`
	BaseCurrency  string     `json:"baseCurrency"`
	QuoteCurrency string     `json:"quoteCurrency"`
	Type          MarketType `json:"type"`
	Enabled       bool       `json:"enabled"`
}

type MarketsMap map[string]Market
