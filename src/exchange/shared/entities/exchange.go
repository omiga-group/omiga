// Code generated by ent, DO NOT EDIT.

package entities

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
)

// Exchange is the model entity for the Exchange schema.
type Exchange struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ExchangeID holds the value of the "exchange_id" field.
	ExchangeID string `json:"exchange_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// YearEstablished holds the value of the "year_established" field.
	YearEstablished int `json:"year_established,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Links holds the value of the "links" field.
	Links map[string]string `json:"links,omitempty"`
	// HasTradingIncentive holds the value of the "has_trading_incentive" field.
	HasTradingIncentive bool `json:"has_trading_incentive,omitempty"`
	// Centralized holds the value of the "centralized" field.
	Centralized bool `json:"centralized,omitempty"`
	// PublicNotice holds the value of the "public_notice" field.
	PublicNotice string `json:"public_notice,omitempty"`
	// AlertNotice holds the value of the "alert_notice" field.
	AlertNotice string `json:"alert_notice,omitempty"`
	// TrustScore holds the value of the "trust_score" field.
	TrustScore int `json:"trust_score,omitempty"`
	// TrustScoreRank holds the value of the "trust_score_rank" field.
	TrustScoreRank int `json:"trust_score_rank,omitempty"`
	// TradeVolume24hBtc holds the value of the "trade_volume_24h_btc" field.
	TradeVolume24hBtc float64 `json:"trade_volume_24h_btc,omitempty"`
	// TradeVolume24hBtcNormalized holds the value of the "trade_volume_24h_btc_normalized" field.
	TradeVolume24hBtcNormalized float64 `json:"trade_volume_24h_btc_normalized,omitempty"`
	// MakerFee holds the value of the "maker_fee" field.
	MakerFee float64 `json:"maker_fee,omitempty"`
	// TakerFee holds the value of the "taker_fee" field.
	TakerFee float64 `json:"taker_fee,omitempty"`
	// SpreadFee holds the value of the "spread_fee" field.
	SpreadFee bool `json:"spread_fee,omitempty"`
	// SupportAPI holds the value of the "support_api" field.
	SupportAPI bool `json:"support_api,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExchangeQuery when eager-loading is set.
	Edges ExchangeEdges `json:"edges"`
}

// ExchangeEdges holds the relations/edges for other nodes in the graph.
type ExchangeEdges struct {
	// Ticker holds the value of the ticker edge.
	Ticker []*Ticker `json:"ticker,omitempty"`
	// TradingPair holds the value of the trading_pair edge.
	TradingPair []*TradingPair `json:"trading_pair,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedTicker      map[string][]*Ticker
	namedTradingPair map[string][]*TradingPair
}

// TickerOrErr returns the Ticker value or an error if the edge
// was not loaded in eager-loading.
func (e ExchangeEdges) TickerOrErr() ([]*Ticker, error) {
	if e.loadedTypes[0] {
		return e.Ticker, nil
	}
	return nil, &NotLoadedError{edge: "ticker"}
}

// TradingPairOrErr returns the TradingPair value or an error if the edge
// was not loaded in eager-loading.
func (e ExchangeEdges) TradingPairOrErr() ([]*TradingPair, error) {
	if e.loadedTypes[1] {
		return e.TradingPair, nil
	}
	return nil, &NotLoadedError{edge: "trading_pair"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exchange) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exchange.FieldLinks:
			values[i] = new([]byte)
		case exchange.FieldHasTradingIncentive, exchange.FieldCentralized, exchange.FieldSpreadFee, exchange.FieldSupportAPI:
			values[i] = new(sql.NullBool)
		case exchange.FieldTradeVolume24hBtc, exchange.FieldTradeVolume24hBtcNormalized, exchange.FieldMakerFee, exchange.FieldTakerFee:
			values[i] = new(sql.NullFloat64)
		case exchange.FieldID, exchange.FieldYearEstablished, exchange.FieldTrustScore, exchange.FieldTrustScoreRank:
			values[i] = new(sql.NullInt64)
		case exchange.FieldExchangeID, exchange.FieldName, exchange.FieldCountry, exchange.FieldImage, exchange.FieldPublicNotice, exchange.FieldAlertNotice:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Exchange", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exchange fields.
func (e *Exchange) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exchange.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case exchange.FieldExchangeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field exchange_id", values[i])
			} else if value.Valid {
				e.ExchangeID = value.String
			}
		case exchange.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case exchange.FieldYearEstablished:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field year_established", values[i])
			} else if value.Valid {
				e.YearEstablished = int(value.Int64)
			}
		case exchange.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				e.Country = value.String
			}
		case exchange.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				e.Image = value.String
			}
		case exchange.FieldLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Links); err != nil {
					return fmt.Errorf("unmarshal field links: %w", err)
				}
			}
		case exchange.FieldHasTradingIncentive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_trading_incentive", values[i])
			} else if value.Valid {
				e.HasTradingIncentive = value.Bool
			}
		case exchange.FieldCentralized:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field centralized", values[i])
			} else if value.Valid {
				e.Centralized = value.Bool
			}
		case exchange.FieldPublicNotice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_notice", values[i])
			} else if value.Valid {
				e.PublicNotice = value.String
			}
		case exchange.FieldAlertNotice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alert_notice", values[i])
			} else if value.Valid {
				e.AlertNotice = value.String
			}
		case exchange.FieldTrustScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trust_score", values[i])
			} else if value.Valid {
				e.TrustScore = int(value.Int64)
			}
		case exchange.FieldTrustScoreRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trust_score_rank", values[i])
			} else if value.Valid {
				e.TrustScoreRank = int(value.Int64)
			}
		case exchange.FieldTradeVolume24hBtc:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field trade_volume_24h_btc", values[i])
			} else if value.Valid {
				e.TradeVolume24hBtc = value.Float64
			}
		case exchange.FieldTradeVolume24hBtcNormalized:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field trade_volume_24h_btc_normalized", values[i])
			} else if value.Valid {
				e.TradeVolume24hBtcNormalized = value.Float64
			}
		case exchange.FieldMakerFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field maker_fee", values[i])
			} else if value.Valid {
				e.MakerFee = value.Float64
			}
		case exchange.FieldTakerFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field taker_fee", values[i])
			} else if value.Valid {
				e.TakerFee = value.Float64
			}
		case exchange.FieldSpreadFee:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field spread_fee", values[i])
			} else if value.Valid {
				e.SpreadFee = value.Bool
			}
		case exchange.FieldSupportAPI:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field support_api", values[i])
			} else if value.Valid {
				e.SupportAPI = value.Bool
			}
		}
	}
	return nil
}

