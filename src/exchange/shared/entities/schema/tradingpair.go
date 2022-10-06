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
		field.String("base").Annotations(entgql.OrderField("base")),
		field.Int("base_precision").Annotations(entgql.OrderField("basePrecision")),
		field.String("counter").Annotations(entgql.OrderField("counter")),
		field.Int("counter_precision").Annotations(entgql.OrderField("counterPrecision")),
	}
}

// Edges of the TradingPair.
func (TradingPair) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exchange", Exchange.Type).
			Ref("trading_pairs").
			Unique().
			Required(),
	}
}

func (TradingPair) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("symbol"),
		index.Fields("base"),
		index.Fields("base_precision"),
		index.Fields("counter"),
		index.Fields("counter_precision"),
	}
}