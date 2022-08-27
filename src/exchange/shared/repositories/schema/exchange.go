package schema

import (
	"entgo.io/ent"
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
		field.Bool("has_trading_incentive"),
		field.Bool("centralized"),
		field.String("public_notice"),
		field.String("alert_notice"),
		field.Int("trust_score"),
		field.Int("trust_score_rank"),
		field.Float("trade_volume_24h_btc"),
		field.Float("trade_volume_24h_btc_normalized"),
	}
}

// Edges of the Exchange.
func (Exchange) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Exchange) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("exchange_id"),
	}
}
