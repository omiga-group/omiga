// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/venue/shared/entities/currency"
	"github.com/omiga-group/omiga/src/venue/shared/entities/internal"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
	"github.com/omiga-group/omiga/src/venue/shared/entities/tradingpair"
)

// CurrencyUpdate is the builder for updating Currency entities.
type CurrencyUpdate struct {
	config
	hooks     []Hook
	mutation  *CurrencyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CurrencyUpdate builder.
func (cu *CurrencyUpdate) Where(ps ...predicate.Currency) *CurrencyUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetSymbol sets the "symbol" field.
func (cu *CurrencyUpdate) SetSymbol(s string) *CurrencyUpdate {
	cu.mutation.SetSymbol(s)
	return cu
}

// SetName sets the "name" field.
func (cu *CurrencyUpdate) SetName(s string) *CurrencyUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CurrencyUpdate) SetNillableName(s *string) *CurrencyUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *CurrencyUpdate) ClearName() *CurrencyUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetType sets the "type" field.
func (cu *CurrencyUpdate) SetType(c currency.Type) *CurrencyUpdate {
	cu.mutation.SetType(c)
	return cu
}

// AddCurrencyBaseIDs adds the "currency_base" edge to the TradingPair entity by IDs.
func (cu *CurrencyUpdate) AddCurrencyBaseIDs(ids ...int) *CurrencyUpdate {
	cu.mutation.AddCurrencyBaseIDs(ids...)
	return cu
}

// AddCurrencyBase adds the "currency_base" edges to the TradingPair entity.
func (cu *CurrencyUpdate) AddCurrencyBase(t ...*TradingPair) *CurrencyUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddCurrencyBaseIDs(ids...)
}

// AddCurrencyCounterIDs adds the "currency_counter" edge to the TradingPair entity by IDs.
func (cu *CurrencyUpdate) AddCurrencyCounterIDs(ids ...int) *CurrencyUpdate {
	cu.mutation.AddCurrencyCounterIDs(ids...)
	return cu
}

// AddCurrencyCounter adds the "currency_counter" edges to the TradingPair entity.
func (cu *CurrencyUpdate) AddCurrencyCounter(t ...*TradingPair) *CurrencyUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddCurrencyCounterIDs(ids...)
}

// Mutation returns the CurrencyMutation object of the builder.
func (cu *CurrencyUpdate) Mutation() *CurrencyMutation {
	return cu.mutation
}

// ClearCurrencyBase clears all "currency_base" edges to the TradingPair entity.
func (cu *CurrencyUpdate) ClearCurrencyBase() *CurrencyUpdate {
	cu.mutation.ClearCurrencyBase()
	return cu
}

// RemoveCurrencyBaseIDs removes the "currency_base" edge to TradingPair entities by IDs.
func (cu *CurrencyUpdate) RemoveCurrencyBaseIDs(ids ...int) *CurrencyUpdate {
	cu.mutation.RemoveCurrencyBaseIDs(ids...)
	return cu
}

// RemoveCurrencyBase removes "currency_base" edges to TradingPair entities.
func (cu *CurrencyUpdate) RemoveCurrencyBase(t ...*TradingPair) *CurrencyUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveCurrencyBaseIDs(ids...)
}

// ClearCurrencyCounter clears all "currency_counter" edges to the TradingPair entity.
func (cu *CurrencyUpdate) ClearCurrencyCounter() *CurrencyUpdate {
	cu.mutation.ClearCurrencyCounter()
	return cu
}

// RemoveCurrencyCounterIDs removes the "currency_counter" edge to TradingPair entities by IDs.
func (cu *CurrencyUpdate) RemoveCurrencyCounterIDs(ids ...int) *CurrencyUpdate {
	cu.mutation.RemoveCurrencyCounterIDs(ids...)
	return cu
}

