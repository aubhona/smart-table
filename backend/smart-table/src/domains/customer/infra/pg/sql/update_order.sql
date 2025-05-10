-- name: UpdateOrder :exec
--
-- args:
-- $1 - JSONB (PgOrder)

UPDATE smart_table_customer.orders
SET
    customers_uuid = input.customers_uuid,
    host_user_uuid = input.host_user_uuid,
    status = input.status,
    resolution = input.resolution
FROM jsonb_to_record($1::jsonb) AS input(
  uuid            UUID,
  room_code       TEXT,
  table_id        TEXT,
  customers_uuid  UUID[],
  host_user_uuid  UUID,
  status          TEXT,
  resolution      TEXT,
  created_at      TIMESTAMPTZ,
  updated_at      TIMESTAMPTZ
)
WHERE orders.uuid = input.uuid;
