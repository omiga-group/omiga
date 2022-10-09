// Code generated by ent, DO NOT EDIT.

package coin

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/internal"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymbol), v))
	})
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.Coin {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSymbol), v...))
	})
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.Coin {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSymbol), v...))
	})
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymbol), v))
	})
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymbol), v))
	})
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymbol), v))
	})
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymbol), v))
	})
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymbol), v))
	})
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymbol), v))
	})
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymbol), v))
	})
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymbol), v))
	})
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymbol), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Coin {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Coin {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasCoinBase applies the HasEdge predicate on the "coin_base" edge.
func HasCoinBase() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CoinBaseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CoinBaseTable, CoinBaseColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCoinBaseWith applies the HasEdge predicate on the "coin_base" edge with a given conditions (other predicates).
func HasCoinBaseWith(preds ...predicate.TradingPair) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CoinBaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CoinBaseTable, CoinBaseColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCoinCounter applies the HasEdge predicate on the "coin_counter" edge.
func HasCoinCounter() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CoinCounterTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CoinCounterTable, CoinCounterColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCoinCounterWith applies the HasEdge predicate on the "coin_counter" edge with a given conditions (other predicates).
func HasCoinCounterWith(preds ...predicate.TradingPair) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CoinCounterInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CoinCounterTable, CoinCounterColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
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
func Not(p predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		p(s.Not())
	})
}
