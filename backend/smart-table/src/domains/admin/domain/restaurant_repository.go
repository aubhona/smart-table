package domain

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/smart-table/src/utils"
)

type RestaurantRepository interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context, tx pgx.Tx) error

	Save(ctx context.Context, tx pgx.Tx, restaurant utils.SharedRef[Restaurant]) error

	CheckNameExist(ctx context.Context, name string) (bool, error)
}
