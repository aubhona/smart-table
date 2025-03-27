package domain

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/smart-table/src/utils"
)

type UserRepository interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context, tx pgx.Tx) error

	Save(ctx context.Context, tx pgx.Tx, user utils.SharedRef[User]) error

	FindUser(ctx context.Context, userLogin string) (utils.SharedRef[User], error)
	CheckLoginOrTgLoginExist(ctx context.Context, userLogin string, tgLogin string) (bool, error)
}
