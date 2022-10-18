// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/venue/shared/entities/internal"
	"github.com/omiga-group/omiga/src/venue/shared/entities/market"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
	"github.com/omiga-group/omiga/src/venue/shared/entities/tradingpair"
	"github.com/omiga-group/omiga/src/venue/shared/entities/venue"
)

// MarketUpdate is the builder for updating Market entities.
type MarketUpdate struct {
	config
	hooks     []Hook
	mutation  *MarketMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MarketUpdate builder.
func (mu *MarketUpdate) Where(ps ...predicate.Market) *MarketUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MarketUpdate) SetName(s string) *MarketUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetType sets the "type" field.
func (mu *MarketUpdate) SetType(m market.Type) *MarketUpdate {
	mu.mutation.SetType(m)
	return mu
}

// SetVenueID sets the "venue" edge to the Venue entity by ID.
func (mu *MarketUpdate) SetVenueID(id int) *MarketUpdate {
	mu.mutation.SetVenueID(id)
	return mu
}

// SetVenue sets the "venue" edge to the Venue entity.
func (mu *MarketUpdate) SetVenue(v *Venue) *MarketUpdate {
	return mu.SetVenueID(v.ID)
}

// AddTradingPairIDs adds the "trading_pair" edge to the TradingPair entity by IDs.
func (mu *MarketUpdate) AddTradingPairIDs(ids ...int) *MarketUpdate {
	mu.mutation.AddTradingPairIDs(ids...)
	return mu
}

// AddTradingPair adds the "trading_pair" edges to the TradingPair entity.
func (mu *MarketUpdate) AddTradingPair(t ...*TradingPair) *MarketUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.AddTradingPairIDs(ids...)
}

// Mutation returns the MarketMutation object of the builder.
func (mu *MarketUpdate) Mutation() *MarketMutation {
	return mu.mutation
}

// ClearVenue clears the "venue" edge to the Venue entity.
func (mu *MarketUpdate) ClearVenue() *MarketUpdate {
	mu.mutation.ClearVenue()
	return mu
}

// ClearTradingPair clears all "trading_pair" edges to the TradingPair entity.
func (mu *MarketUpdate) ClearTradingPair() *MarketUpdate {
	mu.mutation.ClearTradingPair()
	return mu
}

// RemoveTradingPairIDs removes the "trading_pair" edge to TradingPair entities by IDs.
func (mu *MarketUpdate) RemoveTradingPairIDs(ids ...int) *MarketUpdate {
	mu.mutation.RemoveTradingPairIDs(ids...)
	return mu
}

// RemoveTradingPair removes "trading_pair" edges to TradingPair entities.
func (mu *MarketUpdate) RemoveTradingPair(t ...*TradingPair) *MarketUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.RemoveTradingPairIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MarketUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MarketMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MarketUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MarketUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MarketUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MarketUpdate) check() error {
	if v, ok := mu.mutation.GetType(); ok {
		if err := market.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`entities: validator failed for field "Market.type": %w`, err)}
		}
	}
	if _, ok := mu.mutation.VenueID(); mu.mutation.VenueCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Market.venue"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MarketUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MarketUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MarketUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   market.Table,
			Columns: market.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: market.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: market.FieldName,
		})
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: market.FieldType,
		})
	}
	if mu.mutation.VenueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   market.VenueTable,
			Columns: []string{market.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: venue.FieldID,
				},
			},
		}
		edge.Schema = mu.schemaConfig.Market
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.VenueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   market.VenueTable,
			Columns: []string{market.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: venue.FieldID,
				},
			},
		}
		edge.Schema = mu.schemaConfig.Market
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.TradingPairCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   market.TradingPairTable,
			Columns: market.TradingPairPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = mu.schemaConfig.MarketTradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedTradingPairIDs(); len(nodes) > 0 && !mu.mutation.TradingPairCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   market.TradingPairTable,
			Columns: market.TradingPairPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = mu.schemaConfig.MarketTradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.TradingPairIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   market.TradingPairTable,
			Columns: market.TradingPairPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = mu.schemaConfig.MarketTradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = mu.schemaConfig.Market
	ctx = internal.NewSchemaConfigContext(ctx, mu.schemaConfig)
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{market.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MarketUpdateOne is the builder for updating a single Market entity.
type MarketUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MarketMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (muo *MarketUpdateOne) SetName(s string) *MarketUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetType sets the "type" field.
func (muo *MarketUpdateOne) SetType(m market.Type) *MarketUpdateOne {
	muo.mutation.SetType(m)
	return muo
}

// SetVenueID sets the "venue" edge to the Venue entity by ID.
func (muo *MarketUpdateOne) SetVenueID(id int) *MarketUpdateOne {
	muo.mutation.SetVenueID(id)
	return muo
}

// SetVenue sets the "venue" edge to the Venue entity.
func (muo *MarketUpdateOne) SetVenue(v *Venue) *MarketUpdateOne {
	return muo.SetVenueID(v.ID)
}

// AddTradingPairIDs adds the "trading_pair" edge to the TradingPair entity by IDs.
func (muo *MarketUpdateOne) AddTradingPairIDs(ids ...int) *MarketUpdateOne {
	muo.mutation.AddTradingPairIDs(ids...)
	return muo
}

// AddTradingPair adds the "trading_pair" edges to the TradingPair entity.
func (muo *MarketUpdateOne) AddTradingPair(t ...*TradingPair) *MarketUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.AddTradingPairIDs(ids...)
}

