package schema

import (
	"entgo.io/contrib/entgql"
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
		field.String("exchange_id").Annotations(entgql.OrderField("exchangeId")),
		field.String("name").Optional().Annotations(entgql.OrderField("name")),
		field.Int("year_established").Optional().Annotations(entgql.OrderField("yearEstablished")),
		field.String("country").Optional().Annotations(entgql.OrderField("country")),
		field.String("image").Optional().Annotations(entgql.OrderField("image")),
		field.JSON("links", map[string]string{}).Optional(),
		field.Bool("has_trading_incentive").Optional().Annotations(entgql.OrderField("hasTradingIncentive")),
		field.Bool("centralized").Optional().Annotations(entgql.OrderField("centralized")),
		field.String("public_notice").Optional().Annotations(entgql.OrderField("publicNotice")),
		field.String("alert_notice").Optional().Annotations(entgql.OrderField("alertNotice")),
		field.Int("trust_score").Optional().Annotations(entgql.OrderField("trustScore")),
		field.Int("trust_score_rank").Optional().Annotations(entgql.OrderField("trustScoreRank")),
		field.Float("trade_volume_24h_btc").Optional().Annotations(entgql.OrderField("tradeVolume24hBtc")),
		field.Float("trade_volume_24h_btc_normalized").Optional().Annotations(entgql.OrderField("tradeVolume24hBtcNormalized")),
		field.Float("maker_fee").Optional().Annotations(entgql.OrderField("makerFee")),
		field.Float("taker_fee").Optional().Annotations(entgql.OrderField("takerFee")),
		field.Bool("spread_fee").Optional().Annotations(entgql.OrderField("spreadFee")),
		field.Bool("support_api").Optional().Annotations(entgql.OrderField("supportAPI")),
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
		index.Fields("name"),
		index.Fields("year_established"),
		index.Fields("country"),
		index.Fields("image"),
		index.Fields("has_trading_incentive"),
		index.Fields("centralized"),
		index.Fields("public_notice"),
		index.Fields("alert_notice"),
		index.Fields("trust_score"),
		index.Fields("trust_score_rank"),
		index.Fields("trade_volume_24h_btc"),
		index.Fields("trade_volume_24h_btc_normalized"),
		index.Fields("maker_fee"),
		index.Fields("taker_fee"),
		index.Fields("spread_fee"),
		index.Fields("support_api"),
	}
}
