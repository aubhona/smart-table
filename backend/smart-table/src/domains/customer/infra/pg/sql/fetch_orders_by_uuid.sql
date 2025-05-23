-- name: FetchOrders :many
--
-- args:
-- $1 - UUID[] (order_uuids)

SELECT 
    jsonb_build_object(
        'order', to_jsonb(o),
        'items', (
            SELECT COALESCE(
                (
                    SELECT
                        jsonb_agg(to_jsonb(i) ORDER BY i.name)
                    FROM
                        smart_table_customer.items i
                    WHERE
                        i.order_uuid = o.uuid
                ),
                '[]' :: jsonb
            )
        ),
        'customers', (
            SELECT COALESCE(
                (
                    SELECT
                        jsonb_agg(to_jsonb(c) ORDER BY c.tg_id)
                    FROM
                        smart_table_customer.customers c
                    WHERE
                        c.uuid = ANY(o.customers_uuid)
                ),
                '[]' :: jsonb
            )
        )
) AS order_data
FROM
  smart_table_customer.orders AS o
WHERE
  o.uuid = ANY($1::UUID[])
ORDER BY o.uuid;
