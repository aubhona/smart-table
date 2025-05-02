-- name: FetchRestaurantByUUID :one
--
-- args:
-- $1 - UUID

SELECT
    to_jsonb(r)
FROM
    smart_table_admin.restaurants r
WHERE
    r.uuid = $1::UUID;
