// Code generated by ent, DO NOT EDIT.

package tradingpair

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/internal"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// BasePriceMinPrecision applies equality check predicate on the "base_price_min_precision" field. It's identical to BasePriceMinPrecisionEQ.
func BasePriceMinPrecision(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBasePriceMinPrecision), v))
	})
}

// BasePriceMaxPrecision applies equality check predicate on the "base_price_max_precision" field. It's identical to BasePriceMaxPrecisionEQ.
func BasePriceMaxPrecision(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBasePriceMaxPrecision), v))
	})
}

// BaseQuantityMinPrecision applies equality check predicate on the "base_quantity_min_precision" field. It's identical to BaseQuantityMinPrecisionEQ.
func BaseQuantityMinPrecision(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaseQuantityMinPrecision), v))
	})
}

// BaseQuantityMaxPrecision applies equality check predicate on the "base_quantity_max_precision" field. It's identical to BaseQuantityMaxPrecisionEQ.
func BaseQuantityMaxPrecision(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaseQuantityMaxPrecision), v))
	})
}

// CounterPriceMinPrecision applies equality check predicate on the "counter_price_min_precision" field. It's identical to CounterPriceMinPrecisionEQ.
func CounterPriceMinPrecision(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounterPriceMinPrecision), v))
	})
}

// CounterPriceMaxPrecision applies equality check predicate on the "counter_price_max_precision" field. It's identical to CounterPriceMaxPrecisionEQ.
func CounterPriceMaxPrecision(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounterPriceMaxPrecision), v))
	})
}

// CounterQuantityMinPrecision applies equality check predicate on the "counter_quantity_min_precision" field. It's identical to CounterQuantityMinPrecisionEQ.
func CounterQuantityMinPrecision(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounterQuantityMinPrecision), v))
	})
}

// CounterQuantityMaxPrecision applies equality check predicate on the "counter_quantity_max_precision" field. It's identical to CounterQuantityMaxPrecisionEQ.
func CounterQuantityMaxPrecision(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounterQuantityMaxPrecision), v))
	})
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymbol), v))
	})
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSymbol), v...))
	})
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSymbol), v...))
	})
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymbol), v))
	})
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymbol), v))
	})
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymbol), v))
	})
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymbol), v))
	})
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymbol), v))
	})
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymbol), v))
	})
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymbol), v))
	})
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymbol), v))
	})
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymbol), v))
	})
}

// BasePriceMinPrecisionEQ applies the EQ predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBasePriceMinPrecision), v))
	})
}

// BasePriceMinPrecisionNEQ applies the NEQ predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionNEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBasePriceMinPrecision), v))
	})
}

// BasePriceMinPrecisionIn applies the In predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBasePriceMinPrecision), v...))
	})
}

// BasePriceMinPrecisionNotIn applies the NotIn predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionNotIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBasePriceMinPrecision), v...))
	})
}

// BasePriceMinPrecisionGT applies the GT predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionGT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBasePriceMinPrecision), v))
	})
}

// BasePriceMinPrecisionGTE applies the GTE predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionGTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBasePriceMinPrecision), v))
	})
}

// BasePriceMinPrecisionLT applies the LT predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionLT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBasePriceMinPrecision), v))
	})
}

// BasePriceMinPrecisionLTE applies the LTE predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionLTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBasePriceMinPrecision), v))
	})
}

// BasePriceMinPrecisionIsNil applies the IsNil predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionIsNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBasePriceMinPrecision)))
	})
}

// BasePriceMinPrecisionNotNil applies the NotNil predicate on the "base_price_min_precision" field.
func BasePriceMinPrecisionNotNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBasePriceMinPrecision)))
	})
}

// BasePriceMaxPrecisionEQ applies the EQ predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBasePriceMaxPrecision), v))
	})
}

// BasePriceMaxPrecisionNEQ applies the NEQ predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionNEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBasePriceMaxPrecision), v))
	})
}

// BasePriceMaxPrecisionIn applies the In predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBasePriceMaxPrecision), v...))
	})
}

// BasePriceMaxPrecisionNotIn applies the NotIn predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionNotIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBasePriceMaxPrecision), v...))
	})
}

// BasePriceMaxPrecisionGT applies the GT predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionGT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBasePriceMaxPrecision), v))
	})
}

// BasePriceMaxPrecisionGTE applies the GTE predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionGTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBasePriceMaxPrecision), v))
	})
}

// BasePriceMaxPrecisionLT applies the LT predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionLT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBasePriceMaxPrecision), v))
	})
}

// BasePriceMaxPrecisionLTE applies the LTE predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionLTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBasePriceMaxPrecision), v))
	})
}

// BasePriceMaxPrecisionIsNil applies the IsNil predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionIsNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBasePriceMaxPrecision)))
	})
}

// BasePriceMaxPrecisionNotNil applies the NotNil predicate on the "base_price_max_precision" field.
func BasePriceMaxPrecisionNotNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBasePriceMaxPrecision)))
	})
}

// BaseQuantityMinPrecisionEQ applies the EQ predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaseQuantityMinPrecision), v))
	})
}

