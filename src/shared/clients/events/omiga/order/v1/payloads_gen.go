
// Code generated by go-omiga-template, DO NOT EDIT.

package orderv1

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type ID struct {
  UUID uuid.UUID
}

    
    // OrderEvent represents a OrderEvent model.
type OrderEvent struct {
  Metadata Metadata `json:"metadata"`
  Data Data `json:"data"`
}
    
    
    // Metadata represents a Metadata model.
type Metadata struct {
  Id ID `json:"id"` // The unique message ID
  Source string `json:"source"` // undefined
  Type AnonymousSchema3 `json:"type"`
  Subject string `json:"subject"` // undefined
  Time time.Time `json:"time"` // undefined
  CorrelationId ID `json:"correlationId"` // undefined
  Actor string `json:"actor"` // undefined
}
    
    
    // AnonymousSchema3 represents an enum of string.
type AnonymousSchema3 string

const (
  AnonymousSchema3OrderSubmitted AnonymousSchema3 = "orderSubmitted"
  AnonymousSchema3OrderCancel = "orderCancel"
)
    
    
    // Data represents a Data model.
type Data struct {
  BeforeState *Order `json:"beforeState",omitempty`
  AfterState Order `json:"afterState"`
}
    
    
    // Order represents a Order model.
type Order struct {
  Id ID `json:"id"` // The unique order ID
  OrderDetails OrderDetails `json:"orderDetails"`
  User *User `json:"user",omitempty`
  PreferredExchanges *AnonymousSchema23 `json:"preferredExchanges",omitempty`
  AdditionalProperties *[]interface{} `json:"additionalProperties",omitempty` // undefined
}
    
    
    // OrderDetails represents a OrderDetails model.
type OrderDetails struct {
  Id ID `json:"id"` // The unique order ID
  BaseCurrency *AnonymousSchema10 `json:"baseCurrency",omitempty`
  CounterCurrency *AnonymousSchema10 `json:"counterCurrency",omitempty`
  Type *AnonymousSchema15 `json:"type",omitempty`
  Side *AnonymousSchema16 `json:"side",omitempty`
  Quantity *AnonymousSchema17 `json:"quantity",omitempty`
  Price *AnonymousSchema17 `json:"price",omitempty`
  AdditionalProperties *[]interface{} `json:"additionalProperties",omitempty` // undefined
}
    
    
    // AnonymousSchema10 represents a AnonymousSchema10 model.
type AnonymousSchema10 struct {
  Name string `json:"name"` // undefined
  Code string `json:"code"` // undefined
  MaxPrecision int `json:"maxPrecision"` // undefined
  Digital bool `json:"digital"` // undefined
}
    
    
    // AnonymousSchema15 represents an enum of string.
type AnonymousSchema15 string

const (
  AnonymousSchema15Instant AnonymousSchema15 = "instant"
  AnonymousSchema15Market = "market"
  AnonymousSchema15Limit = "limit"
  AnonymousSchema15Stop = "stop"
  AnonymousSchema15TrailingStop = "trailing_stop"
)
    
    
    // AnonymousSchema16 represents an enum of string.
type AnonymousSchema16 string

const (
  AnonymousSchema16Bid AnonymousSchema16 = "bid"
  AnonymousSchema16Ask = "ask"
)
    
    
    // AnonymousSchema17 represents a AnonymousSchema17 model.
type AnonymousSchema17 struct {
  Amount int `json:"amount"` // undefined
  Scale int `json:"scale"` // undefined
  Currency AnonymousSchema10 `json:"currency"`
}
    
    
    // User represents a User model.
type User struct {
  Id ID `json:"id"` // undefined
  Created *time.Time `json:"created",omitempty` // undefined
  Updated *time.Time `json:"updated",omitempty` // undefined
  Type *UserType `json:"type",omitempty`
  AdditionalProperties *[]interface{} `json:"additionalProperties",omitempty` // undefined
}
    
    
    // UserType represents an enum of string.
type UserType string

const (
  UserTypeRetail UserType = "retail"
  UserTypeInstitution = "institution"
)
    
    
    // AnonymousSchema23 represents a AnonymousSchema23 model.
type AnonymousSchema23 struct {
  Id *string `json:"id",omitempty` // The unique ID of the supported exchange
}
    
func (i *ID) UnmarshalJSON(b []byte) error {
  if parsedUuid, err := uuid.Parse(strings.Trim(string(b), "\"")); err == nil {
    i.UUID = parsedUuid
  } else {
    return err
  }

  return nil
}
