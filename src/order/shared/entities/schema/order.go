package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/order/shared/models"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("order_details", models.OrderDetails{}),
		field.JSON("preferred_exchanges", []models.Exchange{}),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{}
}
