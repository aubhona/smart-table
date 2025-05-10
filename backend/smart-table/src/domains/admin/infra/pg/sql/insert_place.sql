-- name: InsertPlace :one
-- 
-- args:
-- $1 - JSONB (PgPlace)

INSERT INTO smart_table_admin.places (
    uuid,
    restaurant_uuid,
    address,
    opening_time,
    closing_time,
    table_count,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.restaurant_uuid,
    input.address,
    input.opening_time,
    input.closing_time,
    input.table_count,
    input.created_at,
    input.updated_at
FROM jsonb_to_record($1::jsonb) AS input(
    uuid            UUID,
    restaurant_uuid UUID,
    address         TEXT,
    opening_time    TIME,
    closing_time    TIME,
    table_count     INT,
    created_at      TIMESTAMPTZ,
    updated_at      TIMESTAMPTZ
)             
RETURNING uuid;
