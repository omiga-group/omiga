DO
$do$
BEGIN
	IF NOT EXISTS (SELECT 1
       FROM pg_catalog.pg_constraint con
            INNER JOIN pg_catalog.pg_class rel
                       ON rel.oid = con.conrelid
            INNER JOIN pg_catalog.pg_namespace nsp
                       ON nsp.oid = connamespace
       WHERE nsp.nspname = 'public'
             AND rel.relname = 'currencies'
             AND con.conname = 'currencies_symbol_key') THEN
		ALTER TABLE public.currencies 
		ADD CONSTRAINT currencies_symbol_key
		UNIQUE (symbol); 
	END IF;

	IF NOT EXISTS (SELECT 1
       FROM pg_catalog.pg_constraint con
            INNER JOIN pg_catalog.pg_class rel
                       ON rel.oid = con.conrelid
            INNER JOIN pg_catalog.pg_namespace nsp
                       ON nsp.oid = connamespace
       WHERE nsp.nspname = 'public'
             AND rel.relname = 'venues'
             AND con.conname = 'venues_venue_id_key') THEN
		ALTER TABLE public.venues 
		ADD CONSTRAINT venues_venue_id_key
		UNIQUE (venue_id); 
	END IF;

	IF NOT EXISTS (SELECT 1
       FROM pg_catalog.pg_constraint con
            INNER JOIN pg_catalog.pg_class rel
                       ON rel.oid = con.conrelid
            INNER JOIN pg_catalog.pg_namespace nsp
                       ON nsp.oid = connamespace
       WHERE nsp.nspname = 'public'
             AND rel.relname = 'tickers'
             AND con.conname = 'tickers_base_counter_venue_ticker_key') THEN
		ALTER TABLE public.tickers 
		ADD CONSTRAINT tickers_base_counter_venue_ticker_key
		UNIQUE (base, counter, venue_ticker); 
	END IF;

	IF NOT EXISTS (SELECT 1
       FROM pg_catalog.pg_constraint con
            INNER JOIN pg_catalog.pg_class rel
                       ON rel.oid = con.conrelid
            INNER JOIN pg_catalog.pg_namespace nsp
                       ON nsp.oid = connamespace
       WHERE nsp.nspname = 'public'
             AND rel.relname = 'trading_pairs'
             AND con.conname = 'trading_pairs_symbol_venue_trading_pair_key') THEN
		ALTER TABLE public.trading_pairs 
		ADD CONSTRAINT trading_pairs_symbol_venue_trading_pair_key
		UNIQUE (symbol, venue_trading_pair); 
	END IF;
END
$do$
