-- name: FetchPlaceListByRestaurantUUID :one
--
-- args:
-- $1 - UUID

SELECT
    jsonb_agg(to_jsonb(p))
FROM
    smart_table_admin.places p
WHERE
    p.restaurant_uuid = $1::UUID;
