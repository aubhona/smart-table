-- name: InsertOrder :exec
--
-- args:
-- $1 - JSONB (PgOrder)

INSERT INTO smart_table_customer.orders (
    uuid,
    room_code,
    table_id,
    customers_uuid,
    host_user_uuid,
    status,
    resolution,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.room_code,
    input.table_id,
    input.customers_uuid,
    input.host_user_uuid,
    input.status,
    input.resolution,
    input.created_at,
    input.updated_at
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
);