// BaseQuantityMinPrecisionNEQ applies the NEQ predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionNEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaseQuantityMinPrecision), v))
	})
}

// BaseQuantityMinPrecisionIn applies the In predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBaseQuantityMinPrecision), v...))
	})
}

// BaseQuantityMinPrecisionNotIn applies the NotIn predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionNotIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBaseQuantityMinPrecision), v...))
	})
}

// BaseQuantityMinPrecisionGT applies the GT predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionGT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaseQuantityMinPrecision), v))
	})
}

// BaseQuantityMinPrecisionGTE applies the GTE predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionGTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaseQuantityMinPrecision), v))
	})
}

// BaseQuantityMinPrecisionLT applies the LT predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionLT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaseQuantityMinPrecision), v))
	})
}

// BaseQuantityMinPrecisionLTE applies the LTE predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionLTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaseQuantityMinPrecision), v))
	})
}

// BaseQuantityMinPrecisionIsNil applies the IsNil predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionIsNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBaseQuantityMinPrecision)))
	})
}

// BaseQuantityMinPrecisionNotNil applies the NotNil predicate on the "base_quantity_min_precision" field.
func BaseQuantityMinPrecisionNotNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBaseQuantityMinPrecision)))
	})
}

// BaseQuantityMaxPrecisionEQ applies the EQ predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBaseQuantityMaxPrecision), v))
	})
}

// BaseQuantityMaxPrecisionNEQ applies the NEQ predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionNEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBaseQuantityMaxPrecision), v))
	})
}

// BaseQuantityMaxPrecisionIn applies the In predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBaseQuantityMaxPrecision), v...))
	})
}

// BaseQuantityMaxPrecisionNotIn applies the NotIn predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionNotIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBaseQuantityMaxPrecision), v...))
	})
}

// BaseQuantityMaxPrecisionGT applies the GT predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionGT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBaseQuantityMaxPrecision), v))
	})
}

// BaseQuantityMaxPrecisionGTE applies the GTE predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionGTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBaseQuantityMaxPrecision), v))
	})
}

// BaseQuantityMaxPrecisionLT applies the LT predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionLT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBaseQuantityMaxPrecision), v))
	})
}

// BaseQuantityMaxPrecisionLTE applies the LTE predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionLTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBaseQuantityMaxPrecision), v))
	})
}

// BaseQuantityMaxPrecisionIsNil applies the IsNil predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionIsNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBaseQuantityMaxPrecision)))
	})
}

// BaseQuantityMaxPrecisionNotNil applies the NotNil predicate on the "base_quantity_max_precision" field.
func BaseQuantityMaxPrecisionNotNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBaseQuantityMaxPrecision)))
	})
}

// CounterPriceMinPrecisionEQ applies the EQ predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounterPriceMinPrecision), v))
	})
}

// CounterPriceMinPrecisionNEQ applies the NEQ predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionNEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCounterPriceMinPrecision), v))
	})
}

// CounterPriceMinPrecisionIn applies the In predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCounterPriceMinPrecision), v...))
	})
}

// CounterPriceMinPrecisionNotIn applies the NotIn predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionNotIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCounterPriceMinPrecision), v...))
	})
}

// CounterPriceMinPrecisionGT applies the GT predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionGT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCounterPriceMinPrecision), v))
	})
}

// CounterPriceMinPrecisionGTE applies the GTE predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionGTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCounterPriceMinPrecision), v))
	})
}

// CounterPriceMinPrecisionLT applies the LT predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionLT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCounterPriceMinPrecision), v))
	})
}

// CounterPriceMinPrecisionLTE applies the LTE predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionLTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCounterPriceMinPrecision), v))
	})
}

// CounterPriceMinPrecisionIsNil applies the IsNil predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionIsNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCounterPriceMinPrecision)))
	})
}

// CounterPriceMinPrecisionNotNil applies the NotNil predicate on the "counter_price_min_precision" field.
func CounterPriceMinPrecisionNotNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCounterPriceMinPrecision)))
	})
}

// CounterPriceMaxPrecisionEQ applies the EQ predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounterPriceMaxPrecision), v))
	})
}

// CounterPriceMaxPrecisionNEQ applies the NEQ predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionNEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCounterPriceMaxPrecision), v))
	})
}

// CounterPriceMaxPrecisionIn applies the In predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCounterPriceMaxPrecision), v...))
	})
}

// CounterPriceMaxPrecisionNotIn applies the NotIn predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionNotIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCounterPriceMaxPrecision), v...))
	})
}

// CounterPriceMaxPrecisionGT applies the GT predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionGT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCounterPriceMaxPrecision), v))
	})
}

// CounterPriceMaxPrecisionGTE applies the GTE predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionGTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCounterPriceMaxPrecision), v))
	})
}

// CounterPriceMaxPrecisionLT applies the LT predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionLT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCounterPriceMaxPrecision), v))
	})
}

