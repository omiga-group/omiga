// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/internal"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/outbox"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/predicate"
)

// OutboxUpdate is the builder for updating Outbox entities.
type OutboxUpdate struct {
	config
	hooks     []Hook
	mutation  *OutboxMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OutboxUpdate builder.
func (ou *OutboxUpdate) Where(ps ...predicate.Outbox) *OutboxUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetTimestamp sets the "timestamp" field.
func (ou *OutboxUpdate) SetTimestamp(t time.Time) *OutboxUpdate {
	ou.mutation.SetTimestamp(t)
	return ou
}

// SetTopic sets the "topic" field.
func (ou *OutboxUpdate) SetTopic(s string) *OutboxUpdate {
	ou.mutation.SetTopic(s)
	return ou
}

// SetKey sets the "key" field.
func (ou *OutboxUpdate) SetKey(s string) *OutboxUpdate {
	ou.mutation.SetKey(s)
	return ou
}

// SetPayload sets the "payload" field.
func (ou *OutboxUpdate) SetPayload(b []byte) *OutboxUpdate {
	ou.mutation.SetPayload(b)
	return ou
}

// SetHeaders sets the "headers" field.
func (ou *OutboxUpdate) SetHeaders(m map[string]string) *OutboxUpdate {
	ou.mutation.SetHeaders(m)
	return ou
}

// SetRetryCount sets the "retry_count" field.
func (ou *OutboxUpdate) SetRetryCount(i int) *OutboxUpdate {
	ou.mutation.ResetRetryCount()
	ou.mutation.SetRetryCount(i)
	return ou
}

// AddRetryCount adds i to the "retry_count" field.
func (ou *OutboxUpdate) AddRetryCount(i int) *OutboxUpdate {
	ou.mutation.AddRetryCount(i)
	return ou
}

// SetStatus sets the "status" field.
func (ou *OutboxUpdate) SetStatus(o outbox.Status) *OutboxUpdate {
	ou.mutation.SetStatus(o)
	return ou
}

// SetLastRetry sets the "last_retry" field.
func (ou *OutboxUpdate) SetLastRetry(t time.Time) *OutboxUpdate {
	ou.mutation.SetLastRetry(t)
	return ou
}

// SetNillableLastRetry sets the "last_retry" field if the given value is not nil.
func (ou *OutboxUpdate) SetNillableLastRetry(t *time.Time) *OutboxUpdate {
	if t != nil {
		ou.SetLastRetry(*t)
	}
	return ou
}

// ClearLastRetry clears the value of the "last_retry" field.
func (ou *OutboxUpdate) ClearLastRetry() *OutboxUpdate {
	ou.mutation.ClearLastRetry()
	return ou
}

// SetProcessingErrors sets the "processing_errors" field.
func (ou *OutboxUpdate) SetProcessingErrors(s []string) *OutboxUpdate {
	ou.mutation.SetProcessingErrors(s)
	return ou
}

// ClearProcessingErrors clears the value of the "processing_errors" field.
func (ou *OutboxUpdate) ClearProcessingErrors() *OutboxUpdate {
	ou.mutation.ClearProcessingErrors()
	return ou
}

// Mutation returns the OutboxMutation object of the builder.
func (ou *OutboxUpdate) Mutation() *OutboxMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OutboxUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutboxMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OutboxUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OutboxUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OutboxUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OutboxUpdate) check() error {
	if v, ok := ou.mutation.Status(); ok {
		if err := outbox.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entities: validator failed for field "Outbox.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ou *OutboxUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OutboxUpdate {
	ou.modifiers = append(ou.modifiers, modifiers...)
	return ou
}

func (ou *OutboxUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outbox.Table,
			Columns: outbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outbox.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: outbox.FieldTimestamp,
		})
	}
	if value, ok := ou.mutation.Topic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outbox.FieldTopic,
		})
	}
	if value, ok := ou.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outbox.FieldKey,
		})
	}
	if value, ok := ou.mutation.Payload(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: outbox.FieldPayload,
		})
	}
	if value, ok := ou.mutation.Headers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: outbox.FieldHeaders,
		})
	}
	if value, ok := ou.mutation.RetryCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: outbox.FieldRetryCount,
		})
	}
	if value, ok := ou.mutation.AddedRetryCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: outbox.FieldRetryCount,
		})
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: outbox.FieldStatus,
		})
	}
	if value, ok := ou.mutation.LastRetry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: outbox.FieldLastRetry,
		})
	}
	if ou.mutation.LastRetryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: outbox.FieldLastRetry,
		})
	}
	if value, ok := ou.mutation.ProcessingErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: outbox.FieldProcessingErrors,
		})
	}
	if ou.mutation.ProcessingErrorsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: outbox.FieldProcessingErrors,
		})
	}
	_spec.Node.Schema = ou.schemaConfig.Outbox
	ctx = internal.NewSchemaConfigContext(ctx, ou.schemaConfig)
	_spec.Modifiers = ou.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outbox.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OutboxUpdateOne is the builder for updating a single Outbox entity.
type OutboxUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OutboxMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTimestamp sets the "timestamp" field.
func (ouo *OutboxUpdateOne) SetTimestamp(t time.Time) *OutboxUpdateOne {
	ouo.mutation.SetTimestamp(t)
	return ouo
}