// RemoveCurrencyCounter removes "currency_counter" edges to TradingPair entities.
func (cu *CurrencyUpdate) RemoveCurrencyCounter(t ...*TradingPair) *CurrencyUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveCurrencyCounterIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CurrencyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CurrencyMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CurrencyUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CurrencyUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CurrencyUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CurrencyUpdate) check() error {
	if v, ok := cu.mutation.GetType(); ok {
		if err := currency.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`entities: validator failed for field "Currency.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CurrencyUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CurrencyUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CurrencyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   currency.Table,
			Columns: currency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: currency.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Symbol(); ok {
		_spec.SetField(currency.FieldSymbol, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(currency.FieldName, field.TypeString, value)
	}
	if cu.mutation.NameCleared() {
		_spec.ClearField(currency.FieldName, field.TypeString)
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(currency.FieldType, field.TypeEnum, value)
	}
	if cu.mutation.CurrencyBaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyBaseTable,
			Columns: []string{currency.CurrencyBaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCurrencyBaseIDs(); len(nodes) > 0 && !cu.mutation.CurrencyBaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyBaseTable,
			Columns: []string{currency.CurrencyBaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CurrencyBaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyBaseTable,
			Columns: []string{currency.CurrencyBaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CurrencyCounterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyCounterTable,
			Columns: []string{currency.CurrencyCounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCurrencyCounterIDs(); len(nodes) > 0 && !cu.mutation.CurrencyCounterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyCounterTable,
			Columns: []string{currency.CurrencyCounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CurrencyCounterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyCounterTable,
			Columns: []string{currency.CurrencyCounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cu.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = cu.schemaConfig.Currency
	ctx = internal.NewSchemaConfigContext(ctx, cu.schemaConfig)
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{currency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CurrencyUpdateOne is the builder for updating a single Currency entity.
type CurrencyUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CurrencyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetSymbol sets the "symbol" field.
func (cuo *CurrencyUpdateOne) SetSymbol(s string) *CurrencyUpdateOne {
	cuo.mutation.SetSymbol(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CurrencyUpdateOne) SetName(s string) *CurrencyUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CurrencyUpdateOne) SetNillableName(s *string) *CurrencyUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *CurrencyUpdateOne) ClearName() *CurrencyUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetType sets the "type" field.
func (cuo *CurrencyUpdateOne) SetType(c currency.Type) *CurrencyUpdateOne {
	cuo.mutation.SetType(c)
	return cuo
}

// AddCurrencyBaseIDs adds the "currency_base" edge to the TradingPair entity by IDs.
func (cuo *CurrencyUpdateOne) AddCurrencyBaseIDs(ids ...int) *CurrencyUpdateOne {
	cuo.mutation.AddCurrencyBaseIDs(ids...)
	return cuo
}

// AddCurrencyBase adds the "currency_base" edges to the TradingPair entity.
func (cuo *CurrencyUpdateOne) AddCurrencyBase(t ...*TradingPair) *CurrencyUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddCurrencyBaseIDs(ids...)
}

// AddCurrencyCounterIDs adds the "currency_counter" edge to the TradingPair entity by IDs.
func (cuo *CurrencyUpdateOne) AddCurrencyCounterIDs(ids ...int) *CurrencyUpdateOne {
	cuo.mutation.AddCurrencyCounterIDs(ids...)
	return cuo
}

// AddCurrencyCounter adds the "currency_counter" edges to the TradingPair entity.
func (cuo *CurrencyUpdateOne) AddCurrencyCounter(t ...*TradingPair) *CurrencyUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddCurrencyCounterIDs(ids...)
}

// Mutation returns the CurrencyMutation object of the builder.
func (cuo *CurrencyUpdateOne) Mutation() *CurrencyMutation {
	return cuo.mutation
}

// ClearCurrencyBase clears all "currency_base" edges to the TradingPair entity.
func (cuo *CurrencyUpdateOne) ClearCurrencyBase() *CurrencyUpdateOne {
	cuo.mutation.ClearCurrencyBase()
	return cuo
}

// RemoveCurrencyBaseIDs removes the "currency_base" edge to TradingPair entities by IDs.
func (cuo *CurrencyUpdateOne) RemoveCurrencyBaseIDs(ids ...int) *CurrencyUpdateOne {
	cuo.mutation.RemoveCurrencyBaseIDs(ids...)
	return cuo
}

// RemoveCurrencyBase removes "currency_base" edges to TradingPair entities.
func (cuo *CurrencyUpdateOne) RemoveCurrencyBase(t ...*TradingPair) *CurrencyUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveCurrencyBaseIDs(ids...)
}

// ClearCurrencyCounter clears all "currency_counter" edges to the TradingPair entity.
func (cuo *CurrencyUpdateOne) ClearCurrencyCounter() *CurrencyUpdateOne {
	cuo.mutation.ClearCurrencyCounter()
	return cuo
}

// RemoveCurrencyCounterIDs removes the "currency_counter" edge to TradingPair entities by IDs.
func (cuo *CurrencyUpdateOne) RemoveCurrencyCounterIDs(ids ...int) *CurrencyUpdateOne {
	cuo.mutation.RemoveCurrencyCounterIDs(ids...)
	return cuo
}

// RemoveCurrencyCounter removes "currency_counter" edges to TradingPair entities.
func (cuo *CurrencyUpdateOne) RemoveCurrencyCounter(t ...*TradingPair) *CurrencyUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveCurrencyCounterIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CurrencyUpdateOne) Select(field string, fields ...string) *CurrencyUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Currency entity.
func (cuo *CurrencyUpdateOne) Save(ctx context.Context) (*Currency, error) {
	return withHooks[*Currency, CurrencyMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CurrencyUpdateOne) SaveX(ctx context.Context) *Currency {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CurrencyUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CurrencyUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CurrencyUpdateOne) check() error {
	if v, ok := cuo.mutation.GetType(); ok {
		if err := currency.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`entities: validator failed for field "Currency.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CurrencyUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CurrencyUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CurrencyUpdateOne) sqlSave(ctx context.Context) (_node *Currency, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   currency.Table,
			Columns: currency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: currency.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Currency.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, currency.FieldID)
		for _, f := range fields {
			if !currency.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != currency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Symbol(); ok {
		_spec.SetField(currency.FieldSymbol, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(currency.FieldName, field.TypeString, value)
	}
	if cuo.mutation.NameCleared() {
		_spec.ClearField(currency.FieldName, field.TypeString)
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(currency.FieldType, field.TypeEnum, value)
	}
	if cuo.mutation.CurrencyBaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyBaseTable,
			Columns: []string{currency.CurrencyBaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCurrencyBaseIDs(); len(nodes) > 0 && !cuo.mutation.CurrencyBaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyBaseTable,
			Columns: []string{currency.CurrencyBaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CurrencyBaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyBaseTable,
			Columns: []string{currency.CurrencyBaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CurrencyCounterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyCounterTable,
			Columns: []string{currency.CurrencyCounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.TradingPair
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCurrencyCounterIDs(); len(nodes) > 0 && !cuo.mutation.CurrencyCounterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyCounterTable,
			Columns: []string{currency.CurrencyCounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CurrencyCounterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   currency.CurrencyCounterTable,
			Columns: []string{currency.CurrencyCounterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tradingpair.FieldID,
				},
			},
		}
		edge.Schema = cuo.schemaConfig.TradingPair
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = cuo.schemaConfig.Currency
	ctx = internal.NewSchemaConfigContext(ctx, cuo.schemaConfig)
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Currency{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{currency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
