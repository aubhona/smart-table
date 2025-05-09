-- name: UpsertMenuDishes :exec
--
-- args:
-- $1 - JSONB ([]PgMenuDish)

INSERT INTO smart_table_admin.menu_dishes (
    uuid,
    dish_uuid,
    place_uuid,
    price,
    exist,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.dish_uuid,
    input.place_uuid,
    input.price,
    input.exist,
    input.created_at,
    input.updated_at
FROM jsonb_to_recordset($1::jsonb) AS input(
    uuid         UUID,
    dish_uuid    UUID,
    place_uuid   UUID,
    price        DECIMAL,
    exist        BOOLEAN,
    created_at   TIMESTAMPTZ,
    updated_at   TIMESTAMPTZ
)
ON CONFLICT (uuid) DO UPDATE SET
    exist = EXCLUDED.exist;
