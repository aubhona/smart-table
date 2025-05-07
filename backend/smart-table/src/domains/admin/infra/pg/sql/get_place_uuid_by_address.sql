-- name: GetPlaceUUIDByAddress :one
--
-- args:
-- $1 - TEXT (address)
-- S2 - UUID (restaurant_uuid)

SELECT p.uuid
FROM
    smart_table_admin.places p
WHERE
    p.address = $1::TEXT
    AND p.restaurant_uuid = $2::UUID;
