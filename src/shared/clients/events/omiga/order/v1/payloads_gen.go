
// Code generated by go-omiga-template, DO NOT EDIT.

package orderv1


    import (
      "time"
    
      "github.com/google/uuid"
    )
    
    type ID uuid.UUID

    
    
    // OrderEvent represents a OrderEvent model.
type OrderEvent struct {
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
  Id int `json:"id"`  // The unique order ID
  OrderDetails OrderDetails `json:"orderDetails"` 
  User *User `json:"user,omitempty"` 
  PreferredExchanges []Exchange `json:"preferredExchanges"`  // the preferred list of the supportef exchanged by the user
  AdditionalProperties *[]interface{} `json:"additionalProperties,omitempty"`  // undefined
}
    
    
    // OrderDetails represents a OrderDetails model.
type OrderDetails struct {
  BaseCurrency Currency `json:"baseCurrency"` 
  CounterCurrency Currency `json:"counterCurrency"` 
  Type OrderType `json:"type"` 
  Side OrderSide `json:"side"` 
  Quantity Quantity `json:"quantity"` 
  Price Money `json:"price"` 
  AdditionalProperties *[]interface{} `json:"additionalProperties,omitempty"`  // undefined
}
    
    
    // Currency represents a Currency model.
type Currency struct {
  Code string `json:"code"`  // undefined
  Name string `json:"name"`  // undefined
  MaxPrecision int32 `json:"maxPrecision"`  // undefined
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
    
    
    // Quantity represents a Quantity model.
type Quantity struct {
  Amount int64 `json:"amount"`  // undefined
  Scale int32 `json:"scale"`  // undefined
}
    
    
    // Money represents a Money model.
type Money struct {
  Quantity Quantity `json:"quantity"` 
  Currency Currency `json:"currency"` 
}
    
    
    // User represents a User model.
type User struct {
  Id ID `json:"id"`  // undefined
  Created *time.Time `json:"created,omitempty"`  // undefined
  Updated *time.Time `json:"updated,omitempty"`  // undefined
  Type *UserType `json:"type,omitempty"` 
  AdditionalProperties *[]interface{} `json:"additionalProperties,omitempty"`  // undefined
}
    
    
    
// UserType represents an enum of string.
type UserType string

const (
  UserTypeRetail UserType = "RETAIL"
  UserTypeInstitution UserType = "INSTITUTION"
)
    
    
    // Exchange represents a Exchange model.
type Exchange struct {
  Id string `json:"id"`  // The unique ID of the supported exchange
}
    