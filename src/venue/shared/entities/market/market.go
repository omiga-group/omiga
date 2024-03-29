// Code generated by ent, DO NOT EDIT.

package market

import (
	"fmt"
	"io"
	"strconv"
)

const (
	// Label holds the string label denoting the market type in the database.
	Label = "market"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeVenue holds the string denoting the venue edge name in mutations.
	EdgeVenue = "venue"
	// EdgeTradingPair holds the string denoting the trading_pair edge name in mutations.
	EdgeTradingPair = "trading_pair"
	// Table holds the table name of the market in the database.
	Table = "markets"
	// VenueTable is the table that holds the venue relation/edge.
	VenueTable = "markets"
	// VenueInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenueInverseTable = "venues"
	// VenueColumn is the table column denoting the venue relation/edge.
	VenueColumn = "venue_market"
	// TradingPairTable is the table that holds the trading_pair relation/edge. The primary key declared below.
	TradingPairTable = "market_trading_pair"
	// TradingPairInverseTable is the table name for the TradingPair entity.
	// It exists in this package in order to avoid circular dependency with the "tradingpair" package.
	TradingPairInverseTable = "trading_pairs"
)

// Columns holds all SQL columns for market fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "markets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"venue_market",
}

var (
	// TradingPairPrimaryKey and TradingPairColumn2 are the table columns denoting the
	// primary key for the trading_pair relation (M2M).
	TradingPairPrimaryKey = []string{"market_id", "trading_pair_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeSPOT_TRADING     Type = "SPOT_TRADING"
	TypeMARGIN_TRADING   Type = "MARGIN_TRADING"
	TypeDERIVATIVES      Type = "DERIVATIVES"
	TypeEARN             Type = "EARN"
	TypePERPETUAL        Type = "PERPETUAL"
	TypeFUTURES          Type = "FUTURES"
	TypeWARRANT          Type = "WARRANT"
	TypeOTC              Type = "OTC"
	TypeYIELD            Type = "YIELD"
	TypeP2P              Type = "P2P"
	TypeSTRATEGY_TRADING Type = "STRATEGY_TRADING"
	TypeSWAP_FARMING     Type = "SWAP_FARMING"
	TypeFAN_TOKEN        Type = "FAN_TOKEN"
	TypeETF              Type = "ETF"
	TypeNFT              Type = "NFT"
	TypeSWAP             Type = "SWAP"
	TypeCFD              Type = "CFD"
	TypeLIQUIDITY        Type = "LIQUIDITY"
	TypeFARM             Type = "FARM"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeSPOT_TRADING, TypeMARGIN_TRADING, TypeDERIVATIVES, TypeEARN, TypePERPETUAL, TypeFUTURES, TypeWARRANT, TypeOTC, TypeYIELD, TypeP2P, TypeSTRATEGY_TRADING, TypeSWAP_FARMING, TypeFAN_TOKEN, TypeETF, TypeNFT, TypeSWAP, TypeCFD, TypeLIQUIDITY, TypeFARM:
		return nil
	default:
		return fmt.Errorf("market: invalid enum value for type field: %q", _type)
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
