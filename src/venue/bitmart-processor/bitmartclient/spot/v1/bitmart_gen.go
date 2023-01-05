// Package bitmartspotv1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package bitmartspotv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// GetSymbolsDetailsData defines model for getSymbolsDetailsData.
type GetSymbolsDetailsData struct {
	Symbols              []Symbol               `json:"symbols"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// GetSymbolsDetailsResponse defines model for getSymbolsDetailsResponse.
type GetSymbolsDetailsResponse struct {
	Code                 int                    `json:"code"`
	Data                 GetSymbolsDetailsData  `json:"data"`
	Message              string                 `json:"message"`
	Trace                string                 `json:"trace"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Symbol defines model for symbol.
type Symbol struct {
	BaseCurrency         string                 `json:"base_currency"`
	BaseMinSize          string                 `json:"base_min_size"`
	Expiration           string                 `json:"expiration"`
	MinBuyAmount         string                 `json:"min_buy_amount"`
	MinSellAmount        string                 `json:"min_sell_amount"`
	PriceMaxPrecision    int                    `json:"price_max_precision"`
	PriceMinPrecision    int                    `json:"price_min_precision"`
	QuoteCurrency        string                 `json:"quote_currency"`
	QuoteIncrement       string                 `json:"quote_increment"`
	Symbol               string                 `json:"symbol"`
	SymbolId             int                    `json:"symbol_id"`
	TradeStatus          string                 `json:"trade_status"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Getter for additional properties for GetSymbolsDetailsData. Returns the specified
// element and whether it was found
func (a GetSymbolsDetailsData) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GetSymbolsDetailsData
func (a *GetSymbolsDetailsData) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GetSymbolsDetailsData to handle AdditionalProperties
func (a *GetSymbolsDetailsData) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for GetSymbolsDetailsData to handle AdditionalProperties
func (a GetSymbolsDetailsData) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["symbols"], err = json.Marshal(a.Symbols)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'symbols': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for GetSymbolsDetailsResponse. Returns the specified
// element and whether it was found
func (a GetSymbolsDetailsResponse) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GetSymbolsDetailsResponse
func (a *GetSymbolsDetailsResponse) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GetSymbolsDetailsResponse to handle AdditionalProperties
func (a *GetSymbolsDetailsResponse) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["code"]; found {
		err = json.Unmarshal(raw, &a.Code)
		if err != nil {
			return fmt.Errorf("error reading 'code': %w", err)
		}
		delete(object, "code")
	}

	if raw, found := object["data"]; found {
		err = json.Unmarshal(raw, &a.Data)
		if err != nil {
			return fmt.Errorf("error reading 'data': %w", err)
		}
		delete(object, "data")
	}

	if raw, found := object["message"]; found {
		err = json.Unmarshal(raw, &a.Message)
		if err != nil {
			return fmt.Errorf("error reading 'message': %w", err)
		}
		delete(object, "message")
	}

	if raw, found := object["trace"]; found {
		err = json.Unmarshal(raw, &a.Trace)
		if err != nil {
			return fmt.Errorf("error reading 'trace': %w", err)
		}
		delete(object, "trace")
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

// Override default JSON handling for GetSymbolsDetailsResponse to handle AdditionalProperties
func (a GetSymbolsDetailsResponse) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["code"], err = json.Marshal(a.Code)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'code': %w", err)
	}

	object["data"], err = json.Marshal(a.Data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'data': %w", err)
	}

	object["message"], err = json.Marshal(a.Message)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'message': %w", err)
	}

	object["trace"], err = json.Marshal(a.Trace)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'trace': %w", err)
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

	if raw, found := object["base_currency"]; found {
		err = json.Unmarshal(raw, &a.BaseCurrency)
		if err != nil {
			return fmt.Errorf("error reading 'base_currency': %w", err)
		}
		delete(object, "base_currency")
	}

	if raw, found := object["base_min_size"]; found {
		err = json.Unmarshal(raw, &a.BaseMinSize)
		if err != nil {
			return fmt.Errorf("error reading 'base_min_size': %w", err)
		}
		delete(object, "base_min_size")
	}

	if raw, found := object["expiration"]; found {
		err = json.Unmarshal(raw, &a.Expiration)
		if err != nil {
			return fmt.Errorf("error reading 'expiration': %w", err)
		}
		delete(object, "expiration")
	}

	if raw, found := object["min_buy_amount"]; found {
		err = json.Unmarshal(raw, &a.MinBuyAmount)
		if err != nil {
			return fmt.Errorf("error reading 'min_buy_amount': %w", err)
		}
		delete(object, "min_buy_amount")
	}

	if raw, found := object["min_sell_amount"]; found {
		err = json.Unmarshal(raw, &a.MinSellAmount)
		if err != nil {
			return fmt.Errorf("error reading 'min_sell_amount': %w", err)
		}
		delete(object, "min_sell_amount")
	}

	if raw, found := object["price_max_precision"]; found {
		err = json.Unmarshal(raw, &a.PriceMaxPrecision)
		if err != nil {
			return fmt.Errorf("error reading 'price_max_precision': %w", err)
		}
		delete(object, "price_max_precision")
	}

	if raw, found := object["price_min_precision"]; found {
		err = json.Unmarshal(raw, &a.PriceMinPrecision)
		if err != nil {
			return fmt.Errorf("error reading 'price_min_precision': %w", err)
		}
		delete(object, "price_min_precision")
	}

	if raw, found := object["quote_currency"]; found {
		err = json.Unmarshal(raw, &a.QuoteCurrency)
		if err != nil {
			return fmt.Errorf("error reading 'quote_currency': %w", err)
		}
		delete(object, "quote_currency")
	}

	if raw, found := object["quote_increment"]; found {
		err = json.Unmarshal(raw, &a.QuoteIncrement)
		if err != nil {
			return fmt.Errorf("error reading 'quote_increment': %w", err)
		}
		delete(object, "quote_increment")
	}

	if raw, found := object["symbol"]; found {
		err = json.Unmarshal(raw, &a.Symbol)
		if err != nil {
			return fmt.Errorf("error reading 'symbol': %w", err)
		}
		delete(object, "symbol")
	}

	if raw, found := object["symbol_id"]; found {
		err = json.Unmarshal(raw, &a.SymbolId)
		if err != nil {
			return fmt.Errorf("error reading 'symbol_id': %w", err)
		}
		delete(object, "symbol_id")
	}

	if raw, found := object["trade_status"]; found {
		err = json.Unmarshal(raw, &a.TradeStatus)
		if err != nil {
			return fmt.Errorf("error reading 'trade_status': %w", err)
		}
		delete(object, "trade_status")
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

	object["base_currency"], err = json.Marshal(a.BaseCurrency)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'base_currency': %w", err)
	}

	object["base_min_size"], err = json.Marshal(a.BaseMinSize)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'base_min_size': %w", err)
	}

	object["expiration"], err = json.Marshal(a.Expiration)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'expiration': %w", err)
	}

	object["min_buy_amount"], err = json.Marshal(a.MinBuyAmount)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'min_buy_amount': %w", err)
	}

	object["min_sell_amount"], err = json.Marshal(a.MinSellAmount)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'min_sell_amount': %w", err)
	}

	object["price_max_precision"], err = json.Marshal(a.PriceMaxPrecision)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'price_max_precision': %w", err)
	}

	object["price_min_precision"], err = json.Marshal(a.PriceMinPrecision)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'price_min_precision': %w", err)
	}

	object["quote_currency"], err = json.Marshal(a.QuoteCurrency)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quote_currency': %w", err)
	}

	object["quote_increment"], err = json.Marshal(a.QuoteIncrement)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quote_increment': %w", err)
	}

	object["symbol"], err = json.Marshal(a.Symbol)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'symbol': %w", err)
	}

	object["symbol_id"], err = json.Marshal(a.SymbolId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'symbol_id': %w", err)
	}

	object["trade_status"], err = json.Marshal(a.TradeStatus)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'trade_status': %w", err)
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
	// GetAllSymbolsDetails request
	GetAllSymbolsDetails(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAllSymbolsDetails(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllSymbolsDetailsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAllSymbolsDetailsRequest generates requests for GetAllSymbolsDetails
func NewGetAllSymbolsDetailsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/symbols/details")
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
	// GetAllSymbolsDetails request
	GetAllSymbolsDetailsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllSymbolsDetailsResponse, error)
}

type GetAllSymbolsDetailsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetSymbolsDetailsResponse
}

// Status returns HTTPResponse.Status
func (r GetAllSymbolsDetailsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllSymbolsDetailsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAllSymbolsDetailsWithResponse request returning *GetAllSymbolsDetailsResponse
func (c *ClientWithResponses) GetAllSymbolsDetailsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllSymbolsDetailsResponse, error) {
	rsp, err := c.GetAllSymbolsDetails(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAllSymbolsDetailsResponse(rsp)
}

// ParseGetAllSymbolsDetailsResponse parses an HTTP response from a GetAllSymbolsDetailsWithResponse call
func ParseGetAllSymbolsDetailsResponse(rsp *http.Response) (*GetAllSymbolsDetailsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllSymbolsDetailsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetSymbolsDetailsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
