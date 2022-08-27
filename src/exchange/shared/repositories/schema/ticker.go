package schema

import (
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
		field.String("base"),
		field.String("target"),
		field.JSON("market", models.Market{}).Optional(),
		field.Float("last").Optional(),
		field.Float("volume").Optional(),
		field.JSON("converted_last", models.ConvertedDetails{}).Optional(),
		field.JSON("converted_volume", models.ConvertedDetails{}).Optional(),
		field.String("trust_score").Optional(),
		field.Float("bid_ask_spread_percentage").Optional(),
		field.Time("timestamp").Optional(),
		field.Time("last_traded_at").Optional(),
		field.Time("last_fetch_at").Optional(),
		field.Bool("is_anomaly").Optional(),
		field.Bool("is_stale").Optional(),
		field.String("trade_url").Optional(),
		field.String("token_info_url").Optional(),
		field.String("coin_id").Optional(),
		field.String("target_coin_id").Optional(),
	}
}

// Edges of the Ticker.
func (Ticker) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exchange", Exchange.Type).
			Ref("ticker").
			Unique().
			Required(),
	}
}

func (Ticker) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("base"),
		index.Fields("target"),

		index.Fields("coin_id"),
		index.Fields("target_coin_id"),
	}
}
