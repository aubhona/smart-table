-- name: FetchRestaurantListByOwnerUUID :one
--
-- args:
-- $1 - UUID

SELECT
    jsonb_agg(to_jsonb(r))
FROM
    smart_table_admin.restaurants r
WHERE
    r.owner_uuid = $1::UUID;
