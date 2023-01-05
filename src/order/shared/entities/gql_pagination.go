// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/omiga-group/omiga/src/order/shared/entities/order"
	"github.com/omiga-group/omiga/src/order/shared/entities/outbox"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// OrderEdge is the edge representation of Order.
type OrderEdge struct {
	Node   *Order `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// OrderConnection is the connection containing edges to Order.
type OrderConnection struct {
	Edges      []*OrderEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *OrderConnection) build(nodes []*Order, pager *orderPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Order
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Order {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Order {
			return nodes[i]
		}
	}
	c.Edges = make([]*OrderEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &OrderEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// OrderPaginateOption enables pagination customization.
type OrderPaginateOption func(*orderPager) error

// WithOrderOrder configures pagination ordering.
func WithOrderOrder(order *OrderOrder) OrderPaginateOption {
	if order == nil {
		order = DefaultOrderOrder
	}
	o := *order
	return func(pager *orderPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOrderOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOrderFilter configures pagination filter.
func WithOrderFilter(filter func(*OrderQuery) (*OrderQuery, error)) OrderPaginateOption {
	return func(pager *orderPager) error {
		if filter == nil {
			return errors.New("OrderQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type orderPager struct {
	order  *OrderOrder
	filter func(*OrderQuery) (*OrderQuery, error)
}

func newOrderPager(opts []OrderPaginateOption) (*orderPager, error) {
	pager := &orderPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOrderOrder
	}
	return pager, nil
}

func (p *orderPager) applyFilter(query *OrderQuery) (*OrderQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *orderPager) toCursor(o *Order) Cursor {
	return p.order.Field.toCursor(o)
}

func (p *orderPager) applyCursors(query *OrderQuery, after, before *Cursor) *OrderQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultOrderOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *orderPager) applyOrder(query *OrderQuery, reverse bool) *OrderQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultOrderOrder.Field {
		query = query.Order(direction.orderFunc(DefaultOrderOrder.Field.field))
	}
	return query
}

func (p *orderPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultOrderOrder.Field {
			b.Comma().Ident(DefaultOrderOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Order.
func (o *OrderQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OrderPaginateOption,
) (*OrderConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOrderPager(opts)
	if err != nil {
		return nil, err
	}
	if o, err = pager.applyFilter(o); err != nil {
		return nil, err
	}
	conn := &OrderConnection{Edges: []*OrderEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = o.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	o = pager.applyCursors(o, after, before)
	o = pager.applyOrder(o, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		o.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := o.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := o.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// OrderOrderField defines the ordering field of Order.
type OrderOrderField struct {
	field    string
	toCursor func(*Order) Cursor
}

// OrderOrder defines the ordering of Order.
type OrderOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *OrderOrderField `json:"field"`
}

// DefaultOrderOrder is the default ordering of Order.
var DefaultOrderOrder = &OrderOrder{
	Direction: OrderDirectionAsc,
	Field: &OrderOrderField{
		field: order.FieldID,
		toCursor: func(o *Order) Cursor {
			return Cursor{ID: o.ID}
		},
	},
}

// ToEdge converts Order into OrderEdge.
func (o *Order) ToEdge(order *OrderOrder) *OrderEdge {
	if order == nil {
		order = DefaultOrderOrder
	}
	return &OrderEdge{
		Node:   o,
		Cursor: order.Field.toCursor(o),
	}
}

// OutboxEdge is the edge representation of Outbox.
type OutboxEdge struct {
	Node   *Outbox `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// OutboxConnection is the connection containing edges to Outbox.
type OutboxConnection struct {
	Edges      []*OutboxEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *OutboxConnection) build(nodes []*Outbox, pager *outboxPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Outbox
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Outbox {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Outbox {
			return nodes[i]
		}
	}
	c.Edges = make([]*OutboxEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &OutboxEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// OutboxPaginateOption enables pagination customization.
type OutboxPaginateOption func(*outboxPager) error

// WithOutboxOrder configures pagination ordering.
func WithOutboxOrder(order *OutboxOrder) OutboxPaginateOption {
	if order == nil {
		order = DefaultOutboxOrder
	}
	o := *order
	return func(pager *outboxPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOutboxOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOutboxFilter configures pagination filter.
func WithOutboxFilter(filter func(*OutboxQuery) (*OutboxQuery, error)) OutboxPaginateOption {
	return func(pager *outboxPager) error {
		if filter == nil {
			return errors.New("OutboxQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type outboxPager struct {
	order  *OutboxOrder
	filter func(*OutboxQuery) (*OutboxQuery, error)
}

func newOutboxPager(opts []OutboxPaginateOption) (*outboxPager, error) {
	pager := &outboxPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOutboxOrder
	}
	return pager, nil
}

func (p *outboxPager) applyFilter(query *OutboxQuery) (*OutboxQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *outboxPager) toCursor(o *Outbox) Cursor {
	return p.order.Field.toCursor(o)
}

func (p *outboxPager) applyCursors(query *OutboxQuery, after, before *Cursor) *OutboxQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultOutboxOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *outboxPager) applyOrder(query *OutboxQuery, reverse bool) *OutboxQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultOutboxOrder.Field {
		query = query.Order(direction.orderFunc(DefaultOutboxOrder.Field.field))
	}
	return query
}

func (p *outboxPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultOutboxOrder.Field {
			b.Comma().Ident(DefaultOutboxOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Outbox.
func (o *OutboxQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OutboxPaginateOption,
) (*OutboxConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOutboxPager(opts)
	if err != nil {
		return nil, err
	}
	if o, err = pager.applyFilter(o); err != nil {
		return nil, err
	}
	conn := &OutboxConnection{Edges: []*OutboxEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = o.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	o = pager.applyCursors(o, after, before)
	o = pager.applyOrder(o, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		o.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := o.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := o.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// OutboxOrderField defines the ordering field of Outbox.
type OutboxOrderField struct {
	field    string
	toCursor func(*Outbox) Cursor
}

// OutboxOrder defines the ordering of Outbox.
type OutboxOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *OutboxOrderField `json:"field"`
}

// DefaultOutboxOrder is the default ordering of Outbox.
var DefaultOutboxOrder = &OutboxOrder{
	Direction: OrderDirectionAsc,
	Field: &OutboxOrderField{
		field: outbox.FieldID,
		toCursor: func(o *Outbox) Cursor {
			return Cursor{ID: o.ID}
		},
	},
}

// ToEdge converts Outbox into OutboxEdge.
func (o *Outbox) ToEdge(order *OutboxOrder) *OutboxEdge {
	if order == nil {
		order = DefaultOutboxOrder
	}
	return &OutboxEdge{
		Node:   o,
		Cursor: order.Field.toCursor(o),
	}
}
