package domain

import (
	"context"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
)

type UserRepository interface {
	Save(ctx context.Context, user utils.SharedRef[User]) error
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error

	FindUser(ctx context.Context, userLogin string) (utils.SharedRef[User], error)
	CheckLoginOrTgLoginExist(ctx context.Context, userLogin string, tgLogin string) (bool, error)
}
