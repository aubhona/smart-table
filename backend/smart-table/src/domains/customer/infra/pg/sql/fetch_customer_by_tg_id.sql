-- name: FetchCustomerByTgId :one
--
-- args:
-- $1 - TEXT

SELECT
    to_jsonb(c)
FROM
    "smart-table.customers" c
WHERE
    tg_id = $1::TEXT;
