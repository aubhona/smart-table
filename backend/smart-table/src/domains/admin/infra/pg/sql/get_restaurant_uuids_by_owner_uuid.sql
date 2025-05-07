-- name: GetRestaurantUUIDsByOwnerUUID :many
--
-- args:
-- $1 - UUID (owner_uuid)

SELECT
    r.uuid
FROM
    smart_table_admin.restaurants r
WHERE
    r.owner_uuid = $1::UUID;
