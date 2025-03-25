BEGIN;

CREATE SCHEMA IF NOT EXISTS smart_table_customer;

CREATE TABLE IF NOT EXISTS smart_table_customer.customers (
   "uuid" UUID PRIMARY KEY NOT NULL,
   "tg_id" TEXT NOT NULL,
   "tg_login" TEXT NOT NULL UNIQUE,
   "avatar_link" TEXT NOT NULL,
   "chat_id" TEXT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS smart_table_customer.orders (
    "uuid" UUID PRIMARY KEY NOT NULL,
    "room_code" TEXT NOT NULL,
    "table_id" TEXT NOT NULL,
    "customers_uuid" UUID[] NOT NULL,
    "host_user_uuid" UUID NOT NULL,
    "status" TEXT NOT NULL,
    "resolution" TEXT,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS smart_table_customer.items (
   "uuid" UUID PRIMARY KEY NOT NULL,
   "order_uuid" UUID NOT NULL,
   "comment" TEXT,
   "status" TEXT NOT NULL,
   "resolution" TEXT,
   "name" TEXT NOT NULL,
   "description" TEXT NOT NULL,
   "picture_link" TEXT NOT NULL,
   "weight" INT NOT NULL,
   "category" TEXT NOT NULL,
   "price" DECIMAL NOT NULL,
   "customer_uuid" UUID NOT NULL,
   "dish_uuid" UUID NOT NULL,
   "is_draft" BOOLEAN NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE OR REPLACE FUNCTION set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
DECLARE tbl TEXT;
BEGIN
FOR tbl IN SELECT tablename FROM pg_tables WHERE schemaname = 'smart_table_customer' LOOP
        EXECUTE format(
            'CREATE TRIGGER trigger_%s_updated_at
            BEFORE UPDATE ON %s
            FOR EACH ROW
            EXECUTE FUNCTION set_timestamp();', tbl, tbl
        );
END LOOP;
END $$;

ALTER TABLE smart_table_customer.orders ADD CONSTRAINT fk_orders_customers FOREIGN KEY ("customers_uuid") REFERENCES smart_table_customer.customers ("uuid");
ALTER TABLE smart_table_customer.items ADD CONSTRAINT fk_items_orders FOREIGN KEY ("order_uuid") REFERENCES smart_table_customer.orders ("uuid");
ALTER TABLE smart_table_customer.items ADD CONSTRAINT fk_items_customers FOREIGN KEY ("customer_uuid") REFERENCES smart_table_customer.customers ("uuid");

COMMIT;