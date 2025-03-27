package pg

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smart-table/src/domains/admin/domain"
	domain_errors "github.com/smart-table/src/domains/admin/domain/errors"
	db "github.com/smart-table/src/domains/admin/infra/pg/codegen"
	"github.com/smart-table/src/domains/admin/infra/pg/mapper"
	"github.com/smart-table/src/utils"
)

type UserRepository struct {
	coonPool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool}
}

func (o *UserRepository) Save(ctx context.Context, tx pgx.Tx, user utils.SharedRef[domain.User]) error {
	queries := db.New(o.coonPool).WithTx(tx)

	pgUser, err := mapper.ConvertToPgUser(user)
	if err != nil {
		return err
	}

	_, err = queries.InsertUser(ctx, pgUser)

	return err
}

func (o *UserRepository) Begin(ctx context.Context) (pgx.Tx, error) {
	return o.coonPool.Begin(ctx)
}

func (o *UserRepository) Commit(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}

func (o *UserRepository) FindUser(ctx context.Context, login string) (utils.SharedRef[domain.User], error) {
	queries := db.New(o.coonPool)

	pgResult, err := queries.FetchUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.User]{}, domain_errors.UserNotFoundByLogin{Login: login}
		}

		return utils.SharedRef[domain.User]{}, err
	}

	user, err := mapper.ConvertPgUserToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.User]{}, err
	}

	return user, nil
}

func (o *UserRepository) CheckLoginOrTgLoginExist(ctx context.Context, login, tgLogin string) (bool, error) {
	queries := db.New(o.coonPool)

	params := db.CheckLoginOrTgLoginExistParams{
		Column1: login,
		Column2: tgLogin,
	}

	userExists, err := queries.CheckLoginOrTgLoginExist(ctx, params)
	if err != nil {
		return false, err
	}

	return userExists, nil
}
