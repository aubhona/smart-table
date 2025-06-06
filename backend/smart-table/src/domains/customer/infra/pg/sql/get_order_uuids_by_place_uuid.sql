-- name: GetOrderUUIDsByPlaceUUID :many
--
-- args:
-- $1 - UUID (place_uuid)
-- $2 - boolean (is_active)

SELECT
    uuid
FROM
    smart_table_customer.orders
WHERE
    split_part(table_id, '_', 1)::UUID = $1::UUID 
    AND (
        ($2 = true AND resolution IS NULL) OR
        ($2 = false AND resolution IS NOT NULL)
    );
