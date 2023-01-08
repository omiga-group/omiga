// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/omiga-group/omiga/src/order/shared/entities/migrate"

	"github.com/omiga-group/omiga/src/order/shared/entities/order"
	"github.com/omiga-group/omiga/src/order/shared/entities/outbox"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// Outbox is the client for interacting with the Outbox builders.
	Outbox *OutboxClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Order = NewOrderClient(c.config)
	c.Outbox = NewOutboxClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("entities: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("entities: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Order:  NewOrderClient(cfg),
		Outbox: NewOutboxClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Order:  NewOrderClient(cfg),
		Outbox: NewOutboxClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Order.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Order.Use(hooks...)
	c.Outbox.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Order.Intercept(interceptors...)
	c.Outbox.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *OrderMutation:
		return c.Order.mutate(ctx, m)
	case *OutboxMutation:
		return c.Outbox.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("entities: unknown mutation type %T", m)
	}
}

// OrderClient is a client for the Order schema.
type OrderClient struct {
	config
}

// NewOrderClient returns a client for the Order from the given config.
func NewOrderClient(c config) *OrderClient {
	return &OrderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `order.Hooks(f(g(h())))`.
func (c *OrderClient) Use(hooks ...Hook) {
	c.hooks.Order = append(c.hooks.Order, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `order.Intercept(f(g(h())))`.
func (c *OrderClient) Intercept(interceptors ...Interceptor) {
	c.inters.Order = append(c.inters.Order, interceptors...)
}

// Create returns a builder for creating a Order entity.
func (c *OrderClient) Create() *OrderCreate {
	mutation := newOrderMutation(c.config, OpCreate)
	return &OrderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Order entities.
func (c *OrderClient) CreateBulk(builders ...*OrderCreate) *OrderCreateBulk {
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Order.
func (c *OrderClient) Update() *OrderUpdate {
	mutation := newOrderMutation(c.config, OpUpdate)
	return &OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderClient) UpdateOne(o *Order) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrder(o))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderClient) UpdateOneID(id int) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrderID(id))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Order.
func (c *OrderClient) Delete() *OrderDelete {
	mutation := newOrderMutation(c.config, OpDelete)
	return &OrderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderClient) DeleteOne(o *Order) *OrderDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderClient) DeleteOneID(id int) *OrderDeleteOne {
	builder := c.Delete().Where(order.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderDeleteOne{builder}
}

// Query returns a query builder for Order.
func (c *OrderClient) Query() *OrderQuery {
	return &OrderQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a Order entity by its id.
func (c *OrderClient) Get(ctx context.Context, id int) (*Order, error) {
	return c.Query().Where(order.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderClient) GetX(ctx context.Context, id int) *Order {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OrderClient) Hooks() []Hook {
	return c.hooks.Order
}

// Interceptors returns the client interceptors.
func (c *OrderClient) Interceptors() []Interceptor {
	return c.inters.Order
}

func (c *OrderClient) mutate(ctx context.Context, m *OrderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entities: unknown Order mutation op: %q", m.Op())
	}
}

// OutboxClient is a client for the Outbox schema.
type OutboxClient struct {
	config
}

// NewOutboxClient returns a client for the Outbox from the given config.
func NewOutboxClient(c config) *OutboxClient {
	return &OutboxClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outbox.Hooks(f(g(h())))`.
func (c *OutboxClient) Use(hooks ...Hook) {
	c.hooks.Outbox = append(c.hooks.Outbox, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `outbox.Intercept(f(g(h())))`.
func (c *OutboxClient) Intercept(interceptors ...Interceptor) {
	c.inters.Outbox = append(c.inters.Outbox, interceptors...)
}

// Create returns a builder for creating a Outbox entity.
func (c *OutboxClient) Create() *OutboxCreate {
	mutation := newOutboxMutation(c.config, OpCreate)
	return &OutboxCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Outbox entities.
func (c *OutboxClient) CreateBulk(builders ...*OutboxCreate) *OutboxCreateBulk {
	return &OutboxCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Outbox.
func (c *OutboxClient) Update() *OutboxUpdate {
	mutation := newOutboxMutation(c.config, OpUpdate)
	return &OutboxUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutboxClient) UpdateOne(o *Outbox) *OutboxUpdateOne {
	mutation := newOutboxMutation(c.config, OpUpdateOne, withOutbox(o))
	return &OutboxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutboxClient) UpdateOneID(id int) *OutboxUpdateOne {
	mutation := newOutboxMutation(c.config, OpUpdateOne, withOutboxID(id))
	return &OutboxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Outbox.
func (c *OutboxClient) Delete() *OutboxDelete {
	mutation := newOutboxMutation(c.config, OpDelete)
	return &OutboxDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OutboxClient) DeleteOne(o *Outbox) *OutboxDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OutboxClient) DeleteOneID(id int) *OutboxDeleteOne {
	builder := c.Delete().Where(outbox.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutboxDeleteOne{builder}
}

// Query returns a query builder for Outbox.
func (c *OutboxClient) Query() *OutboxQuery {
	return &OutboxQuery{
		config: c.config,
		inters: c.Interceptors(),
	}
}

// Get returns a Outbox entity by its id.
func (c *OutboxClient) Get(ctx context.Context, id int) (*Outbox, error) {
	return c.Query().Where(outbox.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutboxClient) GetX(ctx context.Context, id int) *Outbox {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OutboxClient) Hooks() []Hook {
	return c.hooks.Outbox
}

// Interceptors returns the client interceptors.
func (c *OutboxClient) Interceptors() []Interceptor {
	return c.inters.Outbox
}

func (c *OutboxClient) mutate(ctx context.Context, m *OutboxMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OutboxCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OutboxUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OutboxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OutboxDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entities: unknown Outbox mutation op: %q", m.Op())
	}
}