// CounterPriceMaxPrecisionLTE applies the LTE predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionLTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCounterPriceMaxPrecision), v))
	})
}

// CounterPriceMaxPrecisionIsNil applies the IsNil predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionIsNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCounterPriceMaxPrecision)))
	})
}

// CounterPriceMaxPrecisionNotNil applies the NotNil predicate on the "counter_price_max_precision" field.
func CounterPriceMaxPrecisionNotNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCounterPriceMaxPrecision)))
	})
}

// CounterQuantityMinPrecisionEQ applies the EQ predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounterQuantityMinPrecision), v))
	})
}

// CounterQuantityMinPrecisionNEQ applies the NEQ predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionNEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCounterQuantityMinPrecision), v))
	})
}

// CounterQuantityMinPrecisionIn applies the In predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCounterQuantityMinPrecision), v...))
	})
}

// CounterQuantityMinPrecisionNotIn applies the NotIn predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionNotIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCounterQuantityMinPrecision), v...))
	})
}

// CounterQuantityMinPrecisionGT applies the GT predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionGT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCounterQuantityMinPrecision), v))
	})
}

// CounterQuantityMinPrecisionGTE applies the GTE predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionGTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCounterQuantityMinPrecision), v))
	})
}

// CounterQuantityMinPrecisionLT applies the LT predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionLT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCounterQuantityMinPrecision), v))
	})
}

// CounterQuantityMinPrecisionLTE applies the LTE predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionLTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCounterQuantityMinPrecision), v))
	})
}

// CounterQuantityMinPrecisionIsNil applies the IsNil predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionIsNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCounterQuantityMinPrecision)))
	})
}

// CounterQuantityMinPrecisionNotNil applies the NotNil predicate on the "counter_quantity_min_precision" field.
func CounterQuantityMinPrecisionNotNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCounterQuantityMinPrecision)))
	})
}

// CounterQuantityMaxPrecisionEQ applies the EQ predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCounterQuantityMaxPrecision), v))
	})
}

// CounterQuantityMaxPrecisionNEQ applies the NEQ predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionNEQ(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCounterQuantityMaxPrecision), v))
	})
}

// CounterQuantityMaxPrecisionIn applies the In predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCounterQuantityMaxPrecision), v...))
	})
}

// CounterQuantityMaxPrecisionNotIn applies the NotIn predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionNotIn(vs ...int) predicate.TradingPair {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCounterQuantityMaxPrecision), v...))
	})
}

// CounterQuantityMaxPrecisionGT applies the GT predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionGT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCounterQuantityMaxPrecision), v))
	})
}

// CounterQuantityMaxPrecisionGTE applies the GTE predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionGTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCounterQuantityMaxPrecision), v))
	})
}

// CounterQuantityMaxPrecisionLT applies the LT predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionLT(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCounterQuantityMaxPrecision), v))
	})
}

// CounterQuantityMaxPrecisionLTE applies the LTE predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionLTE(v int) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCounterQuantityMaxPrecision), v))
	})
}

// CounterQuantityMaxPrecisionIsNil applies the IsNil predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionIsNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCounterQuantityMaxPrecision)))
	})
}

// CounterQuantityMaxPrecisionNotNil applies the NotNil predicate on the "counter_quantity_max_precision" field.
func CounterQuantityMaxPrecisionNotNil() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCounterQuantityMaxPrecision)))
	})
}

// HasExchange applies the HasEdge predicate on the "exchange" edge.
func HasExchange() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ExchangeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExchangeTable, ExchangeColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Exchange
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExchangeWith applies the HasEdge predicate on the "exchange" edge with a given conditions (other predicates).
func HasExchangeWith(preds ...predicate.Exchange) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ExchangeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExchangeTable, ExchangeColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Exchange
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBase applies the HasEdge predicate on the "base" edge.
func HasBase() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BaseTable, BaseColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Currency
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBaseWith applies the HasEdge predicate on the "base" edge with a given conditions (other predicates).
func HasBaseWith(preds ...predicate.Currency) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BaseTable, BaseColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Currency
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCounter applies the HasEdge predicate on the "counter" edge.
func HasCounter() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CounterTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CounterTable, CounterColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Currency
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCounterWith applies the HasEdge predicate on the "counter" edge with a given conditions (other predicates).
func HasCounterWith(preds ...predicate.Currency) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CounterInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CounterTable, CounterColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Currency
		step.Edge.Schema = schemaConfig.TradingPair
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMarket applies the HasEdge predicate on the "market" edge.
func HasMarket() predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MarketTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MarketTable, MarketPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Market
		step.Edge.Schema = schemaConfig.MarketTradingPair
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMarketWith applies the HasEdge predicate on the "market" edge with a given conditions (other predicates).
func HasMarketWith(preds ...predicate.Market) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MarketInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MarketTable, MarketPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Market
		step.Edge.Schema = schemaConfig.MarketTradingPair
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TradingPair) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TradingPair) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
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
func Not(p predicate.TradingPair) predicate.TradingPair {
	return predicate.TradingPair(func(s *sql.Selector) {
		p(s.Not())
	})
}
