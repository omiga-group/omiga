package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/order/shared/models"
)

// OrderBook holds the schema definition for the OrderBook entity.
type OrderBook struct {
	ent.Schema
}

// Fields of the OrderBook.
func (OrderBook) Fields() []ent.Field {
	return []ent.Field{
		field.String("exchange_id"),
		field.Time("last_updated"),
		field.JSON("order_book", models.OrderBook{}),
	}
}

// Edges of the OrderBook.
func (OrderBook) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (OrderBook) Indexes() []ent.Index {
	return []ent.Index{}
}
