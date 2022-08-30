package models

import (
	"time"

	"github.com/omiga-group/omiga/src/exchange/shared/models"
)

type Ticker struct {
	Base                   string
	BidAskSpreadPercentage float64
	CoinId                 string
	ConvertedLast          models.ConvertedDetails
	ConvertedVolume        models.ConvertedDetails
	IsAnomaly              bool
	IsStale                bool
	Last                   float64
	LastFetchAt            time.Time
	LastTradedAt           time.Time
	Market                 models.Market
	Target                 string
	TargetCoinId           string
	Timestamp              time.Time
	TokenInfoUrl           *string
	TradeUrl               string
	TrustScore             string
	Volume                 float64
}

type Exchange struct {
	Id                          int
	ExchangeId                  string
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
}
