// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: fetch_user_by_login.sql

package db

import (
	"context"
)

const fetchUserByLogin = `-- name: FetchUserByLogin :one

SELECT
    to_jsonb(u)
FROM
    smart_table_admin.users u
WHERE
    u.login = $1::TEXT
`

// args:
// $1 - TEXT
func (q *Queries) FetchUserByLogin(ctx context.Context, dollar_1 string) ([]byte, error) {
	row := q.db.QueryRow(ctx, fetchUserByLogin, dollar_1)
	var to_jsonb []byte
	err := row.Scan(&to_jsonb)
	return to_jsonb, err
}
