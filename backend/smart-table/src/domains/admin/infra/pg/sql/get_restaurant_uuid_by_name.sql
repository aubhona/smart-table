-- name: GetRestaurantUUIDByName :one
--
-- args:
-- $1 - TEXT (name)

SELECT r.uuid
FROM
    smart_table_admin.restaurants r
WHERE
    r.name = $1::TEXT;
