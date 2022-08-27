package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Exchange holds the schema definition for the Exchange entity.
type Exchange struct {
	ent.Schema
}

// Fields of the Exchange.
func (Exchange) Fields() []ent.Field {
	return []ent.Field{
		field.String("exchange_id"),
		field.String("name").Optional(),
		field.Int("year_established").Optional(),
		field.String("country").Optional(),
		field.String("image").Optional(),
		field.JSON("links", map[string]string{}).Optional(),
		field.Bool("has_trading_incentive").Optional(),
		field.Bool("centralized").Optional(),
		field.String("public_notice").Optional(),
		field.String("alert_notice").Optional(),
		field.Int("trust_score").Optional(),
		field.Int("trust_score_rank").Optional(),
		field.Float("trade_volume_24h_btc").Optional(),
		field.Float("trade_volume_24h_btc_normalized").Optional(),
	}
}

// Edges of the Exchange.
func (Exchange) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ticker", Ticker.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Exchange) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("exchange_id"),
	}
}
