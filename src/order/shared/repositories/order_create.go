// Code generated by ent, DO NOT EDIT.

package repositories

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/order/shared/repositories/order"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("repositories: uninitialized hook (forgotten import repositories/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Order)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: order.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		}
	)
	_spec.Schema = oc.schemaConfig.Order
	_spec.OnConflict = oc.conflict
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Order.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (oc *OrderCreate) OnConflict(opts ...sql.ConflictOption) *OrderUpsertOne {
	oc.conflict = opts
	return &OrderUpsertOne{
		create: oc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (oc *OrderCreate) OnConflictColumns(columns ...string) *OrderUpsertOne {
	oc.conflict = append(oc.conflict, sql.ConflictColumns(columns...))
	return &OrderUpsertOne{
		create: oc,
	}
}

type (
	// OrderUpsertOne is the builder for "upsert"-ing
	//  one Order node.
	OrderUpsertOne struct {
		create *OrderCreate
	}

	// OrderUpsert is the "OnConflict" setter.
	OrderUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *OrderUpsertOne) UpdateNewValues() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Order.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *OrderUpsertOne) Ignore() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderUpsertOne) DoNothing() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderCreate.OnConflict
// documentation for more info.
func (u *OrderUpsertOne) Update(set func(*OrderUpsert)) *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *OrderUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("repositories: missing options for OrderCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
	conflict []sql.ConflictOption
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Order.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (ocb *OrderCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderUpsertBulk {
	ocb.conflict = opts
	return &OrderUpsertBulk{
		create: ocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ocb *OrderCreateBulk) OnConflictColumns(columns ...string) *OrderUpsertBulk {
	ocb.conflict = append(ocb.conflict, sql.ConflictColumns(columns...))
	return &OrderUpsertBulk{
		create: ocb,
	}
}

// OrderUpsertBulk is the builder for "upsert"-ing
// a bulk of Order nodes.
type OrderUpsertBulk struct {
	create *OrderCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *OrderUpsertBulk) UpdateNewValues() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *OrderUpsertBulk) Ignore() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderUpsertBulk) DoNothing() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderCreateBulk.OnConflict
// documentation for more info.
func (u *OrderUpsertBulk) Update(set func(*OrderUpsert)) *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *OrderUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("repositories: OnConflict was set for builder %d. Set it on the OrderCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("repositories: missing options for OrderCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
