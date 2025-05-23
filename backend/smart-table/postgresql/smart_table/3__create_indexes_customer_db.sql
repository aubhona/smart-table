-- +goose Up
-- +goose StatementBegin
BEGIN;

CREATE INDEX IF NOT EXISTS idx_orders_table_id_resolution
  ON smart_table_customer.orders (table_id, resolution);

CREATE INDEX IF NOT EXISTS idx_orders_place_uuid
  ON smart_table_customer.orders ((split_part(table_id, '_', 1)::UUID));

CREATE INDEX IF NOT EXISTS idx_orders_customers_uuid
  ON smart_table_customer.orders
  USING GIN (customers_uuid);

CREATE INDEX IF NOT EXISTS idx_items_order_uuid
  ON smart_table_customer.items (order_uuid);

CREATE INDEX IF NOT EXISTS idx_items_customer_uuid
  ON smart_table_customer.items (customer_uuid);

CREATE UNIQUE INDEX IF NOT EXISTS idx_customers_tg_id
  ON smart_table_customer.customers (tg_id); 

END;
-- +goose StatementEnd