// Mutation returns the MarketMutation object of the builder.
func (muo *MarketUpdateOne) Mutation() *MarketMutation {
	return muo.mutation
}

// ClearVenue clears the "venue" edge to the Venue entity.
func (muo *MarketUpdateOne) ClearVenue() *MarketUpdateOne {
	muo.mutation.ClearVenue()
	return muo
}

// ClearTradingPair clears all "trading_pair" edges to the TradingPair entity.
func (muo *MarketUpdateOne) ClearTradingPair() *MarketUpdateOne {
	muo.mutation.ClearTradingPair()
	return muo
}

// RemoveTradingPairIDs removes the "trading_pair" edge to TradingPair entities by IDs.
func (muo *MarketUpdateOne) RemoveTradingPairIDs(ids ...int) *MarketUpdateOne {
	muo.mutation.RemoveTradingPairIDs(ids...)
	return muo
}

// RemoveTradingPair removes "trading_pair" edges to TradingPair entities.
func (muo *MarketUpdateOne) RemoveTradingPair(t ...*TradingPair) *MarketUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.RemoveTradingPairIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MarketUpdateOne) Select(field string, fields ...string) *MarketUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Market entity.
func (muo *MarketUpdateOne) Save(ctx context.Context) (*Market, error) {
	var (
		err  error
		node *Market
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MarketMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Market)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MarketMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MarketUpdateOne) SaveX(ctx context.Context) *Market {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MarketUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MarketUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MarketUpdateOne) check() error {
	if v, ok := muo.mutation.GetType(); ok {
		if err := market.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`entities: validator failed for field "Market.type": %w`, err)}
		}
	}
	if _, ok := muo.mutation.VenueID(); muo.mutation.VenueCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "Market.venue"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MarketUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MarketUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MarketUpdateOne) sqlSave(ctx context.Context) (_node *Market, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   market.Table,
			Columns: market.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: market.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Market.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, market.FieldID)
		for _, f := range fields {
			if !market.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != market.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: market.FieldName,
		})
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: market.FieldType,
		})
	}
	if muo.mutation.VenueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   market.VenueTable,
			Columns: []string{market.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: venue.FieldID,
				},
			},
		}
		edge.Schema = muo.schemaConfig.Market
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.VenueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   market.VenueTable,
			Columns: []string{market.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: venue.FieldID,
				},
			},
		}
		edge.Schema = muo.schemaConfig.Market
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.TradingPairCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   market.TradingPairTable,
			Columns: market.TradingPairPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = muo.schemaConfig.MarketTradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedTradingPairIDs(); len(nodes) > 0 && !muo.mutation.TradingPairCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   market.TradingPairTable,
			Columns: market.TradingPairPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = muo.schemaConfig.MarketTradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.TradingPairIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   market.TradingPairTable,
			Columns: market.TradingPairPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = muo.schemaConfig.MarketTradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = muo.schemaConfig.Market
	ctx = internal.NewSchemaConfigContext(ctx, muo.schemaConfig)
	_spec.AddModifiers(muo.modifiers...)
	_node = &Market{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{market.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
