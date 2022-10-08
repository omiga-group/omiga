package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Coin holds the schema definition for the Coin entity.
type Coin struct {
	ent.Schema
}

// Fields of the Coin.
func (Coin) Fields() []ent.Field {
	return []ent.Field{
		field.String("symbol").Annotations(entgql.OrderField("symbol")),
		field.String("name").Optional().Annotations(entgql.OrderField("name")),
	}
}

// Edges of the Coin.
func (Coin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coin_base", TradingPair.Type),
		edge.To("coin_counter", TradingPair.Type),
	}
}

func (Coin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("symbol"),
		index.Fields("name"),
	}
}
