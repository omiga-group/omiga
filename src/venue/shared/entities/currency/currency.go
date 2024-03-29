// Code generated by ent, DO NOT EDIT.

package currency

import (
	"fmt"
	"io"
	"strconv"
)

const (
	// Label holds the string label denoting the currency type in the database.
	Label = "currency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSymbol holds the string denoting the symbol field in the database.
	FieldSymbol = "symbol"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeCurrencyBase holds the string denoting the currency_base edge name in mutations.
	EdgeCurrencyBase = "currency_base"
	// EdgeCurrencyCounter holds the string denoting the currency_counter edge name in mutations.
	EdgeCurrencyCounter = "currency_counter"
	// Table holds the table name of the currency in the database.
	Table = "currencies"
	// CurrencyBaseTable is the table that holds the currency_base relation/edge.
	CurrencyBaseTable = "trading_pairs"
	// CurrencyBaseInverseTable is the table name for the TradingPair entity.
	// It exists in this package in order to avoid circular dependency with the "tradingpair" package.
	CurrencyBaseInverseTable = "trading_pairs"
	// CurrencyBaseColumn is the table column denoting the currency_base relation/edge.
	CurrencyBaseColumn = "currency_currency_base"
	// CurrencyCounterTable is the table that holds the currency_counter relation/edge.
	CurrencyCounterTable = "trading_pairs"
	// CurrencyCounterInverseTable is the table name for the TradingPair entity.
	// It exists in this package in order to avoid circular dependency with the "tradingpair" package.
	CurrencyCounterInverseTable = "trading_pairs"
	// CurrencyCounterColumn is the table column denoting the currency_counter relation/edge.
	CurrencyCounterColumn = "currency_currency_counter"
)

// Columns holds all SQL columns for currency fields.
var Columns = []string{
	FieldID,
	FieldSymbol,
	FieldName,
	FieldType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeDIGITAL Type = "DIGITAL"
	TypeFIAT    Type = "FIAT"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeDIGITAL, TypeFIAT:
		return nil
	default:
		return fmt.Errorf("currency: invalid enum value for type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Type(str)
	if err := TypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}
