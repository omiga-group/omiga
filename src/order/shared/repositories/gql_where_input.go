// Code generated by ent, DO NOT EDIT.

package repositories

import (
	"errors"
	"fmt"

	"github.com/omiga-group/omiga/src/order/shared/repositories/order"
	"github.com/omiga-group/omiga/src/order/shared/repositories/predicate"
)

// OrderWhereInput represents a where input for filtering Order queries.
type OrderWhereInput struct {
	Predicates []predicate.Order  `json:"-"`
	Not        *OrderWhereInput   `json:"not,omitempty"`
	Or         []*OrderWhereInput `json:"or,omitempty"`
	And        []*OrderWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *OrderWhereInput) AddPredicates(predicates ...predicate.Order) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the OrderWhereInput filter on the OrderQuery builder.
func (i *OrderWhereInput) Filter(q *OrderQuery) (*OrderQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyOrderWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyOrderWhereInput is returned in case the OrderWhereInput is empty.
var ErrEmptyOrderWhereInput = errors.New("repositories: empty predicate OrderWhereInput")

// P returns a predicate for filtering orders.
// An error is returned if the input is empty or invalid.
func (i *OrderWhereInput) P() (predicate.Order, error) {
	var predicates []predicate.Order
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, order.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Order, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, order.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Order, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, order.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, order.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, order.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, order.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, order.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, order.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, order.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, order.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, order.IDLTE(*i.IDLTE))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyOrderWhereInput
	case 1:
		return predicates[0], nil
	default:
		return order.And(predicates...), nil
	}
}
