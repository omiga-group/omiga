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
	"github.com/omiga-group/omiga/src/exchange/shared/entities/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/internal"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/predicate"
	"github.com/omiga-group/omiga/src/exchange/shared/entities/tradingpair"
)

// TradingPairQuery is the builder for querying TradingPair entities.
type TradingPairQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.TradingPair
	withExchange *ExchangeQuery
	withFKs      bool
	loadTotal    []func(context.Context, []*TradingPair) error
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TradingPairQuery builder.
func (tpq *TradingPairQuery) Where(ps ...predicate.TradingPair) *TradingPairQuery {
	tpq.predicates = append(tpq.predicates, ps...)
	return tpq
}

// Limit adds a limit step to the query.
func (tpq *TradingPairQuery) Limit(limit int) *TradingPairQuery {
	tpq.limit = &limit
	return tpq
}

// Offset adds an offset step to the query.
func (tpq *TradingPairQuery) Offset(offset int) *TradingPairQuery {
	tpq.offset = &offset
	return tpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tpq *TradingPairQuery) Unique(unique bool) *TradingPairQuery {
	tpq.unique = &unique
	return tpq
}

// Order adds an order step to the query.
func (tpq *TradingPairQuery) Order(o ...OrderFunc) *TradingPairQuery {
	tpq.order = append(tpq.order, o...)
	return tpq
}

