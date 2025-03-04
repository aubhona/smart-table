-- name: GetActiveOrderUuidByTableId :one
--
-- args:
-- $1 - TEXT

SELECT
    uuid
FROM
    "smart-table.orders"
WHERE
    table_id = $1::TEXT
    AND resolution IS NULL;
