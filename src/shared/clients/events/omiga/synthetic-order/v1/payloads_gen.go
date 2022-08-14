
// Code generated by go-omiga-template, DO NOT EDIT.

package syntheticorderv1


    import (
      "time"
    
      "github.com/google/uuid"
    )
    
    type ID uuid.UUID

    
    
    // SyntheticOrderEvent represents a SyntheticOrderEvent model.
type SyntheticOrderEvent struct {
  Metadata Metadata `json:"metadata"` 
  Data Data `json:"data"` 
}
    
    
    // Metadata represents a Metadata model.
type Metadata struct {
  Id ID `json:"id"`  // The unique message ID
  Source string `json:"source"`  // undefined
  Type Type `json:"type"` 
  Subject string `json:"subject"`  // undefined
  Time time.Time `json:"time"`  // undefined
  CorrelationId ID `json:"correlationId"`  // undefined
  Actor string `json:"actor"`  // undefined
}
    
    
    
// Type represents an enum of string.
type Type string

const (
  TypeOrderSubmitted Type = "orderSubmitted"
  TypeOrderCancel Type = "orderCancel"
)
    
    
    // Data represents a Data model.
type Data struct {
  BeforeState *Order `json:"beforeState,omitempty"` 
  AfterState Order `json:"afterState"` 
}
    
    
    // Order represents a Order model.
type Order struct {
  Id int `json:"id"`  // The unique synthetic order ID
  Exchange *Exchange `json:"exchange,omitempty"` 
  OrderDetails OrderDetails `json:"orderDetails"` 
  AdditionalProperties *[]interface{} `json:"additionalProperties,omitempty"`  // undefined
}
    
    
    // Exchange represents a Exchange model.
type Exchange struct {
  Id int `json:"id"`  // The unique ID of the supported exchange
}
    
    
    // OrderDetails represents a OrderDetails model.
type OrderDetails struct {
  BaseCurrency Currency `json:"baseCurrency"` 
  CounterCurrency Currency `json:"counterCurrency"` 
  Type OrderType `json:"type"` 
  Side OrderSide `json:"side"` 
  Quantity Money `json:"quantity"` 
  Price Money `json:"price"` 
  AdditionalProperties *[]interface{} `json:"additionalProperties,omitempty"`  // undefined
}
    
    
    // Currency represents a Currency model.
type Currency struct {
  Name string `json:"name"`  // undefined
  Code string `json:"code"`  // undefined
  MaxPrecision int `json:"maxPrecision"`  // undefined
  Digital bool `json:"digital"`  // undefined
}
    
    
    
// OrderType represents an enum of string.
type OrderType string

const (
  OrderTypeInstant OrderType = "INSTANT"
  OrderTypeMarket OrderType = "MARKET"
  OrderTypeLimit OrderType = "LIMIT"
  OrderTypeStop OrderType = "STOP"
  OrderTypeTrailingStop OrderType = "TRAILING_STOP"
)
    
    
    
// OrderSide represents an enum of string.
type OrderSide string

const (
  OrderSideBid OrderSide = "BID"
  OrderSideAsk OrderSide = "ASK"
)
    
    
    // Money represents a Money model.
type Money struct {
  Amount int `json:"amount"`  // undefined
  Scale int `json:"scale"`  // undefined
  Currency Currency `json:"currency"` 
}
    