// QueryExchange chains the current query on the "exchange" edge.
func (tpq *TradingPairQuery) QueryExchange() *ExchangeQuery {
	query := &ExchangeQuery{config: tpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tradingpair.Table, tradingpair.FieldID, selector),
			sqlgraph.To(exchange.Table, exchange.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tradingpair.ExchangeTable, tradingpair.ExchangeColumn),
		)
		schemaConfig := tpq.schemaConfig
		step.To.Schema = schemaConfig.Exchange
		step.Edge.Schema = schemaConfig.TradingPair
		fromU = sqlgraph.SetNeighbors(tpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TradingPair entity from the query.
// Returns a *NotFoundError when no TradingPair was found.
func (tpq *TradingPairQuery) First(ctx context.Context) (*TradingPair, error) {
	nodes, err := tpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tradingpair.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tpq *TradingPairQuery) FirstX(ctx context.Context) *TradingPair {
	node, err := tpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TradingPair ID from the query.
// Returns a *NotFoundError when no TradingPair ID was found.
func (tpq *TradingPairQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tradingpair.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tpq *TradingPairQuery) FirstIDX(ctx context.Context) int {
	id, err := tpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TradingPair entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TradingPair entity is found.
// Returns a *NotFoundError when no TradingPair entities are found.
func (tpq *TradingPairQuery) Only(ctx context.Context) (*TradingPair, error) {
	nodes, err := tpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tradingpair.Label}
	default:
		return nil, &NotSingularError{tradingpair.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tpq *TradingPairQuery) OnlyX(ctx context.Context) *TradingPair {
	node, err := tpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TradingPair ID in the query.
// Returns a *NotSingularError when more than one TradingPair ID is found.
// Returns a *NotFoundError when no entities are found.
func (tpq *TradingPairQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tradingpair.Label}
	default:
		err = &NotSingularError{tradingpair.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tpq *TradingPairQuery) OnlyIDX(ctx context.Context) int {
	id, err := tpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TradingPairs.
func (tpq *TradingPairQuery) All(ctx context.Context) ([]*TradingPair, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tpq *TradingPairQuery) AllX(ctx context.Context) []*TradingPair {
	nodes, err := tpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TradingPair IDs.
func (tpq *TradingPairQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tpq.Select(tradingpair.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tpq *TradingPairQuery) IDsX(ctx context.Context) []int {
	ids, err := tpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tpq *TradingPairQuery) Count(ctx context.Context) (int, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tpq *TradingPairQuery) CountX(ctx context.Context) int {
	count, err := tpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tpq *TradingPairQuery) Exist(ctx context.Context) (bool, error) {
	if err := tpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tpq *TradingPairQuery) ExistX(ctx context.Context) bool {
	exist, err := tpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TradingPairQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tpq *TradingPairQuery) Clone() *TradingPairQuery {
	if tpq == nil {
		return nil
	}
	return &TradingPairQuery{
		config:       tpq.config,
		limit:        tpq.limit,
		offset:       tpq.offset,
		order:        append([]OrderFunc{}, tpq.order...),
		predicates:   append([]predicate.TradingPair{}, tpq.predicates...),
		withExchange: tpq.withExchange.Clone(),
		// clone intermediate query.
		sql:    tpq.sql.Clone(),
		path:   tpq.path,
		unique: tpq.unique,
	}
}

// WithExchange tells the query-builder to eager-load the nodes that are connected to
// the "exchange" edge. The optional arguments are used to configure the query builder of the edge.
func (tpq *TradingPairQuery) WithExchange(opts ...func(*ExchangeQuery)) *TradingPairQuery {
	query := &ExchangeQuery{config: tpq.config}
	for _, opt := range opts {
		opt(query)
	}
	tpq.withExchange = query
	return tpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Symbol string `json:"symbol,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TradingPair.Query().
//		GroupBy(tradingpair.FieldSymbol).
//		Aggregate(entities.Count()).
//		Scan(ctx, &v)
func (tpq *TradingPairQuery) GroupBy(field string, fields ...string) *TradingPairGroupBy {
	grbuild := &TradingPairGroupBy{config: tpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tpq.sqlQuery(ctx), nil
	}
	grbuild.label = tradingpair.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Symbol string `json:"symbol,omitempty"`
//	}
//
//	client.TradingPair.Query().
//		Select(tradingpair.FieldSymbol).
//		Scan(ctx, &v)
func (tpq *TradingPairQuery) Select(fields ...string) *TradingPairSelect {
	tpq.fields = append(tpq.fields, fields...)
	selbuild := &TradingPairSelect{TradingPairQuery: tpq}
	selbuild.label = tradingpair.Label
	selbuild.flds, selbuild.scan = &tpq.fields, selbuild.Scan
	return selbuild
}

func (tpq *TradingPairQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tpq.fields {
		if !tradingpair.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
		}
	}
	if tpq.path != nil {
		prev, err := tpq.path(ctx)
		if err != nil {
			return err
		}
		tpq.sql = prev
	}
	return nil
}

func (tpq *TradingPairQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TradingPair, error) {
	var (
		nodes       = []*TradingPair{}
		withFKs     = tpq.withFKs
		_spec       = tpq.querySpec()
		loadedTypes = [1]bool{
			tpq.withExchange != nil,
		}
	)
	if tpq.withExchange != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, tradingpair.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TradingPair).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TradingPair{config: tpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tpq.schemaConfig.TradingPair
	ctx = internal.NewSchemaConfigContext(ctx, tpq.schemaConfig)
	if len(tpq.modifiers) > 0 {
		_spec.Modifiers = tpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tpq.withExchange; query != nil {
		if err := tpq.loadExchange(ctx, query, nodes, nil,
			func(n *TradingPair, e *Exchange) { n.Edges.Exchange = e }); err != nil {
			return nil, err
		}
	}
	for i := range tpq.loadTotal {
		if err := tpq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tpq *TradingPairQuery) loadExchange(ctx context.Context, query *ExchangeQuery, nodes []*TradingPair, init func(*TradingPair), assign func(*TradingPair, *Exchange)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TradingPair)
	for i := range nodes {
		if nodes[i].exchange_trading_pairs == nil {
			continue
		}
		fk := *nodes[i].exchange_trading_pairs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(exchange.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "exchange_trading_pairs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tpq *TradingPairQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tpq.querySpec()
	_spec.Node.Schema = tpq.schemaConfig.TradingPair
	ctx = internal.NewSchemaConfigContext(ctx, tpq.schemaConfig)
	if len(tpq.modifiers) > 0 {
		_spec.Modifiers = tpq.modifiers
	}
	_spec.Node.Columns = tpq.fields
	if len(tpq.fields) > 0 {
		_spec.Unique = tpq.unique != nil && *tpq.unique
	}
	return sqlgraph.CountNodes(ctx, tpq.driver, _spec)
}

func (tpq *TradingPairQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entities: check existence: %w", err)
	}
	return n > 0, nil
}

func (tpq *TradingPairQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tradingpair.Table,
			Columns: tradingpair.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tradingpair.FieldID,
			},
		},
		From:   tpq.sql,
		Unique: true,
	}
	if unique := tpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tradingpair.FieldID)
		for i := range fields {
			if fields[i] != tradingpair.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tpq *TradingPairQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tpq.driver.Dialect())
	t1 := builder.Table(tradingpair.Table)
	columns := tpq.fields
	if len(columns) == 0 {
		columns = tradingpair.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tpq.sql != nil {
		selector = tpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tpq.unique != nil && *tpq.unique {
		selector.Distinct()
	}
	t1.Schema(tpq.schemaConfig.TradingPair)
	ctx = internal.NewSchemaConfigContext(ctx, tpq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range tpq.modifiers {
		m(selector)
	}
	for _, p := range tpq.predicates {
		p(selector)
	}
	for _, p := range tpq.order {
		p(selector)
	}
	if offset := tpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tpq *TradingPairQuery) ForUpdate(opts ...sql.LockOption) *TradingPairQuery {
	if tpq.driver.Dialect() == dialect.Postgres {
		tpq.Unique(false)
	}
	tpq.modifiers = append(tpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tpq *TradingPairQuery) ForShare(opts ...sql.LockOption) *TradingPairQuery {
	if tpq.driver.Dialect() == dialect.Postgres {
		tpq.Unique(false)
	}
	tpq.modifiers = append(tpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tpq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tpq *TradingPairQuery) Modify(modifiers ...func(s *sql.Selector)) *TradingPairSelect {
	tpq.modifiers = append(tpq.modifiers, modifiers...)
	return tpq.Select()
}

// TradingPairGroupBy is the group-by builder for TradingPair entities.
type TradingPairGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tpgb *TradingPairGroupBy) Aggregate(fns ...AggregateFunc) *TradingPairGroupBy {
	tpgb.fns = append(tpgb.fns, fns...)
	return tpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tpgb *TradingPairGroupBy) Scan(ctx context.Context, v any) error {
	query, err := tpgb.path(ctx)
	if err != nil {
		return err
	}
	tpgb.sql = query
	return tpgb.sqlScan(ctx, v)
}

func (tpgb *TradingPairGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range tpgb.fields {
		if !tradingpair.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tpgb *TradingPairGroupBy) sqlQuery() *sql.Selector {
	selector := tpgb.sql.Select()
	aggregation := make([]string, 0, len(tpgb.fns))
	for _, fn := range tpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tpgb.fields)+len(tpgb.fns))
		for _, f := range tpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tpgb.fields...)...)
}

// TradingPairSelect is the builder for selecting fields of TradingPair entities.
type TradingPairSelect struct {
	*TradingPairQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tps *TradingPairSelect) Scan(ctx context.Context, v any) error {
	if err := tps.prepareQuery(ctx); err != nil {
		return err
	}
	tps.sql = tps.TradingPairQuery.sqlQuery(ctx)
	return tps.sqlScan(ctx, v)
}

func (tps *TradingPairSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := tps.sql.Query()
	if err := tps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tps *TradingPairSelect) Modify(modifiers ...func(s *sql.Selector)) *TradingPairSelect {
	tps.modifiers = append(tps.modifiers, modifiers...)
	return tps
}