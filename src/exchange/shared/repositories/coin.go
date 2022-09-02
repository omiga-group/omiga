// Code generated by ent, DO NOT EDIT.

package repositories

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/coin"
)

// Coin is the model entity for the Coin schema.
type Coin struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Symbol holds the value of the "symbol" field.
	Symbol string `json:"symbol,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Coin) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coin.FieldID:
			values[i] = new(sql.NullInt64)
		case coin.FieldSymbol, coin.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Coin", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Coin fields.
func (c *Coin) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case coin.FieldSymbol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field symbol", values[i])
			} else if value.Valid {
				c.Symbol = value.String
			}
		case coin.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Coin.
// Note that you need to call Coin.Unwrap() before calling this method if this Coin
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Coin) Update() *CoinUpdateOne {
	return (&CoinClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Coin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Coin) Unwrap() *Coin {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("repositories: Coin is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Coin) String() string {
	var builder strings.Builder
	builder.WriteString("Coin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("symbol=")
	builder.WriteString(c.Symbol)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Coins is a parsable slice of Coin.
type Coins []*Coin

func (c Coins) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}