-- name: UpsertDishes :exec
--
-- args:
-- $1 - JSONB ([]PgDish)

INSERT INTO smart_table_admin.dishes (
    uuid,
    restaurant_uuid,
    name,
    description,
    calories,
    weight,
    picture_key,
    category,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.restaurant_uuid,
    input.name,
    input.description,
    input.calories,
    input.weight,
    input.picture_key,
    input.category,
    input.created_at,
    input.updated_at
FROM jsonb_to_recordset($1::jsonb) AS input(
    uuid            UUID,
    restaurant_uuid UUID,
    name            TEXT,
    description     TEXT,
    calories        INTEGER,
    weight          INTEGER,
    picture_key     TEXT,
    category        TEXT,
    created_at      TIMESTAMPTZ,
    updated_at      TIMESTAMPTZ
)
ON CONFLICT (uuid) DO NOTHING;
