// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: fetch_customer_by_tg_id.sql

package db

import (
	"context"
)

const fetchCustomerByTgId = `-- name: FetchCustomerByTgId :one

SELECT
    to_jsonb(c)
FROM
    "smart-table.customers" c
WHERE
    tg_id = $1::TEXT
`

// args:
// $1 - TEXT
func (q *Queries) FetchCustomerByTgId(ctx context.Context, dollar_1 string) ([]byte, error) {
	row := q.db.QueryRow(ctx, fetchCustomerByTgId, dollar_1)
	var to_jsonb []byte
	err := row.Scan(&to_jsonb)
	return to_jsonb, err
}
