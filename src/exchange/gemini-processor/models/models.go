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

type MarketName string
type MarketNames []MarketName
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

// There are far more fields to gemini market object but for now we only handle those that we actually need, ok!
type Market struct {
	Name          string `json:"symbol"`
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

type MarketsMap map[string]Market
