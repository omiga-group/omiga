// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CurrenciesColumns holds the columns for the "currencies" table.
	CurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "symbol", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"DIGITAL", "FIAT"}},
	}
	// CurrenciesTable holds the schema information for the "currencies" table.
	CurrenciesTable = &schema.Table{
		Name:       "currencies",
		Columns:    CurrenciesColumns,
		PrimaryKey: []*schema.Column{CurrenciesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "currency_symbol",
				Unique:  false,
				Columns: []*schema.Column{CurrenciesColumns[1]},
			},
			{
				Name:    "currency_name",
				Unique:  false,
				Columns: []*schema.Column{CurrenciesColumns[2]},
			},
			{
				Name:    "currency_type",
				Unique:  false,
				Columns: []*schema.Column{CurrenciesColumns[3]},
			},
		},
	}
	// MarketsColumns holds the columns for the "markets" table.
	MarketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"SPORT_TRADING", "MARGIN_TRADING", "DERIVATIVES", "EARN", "PERPETUAL", "FUTURES", "WARRANT", "OTC", "YIELD", "P2P", "STRATEGY_TRADING", "SWAP_FARMING", "FAN_TOKEN", "ETF", "NFT", "SWAP"}},
		{Name: "venue_market", Type: field.TypeInt},
	}
	// MarketsTable holds the schema information for the "markets" table.
	MarketsTable = &schema.Table{
		Name:       "markets",
		Columns:    MarketsColumns,
		PrimaryKey: []*schema.Column{MarketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "markets_venues_market",
				Columns:    []*schema.Column{MarketsColumns[3]},
				RefColumns: []*schema.Column{VenuesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "market_name",
				Unique:  false,
				Columns: []*schema.Column{MarketsColumns[1]},
			},
			{
				Name:    "market_type",
				Unique:  false,
				Columns: []*schema.Column{MarketsColumns[2]},
			},
		},
	}
	// OutboxesColumns holds the columns for the "outboxes" table.
	OutboxesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "topic", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "payload", Type: field.TypeBytes},
		{Name: "headers", Type: field.TypeJSON},
		{Name: "retry_count", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "SUCCEEDED", "FAILED"}},
		{Name: "last_retry", Type: field.TypeTime, Nullable: true},
		{Name: "processing_errors", Type: field.TypeJSON, Nullable: true},
	}
	// OutboxesTable holds the schema information for the "outboxes" table.
	OutboxesTable = &schema.Table{
		Name:       "outboxes",
		Columns:    OutboxesColumns,
		PrimaryKey: []*schema.Column{OutboxesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "outbox_last_retry",
				Unique:  false,
				Columns: []*schema.Column{OutboxesColumns[8]},
			},
			{
				Name:    "outbox_status",
				Unique:  false,
				Columns: []*schema.Column{OutboxesColumns[7]},
			},
		},
	}
	// TickersColumns holds the columns for the "tickers" table.
	TickersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "base", Type: field.TypeString},
		{Name: "base_coin_id", Type: field.TypeString, Nullable: true},
		{Name: "counter", Type: field.TypeString},
		{Name: "counter_coin_id", Type: field.TypeString, Nullable: true},
		{Name: "market", Type: field.TypeJSON, Nullable: true},
		{Name: "last", Type: field.TypeFloat64, Nullable: true},
		{Name: "volume", Type: field.TypeFloat64, Nullable: true},
		{Name: "converted_last", Type: field.TypeJSON, Nullable: true},
		{Name: "converted_volume", Type: field.TypeJSON, Nullable: true},
		{Name: "trust_score", Type: field.TypeString, Nullable: true},
		{Name: "bid_ask_spread_percentage", Type: field.TypeFloat64, Nullable: true},
		{Name: "timestamp", Type: field.TypeTime, Nullable: true},
		{Name: "last_traded_at", Type: field.TypeTime, Nullable: true},
		{Name: "last_fetch_at", Type: field.TypeTime, Nullable: true},
		{Name: "is_anomaly", Type: field.TypeBool, Nullable: true},
		{Name: "is_stale", Type: field.TypeBool, Nullable: true},
		{Name: "trade_url", Type: field.TypeString, Nullable: true},
		{Name: "token_info_url", Type: field.TypeString, Nullable: true},
		{Name: "venue_ticker", Type: field.TypeInt},
	}
	// TickersTable holds the schema information for the "tickers" table.
	TickersTable = &schema.Table{
		Name:       "tickers",
		Columns:    TickersColumns,
		PrimaryKey: []*schema.Column{TickersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tickers_venues_ticker",
				Columns:    []*schema.Column{TickersColumns[19]},
				RefColumns: []*schema.Column{VenuesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "ticker_base",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[1]},
			},
			{
				Name:    "ticker_base_coin_id",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[2]},
			},
			{
				Name:    "ticker_counter",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[3]},
			},
			{
				Name:    "ticker_counter_coin_id",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[4]},
			},
			{
				Name:    "ticker_last",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[6]},
			},
			{
				Name:    "ticker_volume",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[7]},
			},
			{
				Name:    "ticker_trust_score",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[10]},
			},
			{
				Name:    "ticker_bid_ask_spread_percentage",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[11]},
			},
			{
				Name:    "ticker_timestamp",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[12]},
			},
			{
				Name:    "ticker_last_traded_at",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[13]},
			},
			{
				Name:    "ticker_last_fetch_at",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[14]},
			},
			{
				Name:    "ticker_is_anomaly",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[15]},
			},
			{
				Name:    "ticker_is_stale",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[16]},
			},
			{
				Name:    "ticker_trade_url",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[17]},
			},
			{
				Name:    "ticker_token_info_url",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[18]},
			},
		},
	}
	// TradingPairsColumns holds the columns for the "trading_pairs" table.
	TradingPairsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "symbol", Type: field.TypeString},
		{Name: "base_price_min_precision", Type: field.TypeInt, Nullable: true},
		{Name: "base_price_max_precision", Type: field.TypeInt, Nullable: true},
		{Name: "base_quantity_min_precision", Type: field.TypeInt, Nullable: true},
		{Name: "base_quantity_max_precision", Type: field.TypeInt, Nullable: true},
		{Name: "counter_price_min_precision", Type: field.TypeInt, Nullable: true},
		{Name: "counter_price_max_precision", Type: field.TypeInt, Nullable: true},
		{Name: "counter_quantity_min_precision", Type: field.TypeInt, Nullable: true},
		{Name: "counter_quantity_max_precision", Type: field.TypeInt, Nullable: true},
		{Name: "currency_currency_base", Type: field.TypeInt},
		{Name: "currency_currency_counter", Type: field.TypeInt},
		{Name: "venue_trading_pair", Type: field.TypeInt},
	}
	// TradingPairsTable holds the schema information for the "trading_pairs" table.
	TradingPairsTable = &schema.Table{
		Name:       "trading_pairs",
		Columns:    TradingPairsColumns,
		PrimaryKey: []*schema.Column{TradingPairsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "trading_pairs_currencies_currency_base",
				Columns:    []*schema.Column{TradingPairsColumns[10]},
				RefColumns: []*schema.Column{CurrenciesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "trading_pairs_currencies_currency_counter",
				Columns:    []*schema.Column{TradingPairsColumns[11]},
				RefColumns: []*schema.Column{CurrenciesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "trading_pairs_venues_trading_pair",
				Columns:    []*schema.Column{TradingPairsColumns[12]},
				RefColumns: []*schema.Column{VenuesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "tradingpair_symbol",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[1]},
			},
			{
				Name:    "tradingpair_base_price_min_precision",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[2]},
			},
			{
				Name:    "tradingpair_base_price_max_precision",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[3]},
			},
			{
				Name:    "tradingpair_base_quantity_min_precision",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[4]},
			},
			{
				Name:    "tradingpair_base_quantity_max_precision",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[5]},
			},
			{
				Name:    "tradingpair_counter_price_min_precision",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[6]},
			},
			{
				Name:    "tradingpair_counter_price_max_precision",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[7]},
			},
			{
				Name:    "tradingpair_counter_quantity_min_precision",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[8]},
			},
			{
				Name:    "tradingpair_counter_quantity_max_precision",
				Unique:  false,
				Columns: []*schema.Column{TradingPairsColumns[9]},
			},
		},
	}
	// VenuesColumns holds the columns for the "venues" table.
	VenuesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "venue_id", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"EXCHANGE"}},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "year_established", Type: field.TypeInt, Nullable: true},
		{Name: "country", Type: field.TypeString, Nullable: true},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "links", Type: field.TypeJSON, Nullable: true},
		{Name: "has_trading_incentive", Type: field.TypeBool, Nullable: true},
		{Name: "centralized", Type: field.TypeBool, Nullable: true},
		{Name: "public_notice", Type: field.TypeString, Nullable: true},
		{Name: "alert_notice", Type: field.TypeString, Nullable: true},
		{Name: "trust_score", Type: field.TypeInt, Nullable: true},
		{Name: "trust_score_rank", Type: field.TypeInt, Nullable: true},
		{Name: "trade_volume_24h_btc", Type: field.TypeFloat64, Nullable: true},
		{Name: "trade_volume_24h_btc_normalized", Type: field.TypeFloat64, Nullable: true},
		{Name: "maker_fee", Type: field.TypeFloat64, Nullable: true},
		{Name: "taker_fee", Type: field.TypeFloat64, Nullable: true},
		{Name: "spread_fee", Type: field.TypeBool, Nullable: true},
		{Name: "support_api", Type: field.TypeBool, Nullable: true},
	}
	// VenuesTable holds the schema information for the "venues" table.
	VenuesTable = &schema.Table{
		Name:       "venues",
		Columns:    VenuesColumns,
		PrimaryKey: []*schema.Column{VenuesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "venue_venue_id",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[1]},
			},
			{
				Name:    "venue_type",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[2]},
			},
			{
				Name:    "venue_name",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[3]},
			},
			{
				Name:    "venue_year_established",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[4]},
			},
			{
				Name:    "venue_country",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[5]},
			},
			{
				Name:    "venue_image",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[6]},
			},
			{
				Name:    "venue_has_trading_incentive",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[8]},
			},
			{
				Name:    "venue_centralized",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[9]},
			},
			{
				Name:    "venue_public_notice",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[10]},
			},
			{
				Name:    "venue_alert_notice",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[11]},
			},
			{
				Name:    "venue_trust_score",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[12]},
			},
			{
				Name:    "venue_trust_score_rank",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[13]},
			},
			{
				Name:    "venue_trade_volume_24h_btc",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[14]},
			},
			{
				Name:    "venue_trade_volume_24h_btc_normalized",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[15]},
			},
			{
				Name:    "venue_maker_fee",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[16]},
			},
			{
				Name:    "venue_taker_fee",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[17]},
			},
			{
				Name:    "venue_spread_fee",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[18]},
			},
			{
				Name:    "venue_support_api",
				Unique:  false,
				Columns: []*schema.Column{VenuesColumns[19]},
			},
		},
	}
	// MarketTradingPairColumns holds the columns for the "market_trading_pair" table.
	MarketTradingPairColumns = []*schema.Column{
		{Name: "market_id", Type: field.TypeInt},
		{Name: "trading_pair_id", Type: field.TypeInt},
	}
	// MarketTradingPairTable holds the schema information for the "market_trading_pair" table.
	MarketTradingPairTable = &schema.Table{
		Name:       "market_trading_pair",
		Columns:    MarketTradingPairColumns,
		PrimaryKey: []*schema.Column{MarketTradingPairColumns[0], MarketTradingPairColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "market_trading_pair_market_id",
				Columns:    []*schema.Column{MarketTradingPairColumns[0]},
				RefColumns: []*schema.Column{MarketsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "market_trading_pair_trading_pair_id",
				Columns:    []*schema.Column{MarketTradingPairColumns[1]},
				RefColumns: []*schema.Column{TradingPairsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CurrenciesTable,
		MarketsTable,
		OutboxesTable,
		TickersTable,
		TradingPairsTable,
		VenuesTable,
		MarketTradingPairTable,
	}
)

func init() {
	MarketsTable.ForeignKeys[0].RefTable = VenuesTable
	TickersTable.ForeignKeys[0].RefTable = VenuesTable
	TradingPairsTable.ForeignKeys[0].RefTable = CurrenciesTable
	TradingPairsTable.ForeignKeys[1].RefTable = CurrenciesTable
	TradingPairsTable.ForeignKeys[2].RefTable = VenuesTable
	MarketTradingPairTable.ForeignKeys[0].RefTable = MarketsTable
	MarketTradingPairTable.ForeignKeys[1].RefTable = TradingPairsTable
}
