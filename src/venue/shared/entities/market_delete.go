// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/venue/shared/entities/internal"
	"github.com/omiga-group/omiga/src/venue/shared/entities/market"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
)

// MarketDelete is the builder for deleting a Market entity.
type MarketDelete struct {
	config
	hooks    []Hook
	mutation *MarketMutation
}

// Where appends a list predicates to the MarketDelete builder.
func (md *MarketDelete) Where(ps ...predicate.Market) *MarketDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MarketDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MarketMutation](ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MarketDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MarketDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: market.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: market.FieldID,
			},
		},
	}
	_spec.Node.Schema = md.schemaConfig.Market
	ctx = internal.NewSchemaConfigContext(ctx, md.schemaConfig)
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MarketDeleteOne is the builder for deleting a single Market entity.
type MarketDeleteOne struct {
	md *MarketDelete
}

// Where appends a list predicates to the MarketDelete builder.
func (mdo *MarketDeleteOne) Where(ps ...predicate.Market) *MarketDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MarketDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{market.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MarketDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