// QueryTicker queries the "ticker" edge of the Exchange entity.
func (e *Exchange) QueryTicker() *TickerQuery {
	return (&ExchangeClient{config: e.config}).QueryTicker(e)
}

// QueryTradingPair queries the "trading_pair" edge of the Exchange entity.
func (e *Exchange) QueryTradingPair() *TradingPairQuery {
	return (&ExchangeClient{config: e.config}).QueryTradingPair(e)
}

// Update returns a builder for updating this Exchange.
// Note that you need to call Exchange.Unwrap() before calling this method if this Exchange
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Exchange) Update() *ExchangeUpdateOne {
	return (&ExchangeClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Exchange entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Exchange) Unwrap() *Exchange {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("entities: Exchange is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Exchange) String() string {
	var builder strings.Builder
	builder.WriteString("Exchange(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("exchange_id=")
	builder.WriteString(e.ExchangeID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("year_established=")
	builder.WriteString(fmt.Sprintf("%v", e.YearEstablished))
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(e.Country)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(e.Image)
	builder.WriteString(", ")
	builder.WriteString("links=")
	builder.WriteString(fmt.Sprintf("%v", e.Links))
	builder.WriteString(", ")
	builder.WriteString("has_trading_incentive=")
	builder.WriteString(fmt.Sprintf("%v", e.HasTradingIncentive))
	builder.WriteString(", ")
	builder.WriteString("centralized=")
	builder.WriteString(fmt.Sprintf("%v", e.Centralized))
	builder.WriteString(", ")
	builder.WriteString("public_notice=")
	builder.WriteString(e.PublicNotice)
	builder.WriteString(", ")
	builder.WriteString("alert_notice=")
	builder.WriteString(e.AlertNotice)
	builder.WriteString(", ")
	builder.WriteString("trust_score=")
	builder.WriteString(fmt.Sprintf("%v", e.TrustScore))
	builder.WriteString(", ")
	builder.WriteString("trust_score_rank=")
	builder.WriteString(fmt.Sprintf("%v", e.TrustScoreRank))
	builder.WriteString(", ")
	builder.WriteString("trade_volume_24h_btc=")
	builder.WriteString(fmt.Sprintf("%v", e.TradeVolume24hBtc))
	builder.WriteString(", ")
	builder.WriteString("trade_volume_24h_btc_normalized=")
	builder.WriteString(fmt.Sprintf("%v", e.TradeVolume24hBtcNormalized))
	builder.WriteString(", ")
	builder.WriteString("maker_fee=")
	builder.WriteString(fmt.Sprintf("%v", e.MakerFee))
	builder.WriteString(", ")
	builder.WriteString("taker_fee=")
	builder.WriteString(fmt.Sprintf("%v", e.TakerFee))
	builder.WriteString(", ")
	builder.WriteString("spread_fee=")
	builder.WriteString(fmt.Sprintf("%v", e.SpreadFee))
	builder.WriteString(", ")
	builder.WriteString("support_api=")
	builder.WriteString(fmt.Sprintf("%v", e.SupportAPI))
	builder.WriteByte(')')
	return builder.String()
}

// NamedTicker returns the Ticker named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exchange) NamedTicker(name string) ([]*Ticker, error) {
	if e.Edges.namedTicker == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedTicker[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exchange) appendNamedTicker(name string, edges ...*Ticker) {
	if e.Edges.namedTicker == nil {
		e.Edges.namedTicker = make(map[string][]*Ticker)
	}
	if len(edges) == 0 {
		e.Edges.namedTicker[name] = []*Ticker{}
	} else {
		e.Edges.namedTicker[name] = append(e.Edges.namedTicker[name], edges...)
	}
}

// NamedTradingPair returns the TradingPair named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exchange) NamedTradingPair(name string) ([]*TradingPair, error) {
	if e.Edges.namedTradingPair == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedTradingPair[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exchange) appendNamedTradingPair(name string, edges ...*TradingPair) {
	if e.Edges.namedTradingPair == nil {
		e.Edges.namedTradingPair = make(map[string][]*TradingPair)
	}
	if len(edges) == 0 {
		e.Edges.namedTradingPair[name] = []*TradingPair{}
	} else {
		e.Edges.namedTradingPair[name] = append(e.Edges.namedTradingPair[name], edges...)
	}
}

// Exchanges is a parsable slice of Exchange.
type Exchanges []*Exchange

func (e Exchanges) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
