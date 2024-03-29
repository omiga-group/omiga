// Package xtv4 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package xtv4

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Filter defines model for filter.
type Filter struct {
	DurationSeconds      *string                `json:"durationSeconds,omitempty"`
	Filter               string                 `json:"filter"`
	Max                  *string                `json:"max,omitempty"`
	MaxPriceMultiple     *string                `json:"maxPriceMultiple,omitempty"`
	Min                  string                 `json:"min"`
	TickSize             *string                `json:"tickSize,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// GetSymbolsResponse defines model for getSymbolsResponse.
type GetSymbolsResponse struct {
	Mc                   string                 `json:"mc"`
	Rc                   int                    `json:"rc"`
	Result               Result                 `json:"result"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Result defines model for result.
type Result struct {
	Symbols              []Symbol               `json:"symbols"`
	Time                 int                    `json:"time"`
	Version              string                 `json:"version"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Symbol defines model for symbol.
type Symbol struct {
	BaseCurrency           string                 `json:"baseCurrency"`
	BaseCurrencyId         int                    `json:"baseCurrencyId"`
	BaseCurrencyPrecision  int                    `json:"baseCurrencyPrecision"`
	DepthMergePrecision    int                    `json:"depthMergePrecision"`
	DisplayLevel           string                 `json:"displayLevel"`
	DisplayWeight          int                    `json:"displayWeight"`
	Filters                []Filter               `json:"filters"`
	Id                     int                    `json:"id"`
	NextState              *string                `json:"nextState,omitempty"`
	NextStateTime          *int                   `json:"nextStateTime,omitempty"`
	OrderTypes             []string               `json:"orderTypes"`
	Plates                 []int                  `json:"plates"`
	PricePrecision         int                    `json:"pricePrecision"`
	QuantityPrecision      int                    `json:"quantityPrecision"`
	QuoteCurrency          string                 `json:"quoteCurrency"`
	QuoteCurrencyId        int                    `json:"quoteCurrencyId"`
	QuoteCurrencyPrecision int                    `json:"quoteCurrencyPrecision"`
	State                  string                 `json:"state"`
	StateTime              int                    `json:"stateTime"`
	Symbol                 string                 `json:"symbol"`
	TimeInForces           []string               `json:"timeInForces"`
	TradingEnabled         bool                   `json:"tradingEnabled"`
	AdditionalProperties   map[string]interface{} `json:"-"`
}

// Getter for additional properties for Filter. Returns the specified
// element and whether it was found
func (a Filter) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Filter
func (a *Filter) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Filter to handle AdditionalProperties
func (a *Filter) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["durationSeconds"]; found {
		err = json.Unmarshal(raw, &a.DurationSeconds)
		if err != nil {
			return fmt.Errorf("error reading 'durationSeconds': %w", err)
		}
		delete(object, "durationSeconds")
	}

	if raw, found := object["filter"]; found {
		err = json.Unmarshal(raw, &a.Filter)
		if err != nil {
			return fmt.Errorf("error reading 'filter': %w", err)
		}
		delete(object, "filter")
	}

	if raw, found := object["max"]; found {
		err = json.Unmarshal(raw, &a.Max)
		if err != nil {
			return fmt.Errorf("error reading 'max': %w", err)
		}
		delete(object, "max")
	}

	if raw, found := object["maxPriceMultiple"]; found {
		err = json.Unmarshal(raw, &a.MaxPriceMultiple)
		if err != nil {
			return fmt.Errorf("error reading 'maxPriceMultiple': %w", err)
		}
		delete(object, "maxPriceMultiple")
	}

	if raw, found := object["min"]; found {
		err = json.Unmarshal(raw, &a.Min)
		if err != nil {
			return fmt.Errorf("error reading 'min': %w", err)
		}
		delete(object, "min")
	}

	if raw, found := object["tickSize"]; found {
		err = json.Unmarshal(raw, &a.TickSize)
		if err != nil {
			return fmt.Errorf("error reading 'tickSize': %w", err)
		}
		delete(object, "tickSize")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshalling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Filter to handle AdditionalProperties
func (a Filter) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.DurationSeconds != nil {
		object["durationSeconds"], err = json.Marshal(a.DurationSeconds)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'durationSeconds': %w", err)
		}
	}

	object["filter"], err = json.Marshal(a.Filter)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'filter': %w", err)
	}

	if a.Max != nil {
		object["max"], err = json.Marshal(a.Max)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'max': %w", err)
		}
	}

	if a.MaxPriceMultiple != nil {
		object["maxPriceMultiple"], err = json.Marshal(a.MaxPriceMultiple)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'maxPriceMultiple': %w", err)
		}
	}

	object["min"], err = json.Marshal(a.Min)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'min': %w", err)
	}

	if a.TickSize != nil {
		object["tickSize"], err = json.Marshal(a.TickSize)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'tickSize': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for GetSymbolsResponse. Returns the specified
// element and whether it was found
func (a GetSymbolsResponse) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GetSymbolsResponse
func (a *GetSymbolsResponse) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GetSymbolsResponse to handle AdditionalProperties
func (a *GetSymbolsResponse) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["mc"]; found {
		err = json.Unmarshal(raw, &a.Mc)
		if err != nil {
			return fmt.Errorf("error reading 'mc': %w", err)
		}
		delete(object, "mc")
	}

	if raw, found := object["rc"]; found {
		err = json.Unmarshal(raw, &a.Rc)
		if err != nil {
			return fmt.Errorf("error reading 'rc': %w", err)
		}
		delete(object, "rc")
	}

	if raw, found := object["result"]; found {
		err = json.Unmarshal(raw, &a.Result)
		if err != nil {
			return fmt.Errorf("error reading 'result': %w", err)
		}
		delete(object, "result")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshalling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for GetSymbolsResponse to handle AdditionalProperties
func (a GetSymbolsResponse) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["mc"], err = json.Marshal(a.Mc)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'mc': %w", err)
	}

	object["rc"], err = json.Marshal(a.Rc)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'rc': %w", err)
	}

	object["result"], err = json.Marshal(a.Result)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'result': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for Result. Returns the specified
// element and whether it was found
func (a Result) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Result
func (a *Result) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Result to handle AdditionalProperties
func (a *Result) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["symbols"]; found {
		err = json.Unmarshal(raw, &a.Symbols)
		if err != nil {
			return fmt.Errorf("error reading 'symbols': %w", err)
		}
		delete(object, "symbols")
	}

	if raw, found := object["time"]; found {
		err = json.Unmarshal(raw, &a.Time)
		if err != nil {
			return fmt.Errorf("error reading 'time': %w", err)
		}
		delete(object, "time")
	}

	if raw, found := object["version"]; found {
		err = json.Unmarshal(raw, &a.Version)
		if err != nil {
			return fmt.Errorf("error reading 'version': %w", err)
		}
		delete(object, "version")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshalling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Result to handle AdditionalProperties
func (a Result) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["symbols"], err = json.Marshal(a.Symbols)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'symbols': %w", err)
	}

	object["time"], err = json.Marshal(a.Time)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'time': %w", err)
	}

	object["version"], err = json.Marshal(a.Version)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'version': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for Symbol. Returns the specified
// element and whether it was found
func (a Symbol) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Symbol
func (a *Symbol) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Symbol to handle AdditionalProperties
func (a *Symbol) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["baseCurrency"]; found {
		err = json.Unmarshal(raw, &a.BaseCurrency)
		if err != nil {
			return fmt.Errorf("error reading 'baseCurrency': %w", err)
		}
		delete(object, "baseCurrency")
	}

	if raw, found := object["baseCurrencyId"]; found {
		err = json.Unmarshal(raw, &a.BaseCurrencyId)
		if err != nil {
			return fmt.Errorf("error reading 'baseCurrencyId': %w", err)
		}
		delete(object, "baseCurrencyId")
	}

	if raw, found := object["baseCurrencyPrecision"]; found {
		err = json.Unmarshal(raw, &a.BaseCurrencyPrecision)
		if err != nil {
			return fmt.Errorf("error reading 'baseCurrencyPrecision': %w", err)
		}
		delete(object, "baseCurrencyPrecision")
	}

	if raw, found := object["depthMergePrecision"]; found {
		err = json.Unmarshal(raw, &a.DepthMergePrecision)
		if err != nil {
			return fmt.Errorf("error reading 'depthMergePrecision': %w", err)
		}
		delete(object, "depthMergePrecision")
	}

	if raw, found := object["displayLevel"]; found {
		err = json.Unmarshal(raw, &a.DisplayLevel)
		if err != nil {
			return fmt.Errorf("error reading 'displayLevel': %w", err)
		}
		delete(object, "displayLevel")
	}

	if raw, found := object["displayWeight"]; found {
		err = json.Unmarshal(raw, &a.DisplayWeight)
		if err != nil {
			return fmt.Errorf("error reading 'displayWeight': %w", err)
		}
		delete(object, "displayWeight")
	}

	if raw, found := object["filters"]; found {
		err = json.Unmarshal(raw, &a.Filters)
		if err != nil {
			return fmt.Errorf("error reading 'filters': %w", err)
		}
		delete(object, "filters")
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.Id)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["nextState"]; found {
		err = json.Unmarshal(raw, &a.NextState)
		if err != nil {
			return fmt.Errorf("error reading 'nextState': %w", err)
		}
		delete(object, "nextState")
	}

	if raw, found := object["nextStateTime"]; found {
		err = json.Unmarshal(raw, &a.NextStateTime)
		if err != nil {
			return fmt.Errorf("error reading 'nextStateTime': %w", err)
		}
		delete(object, "nextStateTime")
	}

	if raw, found := object["orderTypes"]; found {
		err = json.Unmarshal(raw, &a.OrderTypes)
		if err != nil {
			return fmt.Errorf("error reading 'orderTypes': %w", err)
		}
		delete(object, "orderTypes")
	}

	if raw, found := object["plates"]; found {
		err = json.Unmarshal(raw, &a.Plates)
		if err != nil {
			return fmt.Errorf("error reading 'plates': %w", err)
		}
		delete(object, "plates")
	}

	if raw, found := object["pricePrecision"]; found {
		err = json.Unmarshal(raw, &a.PricePrecision)
		if err != nil {
			return fmt.Errorf("error reading 'pricePrecision': %w", err)
		}
		delete(object, "pricePrecision")
	}

	if raw, found := object["quantityPrecision"]; found {
		err = json.Unmarshal(raw, &a.QuantityPrecision)
		if err != nil {
			return fmt.Errorf("error reading 'quantityPrecision': %w", err)
		}
		delete(object, "quantityPrecision")
	}

	if raw, found := object["quoteCurrency"]; found {
		err = json.Unmarshal(raw, &a.QuoteCurrency)
		if err != nil {
			return fmt.Errorf("error reading 'quoteCurrency': %w", err)
		}
		delete(object, "quoteCurrency")
	}

	if raw, found := object["quoteCurrencyId"]; found {
		err = json.Unmarshal(raw, &a.QuoteCurrencyId)
		if err != nil {
			return fmt.Errorf("error reading 'quoteCurrencyId': %w", err)
		}
		delete(object, "quoteCurrencyId")
	}

	if raw, found := object["quoteCurrencyPrecision"]; found {
		err = json.Unmarshal(raw, &a.QuoteCurrencyPrecision)
		if err != nil {
			return fmt.Errorf("error reading 'quoteCurrencyPrecision': %w", err)
		}
		delete(object, "quoteCurrencyPrecision")
	}

	if raw, found := object["state"]; found {
		err = json.Unmarshal(raw, &a.State)
		if err != nil {
			return fmt.Errorf("error reading 'state': %w", err)
		}
		delete(object, "state")
	}

	if raw, found := object["stateTime"]; found {
		err = json.Unmarshal(raw, &a.StateTime)
		if err != nil {
			return fmt.Errorf("error reading 'stateTime': %w", err)
		}
		delete(object, "stateTime")
	}

	if raw, found := object["symbol"]; found {
		err = json.Unmarshal(raw, &a.Symbol)
		if err != nil {
			return fmt.Errorf("error reading 'symbol': %w", err)
		}
		delete(object, "symbol")
	}

	if raw, found := object["timeInForces"]; found {
		err = json.Unmarshal(raw, &a.TimeInForces)
		if err != nil {
			return fmt.Errorf("error reading 'timeInForces': %w", err)
		}
		delete(object, "timeInForces")
	}

	if raw, found := object["tradingEnabled"]; found {
		err = json.Unmarshal(raw, &a.TradingEnabled)
		if err != nil {
			return fmt.Errorf("error reading 'tradingEnabled': %w", err)
		}
		delete(object, "tradingEnabled")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshalling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Symbol to handle AdditionalProperties
func (a Symbol) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["baseCurrency"], err = json.Marshal(a.BaseCurrency)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'baseCurrency': %w", err)
	}

	object["baseCurrencyId"], err = json.Marshal(a.BaseCurrencyId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'baseCurrencyId': %w", err)
	}

	object["baseCurrencyPrecision"], err = json.Marshal(a.BaseCurrencyPrecision)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'baseCurrencyPrecision': %w", err)
	}

	object["depthMergePrecision"], err = json.Marshal(a.DepthMergePrecision)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'depthMergePrecision': %w", err)
	}

	object["displayLevel"], err = json.Marshal(a.DisplayLevel)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'displayLevel': %w", err)
	}

	object["displayWeight"], err = json.Marshal(a.DisplayWeight)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'displayWeight': %w", err)
	}

	object["filters"], err = json.Marshal(a.Filters)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'filters': %w", err)
	}

	object["id"], err = json.Marshal(a.Id)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	if a.NextState != nil {
		object["nextState"], err = json.Marshal(a.NextState)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'nextState': %w", err)
		}
	}

	if a.NextStateTime != nil {
		object["nextStateTime"], err = json.Marshal(a.NextStateTime)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'nextStateTime': %w", err)
		}
	}

	object["orderTypes"], err = json.Marshal(a.OrderTypes)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'orderTypes': %w", err)
	}

	object["plates"], err = json.Marshal(a.Plates)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'plates': %w", err)
	}

	object["pricePrecision"], err = json.Marshal(a.PricePrecision)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'pricePrecision': %w", err)
	}

	object["quantityPrecision"], err = json.Marshal(a.QuantityPrecision)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quantityPrecision': %w", err)
	}

	object["quoteCurrency"], err = json.Marshal(a.QuoteCurrency)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quoteCurrency': %w", err)
	}

	object["quoteCurrencyId"], err = json.Marshal(a.QuoteCurrencyId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quoteCurrencyId': %w", err)
	}

	object["quoteCurrencyPrecision"], err = json.Marshal(a.QuoteCurrencyPrecision)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quoteCurrencyPrecision': %w", err)
	}

	object["state"], err = json.Marshal(a.State)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'state': %w", err)
	}

	object["stateTime"], err = json.Marshal(a.StateTime)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'stateTime': %w", err)
	}

	object["symbol"], err = json.Marshal(a.Symbol)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'symbol': %w", err)
	}

	object["timeInForces"], err = json.Marshal(a.TimeInForces)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'timeInForces': %w", err)
	}

	object["tradingEnabled"], err = json.Marshal(a.TradingEnabled)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'tradingEnabled': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetAllSymbols request
	GetAllSymbols(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAllSymbols(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllSymbolsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAllSymbolsRequest generates requests for GetAllSymbols
func NewGetAllSymbolsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/public/symbol")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetAllSymbols request
	GetAllSymbolsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllSymbolsResponse, error)
}

type GetAllSymbolsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetSymbolsResponse
}

// Status returns HTTPResponse.Status
func (r GetAllSymbolsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllSymbolsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAllSymbolsWithResponse request returning *GetAllSymbolsResponse
func (c *ClientWithResponses) GetAllSymbolsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllSymbolsResponse, error) {
	rsp, err := c.GetAllSymbols(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAllSymbolsResponse(rsp)
}

// ParseGetAllSymbolsResponse parses an HTTP response from a GetAllSymbolsWithResponse call
func ParseGetAllSymbolsResponse(rsp *http.Response) (*GetAllSymbolsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllSymbolsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetSymbolsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
