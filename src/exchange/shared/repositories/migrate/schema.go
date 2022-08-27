// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ExchangesColumns holds the columns for the "exchanges" table.
	ExchangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "exchange_id", Type: field.TypeString},
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
	}
	// ExchangesTable holds the schema information for the "exchanges" table.
	ExchangesTable = &schema.Table{
		Name:       "exchanges",
		Columns:    ExchangesColumns,
		PrimaryKey: []*schema.Column{ExchangesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "exchange_exchange_id",
				Unique:  false,
				Columns: []*schema.Column{ExchangesColumns[1]},
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
		{Name: "target", Type: field.TypeString},
		{Name: "market", Type: field.TypeJSON},
		{Name: "last", Type: field.TypeFloat64},
		{Name: "volume", Type: field.TypeFloat64},
		{Name: "converted_last", Type: field.TypeJSON},
		{Name: "converted_volume", Type: field.TypeJSON},
		{Name: "trust_score", Type: field.TypeString},
		{Name: "bid_ask_spread_percentage", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "last_traded_at", Type: field.TypeTime},
		{Name: "last_fetch_at", Type: field.TypeTime},
		{Name: "is_anomaly", Type: field.TypeBool},
		{Name: "is_stale", Type: field.TypeBool},
		{Name: "trade_url", Type: field.TypeString, Nullable: true},
		{Name: "token_info_url", Type: field.TypeString, Nullable: true},
		{Name: "coin_id", Type: field.TypeString},
		{Name: "target_coin_id", Type: field.TypeString},
		{Name: "exchange_ticker", Type: field.TypeInt},
	}
	// TickersTable holds the schema information for the "tickers" table.
	TickersTable = &schema.Table{
		Name:       "tickers",
		Columns:    TickersColumns,
		PrimaryKey: []*schema.Column{TickersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tickers_exchanges_ticker",
				Columns:    []*schema.Column{TickersColumns[19]},
				RefColumns: []*schema.Column{ExchangesColumns[0]},
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
				Name:    "ticker_base",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[1]},
			},
			{
				Name:    "ticker_coin_id",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[17]},
			},
			{
				Name:    "ticker_target_coin_id",
				Unique:  false,
				Columns: []*schema.Column{TickersColumns[18]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ExchangesTable,
		OutboxesTable,
		TickersTable,
	}
)

func init() {
	TickersTable.ForeignKeys[0].RefTable = ExchangesTable
}
