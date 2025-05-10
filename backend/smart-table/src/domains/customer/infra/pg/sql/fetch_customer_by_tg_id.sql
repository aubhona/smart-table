-- name: FetchCustomerByTgId :one
--
-- args:
-- $1 - TEXT (tg_id)

SELECT
    to_jsonb(c)
FROM
    smart_table_customer.customers c
WHERE
    tg_id = $1::TEXT;
