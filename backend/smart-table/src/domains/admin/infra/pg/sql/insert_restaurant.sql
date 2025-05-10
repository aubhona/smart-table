-- name: InsertRestaurant :one
--
-- args:
-- $1 - JSONB (PgRestaurant)

INSERT INTO smart_table_admin.restaurants (
    uuid,
    owner_uuid,
    name,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.owner_uuid,
    input.name,
    input.created_at,
    input.updated_at
FROM jsonb_to_record($1::jsonb) AS input(
  uuid        UUID,
  owner_uuid  UUID,
  name        TEXT,
  created_at  TIMESTAMPTZ,
  updated_at  TIMESTAMPTZ
)             
RETURNING uuid;
