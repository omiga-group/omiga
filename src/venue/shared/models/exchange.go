package models

import (
	"time"

	venuerepo "github.com/omiga-group/omiga/src/venue/shared/entities/venue"
)

type Market struct {
	HasTradingIncentive bool
	Identifier          string
	Name                string
}

type ConvertedDetails struct {
	Btc float64
	Eth float64
	Usd float64
}

type Ticker struct {
	Base                   string
	BaseCoinId             string
	Counter                string
	CounterCoinId          string
	BidAskSpreadPercentage float64
	ConvertedLast          ConvertedDetails
	ConvertedVolume        ConvertedDetails
	IsAnomaly              bool
	IsStale                bool
	Last                   float64
	LastFetchAt            time.Time
	LastTradedAt           time.Time
	Market                 Market
	Timestamp              time.Time
	TokenInfoUrl           *string
	TradeUrl               string
	TrustScore             string
	Volume                 float64
}

type Venue struct {
	Id                          int
	VenueId                     string
	Type                        venuerepo.Type
	Name                        string
	YearEstablished             int
	Country                     string
	Image                       string
	Links                       map[string]string
	HasTradingIncentive         bool
	Centralized                 bool
	PublicNotice                string
	AlertNotice                 string
	TrustScore                  int
	TrustScoreRank              int
	TradeVolume24hBtc           float64
	TradeVolume24hBtcNormalized float64
	Tickers                     []Ticker
	MakerFee                    *float64
	TakerFee                    *float64
	SpreadFee                   *bool
	SupportAPI                  *bool
}
