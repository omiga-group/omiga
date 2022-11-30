// Package cryptov2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package cryptov2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// GetInstrumentsResponse defines model for getInstrumentsResponse.
type GetInstrumentsResponse struct {
	Code                 int                    `json:"code"`
	Id                   int                    `json:"id"`
	Method               string                 `json:"method"`
	Result               GetInstrumentsResult   `json:"result"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// GetInstrumentsResult defines model for getInstrumentsResult.
type GetInstrumentsResult struct {
	Instruments          []Instrument           `json:"instruments"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Instrument defines model for instrument.
type Instrument struct {
	BaseCurrency            string                 `json:"base_currency"`
	InstrumentName          string                 `json:"instrument_name"`
	LastUpdateDate          int                    `json:"last_update_date"`
	MarginTradingEnabled    bool                   `json:"margin_trading_enabled"`
	MarginTradingEnabled10x bool                   `json:"margin_trading_enabled_10x"`
	MarginTradingEnabled5x  bool                   `json:"margin_trading_enabled_5x"`
	MaxPrice                string                 `json:"max_price"`
	MaxQuantity             string                 `json:"max_quantity"`
	MinPrice                string                 `json:"min_price"`
	MinQuantity             string                 `json:"min_quantity"`
	PriceDecimals           int                    `json:"price_decimals"`
	QuantityDecimals        int                    `json:"quantity_decimals"`
	QuoteCurrency           string                 `json:"quote_currency"`
	AdditionalProperties    map[string]interface{} `json:"-"`
}

// Getter for additional properties for GetInstrumentsResponse. Returns the specified
// element and whether it was found
func (a GetInstrumentsResponse) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GetInstrumentsResponse
func (a *GetInstrumentsResponse) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GetInstrumentsResponse to handle AdditionalProperties
func (a *GetInstrumentsResponse) UnmarshalJSON(b []byte) error {
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

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.Id)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["method"]; found {
		err = json.Unmarshal(raw, &a.Method)
		if err != nil {
			return fmt.Errorf("error reading 'method': %w", err)
		}
		delete(object, "method")
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

// Override default JSON handling for GetInstrumentsResponse to handle AdditionalProperties
func (a GetInstrumentsResponse) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["code"], err = json.Marshal(a.Code)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'code': %w", err)
	}

	object["id"], err = json.Marshal(a.Id)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	object["method"], err = json.Marshal(a.Method)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'method': %w", err)
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

// Getter for additional properties for GetInstrumentsResult. Returns the specified
// element and whether it was found
func (a GetInstrumentsResult) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for GetInstrumentsResult
func (a *GetInstrumentsResult) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for GetInstrumentsResult to handle AdditionalProperties
func (a *GetInstrumentsResult) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["instruments"]; found {
		err = json.Unmarshal(raw, &a.Instruments)
		if err != nil {
			return fmt.Errorf("error reading 'instruments': %w", err)
		}
		delete(object, "instruments")
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

// Override default JSON handling for GetInstrumentsResult to handle AdditionalProperties
func (a GetInstrumentsResult) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["instruments"], err = json.Marshal(a.Instruments)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'instruments': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for Instrument. Returns the specified
// element and whether it was found
func (a Instrument) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Instrument
func (a *Instrument) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Instrument to handle AdditionalProperties
func (a *Instrument) UnmarshalJSON(b []byte) error {
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

	if raw, found := object["instrument_name"]; found {
		err = json.Unmarshal(raw, &a.InstrumentName)
		if err != nil {
			return fmt.Errorf("error reading 'instrument_name': %w", err)
		}
		delete(object, "instrument_name")
	}

	if raw, found := object["last_update_date"]; found {
		err = json.Unmarshal(raw, &a.LastUpdateDate)
		if err != nil {
			return fmt.Errorf("error reading 'last_update_date': %w", err)
		}
		delete(object, "last_update_date")
	}

	if raw, found := object["margin_trading_enabled"]; found {
		err = json.Unmarshal(raw, &a.MarginTradingEnabled)
		if err != nil {
			return fmt.Errorf("error reading 'margin_trading_enabled': %w", err)
		}
		delete(object, "margin_trading_enabled")
	}

	if raw, found := object["margin_trading_enabled_10x"]; found {
		err = json.Unmarshal(raw, &a.MarginTradingEnabled10x)
		if err != nil {
			return fmt.Errorf("error reading 'margin_trading_enabled_10x': %w", err)
		}
		delete(object, "margin_trading_enabled_10x")
	}

	if raw, found := object["margin_trading_enabled_5x"]; found {
		err = json.Unmarshal(raw, &a.MarginTradingEnabled5x)
		if err != nil {
			return fmt.Errorf("error reading 'margin_trading_enabled_5x': %w", err)
		}
		delete(object, "margin_trading_enabled_5x")
	}

	if raw, found := object["max_price"]; found {
		err = json.Unmarshal(raw, &a.MaxPrice)
		if err != nil {
			return fmt.Errorf("error reading 'max_price': %w", err)
		}
		delete(object, "max_price")
	}

	if raw, found := object["max_quantity"]; found {
		err = json.Unmarshal(raw, &a.MaxQuantity)
		if err != nil {
			return fmt.Errorf("error reading 'max_quantity': %w", err)
		}
		delete(object, "max_quantity")
	}

	if raw, found := object["min_price"]; found {
		err = json.Unmarshal(raw, &a.MinPrice)
		if err != nil {
			return fmt.Errorf("error reading 'min_price': %w", err)
		}
		delete(object, "min_price")
	}

	if raw, found := object["min_quantity"]; found {
		err = json.Unmarshal(raw, &a.MinQuantity)
		if err != nil {
			return fmt.Errorf("error reading 'min_quantity': %w", err)
		}
		delete(object, "min_quantity")
	}

	if raw, found := object["price_decimals"]; found {
		err = json.Unmarshal(raw, &a.PriceDecimals)
		if err != nil {
			return fmt.Errorf("error reading 'price_decimals': %w", err)
		}
		delete(object, "price_decimals")
	}

	if raw, found := object["quantity_decimals"]; found {
		err = json.Unmarshal(raw, &a.QuantityDecimals)
		if err != nil {
			return fmt.Errorf("error reading 'quantity_decimals': %w", err)
		}
		delete(object, "quantity_decimals")
	}

	if raw, found := object["quote_currency"]; found {
		err = json.Unmarshal(raw, &a.QuoteCurrency)
		if err != nil {
			return fmt.Errorf("error reading 'quote_currency': %w", err)
		}
		delete(object, "quote_currency")
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

// Override default JSON handling for Instrument to handle AdditionalProperties
func (a Instrument) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["base_currency"], err = json.Marshal(a.BaseCurrency)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'base_currency': %w", err)
	}

	object["instrument_name"], err = json.Marshal(a.InstrumentName)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'instrument_name': %w", err)
	}

	object["last_update_date"], err = json.Marshal(a.LastUpdateDate)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'last_update_date': %w", err)
	}

	object["margin_trading_enabled"], err = json.Marshal(a.MarginTradingEnabled)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'margin_trading_enabled': %w", err)
	}

	object["margin_trading_enabled_10x"], err = json.Marshal(a.MarginTradingEnabled10x)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'margin_trading_enabled_10x': %w", err)
	}

	object["margin_trading_enabled_5x"], err = json.Marshal(a.MarginTradingEnabled5x)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'margin_trading_enabled_5x': %w", err)
	}

	object["max_price"], err = json.Marshal(a.MaxPrice)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'max_price': %w", err)
	}

	object["max_quantity"], err = json.Marshal(a.MaxQuantity)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'max_quantity': %w", err)
	}

	object["min_price"], err = json.Marshal(a.MinPrice)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'min_price': %w", err)
	}

	object["min_quantity"], err = json.Marshal(a.MinQuantity)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'min_quantity': %w", err)
	}

	object["price_decimals"], err = json.Marshal(a.PriceDecimals)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'price_decimals': %w", err)
	}

	object["quantity_decimals"], err = json.Marshal(a.QuantityDecimals)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quantity_decimals': %w", err)
	}

	object["quote_currency"], err = json.Marshal(a.QuoteCurrency)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'quote_currency': %w", err)
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
	// GetAllInstruments request
	GetAllInstruments(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAllInstruments(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllInstrumentsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAllInstrumentsRequest generates requests for GetAllInstruments
func NewGetAllInstrumentsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/public/get-instruments")
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
	// GetAllInstruments request
	GetAllInstrumentsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllInstrumentsResponse, error)
}

type GetAllInstrumentsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetInstrumentsResponse
}

// Status returns HTTPResponse.Status
func (r GetAllInstrumentsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllInstrumentsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAllInstrumentsWithResponse request returning *GetAllInstrumentsResponse
func (c *ClientWithResponses) GetAllInstrumentsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllInstrumentsResponse, error) {
	rsp, err := c.GetAllInstruments(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAllInstrumentsResponse(rsp)
}

// ParseGetAllInstrumentsResponse parses an HTTP response from a GetAllInstrumentsWithResponse call
func ParseGetAllInstrumentsResponse(rsp *http.Response) (*GetAllInstrumentsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllInstrumentsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetInstrumentsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
