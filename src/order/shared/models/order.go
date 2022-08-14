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
	MaxPrecision int
	Digital      bool
}

type Money struct {
	Amount   int
	Scale    int
	Currency Currency
}

type OrderDetails struct {
	BaseCurrency    Currency
	CounterCurrency Currency
	Type            OrderType
	Side            OrderSide
	Quantity        Money
	Price           Money
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
	Id int
}

type Order struct {
	Id                 int
	OrderDetails       OrderDetails
	User               *User
	PreferredExchanges []Exchange
}
