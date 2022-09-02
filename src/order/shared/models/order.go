package models

import (
	"time"

	"github.com/google/uuid"
)

type ID uuid.UUID

type OrderType string

const (
	OrderTypeInstant      OrderType = "INSTANT"
	OrderTypeMarket       OrderType = "MARKET"
	OrderTypeLimit        OrderType = "LIMIT"
	OrderTypeStop         OrderType = "STOP"
	OrderTypeTrailingStop OrderType = "TRAILING_STOP"
)

type OrderSide string

const (
	OrderSideBid OrderSide = "BID"
	OrderSideAsk OrderSide = "ASK"
)

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

type OrderDetails struct {
	BaseCurrency    Currency
	CounterCurrency Currency
	Type            OrderType
	Side            OrderSide
	Quantity        Quantity
	Price           Quantity
}

type UserType string

const (
	UserTypeRetail      UserType = "RETAIL"
	UserTypeInstitution UserType = "INSTITUTION"
)

type User struct {
	Id      ID
	Created *time.Time
	Updated *time.Time
	Type    *UserType
}

type Exchange struct {
	Id string
}

type Order struct {
	Id                 int
	OrderDetails       OrderDetails
	User               User
	PreferredExchanges []Exchange
}

type OrderBookEntry struct {
	Quantity Quantity
	Price    Quantity
}

type OrderBook struct {
	BaseCurrency    Currency
	CounterCurrency Currency
	Time            time.Time
	Bids            []OrderBookEntry
	Asks            []OrderBookEntry
}
