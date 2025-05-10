-- name: GetPlaceUUIDsByRestaurantUUID :many
--
-- args:
-- $1 - UUID (restaurant_uuid)

SELECT
    p.uuid
FROM
    smart_table_admin.places p
WHERE
    p.restaurant_uuid = $1::UUID;
