package models

import (
	"time"

	"github.com/google/uuid"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
)

type ID uuid.UUID

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
	Type            orderv1.OrderType
	Side            orderv1.OrderSide
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
