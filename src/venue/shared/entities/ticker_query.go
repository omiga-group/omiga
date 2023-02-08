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
	"github.com/omiga-group/omiga/src/venue/shared/entities/internal"
	"github.com/omiga-group/omiga/src/venue/shared/entities/predicate"
	"github.com/omiga-group/omiga/src/venue/shared/entities/ticker"
	"github.com/omiga-group/omiga/src/venue/shared/entities/venue"
)

// TickerQuery is the builder for querying Ticker entities.
type TickerQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Ticker
	withVenue  *VenueQuery
	withFKs    bool
	loadTotal  []func(context.Context, []*Ticker) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TickerQuery builder.
func (tq *TickerQuery) Where(ps ...predicate.Ticker) *TickerQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TickerQuery) Limit(limit int) *TickerQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TickerQuery) Offset(offset int) *TickerQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TickerQuery) Unique(unique bool) *TickerQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TickerQuery) Order(o ...OrderFunc) *TickerQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryVenue chains the current query on the "venue" edge.
func (tq *TickerQuery) QueryVenue() *VenueQuery {
	query := (&VenueClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ticker.Table, ticker.FieldID, selector),
			sqlgraph.To(venue.Table, venue.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticker.VenueTable, ticker.VenueColumn),
		)
		schemaConfig := tq.schemaConfig
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.Ticker
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Ticker entity from the query.
// Returns a *NotFoundError when no Ticker was found.
func (tq *TickerQuery) First(ctx context.Context) (*Ticker, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ticker.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TickerQuery) FirstX(ctx context.Context) *Ticker {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Ticker ID from the query.
// Returns a *NotFoundError when no Ticker ID was found.
func (tq *TickerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ticker.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TickerQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Ticker entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Ticker entity is found.
// Returns a *NotFoundError when no Ticker entities are found.
func (tq *TickerQuery) Only(ctx context.Context) (*Ticker, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ticker.Label}
	default:
		return nil, &NotSingularError{ticker.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TickerQuery) OnlyX(ctx context.Context) *Ticker {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Ticker ID in the query.
// Returns a *NotSingularError when more than one Ticker ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TickerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ticker.Label}
	default:
		err = &NotSingularError{ticker.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TickerQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tickers.
func (tq *TickerQuery) All(ctx context.Context) ([]*Ticker, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Ticker, *TickerQuery]()
	return withInterceptors[[]*Ticker](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TickerQuery) AllX(ctx context.Context) []*Ticker {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Ticker IDs.
func (tq *TickerQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err := tq.Select(ticker.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TickerQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TickerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TickerQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TickerQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TickerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entities: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TickerQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TickerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TickerQuery) Clone() *TickerQuery {
	if tq == nil {
		return nil
	}
	return &TickerQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]OrderFunc{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Ticker{}, tq.predicates...),
		withVenue:  tq.withVenue.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithVenue tells the query-builder to eager-load the nodes that are connected to
// the "venue" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TickerQuery) WithVenue(opts ...func(*VenueQuery)) *TickerQuery {
	query := (&VenueClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withVenue = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Base string `json:"base,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Ticker.Query().
//		GroupBy(ticker.FieldBase).
//		Aggregate(entities.Count()).
//		Scan(ctx, &v)
func (tq *TickerQuery) GroupBy(field string, fields ...string) *TickerGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TickerGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = ticker.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Base string `json:"base,omitempty"`
//	}
//
//	client.Ticker.Query().
//		Select(ticker.FieldBase).
//		Scan(ctx, &v)
func (tq *TickerQuery) Select(fields ...string) *TickerSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TickerSelect{TickerQuery: tq}
	sbuild.label = ticker.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TickerSelect configured with the given aggregations.
func (tq *TickerQuery) Aggregate(fns ...AggregateFunc) *TickerSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TickerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("entities: uninitialized interceptor (forgotten import entities/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !ticker.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TickerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Ticker, error) {
	var (
		nodes       = []*Ticker{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withVenue != nil,
		}
	)
	if tq.withVenue != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ticker.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Ticker).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Ticker{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tq.schemaConfig.Ticker
	ctx = internal.NewSchemaConfigContext(ctx, tq.schemaConfig)
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withVenue; query != nil {
		if err := tq.loadVenue(ctx, query, nodes, nil,
			func(n *Ticker, e *Venue) { n.Edges.Venue = e }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TickerQuery) loadVenue(ctx context.Context, query *VenueQuery, nodes []*Ticker, init func(*Ticker), assign func(*Ticker, *Venue)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Ticker)
	for i := range nodes {
		if nodes[i].venue_ticker == nil {
			continue
		}
		fk := *nodes[i].venue_ticker
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(venue.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "venue_ticker" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TickerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Schema = tq.schemaConfig.Ticker
	ctx = internal.NewSchemaConfigContext(ctx, tq.schemaConfig)
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TickerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ticker.Table,
			Columns: ticker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ticker.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ticker.FieldID)
		for i := range fields {
			if fields[i] != ticker.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TickerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(ticker.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = ticker.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(tq.schemaConfig.Ticker)
	ctx = internal.NewSchemaConfigContext(ctx, tq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range tq.modifiers {
		m(selector)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tq *TickerQuery) ForUpdate(opts ...sql.LockOption) *TickerQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tq *TickerQuery) ForShare(opts ...sql.LockOption) *TickerQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tq *TickerQuery) Modify(modifiers ...func(s *sql.Selector)) *TickerSelect {
	tq.modifiers = append(tq.modifiers, modifiers...)
	return tq.Select()
}

// TickerGroupBy is the group-by builder for Ticker entities.
type TickerGroupBy struct {
	selector
	build *TickerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TickerGroupBy) Aggregate(fns ...AggregateFunc) *TickerGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TickerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TickerQuery, *TickerGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TickerGroupBy) sqlScan(ctx context.Context, root *TickerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TickerSelect is the builder for selecting fields of Ticker entities.
type TickerSelect struct {
	*TickerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TickerSelect) Aggregate(fns ...AggregateFunc) *TickerSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TickerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TickerQuery, *TickerSelect](ctx, ts.TickerQuery, ts, ts.inters, v)
}

func (ts *TickerSelect) sqlScan(ctx context.Context, root *TickerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ts *TickerSelect) Modify(modifiers ...func(s *sql.Selector)) *TickerSelect {
	ts.modifiers = append(ts.modifiers, modifiers...)
	return ts
}
