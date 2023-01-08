// Code generated by ent, DO NOT EDIT.

package market

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/omiga-group/omiga/src/venue/shared/entities/internal"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Market {
	return predicate.Market(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Market {
	return predicate.Market(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Market {
	return predicate.Market(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Market {
	return predicate.Market(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Market {
	return predicate.Market(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Market {
	return predicate.Market(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Market {
	return predicate.Market(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Market {
	return predicate.Market(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Market {
	return predicate.Market(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Market {
	return predicate.Market(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Market {
	return predicate.Market(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Market {
	return predicate.Market(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Market {
	return predicate.Market(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Market {
	return predicate.Market(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Market {
	return predicate.Market(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Market {
	return predicate.Market(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Market {
	return predicate.Market(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Market {
	return predicate.Market(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Market {
	return predicate.Market(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Market {
	return predicate.Market(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Market {
	return predicate.Market(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Market {
	return predicate.Market(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Market {
	return predicate.Market(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Market {
	return predicate.Market(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Market {
	return predicate.Market(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Market {
	return predicate.Market(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Market {
	return predicate.Market(sql.FieldNotIn(FieldType, vs...))
}

// HasVenue applies the HasEdge predicate on the "venue" edge.
func HasVenue() predicate.Market {
	return predicate.Market(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VenueTable, VenueColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.Market
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVenueWith applies the HasEdge predicate on the "venue" edge with a given conditions (other predicates).
func HasVenueWith(preds ...predicate.Venue) predicate.Market {
	return predicate.Market(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VenueInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VenueTable, VenueColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.Market
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTradingPair applies the HasEdge predicate on the "trading_pair" edge.
func HasTradingPair() predicate.Market {
	return predicate.Market(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TradingPairTable, TradingPairPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.MarketTradingPair
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTradingPairWith applies the HasEdge predicate on the "trading_pair" edge with a given conditions (other predicates).
func HasTradingPairWith(preds ...predicate.TradingPair) predicate.Market {
	return predicate.Market(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TradingPairInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TradingPairTable, TradingPairPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.MarketTradingPair
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Market) predicate.Market {
	return predicate.Market(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Market) predicate.Market {
	return predicate.Market(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Market) predicate.Market {
	return predicate.Market(func(s *sql.Selector) {
		p(s.Not())
	})
}
