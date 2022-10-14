package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TradingPair holds the schema definition for the TradingPair entity.
type TradingPair struct {
	ent.Schema
}

// Fields of the TradingPair.
func (TradingPair) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").Annotations(entgql.OrderField("symbol")),

		field.Int("base_price_min_precision").Optional().Annotations(entgql.OrderField("basePriceMinPrecision")),
		field.Int("base_price_max_precision").Optional().Annotations(entgql.OrderField("basePriceMaxPrecision")),
		field.Int("base_quantity_min_precision").Optional().Annotations(entgql.OrderField("baseQuantityMinPrecision")),
		field.Int("base_quantity_max_precision").Optional().Annotations(entgql.OrderField("baseQuantityMaxPrecision")),

		field.Int("counter_price_min_precision").Optional().Annotations(entgql.OrderField("counterPriceMinPrecision")),
		field.Int("counter_price_max_precision").Optional().Annotations(entgql.OrderField("counterPriceMaxPrecision")),
		field.Int("counter_quantity_min_precision").Optional().Annotations(entgql.OrderField("counterQuantityMinPrecision")),
		field.Int("counter_quantity_max_precision").Optional().Annotations(entgql.OrderField("counterQuantityMaxPrecision")),
	}
}

// Edges of the TradingPair.
func (TradingPair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exchange", Exchange.Type).
			Ref("trading_pair").
			Unique().
			Required(),
		edge.From("base", Currency.Type).
			Ref("currency_base").
			Unique().
			Required(),
		edge.From("counter", Currency.Type).
			Ref("currency_counter").
			Unique().
			Required(),
		edge.From("market", Market.Type).
			Ref("trading_pair"),
	}
}

func (TradingPair) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("symbol"),

		index.Fields("base_price_min_precision"),
		index.Fields("base_price_max_precision"),
		index.Fields("base_quantity_min_precision"),
		index.Fields("base_quantity_max_precision"),

		index.Fields("counter_price_min_precision"),
		index.Fields("counter_price_max_precision"),
		index.Fields("counter_quantity_min_precision"),
		index.Fields("counter_quantity_max_precision"),
	}
}
