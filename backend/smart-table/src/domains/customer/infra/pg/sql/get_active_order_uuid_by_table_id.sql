-- name: GetActiveOrderUuidByTableId :one
--
-- args:
-- $1 - TEXT

SELECT
    uuid
FROM
    smart_table_customer.orders
WHERE
    table_id = $1::TEXT
    AND resolution IS NULL;
