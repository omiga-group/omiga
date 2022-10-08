// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/coin"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/internal"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/predicate"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/tradingpair"
)

// TradingPairUpdate is the builder for updating TradingPair entities.
type TradingPairUpdate struct {
	config
	hooks     []Hook
	mutation  *TradingPairMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TradingPairUpdate builder.
func (tpu *TradingPairUpdate) Where(ps ...predicate.TradingPair) *TradingPairUpdate {
	tpu.mutation.Where(ps...)
	return tpu
}

// SetSymbol sets the "symbol" field.
func (tpu *TradingPairUpdate) SetSymbol(s string) *TradingPairUpdate {
	tpu.mutation.SetSymbol(s)
	return tpu
}

// SetBasePrecision sets the "base_precision" field.
func (tpu *TradingPairUpdate) SetBasePrecision(i int) *TradingPairUpdate {
	tpu.mutation.ResetBasePrecision()
	tpu.mutation.SetBasePrecision(i)
	return tpu
}

// AddBasePrecision adds i to the "base_precision" field.
func (tpu *TradingPairUpdate) AddBasePrecision(i int) *TradingPairUpdate {
	tpu.mutation.AddBasePrecision(i)
	return tpu
}

// SetCounterPrecision sets the "counter_precision" field.
func (tpu *TradingPairUpdate) SetCounterPrecision(i int) *TradingPairUpdate {
	tpu.mutation.ResetCounterPrecision()
	tpu.mutation.SetCounterPrecision(i)
	return tpu
}

// AddCounterPrecision adds i to the "counter_precision" field.
func (tpu *TradingPairUpdate) AddCounterPrecision(i int) *TradingPairUpdate {
	tpu.mutation.AddCounterPrecision(i)
	return tpu
}

// SetExchangeID sets the "exchange" edge to the Exchange entity by ID.
func (tpu *TradingPairUpdate) SetExchangeID(id int) *TradingPairUpdate {
	tpu.mutation.SetExchangeID(id)
	return tpu
}

// SetExchange sets the "exchange" edge to the Exchange entity.
func (tpu *TradingPairUpdate) SetExchange(e *Exchange) *TradingPairUpdate {
	return tpu.SetExchangeID(e.ID)
}

// SetBaseID sets the "base" edge to the Coin entity by ID.
func (tpu *TradingPairUpdate) SetBaseID(id int) *TradingPairUpdate {
	tpu.mutation.SetBaseID(id)
	return tpu
}

// SetBase sets the "base" edge to the Coin entity.
func (tpu *TradingPairUpdate) SetBase(c *Coin) *TradingPairUpdate {
	return tpu.SetBaseID(c.ID)
}

// SetCounterID sets the "counter" edge to the Coin entity by ID.
func (tpu *TradingPairUpdate) SetCounterID(id int) *TradingPairUpdate {
	tpu.mutation.SetCounterID(id)
	return tpu
}

// SetCounter sets the "counter" edge to the Coin entity.
func (tpu *TradingPairUpdate) SetCounter(c *Coin) *TradingPairUpdate {
	return tpu.SetCounterID(c.ID)
}

// Mutation returns the TradingPairMutation object of the builder.
func (tpu *TradingPairUpdate) Mutation() *TradingPairMutation {
	return tpu.mutation
}

// ClearExchange clears the "exchange" edge to the Exchange entity.
func (tpu *TradingPairUpdate) ClearExchange() *TradingPairUpdate {
	tpu.mutation.ClearExchange()
	return tpu
}

// ClearBase clears the "base" edge to the Coin entity.
func (tpu *TradingPairUpdate) ClearBase() *TradingPairUpdate {
	tpu.mutation.ClearBase()
	return tpu
}

// ClearCounter clears the "counter" edge to the Coin entity.
func (tpu *TradingPairUpdate) ClearCounter() *TradingPairUpdate {
	tpu.mutation.ClearCounter()
	return tpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tpu *TradingPairUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tpu.hooks) == 0 {
		if err = tpu.check(); err != nil {
			return 0, err
		}
		affected, err = tpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TradingPairMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tpu.check(); err != nil {
				return 0, err
			}
			tpu.mutation = mutation
			affected, err = tpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tpu.hooks) - 1; i >= 0; i-- {
			if tpu.hooks[i] == nil {
				return 0, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = tpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tpu *TradingPairUpdate) SaveX(ctx context.Context) int {
	affected, err := tpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tpu *TradingPairUpdate) Exec(ctx context.Context) error {
	_, err := tpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpu *TradingPairUpdate) ExecX(ctx context.Context) {
	if err := tpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpu *TradingPairUpdate) check() error {
	if _, ok := tpu.mutation.ExchangeID(); tpu.mutation.ExchangeCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "TradingPair.exchange"`)
	}
	if _, ok := tpu.mutation.BaseID(); tpu.mutation.BaseCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "TradingPair.base"`)
	}
	if _, ok := tpu.mutation.CounterID(); tpu.mutation.CounterCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "TradingPair.counter"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tpu *TradingPairUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TradingPairUpdate {
	tpu.modifiers = append(tpu.modifiers, modifiers...)
	return tpu
}

func (tpu *TradingPairUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tradingpair.Table,
			Columns: tradingpair.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tradingpair.FieldID,
			},
		},
	}
	if ps := tpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpu.mutation.Symbol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tradingpair.FieldSymbol,
		})
	}
	if value, ok := tpu.mutation.BasePrecision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tradingpair.FieldBasePrecision,
		})
	}
	if value, ok := tpu.mutation.AddedBasePrecision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tradingpair.FieldBasePrecision,
		})
	}
	if value, ok := tpu.mutation.CounterPrecision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tradingpair.FieldCounterPrecision,
		})
	}
	if value, ok := tpu.mutation.AddedCounterPrecision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tradingpair.FieldCounterPrecision,
		})
	}
	if tpu.mutation.ExchangeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.ExchangeTable,
			Columns: []string{tradingpair.ExchangeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: exchange.FieldID,
				},
			},
		}
		edge.Schema = tpu.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpu.mutation.ExchangeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.ExchangeTable,
			Columns: []string{tradingpair.ExchangeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: exchange.FieldID,
				},
			},
		}
		edge.Schema = tpu.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tpu.mutation.BaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.BaseTable,
			Columns: []string{tradingpair.BaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		edge.Schema = tpu.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpu.mutation.BaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.BaseTable,
			Columns: []string{tradingpair.BaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		edge.Schema = tpu.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tpu.mutation.CounterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.CounterTable,
			Columns: []string{tradingpair.CounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		edge.Schema = tpu.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpu.mutation.CounterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.CounterTable,
			Columns: []string{tradingpair.CounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		edge.Schema = tpu.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tpu.schemaConfig.TradingPair
	ctx = internal.NewSchemaConfigContext(ctx, tpu.schemaConfig)
	_spec.Modifiers = tpu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, tpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tradingpair.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TradingPairUpdateOne is the builder for updating a single TradingPair entity.
type TradingPairUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TradingPairMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetSymbol sets the "symbol" field.
func (tpuo *TradingPairUpdateOne) SetSymbol(s string) *TradingPairUpdateOne {
	tpuo.mutation.SetSymbol(s)
	return tpuo
}

// SetBasePrecision sets the "base_precision" field.
func (tpuo *TradingPairUpdateOne) SetBasePrecision(i int) *TradingPairUpdateOne {
	tpuo.mutation.ResetBasePrecision()
	tpuo.mutation.SetBasePrecision(i)
	return tpuo
}

// AddBasePrecision adds i to the "base_precision" field.
func (tpuo *TradingPairUpdateOne) AddBasePrecision(i int) *TradingPairUpdateOne {
	tpuo.mutation.AddBasePrecision(i)
	return tpuo
}

// SetCounterPrecision sets the "counter_precision" field.
func (tpuo *TradingPairUpdateOne) SetCounterPrecision(i int) *TradingPairUpdateOne {
	tpuo.mutation.ResetCounterPrecision()
	tpuo.mutation.SetCounterPrecision(i)
	return tpuo
}

// AddCounterPrecision adds i to the "counter_precision" field.
func (tpuo *TradingPairUpdateOne) AddCounterPrecision(i int) *TradingPairUpdateOne {
	tpuo.mutation.AddCounterPrecision(i)
	return tpuo
}

// SetExchangeID sets the "exchange" edge to the Exchange entity by ID.
func (tpuo *TradingPairUpdateOne) SetExchangeID(id int) *TradingPairUpdateOne {
	tpuo.mutation.SetExchangeID(id)
	return tpuo
}

// SetExchange sets the "exchange" edge to the Exchange entity.
func (tpuo *TradingPairUpdateOne) SetExchange(e *Exchange) *TradingPairUpdateOne {
	return tpuo.SetExchangeID(e.ID)
}

// SetBaseID sets the "base" edge to the Coin entity by ID.
func (tpuo *TradingPairUpdateOne) SetBaseID(id int) *TradingPairUpdateOne {
	tpuo.mutation.SetBaseID(id)
	return tpuo
}

// SetBase sets the "base" edge to the Coin entity.
func (tpuo *TradingPairUpdateOne) SetBase(c *Coin) *TradingPairUpdateOne {
	return tpuo.SetBaseID(c.ID)
}

// SetCounterID sets the "counter" edge to the Coin entity by ID.
func (tpuo *TradingPairUpdateOne) SetCounterID(id int) *TradingPairUpdateOne {
	tpuo.mutation.SetCounterID(id)
	return tpuo
}

// SetCounter sets the "counter" edge to the Coin entity.
func (tpuo *TradingPairUpdateOne) SetCounter(c *Coin) *TradingPairUpdateOne {
	return tpuo.SetCounterID(c.ID)
}

// Mutation returns the TradingPairMutation object of the builder.
func (tpuo *TradingPairUpdateOne) Mutation() *TradingPairMutation {
	return tpuo.mutation
}

// ClearExchange clears the "exchange" edge to the Exchange entity.
func (tpuo *TradingPairUpdateOne) ClearExchange() *TradingPairUpdateOne {
	tpuo.mutation.ClearExchange()
	return tpuo
}

// ClearBase clears the "base" edge to the Coin entity.
func (tpuo *TradingPairUpdateOne) ClearBase() *TradingPairUpdateOne {
	tpuo.mutation.ClearBase()
	return tpuo
}

// ClearCounter clears the "counter" edge to the Coin entity.
func (tpuo *TradingPairUpdateOne) ClearCounter() *TradingPairUpdateOne {
	tpuo.mutation.ClearCounter()
	return tpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tpuo *TradingPairUpdateOne) Select(field string, fields ...string) *TradingPairUpdateOne {
	tpuo.fields = append([]string{field}, fields...)
	return tpuo
}

// Save executes the query and returns the updated TradingPair entity.
func (tpuo *TradingPairUpdateOne) Save(ctx context.Context) (*TradingPair, error) {
	var (
		err  error
		node *TradingPair
	)
	if len(tpuo.hooks) == 0 {
		if err = tpuo.check(); err != nil {
			return nil, err
		}
		node, err = tpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TradingPairMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tpuo.check(); err != nil {
				return nil, err
			}
			tpuo.mutation = mutation
			node, err = tpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tpuo.hooks) - 1; i >= 0; i-- {
			if tpuo.hooks[i] == nil {
				return nil, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = tpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TradingPair)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TradingPairMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tpuo *TradingPairUpdateOne) SaveX(ctx context.Context) *TradingPair {
	node, err := tpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tpuo *TradingPairUpdateOne) Exec(ctx context.Context) error {
	_, err := tpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpuo *TradingPairUpdateOne) ExecX(ctx context.Context) {
	if err := tpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpuo *TradingPairUpdateOne) check() error {
	if _, ok := tpuo.mutation.ExchangeID(); tpuo.mutation.ExchangeCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "TradingPair.exchange"`)
	}
	if _, ok := tpuo.mutation.BaseID(); tpuo.mutation.BaseCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "TradingPair.base"`)
	}
	if _, ok := tpuo.mutation.CounterID(); tpuo.mutation.CounterCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "TradingPair.counter"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tpuo *TradingPairUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TradingPairUpdateOne {
	tpuo.modifiers = append(tpuo.modifiers, modifiers...)
	return tpuo
}

func (tpuo *TradingPairUpdateOne) sqlSave(ctx context.Context) (_node *TradingPair, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tradingpair.Table,
			Columns: tradingpair.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tradingpair.FieldID,
			},
		},
	}
	id, ok := tpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "TradingPair.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tradingpair.FieldID)
		for _, f := range fields {
			if !tradingpair.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != tradingpair.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpuo.mutation.Symbol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tradingpair.FieldSymbol,
		})
	}
	if value, ok := tpuo.mutation.BasePrecision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tradingpair.FieldBasePrecision,
		})
	}
	if value, ok := tpuo.mutation.AddedBasePrecision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tradingpair.FieldBasePrecision,
		})
	}
	if value, ok := tpuo.mutation.CounterPrecision(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tradingpair.FieldCounterPrecision,
		})
	}
	if value, ok := tpuo.mutation.AddedCounterPrecision(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tradingpair.FieldCounterPrecision,
		})
	}
	if tpuo.mutation.ExchangeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.ExchangeTable,
			Columns: []string{tradingpair.ExchangeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: exchange.FieldID,
				},
			},
		}
		edge.Schema = tpuo.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpuo.mutation.ExchangeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.ExchangeTable,
			Columns: []string{tradingpair.ExchangeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: exchange.FieldID,
				},
			},
		}
		edge.Schema = tpuo.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tpuo.mutation.BaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.BaseTable,
			Columns: []string{tradingpair.BaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		edge.Schema = tpuo.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpuo.mutation.BaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.BaseTable,
			Columns: []string{tradingpair.BaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		edge.Schema = tpuo.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tpuo.mutation.CounterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.CounterTable,
			Columns: []string{tradingpair.CounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		edge.Schema = tpuo.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpuo.mutation.CounterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tradingpair.CounterTable,
			Columns: []string{tradingpair.CounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coin.FieldID,
				},
			},
		}
		edge.Schema = tpuo.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tpuo.schemaConfig.TradingPair
	ctx = internal.NewSchemaConfigContext(ctx, tpuo.schemaConfig)
	_spec.Modifiers = tpuo.modifiers
	_node = &TradingPair{config: tpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tradingpair.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
