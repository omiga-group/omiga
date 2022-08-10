package schema

import (
	"entgo.io/ent"
	"github.com/omiga-group/omiga/src/shared/enterprise/outbox"
)

// Outbox holds the schema definition for the Outbox entity.
type Outbox struct {
	ent.Schema
}

// Fields of the Outbox.
func (Outbox) Fields() []ent.Field {
	return outbox.Outbox{}.Fields()
}

// Edges of the Outbox.
func (Outbox) Edges() []ent.Edge {
	return outbox.Outbox{}.Edges()
}

func (Outbox) Indexes() []ent.Index {
	return outbox.Outbox{}.Indexes()
}
