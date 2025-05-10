-- name: FetchCustomer :one
--
-- args:
-- $1 - UUID (customer_uuid)

SELECT
    to_jsonb(c)
FROM
    smart_table_customer.customers c
WHERE
    uuid = $1::UUID;
