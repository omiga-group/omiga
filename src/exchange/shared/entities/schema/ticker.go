package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
)

// Ticker holds the schema definition for the Ticker entity.
type Ticker struct {
	ent.Schema
}

// Fields of the Ticker.
func (Ticker) Fields() []ent.Field {
	return []ent.Field{
		field.String("base").Annotations(entgql.OrderField("base")),
		field.String("base_coin_id").Optional().Annotations(entgql.OrderField("baseCoinId")),
		field.String("counter").Annotations(entgql.OrderField("counter")),
		field.String("counter_coin_id").Optional().Annotations(entgql.OrderField("counterCoinId")),
		field.JSON("market", models.Market{}).Optional(),
		field.Float("last").Optional().Annotations(entgql.OrderField("last")),
		field.Float("volume").Optional().Annotations(entgql.OrderField("volume")),
		field.JSON("converted_last", models.ConvertedDetails{}).Optional(),
		field.JSON("converted_volume", models.ConvertedDetails{}).Optional(),
		field.String("trust_score").Optional().Annotations(entgql.OrderField("trustScore")),
		field.Float("bid_ask_spread_percentage").Optional().Annotations(entgql.OrderField("bidAskSpreadPercentage")),
		field.Time("timestamp").Optional().Annotations(entgql.OrderField("timestamp")),
		field.Time("last_traded_at").Optional().Annotations(entgql.OrderField("lastTradedAt")),
		field.Time("last_fetch_at").Optional().Annotations(entgql.OrderField("lastFetchAt")),
		field.Bool("is_anomaly").Optional().Annotations(entgql.OrderField("isAnomaly")),
		field.Bool("is_stale").Optional().Annotations(entgql.OrderField("isStale")),
		field.String("trade_url").Optional().Annotations(entgql.OrderField("tradeUrl")),
		field.String("token_info_url").Optional().Annotations(entgql.OrderField("tokenInfoUrl")),
	}
}

// Edges of the Ticker.
func (Ticker) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("venue", Venue.Type).
			Ref("ticker").
			Unique().
			Required(),
	}
}

func (Ticker) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("base"),
		index.Fields("base_coin_id"),
		index.Fields("counter"),
		index.Fields("counter_coin_id"),
		index.Fields("last"),
		index.Fields("volume"),
		index.Fields("trust_score"),
		index.Fields("bid_ask_spread_percentage"),
		index.Fields("timestamp"),
		index.Fields("last_traded_at"),
		index.Fields("last_fetch_at"),
		index.Fields("is_anomaly"),
		index.Fields("is_stale"),
		index.Fields("trade_url"),
		index.Fields("token_info_url"),
	}
}
