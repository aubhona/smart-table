// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: insert_order.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const insertOrder = `-- name: InsertOrder :one

INSERT INTO smart_table_customer.orders (
    uuid,
    room_code,
    table_id,
    customers_uuid,
    host_user_uuid,
    status,
    resolution,
    created_at,
    updated_at
)
SELECT
    input.uuid,
    input.room_code,
    input.table_id,
    input.customers_uuid,
    input.host_user_uuid,
    input.status,
    input.resolution,
    input.created_at,
    input.updated_at
FROM jsonb_to_record($1::jsonb) AS input(
  uuid            UUID,
  room_code       TEXT,
  table_id        TEXT,
  customers_uuid  UUID[],
  host_user_uuid  UUID,
  status          TEXT,
  resolution      TEXT,
  created_at      TIMESTAMPTZ,
  updated_at      TIMESTAMPTZ
)
RETURNING uuid
`

// args:
// $1 - JSONB
func (q *Queries) InsertOrder(ctx context.Context, dollar_1 []byte) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertOrder, dollar_1)
	var uuid uuid.UUID
	err := row.Scan(&uuid)
	return uuid, err
}
