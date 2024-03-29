-- create "orders" table
CREATE TABLE "orders" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "order_details" jsonb NOT NULL, "preferred_exchanges" jsonb NOT NULL, PRIMARY KEY ("id"));
-- create "outboxes" table
CREATE TABLE "outboxes" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "timestamp" timestamptz NOT NULL, "topic" character varying NOT NULL, "key" character varying NOT NULL, "payload" bytea NOT NULL, "headers" jsonb NOT NULL, "retry_count" bigint NOT NULL, "status" character varying NOT NULL, "last_retry" timestamptz NULL, "processing_errors" jsonb NULL, PRIMARY KEY ("id"));
-- create index "outbox_last_retry" to table: "outboxes"
CREATE INDEX "outbox_last_retry" ON "outboxes" ("last_retry");
-- create index "outbox_status" to table: "outboxes"
CREATE INDEX "outbox_status" ON "outboxes" ("status");
