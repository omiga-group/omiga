// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/omiga-group/omiga/src/venue/shared/entities/migrate"

	"github.com/omiga-group/omiga/src/venue/shared/entities/currency"
	"github.com/omiga-group/omiga/src/venue/shared/entities/market"
	"github.com/omiga-group/omiga/src/venue/shared/entities/outbox"
	"github.com/omiga-group/omiga/src/venue/shared/entities/ticker"
	"github.com/omiga-group/omiga/src/venue/shared/entities/tradingpair"
	"github.com/omiga-group/omiga/src/venue/shared/entities/venue"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Currency is the client for interacting with the Currency builders.
	Currency *CurrencyClient
	// Market is the client for interacting with the Market builders.
	Market *MarketClient
	// Outbox is the client for interacting with the Outbox builders.
	Outbox *OutboxClient
	// Ticker is the client for interacting with the Ticker builders.
	Ticker *TickerClient
	// TradingPair is the client for interacting with the TradingPair builders.
	TradingPair *TradingPairClient
	// Venue is the client for interacting with the Venue builders.
	Venue *VenueClient
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
	c.Currency = NewCurrencyClient(c.config)
	c.Market = NewMarketClient(c.config)
	c.Outbox = NewOutboxClient(c.config)
	c.Ticker = NewTickerClient(c.config)
	c.TradingPair = NewTradingPairClient(c.config)
	c.Venue = NewVenueClient(c.config)
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
		ctx:         ctx,
		config:      cfg,
		Currency:    NewCurrencyClient(cfg),
		Market:      NewMarketClient(cfg),
		Outbox:      NewOutboxClient(cfg),
		Ticker:      NewTickerClient(cfg),
		TradingPair: NewTradingPairClient(cfg),
		Venue:       NewVenueClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		Currency:    NewCurrencyClient(cfg),
		Market:      NewMarketClient(cfg),
		Outbox:      NewOutboxClient(cfg),
		Ticker:      NewTickerClient(cfg),
		TradingPair: NewTradingPairClient(cfg),
		Venue:       NewVenueClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Currency.
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
	c.Currency.Use(hooks...)
	c.Market.Use(hooks...)
	c.Outbox.Use(hooks...)
	c.Ticker.Use(hooks...)
	c.TradingPair.Use(hooks...)
	c.Venue.Use(hooks...)
}

// CurrencyClient is a client for the Currency schema.
type CurrencyClient struct {
	config
}

