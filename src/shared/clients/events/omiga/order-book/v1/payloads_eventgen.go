
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
  TypeOrderBookUpdated Type = "orderBookUpdated"
)
    
    
    // OrderBook represents a OrderBook model.
type OrderBook struct {
  ExchangeId string `json:"exchangeId"`  // The unique ID of the exchange
  BaseCurrency Currency `json:"baseCurrency"` 
  CounterCurrency Currency `json:"counterCurrency"` 
  Bids []OrderBookEntry `json:"bids"`  // undefined
  Asks []OrderBookEntry `json:"asks"`  // undefined
  AdditionalProperties *[]interface{} `json:"additionalProperties,omitempty"`  // undefined
}
    
    
    // Currency represents a Currency model.
type Currency struct {
  Code string `json:"code"`  // undefined
  Name string `json:"name"`  // undefined
  MaxPrecision int32 `json:"maxPrecision"`  // undefined
  Digital bool `json:"digital"`  // undefined
}
    
    
    // OrderBookEntry represents a OrderBookEntry model.
type OrderBookEntry struct {
  Time time.Time `json:"time"`  // The order book entry timestamp
  Quantity Quantity `json:"quantity"` 
  Price Money `json:"price"` 
}
    
    
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
    