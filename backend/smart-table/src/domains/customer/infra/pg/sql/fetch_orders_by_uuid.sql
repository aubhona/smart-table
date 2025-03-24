-- name: FetchOrders :many
--
-- args:
-- $1 - UUID[]

SELECT jsonb_agg(order_data)
FROM (
    SELECT jsonb_build_object(
        'order', to_jsonb(o)
    ) || jsonb_build_object(
        'items',
        (
            COALESCE((
                SELECT jsonb_agg(to_jsonb(i))
                FROM smart_table_customer.items i
                WHERE i.order_uuid = o.uuid
            ), '[]'::jsonb)
        )
    ) || jsonb_build_object(
        'customers',
        (
            COALESCE((
                SELECT jsonb_agg(to_jsonb(c))
                FROM smart_table_customer.customers c
                WHERE c.uuid = ANY(o.customers_uuid)
            ), '[]'::jsonb)
        )
    ) AS order_data
    FROM smart_table_customer.orders AS o
    WHERE o.uuid = ANY($1::UUID[])
) AS orders;

