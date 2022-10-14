package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Currency holds the schema definition for the Currency entity.
type Currency struct {
	ent.Schema
}

// Fields of the Currency.
func (Currency) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").Annotations(entgql.OrderField("symbol")),
		field.String("name").Optional().Annotations(entgql.OrderField("name")),

		field.Enum("type").
			NamedValues(
				"DIGITAL", "DIGITAL",
				"FIAT", "FIAT").
			Annotations(entgql.OrderField("type")),
	}
}

// Edges of the Currency.
func (Currency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("currency_base", TradingPair.Type),
		edge.To("currency_counter", TradingPair.Type),
	}
}

func (Currency) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("symbol"),
		index.Fields("name"),
		index.Fields("type"),
	}
}
