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
		field.JSON("market", models.Market{}),
		field.Float("last"),
		field.Float("volume"),
		field.JSON("converted_last", models.ConvertedDetails{}),
		field.JSON("converted_volume", models.ConvertedDetails{}),
		field.String("trust_score"),
		field.Float("bid_ask_spread_percentage"),
		field.Time("timestamp"),
		field.Time("last_traded_at"),
		field.Time("last_fetch_at"),
		field.Bool("is_anomaly"),
		field.Bool("is_stale"),
		field.String("trade_url").Optional(),
		field.String("token_info_url").Optional(),
		field.String("coin_id"),
		field.String("target_coin_id"),
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
		index.Fields("base"),

		index.Fields("coin_id"),
		index.Fields("target_coin_id"),
	}
}