// NewCurrencyClient returns a client for the Currency from the given config.
func NewCurrencyClient(c config) *CurrencyClient {
	return &CurrencyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `currency.Hooks(f(g(h())))`.
func (c *CurrencyClient) Use(hooks ...Hook) {
	c.hooks.Currency = append(c.hooks.Currency, hooks...)
}

// Create returns a builder for creating a Currency entity.
func (c *CurrencyClient) Create() *CurrencyCreate {
	mutation := newCurrencyMutation(c.config, OpCreate)
	return &CurrencyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Currency entities.
func (c *CurrencyClient) CreateBulk(builders ...*CurrencyCreate) *CurrencyCreateBulk {
	return &CurrencyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Currency.
func (c *CurrencyClient) Update() *CurrencyUpdate {
	mutation := newCurrencyMutation(c.config, OpUpdate)
	return &CurrencyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CurrencyClient) UpdateOne(cu *Currency) *CurrencyUpdateOne {
	mutation := newCurrencyMutation(c.config, OpUpdateOne, withCurrency(cu))
	return &CurrencyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CurrencyClient) UpdateOneID(id int) *CurrencyUpdateOne {
	mutation := newCurrencyMutation(c.config, OpUpdateOne, withCurrencyID(id))
	return &CurrencyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Currency.
func (c *CurrencyClient) Delete() *CurrencyDelete {
	mutation := newCurrencyMutation(c.config, OpDelete)
	return &CurrencyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CurrencyClient) DeleteOne(cu *Currency) *CurrencyDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CurrencyClient) DeleteOneID(id int) *CurrencyDeleteOne {
	builder := c.Delete().Where(currency.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CurrencyDeleteOne{builder}
}

// Query returns a query builder for Currency.
func (c *CurrencyClient) Query() *CurrencyQuery {
	return &CurrencyQuery{
		config: c.config,
	}
}

// Get returns a Currency entity by its id.
func (c *CurrencyClient) Get(ctx context.Context, id int) (*Currency, error) {
	return c.Query().Where(currency.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CurrencyClient) GetX(ctx context.Context, id int) *Currency {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCurrencyBase queries the currency_base edge of a Currency.
func (c *CurrencyClient) QueryCurrencyBase(cu *Currency) *TradingPairQuery {
	query := &TradingPairQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(currency.Table, currency.FieldID, id),
			sqlgraph.To(tradingpair.Table, tradingpair.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, currency.CurrencyBaseTable, currency.CurrencyBaseColumn),
		)
		schemaConfig := cu.schemaConfig
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.TradingPair
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCurrencyCounter queries the currency_counter edge of a Currency.
func (c *CurrencyClient) QueryCurrencyCounter(cu *Currency) *TradingPairQuery {
	query := &TradingPairQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(currency.Table, currency.FieldID, id),
			sqlgraph.To(tradingpair.Table, tradingpair.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, currency.CurrencyCounterTable, currency.CurrencyCounterColumn),
		)
		schemaConfig := cu.schemaConfig
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.TradingPair
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CurrencyClient) Hooks() []Hook {
	return c.hooks.Currency
}

// MarketClient is a client for the Market schema.
type MarketClient struct {
	config
}

// NewMarketClient returns a client for the Market from the given config.
func NewMarketClient(c config) *MarketClient {
	return &MarketClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `market.Hooks(f(g(h())))`.
func (c *MarketClient) Use(hooks ...Hook) {
	c.hooks.Market = append(c.hooks.Market, hooks...)
}

// Create returns a builder for creating a Market entity.
func (c *MarketClient) Create() *MarketCreate {
	mutation := newMarketMutation(c.config, OpCreate)
	return &MarketCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Market entities.
func (c *MarketClient) CreateBulk(builders ...*MarketCreate) *MarketCreateBulk {
	return &MarketCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Market.
func (c *MarketClient) Update() *MarketUpdate {
	mutation := newMarketMutation(c.config, OpUpdate)
	return &MarketUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MarketClient) UpdateOne(m *Market) *MarketUpdateOne {
	mutation := newMarketMutation(c.config, OpUpdateOne, withMarket(m))
	return &MarketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MarketClient) UpdateOneID(id int) *MarketUpdateOne {
	mutation := newMarketMutation(c.config, OpUpdateOne, withMarketID(id))
	return &MarketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Market.
func (c *MarketClient) Delete() *MarketDelete {
	mutation := newMarketMutation(c.config, OpDelete)
	return &MarketDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MarketClient) DeleteOne(m *Market) *MarketDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MarketClient) DeleteOneID(id int) *MarketDeleteOne {
	builder := c.Delete().Where(market.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MarketDeleteOne{builder}
}

// Query returns a query builder for Market.
func (c *MarketClient) Query() *MarketQuery {
	return &MarketQuery{
		config: c.config,
	}
}

// Get returns a Market entity by its id.
func (c *MarketClient) Get(ctx context.Context, id int) (*Market, error) {
	return c.Query().Where(market.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MarketClient) GetX(ctx context.Context, id int) *Market {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVenue queries the venue edge of a Market.
func (c *MarketClient) QueryVenue(m *Market) *VenueQuery {
	query := &VenueQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(market.Table, market.FieldID, id),
			sqlgraph.To(venue.Table, venue.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, market.VenueTable, market.VenueColumn),
		)
		schemaConfig := m.schemaConfig
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.Market
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTradingPair queries the trading_pair edge of a Market.
func (c *MarketClient) QueryTradingPair(m *Market) *TradingPairQuery {
	query := &TradingPairQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(market.Table, market.FieldID, id),
			sqlgraph.To(tradingpair.Table, tradingpair.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, market.TradingPairTable, market.TradingPairPrimaryKey...),
		)
		schemaConfig := m.schemaConfig
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.MarketTradingPair
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MarketClient) Hooks() []Hook {
	return c.hooks.Market
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

// TickerClient is a client for the Ticker schema.
type TickerClient struct {
	config
}

// NewTickerClient returns a client for the Ticker from the given config.
func NewTickerClient(c config) *TickerClient {
	return &TickerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ticker.Hooks(f(g(h())))`.
func (c *TickerClient) Use(hooks ...Hook) {
	c.hooks.Ticker = append(c.hooks.Ticker, hooks...)
}

// Create returns a builder for creating a Ticker entity.
func (c *TickerClient) Create() *TickerCreate {
	mutation := newTickerMutation(c.config, OpCreate)
	return &TickerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ticker entities.
func (c *TickerClient) CreateBulk(builders ...*TickerCreate) *TickerCreateBulk {
	return &TickerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ticker.
func (c *TickerClient) Update() *TickerUpdate {
	mutation := newTickerMutation(c.config, OpUpdate)
	return &TickerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TickerClient) UpdateOne(t *Ticker) *TickerUpdateOne {
	mutation := newTickerMutation(c.config, OpUpdateOne, withTicker(t))
	return &TickerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TickerClient) UpdateOneID(id int) *TickerUpdateOne {
	mutation := newTickerMutation(c.config, OpUpdateOne, withTickerID(id))
	return &TickerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ticker.
func (c *TickerClient) Delete() *TickerDelete {
	mutation := newTickerMutation(c.config, OpDelete)
	return &TickerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TickerClient) DeleteOne(t *Ticker) *TickerDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TickerClient) DeleteOneID(id int) *TickerDeleteOne {
	builder := c.Delete().Where(ticker.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TickerDeleteOne{builder}
}

// Query returns a query builder for Ticker.
func (c *TickerClient) Query() *TickerQuery {
	return &TickerQuery{
		config: c.config,
	}
}

// Get returns a Ticker entity by its id.
func (c *TickerClient) Get(ctx context.Context, id int) (*Ticker, error) {
	return c.Query().Where(ticker.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TickerClient) GetX(ctx context.Context, id int) *Ticker {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVenue queries the venue edge of a Ticker.
func (c *TickerClient) QueryVenue(t *Ticker) *VenueQuery {
	query := &VenueQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticker.Table, ticker.FieldID, id),
			sqlgraph.To(venue.Table, venue.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticker.VenueTable, ticker.VenueColumn),
		)
		schemaConfig := t.schemaConfig
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.Ticker
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TickerClient) Hooks() []Hook {
	return c.hooks.Ticker
}

// TradingPairClient is a client for the TradingPair schema.
type TradingPairClient struct {
	config
}

// NewTradingPairClient returns a client for the TradingPair from the given config.
func NewTradingPairClient(c config) *TradingPairClient {
	return &TradingPairClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tradingpair.Hooks(f(g(h())))`.
func (c *TradingPairClient) Use(hooks ...Hook) {
	c.hooks.TradingPair = append(c.hooks.TradingPair, hooks...)
}

// Create returns a builder for creating a TradingPair entity.
func (c *TradingPairClient) Create() *TradingPairCreate {
	mutation := newTradingPairMutation(c.config, OpCreate)
	return &TradingPairCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TradingPair entities.
func (c *TradingPairClient) CreateBulk(builders ...*TradingPairCreate) *TradingPairCreateBulk {
	return &TradingPairCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TradingPair.
func (c *TradingPairClient) Update() *TradingPairUpdate {
	mutation := newTradingPairMutation(c.config, OpUpdate)
	return &TradingPairUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TradingPairClient) UpdateOne(tp *TradingPair) *TradingPairUpdateOne {
	mutation := newTradingPairMutation(c.config, OpUpdateOne, withTradingPair(tp))
	return &TradingPairUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TradingPairClient) UpdateOneID(id int) *TradingPairUpdateOne {
	mutation := newTradingPairMutation(c.config, OpUpdateOne, withTradingPairID(id))
	return &TradingPairUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TradingPair.
func (c *TradingPairClient) Delete() *TradingPairDelete {
	mutation := newTradingPairMutation(c.config, OpDelete)
	return &TradingPairDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TradingPairClient) DeleteOne(tp *TradingPair) *TradingPairDeleteOne {
	return c.DeleteOneID(tp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TradingPairClient) DeleteOneID(id int) *TradingPairDeleteOne {
	builder := c.Delete().Where(tradingpair.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TradingPairDeleteOne{builder}
}

// Query returns a query builder for TradingPair.
func (c *TradingPairClient) Query() *TradingPairQuery {
	return &TradingPairQuery{
		config: c.config,
	}
}

// Get returns a TradingPair entity by its id.
func (c *TradingPairClient) Get(ctx context.Context, id int) (*TradingPair, error) {
	return c.Query().Where(tradingpair.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TradingPairClient) GetX(ctx context.Context, id int) *TradingPair {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVenue queries the venue edge of a TradingPair.
func (c *TradingPairClient) QueryVenue(tp *TradingPair) *VenueQuery {
	query := &VenueQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tradingpair.Table, tradingpair.FieldID, id),
			sqlgraph.To(venue.Table, venue.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tradingpair.VenueTable, tradingpair.VenueColumn),
		)
		schemaConfig := tp.schemaConfig
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.TradingPair
		fromV = sqlgraph.Neighbors(tp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBase queries the base edge of a TradingPair.
func (c *TradingPairClient) QueryBase(tp *TradingPair) *CurrencyQuery {
	query := &CurrencyQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tradingpair.Table, tradingpair.FieldID, id),
			sqlgraph.To(currency.Table, currency.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tradingpair.BaseTable, tradingpair.BaseColumn),
		)
		schemaConfig := tp.schemaConfig
		step.To.Schema = schemaConfig.Currency
		step.Edge.Schema = schemaConfig.TradingPair
		fromV = sqlgraph.Neighbors(tp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCounter queries the counter edge of a TradingPair.
func (c *TradingPairClient) QueryCounter(tp *TradingPair) *CurrencyQuery {
	query := &CurrencyQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tradingpair.Table, tradingpair.FieldID, id),
			sqlgraph.To(currency.Table, currency.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tradingpair.CounterTable, tradingpair.CounterColumn),
		)
		schemaConfig := tp.schemaConfig
		step.To.Schema = schemaConfig.Currency
		step.Edge.Schema = schemaConfig.TradingPair
		fromV = sqlgraph.Neighbors(tp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMarket queries the market edge of a TradingPair.
func (c *TradingPairClient) QueryMarket(tp *TradingPair) *MarketQuery {
	query := &MarketQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tradingpair.Table, tradingpair.FieldID, id),
			sqlgraph.To(market.Table, market.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tradingpair.MarketTable, tradingpair.MarketPrimaryKey...),
		)
		schemaConfig := tp.schemaConfig
		step.To.Schema = schemaConfig.Market
		step.Edge.Schema = schemaConfig.MarketTradingPair
		fromV = sqlgraph.Neighbors(tp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TradingPairClient) Hooks() []Hook {
	return c.hooks.TradingPair
}

// VenueClient is a client for the Venue schema.
type VenueClient struct {
	config
}

// NewVenueClient returns a client for the Venue from the given config.
func NewVenueClient(c config) *VenueClient {
	return &VenueClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `venue.Hooks(f(g(h())))`.
func (c *VenueClient) Use(hooks ...Hook) {
	c.hooks.Venue = append(c.hooks.Venue, hooks...)
}

// Create returns a builder for creating a Venue entity.
func (c *VenueClient) Create() *VenueCreate {
	mutation := newVenueMutation(c.config, OpCreate)
	return &VenueCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Venue entities.
func (c *VenueClient) CreateBulk(builders ...*VenueCreate) *VenueCreateBulk {
	return &VenueCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Venue.
func (c *VenueClient) Update() *VenueUpdate {
	mutation := newVenueMutation(c.config, OpUpdate)
	return &VenueUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VenueClient) UpdateOne(v *Venue) *VenueUpdateOne {
	mutation := newVenueMutation(c.config, OpUpdateOne, withVenue(v))
	return &VenueUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VenueClient) UpdateOneID(id int) *VenueUpdateOne {
	mutation := newVenueMutation(c.config, OpUpdateOne, withVenueID(id))
	return &VenueUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Venue.
func (c *VenueClient) Delete() *VenueDelete {
	mutation := newVenueMutation(c.config, OpDelete)
	return &VenueDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VenueClient) DeleteOne(v *Venue) *VenueDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VenueClient) DeleteOneID(id int) *VenueDeleteOne {
	builder := c.Delete().Where(venue.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VenueDeleteOne{builder}
}

// Query returns a query builder for Venue.
func (c *VenueClient) Query() *VenueQuery {
	return &VenueQuery{
		config: c.config,
	}
}

// Get returns a Venue entity by its id.
func (c *VenueClient) Get(ctx context.Context, id int) (*Venue, error) {
	return c.Query().Where(venue.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VenueClient) GetX(ctx context.Context, id int) *Venue {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTicker queries the ticker edge of a Venue.
func (c *VenueClient) QueryTicker(v *Venue) *TickerQuery {
	query := &TickerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(venue.Table, venue.FieldID, id),
			sqlgraph.To(ticker.Table, ticker.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, venue.TickerTable, venue.TickerColumn),
		)
		schemaConfig := v.schemaConfig
		step.To.Schema = schemaConfig.Ticker
		step.Edge.Schema = schemaConfig.Ticker
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTradingPair queries the trading_pair edge of a Venue.
func (c *VenueClient) QueryTradingPair(v *Venue) *TradingPairQuery {
	query := &TradingPairQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(venue.Table, venue.FieldID, id),
			sqlgraph.To(tradingpair.Table, tradingpair.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, venue.TradingPairTable, venue.TradingPairColumn),
		)
		schemaConfig := v.schemaConfig
		step.To.Schema = schemaConfig.TradingPair
		step.Edge.Schema = schemaConfig.TradingPair
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMarket queries the market edge of a Venue.
func (c *VenueClient) QueryMarket(v *Venue) *MarketQuery {
	query := &MarketQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(venue.Table, venue.FieldID, id),
			sqlgraph.To(market.Table, market.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, venue.MarketTable, venue.MarketColumn),
		)
		schemaConfig := v.schemaConfig
		step.To.Schema = schemaConfig.Market
		step.Edge.Schema = schemaConfig.Market
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VenueClient) Hooks() []Hook {
	return c.hooks.Venue
}
