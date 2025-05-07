-- name: GetPlaceUUIDsByRestaurantUUID :many
--
-- args:
-- $1 - UUID

SELECT
    p.uuid
FROM
    smart_table_admin.places p
WHERE
    p.restaurant_uuid = $1::UUID;
