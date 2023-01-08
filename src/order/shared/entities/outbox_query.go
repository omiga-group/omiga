// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omiga-group/omiga/src/order/shared/entities/internal"
	"github.com/omiga-group/omiga/src/order/shared/entities/outbox"
	"github.com/omiga-group/omiga/src/order/shared/entities/predicate"
)

// OutboxQuery is the builder for querying Outbox entities.
type OutboxQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Outbox
	loadTotal  []func(context.Context, []*Outbox) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutboxQuery builder.
func (oq *OutboxQuery) Where(ps ...predicate.Outbox) *OutboxQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OutboxQuery) Limit(limit int) *OutboxQuery {
	oq.limit = &limit
	return oq
}

// Offset to start from.
func (oq *OutboxQuery) Offset(offset int) *OutboxQuery {
	oq.offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OutboxQuery) Unique(unique bool) *OutboxQuery {
	oq.unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OutboxQuery) Order(o ...OrderFunc) *OutboxQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// First returns the first Outbox entity from the query.
// Returns a *NotFoundError when no Outbox was found.
func (oq *OutboxQuery) First(ctx context.Context) (*Outbox, error) {
	nodes, err := oq.Limit(1).All(newQueryContext(ctx, TypeOutbox, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outbox.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OutboxQuery) FirstX(ctx context.Context) *Outbox {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Outbox ID from the query.
// Returns a *NotFoundError when no Outbox ID was found.
func (oq *OutboxQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(newQueryContext(ctx, TypeOutbox, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outbox.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OutboxQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Outbox entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Outbox entity is found.
// Returns a *NotFoundError when no Outbox entities are found.
func (oq *OutboxQuery) Only(ctx context.Context) (*Outbox, error) {
	nodes, err := oq.Limit(2).All(newQueryContext(ctx, TypeOutbox, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outbox.Label}
	default:
		return nil, &NotSingularError{outbox.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OutboxQuery) OnlyX(ctx context.Context) *Outbox {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Outbox ID in the query.
// Returns a *NotSingularError when more than one Outbox ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OutboxQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(newQueryContext(ctx, TypeOutbox, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outbox.Label}
	default:
		err = &NotSingularError{outbox.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OutboxQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Outboxes.
func (oq *OutboxQuery) All(ctx context.Context) ([]*Outbox, error) {
	ctx = newQueryContext(ctx, TypeOutbox, "All")
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Outbox, *OutboxQuery]()
	return withInterceptors[[]*Outbox](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OutboxQuery) AllX(ctx context.Context) []*Outbox {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Outbox IDs.
func (oq *OutboxQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeOutbox, "IDs")
	if err := oq.Select(outbox.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OutboxQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OutboxQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeOutbox, "Count")
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OutboxQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OutboxQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OutboxQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeOutbox, "Exist")
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entities: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OutboxQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutboxQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OutboxQuery) Clone() *OutboxQuery {
	if oq == nil {
		return nil
	}
	return &OutboxQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]OrderFunc{}, oq.order...),
		inters:     append([]Interceptor{}, oq.inters...),
		predicates: append([]predicate.Outbox{}, oq.predicates...),
		// clone intermediate query.
		sql:    oq.sql.Clone(),
		path:   oq.path,
		unique: oq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Timestamp time.Time `json:"timestamp,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Outbox.Query().
//		GroupBy(outbox.FieldTimestamp).
//		Aggregate(entities.Count()).
//		Scan(ctx, &v)
func (oq *OutboxQuery) GroupBy(field string, fields ...string) *OutboxGroupBy {
	oq.fields = append([]string{field}, fields...)
	grbuild := &OutboxGroupBy{build: oq}
	grbuild.flds = &oq.fields
	grbuild.label = outbox.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Timestamp time.Time `json:"timestamp,omitempty"`
//	}
//
//	client.Outbox.Query().
//		Select(outbox.FieldTimestamp).
//		Scan(ctx, &v)
func (oq *OutboxQuery) Select(fields ...string) *OutboxSelect {
	oq.fields = append(oq.fields, fields...)
	sbuild := &OutboxSelect{OutboxQuery: oq}
	sbuild.label = outbox.Label
	sbuild.flds, sbuild.scan = &oq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OutboxSelect configured with the given aggregations.
func (oq *OutboxQuery) Aggregate(fns ...AggregateFunc) *OutboxSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OutboxQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("entities: uninitialized interceptor (forgotten import entities/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.fields {
		if !outbox.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OutboxQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Outbox, error) {
	var (
		nodes = []*Outbox{}
		_spec = oq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Outbox).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Outbox{config: oq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = oq.schemaConfig.Outbox
	ctx = internal.NewSchemaConfigContext(ctx, oq.schemaConfig)
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range oq.loadTotal {
		if err := oq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OutboxQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Schema = oq.schemaConfig.Outbox
	ctx = internal.NewSchemaConfigContext(ctx, oq.schemaConfig)
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	_spec.Node.Columns = oq.fields
	if len(oq.fields) > 0 {
		_spec.Unique = oq.unique != nil && *oq.unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OutboxQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outbox.Table,
			Columns: outbox.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: outbox.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if unique := oq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outbox.FieldID)
		for i := range fields {
			if fields[i] != outbox.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OutboxQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(outbox.Table)
	columns := oq.fields
	if len(columns) == 0 {
		columns = outbox.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.unique != nil && *oq.unique {
		selector.Distinct()
	}
	t1.Schema(oq.schemaConfig.Outbox)
	ctx = internal.NewSchemaConfigContext(ctx, oq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range oq.modifiers {
		m(selector)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (oq *OutboxQuery) ForUpdate(opts ...sql.LockOption) *OutboxQuery {
	if oq.driver.Dialect() == dialect.Postgres {
		oq.Unique(false)
	}
	oq.modifiers = append(oq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return oq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (oq *OutboxQuery) ForShare(opts ...sql.LockOption) *OutboxQuery {
	if oq.driver.Dialect() == dialect.Postgres {
		oq.Unique(false)
	}
	oq.modifiers = append(oq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return oq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oq *OutboxQuery) Modify(modifiers ...func(s *sql.Selector)) *OutboxSelect {
	oq.modifiers = append(oq.modifiers, modifiers...)
	return oq.Select()
}

// OutboxGroupBy is the group-by builder for Outbox entities.
type OutboxGroupBy struct {
	selector
	build *OutboxQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OutboxGroupBy) Aggregate(fns ...AggregateFunc) *OutboxGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OutboxGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeOutbox, "GroupBy")
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OutboxQuery, *OutboxGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OutboxGroupBy) sqlScan(ctx context.Context, root *OutboxQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OutboxSelect is the builder for selecting fields of Outbox entities.
type OutboxSelect struct {
	*OutboxQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OutboxSelect) Aggregate(fns ...AggregateFunc) *OutboxSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OutboxSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeOutbox, "Select")
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OutboxQuery, *OutboxSelect](ctx, os.OutboxQuery, os, os.inters, v)
}

func (os *OutboxSelect) sqlScan(ctx context.Context, root *OutboxQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (os *OutboxSelect) Modify(modifiers ...func(s *sql.Selector)) *OutboxSelect {
	os.modifiers = append(os.modifiers, modifiers...)
	return os
}
