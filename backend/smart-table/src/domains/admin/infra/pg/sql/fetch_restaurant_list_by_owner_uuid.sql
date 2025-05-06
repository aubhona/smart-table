-- name: FetchRestaurantListByOwnerUUID :many
--
-- args:
-- $1 - UUID

SELECT
    to_jsonb(r)
FROM
    smart_table_admin.restaurants r
WHERE
    r.owner_uuid = $1::UUID;
