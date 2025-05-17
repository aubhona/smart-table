-- name: DeleteItems :exec
-- args:
-- $1 - JSONB ([]PgItem)

DELETE FROM smart_table_customer.items
WHERE uuid IN (
    SELECT input.uuid
    FROM jsonb_to_recordset($1::jsonb) AS input(
    uuid UUID
    )
);
