// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: get_active_order_uuid_by_table_id.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getActiveOrderUuidByTableId = `-- name: GetActiveOrderUuidByTableId :one

SELECT
    uuid
FROM
    "smart-table.orders"
WHERE
    table_id = $1::TEXT
    AND resolution IS NULL
`

// args:
// $1 - TEXT
func (q *Queries) GetActiveOrderUuidByTableId(ctx context.Context, dollar_1 string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getActiveOrderUuidByTableId, dollar_1)
	var uuid uuid.UUID
	err := row.Scan(&uuid)
	return uuid, err
}
