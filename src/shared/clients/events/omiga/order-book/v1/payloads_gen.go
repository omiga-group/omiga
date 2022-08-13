
// Code generated by go-omiga-template, DO NOT EDIT.

package orderbookv1

import (
	"time"

	"github.com/google/uuid"
)

type ID uuid.UUID

    
    // OrderBookEvent represents a OrderBookEvent model.
type OrderBookEvent struct {
  Metadata Metadata `json:"metadata"`
  Data OrderBook `json:"data"`
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
  AnonymousSchema3OrderBookUpdated AnonymousSchema3 = "orderBookUpdated"
)
    
    
    // OrderBook represents a OrderBook model.
type OrderBook struct {
  Id string `json:"id"` // undefined
  BaseCurrency *AnonymousSchema9 `json:"baseCurrency",omitempty`
  CounterCurrency *AnonymousSchema9 `json:"counterCurrency",omitempty`
  Bids *AnonymousSchema14 `json:"bids",omitempty`
  Asks *AnonymousSchema15 `json:"asks",omitempty`
  AdditionalProperties *[]interface{} `json:"additionalProperties",omitempty` // undefined
}
    
    
    // AnonymousSchema9 represents a AnonymousSchema9 model.
type AnonymousSchema9 struct {
  Name string `json:"name"` // undefined
  Code string `json:"code"` // undefined
  MaxPrecision int `json:"maxPrecision"` // undefined
  Digital bool `json:"digital"` // undefined
}
    
    
    // AnonymousSchema14 represents a AnonymousSchema14 model.
type AnonymousSchema14 struct {
  Quantity *AnonymousSchema16 `json:"quantity",omitempty`
  Price *AnonymousSchema16 `json:"price",omitempty`
}
    
    
    // AnonymousSchema16 represents a AnonymousSchema16 model.
type AnonymousSchema16 struct {
  Amount int `json:"amount"` // undefined
  Scale int `json:"scale"` // undefined
  Currency AnonymousSchema9 `json:"currency"`
}
    
    
    // AnonymousSchema15 represents a AnonymousSchema15 model.
type AnonymousSchema15 struct {
  Quantity *AnonymousSchema16 `json:"quantity",omitempty`
  Price *AnonymousSchema16 `json:"price",omitempty`
}
    