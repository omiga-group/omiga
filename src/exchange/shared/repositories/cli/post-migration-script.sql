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
             AND rel.relname = 'tickers'
             AND con.conname = 'tickers_base_target_exchange_ticker_key') THEN
		ALTER TABLE public.tickers 
		ADD CONSTRAINT tickers_base_target_exchange_ticker_key
		UNIQUE (base, target, exchange_ticker); 
	END IF;

	COMMIT;
END
$do$
