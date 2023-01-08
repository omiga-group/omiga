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
	"github.com/omiga-group/omiga/src/venue/shared/entities/outbox"
)

// OutboxCreate is the builder for creating a Outbox entity.
type OutboxCreate struct {
	config
	mutation *OutboxMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTimestamp sets the "timestamp" field.
func (oc *OutboxCreate) SetTimestamp(t time.Time) *OutboxCreate {
	oc.mutation.SetTimestamp(t)
	return oc
}

// SetTopic sets the "topic" field.
func (oc *OutboxCreate) SetTopic(s string) *OutboxCreate {
	oc.mutation.SetTopic(s)
	return oc
}

// SetKey sets the "key" field.
func (oc *OutboxCreate) SetKey(s string) *OutboxCreate {
	oc.mutation.SetKey(s)
	return oc
}

// SetPayload sets the "payload" field.
func (oc *OutboxCreate) SetPayload(b []byte) *OutboxCreate {
	oc.mutation.SetPayload(b)
	return oc
}

// SetHeaders sets the "headers" field.
func (oc *OutboxCreate) SetHeaders(m map[string]string) *OutboxCreate {
	oc.mutation.SetHeaders(m)
	return oc
}

// SetRetryCount sets the "retry_count" field.
func (oc *OutboxCreate) SetRetryCount(i int) *OutboxCreate {
	oc.mutation.SetRetryCount(i)
	return oc
}

// SetStatus sets the "status" field.
func (oc *OutboxCreate) SetStatus(o outbox.Status) *OutboxCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetLastRetry sets the "last_retry" field.
func (oc *OutboxCreate) SetLastRetry(t time.Time) *OutboxCreate {
	oc.mutation.SetLastRetry(t)
	return oc
}

// SetNillableLastRetry sets the "last_retry" field if the given value is not nil.
func (oc *OutboxCreate) SetNillableLastRetry(t *time.Time) *OutboxCreate {
	if t != nil {
		oc.SetLastRetry(*t)
	}
	return oc
}

// SetProcessingErrors sets the "processing_errors" field.
func (oc *OutboxCreate) SetProcessingErrors(s []string) *OutboxCreate {
	oc.mutation.SetProcessingErrors(s)
	return oc
}

// Mutation returns the OutboxMutation object of the builder.
func (oc *OutboxCreate) Mutation() *OutboxMutation {
	return oc.mutation
}

// Save creates the Outbox in the database.
func (oc *OutboxCreate) Save(ctx context.Context) (*Outbox, error) {
	return withHooks[*Outbox, OutboxMutation](ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OutboxCreate) SaveX(ctx context.Context) *Outbox {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OutboxCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OutboxCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OutboxCreate) check() error {
	if _, ok := oc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`entities: missing required field "Outbox.timestamp"`)}
	}
	if _, ok := oc.mutation.Topic(); !ok {
		return &ValidationError{Name: "topic", err: errors.New(`entities: missing required field "Outbox.topic"`)}
	}
	if _, ok := oc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`entities: missing required field "Outbox.key"`)}
	}
	if _, ok := oc.mutation.Payload(); !ok {
		return &ValidationError{Name: "payload", err: errors.New(`entities: missing required field "Outbox.payload"`)}
	}
	if _, ok := oc.mutation.Headers(); !ok {
		return &ValidationError{Name: "headers", err: errors.New(`entities: missing required field "Outbox.headers"`)}
	}
	if _, ok := oc.mutation.RetryCount(); !ok {
		return &ValidationError{Name: "retry_count", err: errors.New(`entities: missing required field "Outbox.retry_count"`)}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`entities: missing required field "Outbox.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := outbox.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entities: validator failed for field "Outbox.status": %w`, err)}
		}
	}
	return nil
}

