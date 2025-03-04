-- name: FetchCustomer :one
--
-- args:
-- $1 - UUID

SELECT
    to_jsonb(c)
FROM
    "smart-table.customers" c
WHERE
    uuid = $1::UUID;
