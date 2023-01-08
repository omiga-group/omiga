// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/venue/shared/entities/internal"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
	"github.com/omiga-group/omiga/src/venue/shared/entities/venue"
)

// VenueDelete is the builder for deleting a Venue entity.
type VenueDelete struct {
	config
	hooks    []Hook
	mutation *VenueMutation
}

// Where appends a list predicates to the VenueDelete builder.
func (vd *VenueDelete) Where(ps ...predicate.Venue) *VenueDelete {
	vd.mutation.Where(ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VenueDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, VenueMutation](ctx, vd.sqlExec, vd.mutation, vd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VenueDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VenueDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: venue.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: venue.FieldID,
			},
		},
	}
	_spec.Node.Schema = vd.schemaConfig.Venue
	ctx = internal.NewSchemaConfigContext(ctx, vd.schemaConfig)
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vd.mutation.done = true
	return affected, err
}

// VenueDeleteOne is the builder for deleting a single Venue entity.
type VenueDeleteOne struct {
	vd *VenueDelete
}

// Exec executes the deletion query.
func (vdo *VenueDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{venue.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VenueDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
