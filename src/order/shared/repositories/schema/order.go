package schema

import (
	"entgo.io/ent"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{}
}
