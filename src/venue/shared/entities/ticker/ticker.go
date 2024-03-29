// Code generated by ent, DO NOT EDIT.

package ticker

const (
	// Label holds the string label denoting the ticker type in the database.
	Label = "ticker"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBase holds the string denoting the base field in the database.
	FieldBase = "base"
	// FieldBaseCoinID holds the string denoting the base_coin_id field in the database.
	FieldBaseCoinID = "base_coin_id"
	// FieldCounter holds the string denoting the counter field in the database.
	FieldCounter = "counter"
	// FieldCounterCoinID holds the string denoting the counter_coin_id field in the database.
	FieldCounterCoinID = "counter_coin_id"
	// FieldMarket holds the string denoting the market field in the database.
	FieldMarket = "market"
	// FieldLast holds the string denoting the last field in the database.
	FieldLast = "last"
	// FieldVolume holds the string denoting the volume field in the database.
	FieldVolume = "volume"
	// FieldConvertedLast holds the string denoting the converted_last field in the database.
	FieldConvertedLast = "converted_last"
	// FieldConvertedVolume holds the string denoting the converted_volume field in the database.
	FieldConvertedVolume = "converted_volume"
	// FieldTrustScore holds the string denoting the trust_score field in the database.
	FieldTrustScore = "trust_score"
	// FieldBidAskSpreadPercentage holds the string denoting the bid_ask_spread_percentage field in the database.
	FieldBidAskSpreadPercentage = "bid_ask_spread_percentage"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldLastTradedAt holds the string denoting the last_traded_at field in the database.
	FieldLastTradedAt = "last_traded_at"
	// FieldLastFetchAt holds the string denoting the last_fetch_at field in the database.
	FieldLastFetchAt = "last_fetch_at"
	// FieldIsAnomaly holds the string denoting the is_anomaly field in the database.
	FieldIsAnomaly = "is_anomaly"
	// FieldIsStale holds the string denoting the is_stale field in the database.
	FieldIsStale = "is_stale"
	// FieldTradeURL holds the string denoting the trade_url field in the database.
	FieldTradeURL = "trade_url"
	// FieldTokenInfoURL holds the string denoting the token_info_url field in the database.
	FieldTokenInfoURL = "token_info_url"
	// EdgeVenue holds the string denoting the venue edge name in mutations.
	EdgeVenue = "venue"
	// Table holds the table name of the ticker in the database.
	Table = "tickers"
	// VenueTable is the table that holds the venue relation/edge.
	VenueTable = "tickers"
	// VenueInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenueInverseTable = "venues"
	// VenueColumn is the table column denoting the venue relation/edge.
	VenueColumn = "venue_ticker"
)

// Columns holds all SQL columns for ticker fields.
var Columns = []string{
	FieldID,
	FieldBase,
	FieldBaseCoinID,
	FieldCounter,
	FieldCounterCoinID,
	FieldMarket,
	FieldLast,
	FieldVolume,
	FieldConvertedLast,
	FieldConvertedVolume,
	FieldTrustScore,
	FieldBidAskSpreadPercentage,
	FieldTimestamp,
	FieldLastTradedAt,
	FieldLastFetchAt,
	FieldIsAnomaly,
	FieldIsStale,
	FieldTradeURL,
	FieldTokenInfoURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tickers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"venue_ticker",
}

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
