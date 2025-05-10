-- name: GetActiveOrderUUIDByCustomerUUID :one
--
-- args:
-- $1 - UUID

SELECT
    uuid
FROM
    smart_table_customer.orders
WHERE
    $1::UUID = ANY (customers_uuid)
    AND resolution IS NULL;