func (oc *OutboxCreate) sqlSave(ctx context.Context) (*Outbox, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OutboxCreate) createSpec() (*Outbox, *sqlgraph.CreateSpec) {
	var (
		_node = &Outbox{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: outbox.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outbox.FieldID,
			},
		}
	)
	_spec.Schema = oc.schemaConfig.Outbox
	_spec.OnConflict = oc.conflict
	if value, ok := oc.mutation.Timestamp(); ok {
		_spec.SetField(outbox.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := oc.mutation.Topic(); ok {
		_spec.SetField(outbox.FieldTopic, field.TypeString, value)
		_node.Topic = value
	}
	if value, ok := oc.mutation.Key(); ok {
		_spec.SetField(outbox.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := oc.mutation.Payload(); ok {
		_spec.SetField(outbox.FieldPayload, field.TypeBytes, value)
		_node.Payload = value
	}
	if value, ok := oc.mutation.Headers(); ok {
		_spec.SetField(outbox.FieldHeaders, field.TypeJSON, value)
		_node.Headers = value
	}
	if value, ok := oc.mutation.RetryCount(); ok {
		_spec.SetField(outbox.FieldRetryCount, field.TypeInt, value)
		_node.RetryCount = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(outbox.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.LastRetry(); ok {
		_spec.SetField(outbox.FieldLastRetry, field.TypeTime, value)
		_node.LastRetry = value
	}
	if value, ok := oc.mutation.ProcessingErrors(); ok {
		_spec.SetField(outbox.FieldProcessingErrors, field.TypeJSON, value)
		_node.ProcessingErrors = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Outbox.Create().
//		SetTimestamp(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OutboxUpsert) {
//			SetTimestamp(v+v).
//		}).
//		Exec(ctx)
func (oc *OutboxCreate) OnConflict(opts ...sql.ConflictOption) *OutboxUpsertOne {
	oc.conflict = opts
	return &OutboxUpsertOne{
		create: oc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Outbox.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oc *OutboxCreate) OnConflictColumns(columns ...string) *OutboxUpsertOne {
	oc.conflict = append(oc.conflict, sql.ConflictColumns(columns...))
	return &OutboxUpsertOne{
		create: oc,
	}
}

type (
	// OutboxUpsertOne is the builder for "upsert"-ing
	//  one Outbox node.
	OutboxUpsertOne struct {
		create *OutboxCreate
	}

	// OutboxUpsert is the "OnConflict" setter.
	OutboxUpsert struct {
		*sql.UpdateSet
	}
)

// SetTimestamp sets the "timestamp" field.
func (u *OutboxUpsert) SetTimestamp(v time.Time) *OutboxUpsert {
	u.Set(outbox.FieldTimestamp, v)
	return u
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *OutboxUpsert) UpdateTimestamp() *OutboxUpsert {
	u.SetExcluded(outbox.FieldTimestamp)
	return u
}

// SetTopic sets the "topic" field.
func (u *OutboxUpsert) SetTopic(v string) *OutboxUpsert {
	u.Set(outbox.FieldTopic, v)
	return u
}

// UpdateTopic sets the "topic" field to the value that was provided on create.
func (u *OutboxUpsert) UpdateTopic() *OutboxUpsert {
	u.SetExcluded(outbox.FieldTopic)
	return u
}

// SetKey sets the "key" field.
func (u *OutboxUpsert) SetKey(v string) *OutboxUpsert {
	u.Set(outbox.FieldKey, v)
	return u
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *OutboxUpsert) UpdateKey() *OutboxUpsert {
	u.SetExcluded(outbox.FieldKey)
	return u
}

// SetPayload sets the "payload" field.
func (u *OutboxUpsert) SetPayload(v []byte) *OutboxUpsert {
	u.Set(outbox.FieldPayload, v)
	return u
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *OutboxUpsert) UpdatePayload() *OutboxUpsert {
	u.SetExcluded(outbox.FieldPayload)
	return u
}

// SetHeaders sets the "headers" field.
func (u *OutboxUpsert) SetHeaders(v map[string]string) *OutboxUpsert {
	u.Set(outbox.FieldHeaders, v)
	return u
}

// UpdateHeaders sets the "headers" field to the value that was provided on create.
func (u *OutboxUpsert) UpdateHeaders() *OutboxUpsert {
	u.SetExcluded(outbox.FieldHeaders)
	return u
}

// SetRetryCount sets the "retry_count" field.
func (u *OutboxUpsert) SetRetryCount(v int) *OutboxUpsert {
	u.Set(outbox.FieldRetryCount, v)
	return u
}

// UpdateRetryCount sets the "retry_count" field to the value that was provided on create.
func (u *OutboxUpsert) UpdateRetryCount() *OutboxUpsert {
	u.SetExcluded(outbox.FieldRetryCount)
	return u
}

// AddRetryCount adds v to the "retry_count" field.
func (u *OutboxUpsert) AddRetryCount(v int) *OutboxUpsert {
	u.Add(outbox.FieldRetryCount, v)
	return u
}

// SetStatus sets the "status" field.
func (u *OutboxUpsert) SetStatus(v outbox.Status) *OutboxUpsert {
	u.Set(outbox.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OutboxUpsert) UpdateStatus() *OutboxUpsert {
	u.SetExcluded(outbox.FieldStatus)
	return u
}

// SetLastRetry sets the "last_retry" field.
func (u *OutboxUpsert) SetLastRetry(v time.Time) *OutboxUpsert {
	u.Set(outbox.FieldLastRetry, v)
	return u
}

// UpdateLastRetry sets the "last_retry" field to the value that was provided on create.
func (u *OutboxUpsert) UpdateLastRetry() *OutboxUpsert {
	u.SetExcluded(outbox.FieldLastRetry)
	return u
}

// ClearLastRetry clears the value of the "last_retry" field.
func (u *OutboxUpsert) ClearLastRetry() *OutboxUpsert {
	u.SetNull(outbox.FieldLastRetry)
	return u
}

// SetProcessingErrors sets the "processing_errors" field.
func (u *OutboxUpsert) SetProcessingErrors(v []string) *OutboxUpsert {
	u.Set(outbox.FieldProcessingErrors, v)
	return u
}

// UpdateProcessingErrors sets the "processing_errors" field to the value that was provided on create.
func (u *OutboxUpsert) UpdateProcessingErrors() *OutboxUpsert {
	u.SetExcluded(outbox.FieldProcessingErrors)
	return u
}

// ClearProcessingErrors clears the value of the "processing_errors" field.
func (u *OutboxUpsert) ClearProcessingErrors() *OutboxUpsert {
	u.SetNull(outbox.FieldProcessingErrors)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Outbox.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OutboxUpsertOne) UpdateNewValues() *OutboxUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Outbox.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OutboxUpsertOne) Ignore() *OutboxUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OutboxUpsertOne) DoNothing() *OutboxUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OutboxCreate.OnConflict
// documentation for more info.
func (u *OutboxUpsertOne) Update(set func(*OutboxUpsert)) *OutboxUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OutboxUpsert{UpdateSet: update})
	}))
	return u
}

// SetTimestamp sets the "timestamp" field.
func (u *OutboxUpsertOne) SetTimestamp(v time.Time) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetTimestamp(v)
	})
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdateTimestamp() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateTimestamp()
	})
}

