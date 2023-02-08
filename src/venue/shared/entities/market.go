// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/omiga-group/omiga/src/venue/shared/entities/market"
	"github.com/omiga-group/omiga/src/venue/shared/entities/venue"
)

// Market is the model entity for the Market schema.
type Market struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type market.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MarketQuery when eager-loading is set.
	Edges        MarketEdges `json:"edges"`
	venue_market *int
}

// MarketEdges holds the relations/edges for other nodes in the graph.
type MarketEdges struct {
	// Venue holds the value of the venue edge.
	Venue *Venue `json:"venue,omitempty"`
	// TradingPair holds the value of the trading_pair edge.
	TradingPair []*TradingPair `json:"trading_pair,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedTradingPair map[string][]*TradingPair
}

// VenueOrErr returns the Venue value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MarketEdges) VenueOrErr() (*Venue, error) {
	if e.loadedTypes[0] {
		if e.Venue == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: venue.Label}
		}
		return e.Venue, nil
	}
	return nil, &NotLoadedError{edge: "venue"}
}

// TradingPairOrErr returns the TradingPair value or an error if the edge
// was not loaded in eager-loading.
func (e MarketEdges) TradingPairOrErr() ([]*TradingPair, error) {
	if e.loadedTypes[1] {
		return e.TradingPair, nil
	}
	return nil, &NotLoadedError{edge: "trading_pair"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Market) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case market.FieldID:
			values[i] = new(sql.NullInt64)
		case market.FieldName, market.FieldType:
			values[i] = new(sql.NullString)
		case market.ForeignKeys[0]: // venue_market
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Market", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Market fields.
func (m *Market) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case market.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case market.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case market.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = market.Type(value.String)
			}
		case market.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field venue_market", value)
			} else if value.Valid {
				m.venue_market = new(int)
				*m.venue_market = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryVenue queries the "venue" edge of the Market entity.
func (m *Market) QueryVenue() *VenueQuery {
	return NewMarketClient(m.config).QueryVenue(m)
}

// QueryTradingPair queries the "trading_pair" edge of the Market entity.
func (m *Market) QueryTradingPair() *TradingPairQuery {
	return NewMarketClient(m.config).QueryTradingPair(m)
}

// Update returns a builder for updating this Market.
// Note that you need to call Market.Unwrap() before calling this method if this Market
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Market) Update() *MarketUpdateOne {
	return NewMarketClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Market entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Market) Unwrap() *Market {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("entities: Market is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Market) String() string {
	var builder strings.Builder
	builder.WriteString("Market(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", m.Type))
	builder.WriteByte(')')
	return builder.String()
}

// NamedTradingPair returns the TradingPair named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Market) NamedTradingPair(name string) ([]*TradingPair, error) {
	if m.Edges.namedTradingPair == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedTradingPair[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Market) appendNamedTradingPair(name string, edges ...*TradingPair) {
	if m.Edges.namedTradingPair == nil {
		m.Edges.namedTradingPair = make(map[string][]*TradingPair)
	}
	if len(edges) == 0 {
		m.Edges.namedTradingPair[name] = []*TradingPair{}
	} else {
		m.Edges.namedTradingPair[name] = append(m.Edges.namedTradingPair[name], edges...)
	}
}

// Markets is a parsable slice of Market.
type Markets []*Market

func (m Markets) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
