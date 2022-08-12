// Code generated by ent, DO NOT EDIT.

package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/omiga-group/omiga/src/exchange/shared/repositories/migrate"

	"github.com/omiga-group/omiga/src/exchange/shared/repositories/exchange"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories/outbox"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Exchange is the client for interacting with the Exchange builders.
	Exchange *ExchangeClient
	// Outbox is the client for interacting with the Outbox builders.
	Outbox *OutboxClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Exchange = NewExchangeClient(c.config)
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
		return nil, errors.New("repositories: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("repositories: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Exchange: NewExchangeClient(cfg),
		Outbox:   NewOutboxClient(cfg),
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
		ctx:      ctx,
		config:   cfg,
		Exchange: NewExchangeClient(cfg),
		Outbox:   NewOutboxClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Exchange.
//		Query().
//		Count(ctx)
//
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
	c.Exchange.Use(hooks...)
	c.Outbox.Use(hooks...)
}

// ExchangeClient is a client for the Exchange schema.
type ExchangeClient struct {
	config
}

// NewExchangeClient returns a client for the Exchange from the given config.
func NewExchangeClient(c config) *ExchangeClient {
	return &ExchangeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `exchange.Hooks(f(g(h())))`.
func (c *ExchangeClient) Use(hooks ...Hook) {
	c.hooks.Exchange = append(c.hooks.Exchange, hooks...)
}

// Create returns a builder for creating a Exchange entity.
func (c *ExchangeClient) Create() *ExchangeCreate {
	mutation := newExchangeMutation(c.config, OpCreate)
	return &ExchangeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Exchange entities.
func (c *ExchangeClient) CreateBulk(builders ...*ExchangeCreate) *ExchangeCreateBulk {
	return &ExchangeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Exchange.
func (c *ExchangeClient) Update() *ExchangeUpdate {
	mutation := newExchangeMutation(c.config, OpUpdate)
	return &ExchangeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ExchangeClient) UpdateOne(e *Exchange) *ExchangeUpdateOne {
	mutation := newExchangeMutation(c.config, OpUpdateOne, withExchange(e))
	return &ExchangeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ExchangeClient) UpdateOneID(id int) *ExchangeUpdateOne {
	mutation := newExchangeMutation(c.config, OpUpdateOne, withExchangeID(id))
	return &ExchangeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Exchange.
func (c *ExchangeClient) Delete() *ExchangeDelete {
	mutation := newExchangeMutation(c.config, OpDelete)
	return &ExchangeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ExchangeClient) DeleteOne(e *Exchange) *ExchangeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ExchangeClient) DeleteOneID(id int) *ExchangeDeleteOne {
	builder := c.Delete().Where(exchange.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ExchangeDeleteOne{builder}
}

// Query returns a query builder for Exchange.
func (c *ExchangeClient) Query() *ExchangeQuery {
	return &ExchangeQuery{
		config: c.config,
	}
}

// Get returns a Exchange entity by its id.
func (c *ExchangeClient) Get(ctx context.Context, id int) (*Exchange, error) {
	return c.Query().Where(exchange.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ExchangeClient) GetX(ctx context.Context, id int) *Exchange {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ExchangeClient) Hooks() []Hook {
	return c.hooks.Exchange
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

// DeleteOne returns a builder for deleting the given entity by its id.
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