// SetTopic sets the "topic" field.
func (ouo *OutboxUpdateOne) SetTopic(s string) *OutboxUpdateOne {
	ouo.mutation.SetTopic(s)
	return ouo
}

// SetKey sets the "key" field.
func (ouo *OutboxUpdateOne) SetKey(s string) *OutboxUpdateOne {
	ouo.mutation.SetKey(s)
	return ouo
}

// SetPayload sets the "payload" field.
func (ouo *OutboxUpdateOne) SetPayload(b []byte) *OutboxUpdateOne {
	ouo.mutation.SetPayload(b)
	return ouo
}

// SetHeaders sets the "headers" field.
func (ouo *OutboxUpdateOne) SetHeaders(m map[string]string) *OutboxUpdateOne {
	ouo.mutation.SetHeaders(m)
	return ouo
}

// SetRetryCount sets the "retry_count" field.
func (ouo *OutboxUpdateOne) SetRetryCount(i int) *OutboxUpdateOne {
	ouo.mutation.ResetRetryCount()
	ouo.mutation.SetRetryCount(i)
	return ouo
}

// AddRetryCount adds i to the "retry_count" field.
func (ouo *OutboxUpdateOne) AddRetryCount(i int) *OutboxUpdateOne {
	ouo.mutation.AddRetryCount(i)
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OutboxUpdateOne) SetStatus(o outbox.Status) *OutboxUpdateOne {
	ouo.mutation.SetStatus(o)
	return ouo
}

// SetLastRetry sets the "last_retry" field.
func (ouo *OutboxUpdateOne) SetLastRetry(t time.Time) *OutboxUpdateOne {
	ouo.mutation.SetLastRetry(t)
	return ouo
}

// SetNillableLastRetry sets the "last_retry" field if the given value is not nil.
func (ouo *OutboxUpdateOne) SetNillableLastRetry(t *time.Time) *OutboxUpdateOne {
	if t != nil {
		ouo.SetLastRetry(*t)
	}
	return ouo
}

// ClearLastRetry clears the value of the "last_retry" field.
func (ouo *OutboxUpdateOne) ClearLastRetry() *OutboxUpdateOne {
	ouo.mutation.ClearLastRetry()
	return ouo
}

// SetProcessingErrors sets the "processing_errors" field.
func (ouo *OutboxUpdateOne) SetProcessingErrors(s []string) *OutboxUpdateOne {
	ouo.mutation.SetProcessingErrors(s)
	return ouo
}

// ClearProcessingErrors clears the value of the "processing_errors" field.
func (ouo *OutboxUpdateOne) ClearProcessingErrors() *OutboxUpdateOne {
	ouo.mutation.ClearProcessingErrors()
	return ouo
}

// Mutation returns the OutboxMutation object of the builder.
func (ouo *OutboxUpdateOne) Mutation() *OutboxMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OutboxUpdateOne) Select(field string, fields ...string) *OutboxUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Outbox entity.
func (ouo *OutboxUpdateOne) Save(ctx context.Context) (*Outbox, error) {
	var (
		err  error
		node *Outbox
	)
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutboxMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("entities: uninitialized hook (forgotten import entities/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Outbox)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OutboxMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OutboxUpdateOne) SaveX(ctx context.Context) *Outbox {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OutboxUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OutboxUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OutboxUpdateOne) check() error {
	if v, ok := ouo.mutation.Status(); ok {
		if err := outbox.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entities: validator failed for field "Outbox.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ouo *OutboxUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OutboxUpdateOne {
	ouo.modifiers = append(ouo.modifiers, modifiers...)
	return ouo
}

func (ouo *OutboxUpdateOne) sqlSave(ctx context.Context) (_node *Outbox, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outbox.Table,
			Columns: outbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outbox.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "Outbox.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outbox.FieldID)
		for _, f := range fields {
			if !outbox.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != outbox.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: outbox.FieldTimestamp,
		})
	}
	if value, ok := ouo.mutation.Topic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outbox.FieldTopic,
		})
	}
	if value, ok := ouo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: outbox.FieldKey,
		})
	}
	if value, ok := ouo.mutation.Payload(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: outbox.FieldPayload,
		})
	}
	if value, ok := ouo.mutation.Headers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: outbox.FieldHeaders,
		})
	}
	if value, ok := ouo.mutation.RetryCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: outbox.FieldRetryCount,
		})
	}
	if value, ok := ouo.mutation.AddedRetryCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: outbox.FieldRetryCount,
		})
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: outbox.FieldStatus,
		})
	}
	if value, ok := ouo.mutation.LastRetry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: outbox.FieldLastRetry,
		})
	}
	if ouo.mutation.LastRetryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: outbox.FieldLastRetry,
		})
	}
	if value, ok := ouo.mutation.ProcessingErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: outbox.FieldProcessingErrors,
		})
	}
	if ouo.mutation.ProcessingErrorsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: outbox.FieldProcessingErrors,
		})
	}
	_spec.Node.Schema = ouo.schemaConfig.Outbox
	ctx = internal.NewSchemaConfigContext(ctx, ouo.schemaConfig)
	_spec.Modifiers = ouo.modifiers
	_node = &Outbox{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outbox.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}