-- create "outboxes" table
CREATE TABLE "outboxes" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "timestamp" timestamptz NOT NULL, "topic" character varying NOT NULL, "key" character varying NOT NULL, "payload" bytea NOT NULL, "headers" jsonb NOT NULL, "retry_count" bigint NOT NULL, "status" character varying NOT NULL, "last_retry" timestamptz NULL, "processing_errors" jsonb NULL, PRIMARY KEY ("id"));
-- create index "outbox_last_retry" to table: "outboxes"
CREATE INDEX "outbox_last_retry" ON "outboxes" ("last_retry");
-- create index "outbox_status" to table: "outboxes"
CREATE INDEX "outbox_status" ON "outboxes" ("status");
-- create "venues" table
CREATE TABLE "venues" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "venue_id" character varying NOT NULL, "type" character varying NOT NULL, "name" character varying NULL, "year_established" bigint NULL, "country" character varying NULL, "image" character varying NULL, "links" jsonb NULL, "has_trading_incentive" boolean NULL, "centralized" boolean NULL, "public_notice" character varying NULL, "alert_notice" character varying NULL, "trust_score" bigint NULL, "trust_score_rank" bigint NULL, "trade_volume_24h_btc" double precision NULL, "trade_volume_24h_btc_normalized" double precision NULL, "maker_fee" double precision NULL, "taker_fee" double precision NULL, "spread_fee" boolean NULL, "support_api" boolean NULL, PRIMARY KEY ("id"));
-- create index "venue_venue_id" to table: "venues"
CREATE INDEX "venue_venue_id" ON "venues" ("venue_id");
-- create index "venue_type" to table: "venues"
CREATE INDEX "venue_type" ON "venues" ("type");
-- create index "venue_name" to table: "venues"
CREATE INDEX "venue_name" ON "venues" ("name");
-- create index "venue_year_established" to table: "venues"
CREATE INDEX "venue_year_established" ON "venues" ("year_established");
-- create index "venue_country" to table: "venues"
CREATE INDEX "venue_country" ON "venues" ("country");
-- create index "venue_image" to table: "venues"
CREATE INDEX "venue_image" ON "venues" ("image");
-- create index "venue_has_trading_incentive" to table: "venues"
CREATE INDEX "venue_has_trading_incentive" ON "venues" ("has_trading_incentive");
-- create index "venue_centralized" to table: "venues"
CREATE INDEX "venue_centralized" ON "venues" ("centralized");
-- create index "venue_public_notice" to table: "venues"
CREATE INDEX "venue_public_notice" ON "venues" ("public_notice");
-- create index "venue_alert_notice" to table: "venues"
CREATE INDEX "venue_alert_notice" ON "venues" ("alert_notice");
-- create index "venue_trust_score" to table: "venues"
CREATE INDEX "venue_trust_score" ON "venues" ("trust_score");
-- create index "venue_trust_score_rank" to table: "venues"
CREATE INDEX "venue_trust_score_rank" ON "venues" ("trust_score_rank");
-- create index "venue_trade_volume_24h_btc" to table: "venues"
CREATE INDEX "venue_trade_volume_24h_btc" ON "venues" ("trade_volume_24h_btc");
-- create index "venue_trade_volume_24h_btc_normalized" to table: "venues"
CREATE INDEX "venue_trade_volume_24h_btc_normalized" ON "venues" ("trade_volume_24h_btc_normalized");
-- create index "venue_maker_fee" to table: "venues"
CREATE INDEX "venue_maker_fee" ON "venues" ("maker_fee");
-- create index "venue_taker_fee" to table: "venues"
CREATE INDEX "venue_taker_fee" ON "venues" ("taker_fee");
-- create index "venue_spread_fee" to table: "venues"
CREATE INDEX "venue_spread_fee" ON "venues" ("spread_fee");
-- create index "venue_support_api" to table: "venues"
CREATE INDEX "venue_support_api" ON "venues" ("support_api");
-- create "markets" table
CREATE TABLE "markets" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "type" character varying NOT NULL, "venue_market" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "markets_venues_market" FOREIGN KEY ("venue_market") REFERENCES "venues" ("id") ON DELETE CASCADE);
-- create index "market_name" to table: "markets"
CREATE INDEX "market_name" ON "markets" ("name");
-- create index "market_type" to table: "markets"
CREATE INDEX "market_type" ON "markets" ("type");
-- create "tickers" table
CREATE TABLE "tickers" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "base" character varying NOT NULL, "base_coin_id" character varying NULL, "counter" character varying NOT NULL, "counter_coin_id" character varying NULL, "market" jsonb NULL, "last" double precision NULL, "volume" double precision NULL, "converted_last" jsonb NULL, "converted_volume" jsonb NULL, "trust_score" character varying NULL, "bid_ask_spread_percentage" double precision NULL, "timestamp" timestamptz NULL, "last_traded_at" timestamptz NULL, "last_fetch_at" timestamptz NULL, "is_anomaly" boolean NULL, "is_stale" boolean NULL, "trade_url" character varying NULL, "token_info_url" character varying NULL, "venue_ticker" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "tickers_venues_ticker" FOREIGN KEY ("venue_ticker") REFERENCES "venues" ("id") ON DELETE CASCADE);
-- create index "ticker_base" to table: "tickers"
CREATE INDEX "ticker_base" ON "tickers" ("base");
-- create index "ticker_base_coin_id" to table: "tickers"
CREATE INDEX "ticker_base_coin_id" ON "tickers" ("base_coin_id");
-- create index "ticker_counter" to table: "tickers"
CREATE INDEX "ticker_counter" ON "tickers" ("counter");
-- create index "ticker_counter_coin_id" to table: "tickers"
CREATE INDEX "ticker_counter_coin_id" ON "tickers" ("counter_coin_id");
-- create index "ticker_last" to table: "tickers"
CREATE INDEX "ticker_last" ON "tickers" ("last");
-- create index "ticker_volume" to table: "tickers"
CREATE INDEX "ticker_volume" ON "tickers" ("volume");
-- create index "ticker_trust_score" to table: "tickers"
CREATE INDEX "ticker_trust_score" ON "tickers" ("trust_score");
-- create index "ticker_bid_ask_spread_percentage" to table: "tickers"
CREATE INDEX "ticker_bid_ask_spread_percentage" ON "tickers" ("bid_ask_spread_percentage");
-- create index "ticker_timestamp" to table: "tickers"
CREATE INDEX "ticker_timestamp" ON "tickers" ("timestamp");
-- create index "ticker_last_traded_at" to table: "tickers"
CREATE INDEX "ticker_last_traded_at" ON "tickers" ("last_traded_at");
-- create index "ticker_last_fetch_at" to table: "tickers"
CREATE INDEX "ticker_last_fetch_at" ON "tickers" ("last_fetch_at");
-- create index "ticker_is_anomaly" to table: "tickers"
CREATE INDEX "ticker_is_anomaly" ON "tickers" ("is_anomaly");
-- create index "ticker_is_stale" to table: "tickers"
CREATE INDEX "ticker_is_stale" ON "tickers" ("is_stale");
-- create index "ticker_trade_url" to table: "tickers"
CREATE INDEX "ticker_trade_url" ON "tickers" ("trade_url");
-- create index "ticker_token_info_url" to table: "tickers"
CREATE INDEX "ticker_token_info_url" ON "tickers" ("token_info_url");
-- create "currencies" table
CREATE TABLE "currencies" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "symbol" character varying NOT NULL, "name" character varying NULL, "type" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "currency_symbol" to table: "currencies"
CREATE INDEX "currency_symbol" ON "currencies" ("symbol");
-- create index "currency_name" to table: "currencies"
CREATE INDEX "currency_name" ON "currencies" ("name");
-- create index "currency_type" to table: "currencies"
CREATE INDEX "currency_type" ON "currencies" ("type");
-- create "trading_pairs" table
CREATE TABLE "trading_pairs" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "symbol" character varying NOT NULL, "base_price_min_precision" bigint NULL, "base_price_max_precision" bigint NULL, "base_quantity_min_precision" bigint NULL, "base_quantity_max_precision" bigint NULL, "counter_price_min_precision" bigint NULL, "counter_price_max_precision" bigint NULL, "counter_quantity_min_precision" bigint NULL, "counter_quantity_max_precision" bigint NULL, "currency_currency_base" bigint NOT NULL, "currency_currency_counter" bigint NOT NULL, "venue_trading_pair" bigint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "trading_pairs_currencies_currency_base" FOREIGN KEY ("currency_currency_base") REFERENCES "currencies" ("id") ON DELETE NO ACTION, CONSTRAINT "trading_pairs_currencies_currency_counter" FOREIGN KEY ("currency_currency_counter") REFERENCES "currencies" ("id") ON DELETE NO ACTION, CONSTRAINT "trading_pairs_venues_trading_pair" FOREIGN KEY ("venue_trading_pair") REFERENCES "venues" ("id") ON DELETE CASCADE);
-- create index "tradingpair_symbol" to table: "trading_pairs"
CREATE INDEX "tradingpair_symbol" ON "trading_pairs" ("symbol");
-- create index "tradingpair_base_price_min_precision" to table: "trading_pairs"
CREATE INDEX "tradingpair_base_price_min_precision" ON "trading_pairs" ("base_price_min_precision");
-- create index "tradingpair_base_price_max_precision" to table: "trading_pairs"
CREATE INDEX "tradingpair_base_price_max_precision" ON "trading_pairs" ("base_price_max_precision");
-- create index "tradingpair_base_quantity_min_precision" to table: "trading_pairs"
CREATE INDEX "tradingpair_base_quantity_min_precision" ON "trading_pairs" ("base_quantity_min_precision");
-- create index "tradingpair_base_quantity_max_precision" to table: "trading_pairs"
CREATE INDEX "tradingpair_base_quantity_max_precision" ON "trading_pairs" ("base_quantity_max_precision");
-- create index "tradingpair_counter_price_min_precision" to table: "trading_pairs"
CREATE INDEX "tradingpair_counter_price_min_precision" ON "trading_pairs" ("counter_price_min_precision");
-- create index "tradingpair_counter_price_max_precision" to table: "trading_pairs"
CREATE INDEX "tradingpair_counter_price_max_precision" ON "trading_pairs" ("counter_price_max_precision");
-- create index "tradingpair_counter_quantity_min_precision" to table: "trading_pairs"
CREATE INDEX "tradingpair_counter_quantity_min_precision" ON "trading_pairs" ("counter_quantity_min_precision");
-- create index "tradingpair_counter_quantity_max_precision" to table: "trading_pairs"
CREATE INDEX "tradingpair_counter_quantity_max_precision" ON "trading_pairs" ("counter_quantity_max_precision");
-- create "market_trading_pair" table
CREATE TABLE "market_trading_pair" ("market_id" bigint NOT NULL, "trading_pair_id" bigint NOT NULL, PRIMARY KEY ("market_id", "trading_pair_id"), CONSTRAINT "market_trading_pair_market_id" FOREIGN KEY ("market_id") REFERENCES "markets" ("id") ON DELETE CASCADE, CONSTRAINT "market_trading_pair_trading_pair_id" FOREIGN KEY ("trading_pair_id") REFERENCES "trading_pairs" ("id") ON DELETE CASCADE);


-- Custom scripts
ALTER TABLE public.currencies ADD CONSTRAINT currencies_symbol_key UNIQUE (symbol); 
ALTER TABLE public.venues ADD CONSTRAINT venues_venue_id_key UNIQUE (venue_id); 
ALTER TABLE public.tickers ADD CONSTRAINT tickers_base_counter_venue_ticker_key UNIQUE (base, counter, venue_ticker); 
ALTER TABLE public.trading_pairs ADD CONSTRAINT trading_pairs_symbol_venue_trading_pair_key UNIQUE (symbol, venue_trading_pair); 