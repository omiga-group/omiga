// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/venue/shared/entities/currency"
	"github.com/omiga-group/omiga/src/venue/shared/entities/internal"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
)

// CurrencyDelete is the builder for deleting a Currency entity.
type CurrencyDelete struct {
	config
	hooks    []Hook
	mutation *CurrencyMutation
}

// Where appends a list predicates to the CurrencyDelete builder.
func (cd *CurrencyDelete) Where(ps ...predicate.Currency) *CurrencyDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CurrencyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CurrencyMutation](ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CurrencyDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CurrencyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: currency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: currency.FieldID,
			},
		},
	}
	_spec.Node.Schema = cd.schemaConfig.Currency
	ctx = internal.NewSchemaConfigContext(ctx, cd.schemaConfig)
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// CurrencyDeleteOne is the builder for deleting a single Currency entity.
type CurrencyDeleteOne struct {
	cd *CurrencyDelete
}

// Where appends a list predicates to the CurrencyDelete builder.
func (cdo *CurrencyDeleteOne) Where(ps ...predicate.Currency) *CurrencyDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *CurrencyDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{currency.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CurrencyDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
