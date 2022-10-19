// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type ConvertedDetails struct {
	Btc float64 `json:"btc"`
	Eth float64 `json:"eth"`
	Usd float64 `json:"usd"`
}

type Links struct {
	Website  *string `json:"website"`
	Facebook *string `json:"facebook"`
	Reddit   *string `json:"reddit"`
	Twitter  *string `json:"twitter"`
	Slack    *string `json:"slack"`
	Telegram *string `json:"telegram"`
}

type TickerMarket struct {
	HasTradingIncentive bool    `json:"hasTradingIncentive"`
	Identifier          string  `json:"identifier"`
	Name                *string `json:"name"`
}

type CurrencyType string

const (
	CurrencyTypeDigital CurrencyType = "DIGITAL"
	CurrencyTypeFiat    CurrencyType = "FIAT"
)

var AllCurrencyType = []CurrencyType{
	CurrencyTypeDigital,
	CurrencyTypeFiat,
}

func (e CurrencyType) IsValid() bool {
	switch e {
	case CurrencyTypeDigital, CurrencyTypeFiat:
		return true
	}
	return false
}

func (e CurrencyType) String() string {
	return string(e)
}

func (e *CurrencyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CurrencyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CurrencyType", str)
	}
	return nil
}

func (e CurrencyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MarketType string

const (
	MarketTypeSpotTrading     MarketType = "SPOT_TRADING"
	MarketTypeMarginTrading   MarketType = "MARGIN_TRADING"
	MarketTypeDerivatives     MarketType = "DERIVATIVES"
	MarketTypeEarn            MarketType = "EARN"
	MarketTypePerpetual       MarketType = "PERPETUAL"
	MarketTypeFutures         MarketType = "FUTURES"
	MarketTypeWarrant         MarketType = "WARRANT"
	MarketTypeOtc             MarketType = "OTC"
	MarketTypeYield           MarketType = "YIELD"
	MarketTypeP2p             MarketType = "P2P"
	MarketTypeStrategyTrading MarketType = "STRATEGY_TRADING"
	MarketTypeSwapFarming     MarketType = "SWAP_FARMING"
	MarketTypeFanToken        MarketType = "FAN_TOKEN"
	MarketTypeEtf             MarketType = "ETF"
	MarketTypeNft             MarketType = "NFT"
	MarketTypeSwap            MarketType = "SWAP"
	MarketTypeCfd             MarketType = "CFD"
	MarketTypeLiquidity       MarketType = "LIQUIDITY"
	MarketTypeFarm            MarketType = "FARM"
)

var AllMarketType = []MarketType{
	MarketTypeSpotTrading,
	MarketTypeMarginTrading,
	MarketTypeDerivatives,
	MarketTypeEarn,
	MarketTypePerpetual,
	MarketTypeFutures,
	MarketTypeWarrant,
	MarketTypeOtc,
	MarketTypeYield,
	MarketTypeP2p,
	MarketTypeStrategyTrading,
	MarketTypeSwapFarming,
	MarketTypeFanToken,
	MarketTypeEtf,
	MarketTypeNft,
	MarketTypeSwap,
	MarketTypeCfd,
	MarketTypeLiquidity,
	MarketTypeFarm,
}

func (e MarketType) IsValid() bool {
	switch e {
	case MarketTypeSpotTrading, MarketTypeMarginTrading, MarketTypeDerivatives, MarketTypeEarn, MarketTypePerpetual, MarketTypeFutures, MarketTypeWarrant, MarketTypeOtc, MarketTypeYield, MarketTypeP2p, MarketTypeStrategyTrading, MarketTypeSwapFarming, MarketTypeFanToken, MarketTypeEtf, MarketTypeNft, MarketTypeSwap, MarketTypeCfd, MarketTypeLiquidity, MarketTypeFarm:
		return true
	}
	return false
}

func (e MarketType) String() string {
	return string(e)
}

func (e *MarketType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MarketType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MarketType", str)
	}
	return nil
}

func (e MarketType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OutboxStatus string

const (
	OutboxStatusPending   OutboxStatus = "PENDING"
	OutboxStatusSucceeded OutboxStatus = "SUCCEEDED"
	OutboxStatusFailed    OutboxStatus = "FAILED"
)

var AllOutboxStatus = []OutboxStatus{
	OutboxStatusPending,
	OutboxStatusSucceeded,
	OutboxStatusFailed,
}

func (e OutboxStatus) IsValid() bool {
	switch e {
	case OutboxStatusPending, OutboxStatusSucceeded, OutboxStatusFailed:
		return true
	}
	return false
}

func (e OutboxStatus) String() string {
	return string(e)
}

func (e *OutboxStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OutboxStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OutboxStatus", str)
	}
	return nil
}

func (e OutboxStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VenueType string

const (
	VenueTypeExchange VenueType = "EXCHANGE"
)

var AllVenueType = []VenueType{
	VenueTypeExchange,
}

func (e VenueType) IsValid() bool {
	switch e {
	case VenueTypeExchange:
		return true
	}
	return false
}

func (e VenueType) String() string {
	return string(e)
}

func (e *VenueType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VenueType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VenueType", str)
	}
	return nil
}

func (e VenueType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
