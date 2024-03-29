// Code generated by ent, DO NOT EDIT.

package entities

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/omiga-group/omiga/src/venue/shared/entities/venue"
)

// Venue is the model entity for the Venue schema.
type Venue struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// VenueID holds the value of the "venue_id" field.
	VenueID string `json:"venue_id,omitempty"`
	// Type holds the value of the "type" field.
	Type venue.Type `json:"type,omitempty"`
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
	// The values are being populated by the VenueQuery when eager-loading is set.
	Edges VenueEdges `json:"edges"`
}

// VenueEdges holds the relations/edges for other nodes in the graph.
type VenueEdges struct {
	// Ticker holds the value of the ticker edge.
	Ticker []*Ticker `json:"ticker,omitempty"`
	// TradingPair holds the value of the trading_pair edge.
	TradingPair []*TradingPair `json:"trading_pair,omitempty"`
	// Market holds the value of the market edge.
	Market []*Market `json:"market,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedTicker      map[string][]*Ticker
	namedTradingPair map[string][]*TradingPair
	namedMarket      map[string][]*Market
}

// TickerOrErr returns the Ticker value or an error if the edge
// was not loaded in eager-loading.
func (e VenueEdges) TickerOrErr() ([]*Ticker, error) {
	if e.loadedTypes[0] {
		return e.Ticker, nil
	}
	return nil, &NotLoadedError{edge: "ticker"}
}

// TradingPairOrErr returns the TradingPair value or an error if the edge
// was not loaded in eager-loading.
func (e VenueEdges) TradingPairOrErr() ([]*TradingPair, error) {
	if e.loadedTypes[1] {
		return e.TradingPair, nil
	}
	return nil, &NotLoadedError{edge: "trading_pair"}
}

// MarketOrErr returns the Market value or an error if the edge
// was not loaded in eager-loading.
func (e VenueEdges) MarketOrErr() ([]*Market, error) {
	if e.loadedTypes[2] {
		return e.Market, nil
	}
	return nil, &NotLoadedError{edge: "market"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Venue) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case venue.FieldLinks:
			values[i] = new([]byte)
		case venue.FieldHasTradingIncentive, venue.FieldCentralized, venue.FieldSpreadFee, venue.FieldSupportAPI:
			values[i] = new(sql.NullBool)
		case venue.FieldTradeVolume24hBtc, venue.FieldTradeVolume24hBtcNormalized, venue.FieldMakerFee, venue.FieldTakerFee:
			values[i] = new(sql.NullFloat64)
		case venue.FieldID, venue.FieldYearEstablished, venue.FieldTrustScore, venue.FieldTrustScoreRank:
			values[i] = new(sql.NullInt64)
		case venue.FieldVenueID, venue.FieldType, venue.FieldName, venue.FieldCountry, venue.FieldImage, venue.FieldPublicNotice, venue.FieldAlertNotice:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Venue", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Venue fields.
func (v *Venue) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case venue.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case venue.FieldVenueID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				v.VenueID = value.String
			}
		case venue.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				v.Type = venue.Type(value.String)
			}
		case venue.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				v.Name = value.String
			}
		case venue.FieldYearEstablished:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field year_established", values[i])
			} else if value.Valid {
				v.YearEstablished = int(value.Int64)
			}
		case venue.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				v.Country = value.String
			}
		case venue.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				v.Image = value.String
			}
		case venue.FieldLinks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field links", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &v.Links); err != nil {
					return fmt.Errorf("unmarshal field links: %w", err)
				}
			}
		case venue.FieldHasTradingIncentive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_trading_incentive", values[i])
			} else if value.Valid {
				v.HasTradingIncentive = value.Bool
			}
		case venue.FieldCentralized:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field centralized", values[i])
			} else if value.Valid {
				v.Centralized = value.Bool
			}
		case venue.FieldPublicNotice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_notice", values[i])
			} else if value.Valid {
				v.PublicNotice = value.String
			}
		case venue.FieldAlertNotice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alert_notice", values[i])
			} else if value.Valid {
				v.AlertNotice = value.String
			}
		case venue.FieldTrustScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trust_score", values[i])
			} else if value.Valid {
				v.TrustScore = int(value.Int64)
			}
		case venue.FieldTrustScoreRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trust_score_rank", values[i])
			} else if value.Valid {
				v.TrustScoreRank = int(value.Int64)
			}
		case venue.FieldTradeVolume24hBtc:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field trade_volume_24h_btc", values[i])
			} else if value.Valid {
				v.TradeVolume24hBtc = value.Float64
			}
		case venue.FieldTradeVolume24hBtcNormalized:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field trade_volume_24h_btc_normalized", values[i])
			} else if value.Valid {
				v.TradeVolume24hBtcNormalized = value.Float64
			}
		case venue.FieldMakerFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field maker_fee", values[i])
			} else if value.Valid {
				v.MakerFee = value.Float64
			}
		case venue.FieldTakerFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field taker_fee", values[i])
			} else if value.Valid {
				v.TakerFee = value.Float64
			}
		case venue.FieldSpreadFee:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field spread_fee", values[i])
			} else if value.Valid {
				v.SpreadFee = value.Bool
			}
		case venue.FieldSupportAPI:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field support_api", values[i])
			} else if value.Valid {
				v.SupportAPI = value.Bool
			}
		}
	}
	return nil
}

// QueryTicker queries the "ticker" edge of the Venue entity.
func (v *Venue) QueryTicker() *TickerQuery {
	return NewVenueClient(v.config).QueryTicker(v)
}

// QueryTradingPair queries the "trading_pair" edge of the Venue entity.
func (v *Venue) QueryTradingPair() *TradingPairQuery {
	return NewVenueClient(v.config).QueryTradingPair(v)
}

// QueryMarket queries the "market" edge of the Venue entity.
func (v *Venue) QueryMarket() *MarketQuery {
	return NewVenueClient(v.config).QueryMarket(v)
}

// Update returns a builder for updating this Venue.
// Note that you need to call Venue.Unwrap() before calling this method if this Venue
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Venue) Update() *VenueUpdateOne {
	return NewVenueClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Venue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Venue) Unwrap() *Venue {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("entities: Venue is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Venue) String() string {
	var builder strings.Builder
	builder.WriteString("Venue(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("venue_id=")
	builder.WriteString(v.VenueID)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", v.Type))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(v.Name)
	builder.WriteString(", ")
	builder.WriteString("year_established=")
	builder.WriteString(fmt.Sprintf("%v", v.YearEstablished))
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(v.Country)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(v.Image)
	builder.WriteString(", ")
	builder.WriteString("links=")
	builder.WriteString(fmt.Sprintf("%v", v.Links))
	builder.WriteString(", ")
	builder.WriteString("has_trading_incentive=")
	builder.WriteString(fmt.Sprintf("%v", v.HasTradingIncentive))
	builder.WriteString(", ")
	builder.WriteString("centralized=")
	builder.WriteString(fmt.Sprintf("%v", v.Centralized))
	builder.WriteString(", ")
	builder.WriteString("public_notice=")
	builder.WriteString(v.PublicNotice)
	builder.WriteString(", ")
	builder.WriteString("alert_notice=")
	builder.WriteString(v.AlertNotice)
	builder.WriteString(", ")
	builder.WriteString("trust_score=")
	builder.WriteString(fmt.Sprintf("%v", v.TrustScore))
	builder.WriteString(", ")
	builder.WriteString("trust_score_rank=")
	builder.WriteString(fmt.Sprintf("%v", v.TrustScoreRank))
	builder.WriteString(", ")
	builder.WriteString("trade_volume_24h_btc=")
	builder.WriteString(fmt.Sprintf("%v", v.TradeVolume24hBtc))
	builder.WriteString(", ")
	builder.WriteString("trade_volume_24h_btc_normalized=")
	builder.WriteString(fmt.Sprintf("%v", v.TradeVolume24hBtcNormalized))
	builder.WriteString(", ")
	builder.WriteString("maker_fee=")
	builder.WriteString(fmt.Sprintf("%v", v.MakerFee))
	builder.WriteString(", ")
	builder.WriteString("taker_fee=")
	builder.WriteString(fmt.Sprintf("%v", v.TakerFee))
	builder.WriteString(", ")
	builder.WriteString("spread_fee=")
	builder.WriteString(fmt.Sprintf("%v", v.SpreadFee))
	builder.WriteString(", ")
	builder.WriteString("support_api=")
	builder.WriteString(fmt.Sprintf("%v", v.SupportAPI))
	builder.WriteByte(')')
	return builder.String()
}

// NamedTicker returns the Ticker named value or an error if the edge was not
// loaded in eager-loading with this name.
func (v *Venue) NamedTicker(name string) ([]*Ticker, error) {
	if v.Edges.namedTicker == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := v.Edges.namedTicker[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (v *Venue) appendNamedTicker(name string, edges ...*Ticker) {
	if v.Edges.namedTicker == nil {
		v.Edges.namedTicker = make(map[string][]*Ticker)
	}
	if len(edges) == 0 {
		v.Edges.namedTicker[name] = []*Ticker{}
	} else {
		v.Edges.namedTicker[name] = append(v.Edges.namedTicker[name], edges...)
	}
}

// NamedTradingPair returns the TradingPair named value or an error if the edge was not
// loaded in eager-loading with this name.
func (v *Venue) NamedTradingPair(name string) ([]*TradingPair, error) {
	if v.Edges.namedTradingPair == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := v.Edges.namedTradingPair[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (v *Venue) appendNamedTradingPair(name string, edges ...*TradingPair) {
	if v.Edges.namedTradingPair == nil {
		v.Edges.namedTradingPair = make(map[string][]*TradingPair)
	}
	if len(edges) == 0 {
		v.Edges.namedTradingPair[name] = []*TradingPair{}
	} else {
		v.Edges.namedTradingPair[name] = append(v.Edges.namedTradingPair[name], edges...)
	}
}

// NamedMarket returns the Market named value or an error if the edge was not
// loaded in eager-loading with this name.
func (v *Venue) NamedMarket(name string) ([]*Market, error) {
	if v.Edges.namedMarket == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := v.Edges.namedMarket[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (v *Venue) appendNamedMarket(name string, edges ...*Market) {
	if v.Edges.namedMarket == nil {
		v.Edges.namedMarket = make(map[string][]*Market)
	}
	if len(edges) == 0 {
		v.Edges.namedMarket[name] = []*Market{}
	} else {
		v.Edges.namedMarket[name] = append(v.Edges.namedMarket[name], edges...)
	}
}

// Venues is a parsable slice of Venue.
type Venues []*Venue

func (v Venues) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