// SetTopic sets the "topic" field.
func (u *OutboxUpsertOne) SetTopic(v string) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetTopic(v)
	})
}

// UpdateTopic sets the "topic" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdateTopic() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateTopic()
	})
}

// SetKey sets the "key" field.
func (u *OutboxUpsertOne) SetKey(v string) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdateKey() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateKey()
	})
}

// SetPayload sets the "payload" field.
func (u *OutboxUpsertOne) SetPayload(v []byte) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetPayload(v)
	})
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdatePayload() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdatePayload()
	})
}

// SetHeaders sets the "headers" field.
func (u *OutboxUpsertOne) SetHeaders(v map[string]string) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetHeaders(v)
	})
}

// UpdateHeaders sets the "headers" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdateHeaders() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateHeaders()
	})
}

// SetRetryCount sets the "retry_count" field.
func (u *OutboxUpsertOne) SetRetryCount(v int) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetRetryCount(v)
	})
}

// AddRetryCount adds v to the "retry_count" field.
func (u *OutboxUpsertOne) AddRetryCount(v int) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.AddRetryCount(v)
	})
}

// UpdateRetryCount sets the "retry_count" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdateRetryCount() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateRetryCount()
	})
}

// SetStatus sets the "status" field.
func (u *OutboxUpsertOne) SetStatus(v outbox.Status) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdateStatus() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateStatus()
	})
}

// SetLastRetry sets the "last_retry" field.
func (u *OutboxUpsertOne) SetLastRetry(v time.Time) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetLastRetry(v)
	})
}

// UpdateLastRetry sets the "last_retry" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdateLastRetry() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateLastRetry()
	})
}

// ClearLastRetry clears the value of the "last_retry" field.
func (u *OutboxUpsertOne) ClearLastRetry() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.ClearLastRetry()
	})
}

// SetProcessingErrors sets the "processing_errors" field.
func (u *OutboxUpsertOne) SetProcessingErrors(v []string) *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.SetProcessingErrors(v)
	})
}

// UpdateProcessingErrors sets the "processing_errors" field to the value that was provided on create.
func (u *OutboxUpsertOne) UpdateProcessingErrors() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateProcessingErrors()
	})
}

// ClearProcessingErrors clears the value of the "processing_errors" field.
func (u *OutboxUpsertOne) ClearProcessingErrors() *OutboxUpsertOne {
	return u.Update(func(s *OutboxUpsert) {
		s.ClearProcessingErrors()
	})
}

// Exec executes the query.
func (u *OutboxUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for OutboxCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OutboxUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OutboxUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OutboxUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OutboxCreateBulk is the builder for creating many Outbox entities in bulk.
type OutboxCreateBulk struct {
	config
	builders []*OutboxCreate
	conflict []sql.ConflictOption
}

// Save creates the Outbox entities in the database.
func (ocb *OutboxCreateBulk) Save(ctx context.Context) ([]*Outbox, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Outbox, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OutboxMutation)
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
func (ocb *OutboxCreateBulk) SaveX(ctx context.Context) []*Outbox {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OutboxCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OutboxCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Outbox.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OutboxUpsert) {
//			SetTimestamp(v+v).
//		}).
//		Exec(ctx)
func (ocb *OutboxCreateBulk) OnConflict(opts ...sql.ConflictOption) *OutboxUpsertBulk {
	ocb.conflict = opts
	return &OutboxUpsertBulk{
		create: ocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Outbox.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ocb *OutboxCreateBulk) OnConflictColumns(columns ...string) *OutboxUpsertBulk {
	ocb.conflict = append(ocb.conflict, sql.ConflictColumns(columns...))
	return &OutboxUpsertBulk{
		create: ocb,
	}
}

// OutboxUpsertBulk is the builder for "upsert"-ing
// a bulk of Outbox nodes.
type OutboxUpsertBulk struct {
	create *OutboxCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Outbox.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OutboxUpsertBulk) UpdateNewValues() *OutboxUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Outbox.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OutboxUpsertBulk) Ignore() *OutboxUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OutboxUpsertBulk) DoNothing() *OutboxUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OutboxCreateBulk.OnConflict
// documentation for more info.
func (u *OutboxUpsertBulk) Update(set func(*OutboxUpsert)) *OutboxUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OutboxUpsert{UpdateSet: update})
	}))
	return u
}

// SetTimestamp sets the "timestamp" field.
func (u *OutboxUpsertBulk) SetTimestamp(v time.Time) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetTimestamp(v)
	})
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdateTimestamp() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateTimestamp()
	})
}

// SetTopic sets the "topic" field.
func (u *OutboxUpsertBulk) SetTopic(v string) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetTopic(v)
	})
}

// UpdateTopic sets the "topic" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdateTopic() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateTopic()
	})
}

// SetKey sets the "key" field.
func (u *OutboxUpsertBulk) SetKey(v string) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetKey(v)
	})
}

// UpdateKey sets the "key" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdateKey() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateKey()
	})
}

// SetPayload sets the "payload" field.
func (u *OutboxUpsertBulk) SetPayload(v []byte) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetPayload(v)
	})
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdatePayload() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdatePayload()
	})
}

// SetHeaders sets the "headers" field.
func (u *OutboxUpsertBulk) SetHeaders(v map[string]string) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetHeaders(v)
	})
}

// UpdateHeaders sets the "headers" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdateHeaders() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateHeaders()
	})
}

// SetRetryCount sets the "retry_count" field.
func (u *OutboxUpsertBulk) SetRetryCount(v int) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetRetryCount(v)
	})
}

// AddRetryCount adds v to the "retry_count" field.
func (u *OutboxUpsertBulk) AddRetryCount(v int) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.AddRetryCount(v)
	})
}

// UpdateRetryCount sets the "retry_count" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdateRetryCount() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateRetryCount()
	})
}

// SetStatus sets the "status" field.
func (u *OutboxUpsertBulk) SetStatus(v outbox.Status) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdateStatus() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateStatus()
	})
}

// SetLastRetry sets the "last_retry" field.
func (u *OutboxUpsertBulk) SetLastRetry(v time.Time) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetLastRetry(v)
	})
}

// UpdateLastRetry sets the "last_retry" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdateLastRetry() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateLastRetry()
	})
}

// ClearLastRetry clears the value of the "last_retry" field.
func (u *OutboxUpsertBulk) ClearLastRetry() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.ClearLastRetry()
	})
}

// SetProcessingErrors sets the "processing_errors" field.
func (u *OutboxUpsertBulk) SetProcessingErrors(v []string) *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.SetProcessingErrors(v)
	})
}

// UpdateProcessingErrors sets the "processing_errors" field to the value that was provided on create.
func (u *OutboxUpsertBulk) UpdateProcessingErrors() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.UpdateProcessingErrors()
	})
}

// ClearProcessingErrors clears the value of the "processing_errors" field.
func (u *OutboxUpsertBulk) ClearProcessingErrors() *OutboxUpsertBulk {
	return u.Update(func(s *OutboxUpsert) {
		s.ClearProcessingErrors()
	})
}

// Exec executes the query.
func (u *OutboxUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entities: OnConflict was set for builder %d. Set it on the OutboxCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for OutboxCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OutboxUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
