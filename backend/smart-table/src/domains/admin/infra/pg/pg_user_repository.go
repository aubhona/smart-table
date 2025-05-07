package pg

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smart-table/src/domains/admin/domain"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
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

func (u *UserRepository) Begin() (domain.Transaction, error) {
	ctx := context.Background()
	tx, err := u.coonPool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	return &pgTx{tx: tx, ctx: ctx}, nil
}

func (u *UserRepository) Commit(tx domain.Transaction) error {
	return tx.Commit()
}

func (u *UserRepository) Rollback(tx domain.Transaction) error {
	return tx.Rollback()
}

func (u *UserRepository) Save(tx domain.Transaction, user utils.SharedRef[domain.User]) error {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(u.coonPool).WithTx(trx.tx)

	pgUser, err := mapper.ConvertToPgUser(user)
	if err != nil {
		return err
	}

	_, err = queries.InsertUser(ctx, pgUser)

	return err
}

func (u *UserRepository) FindUserByUUID(uuid uuid.UUID) (utils.SharedRef[domain.User], error) {
	ctx := context.Background()
	queries := db.New(u.coonPool)

	pgResult, err := queries.FetchUserByUUID(ctx, uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.User]{}, domainErrors.UserNotFoundByUUID{UUID: uuid}
		}

		return utils.SharedRef[domain.User]{}, err
	}

	user, err := mapper.ConvertPgUserToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.User]{}, err
	}

	return user, nil
}

func (u *UserRepository) FindUserByLogin(userLogin string) (utils.SharedRef[domain.User], error) {
	ctx := context.Background()
	queries := db.New(u.coonPool)

	pgResult, err := queries.FetchUserByLogin(ctx, userLogin)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.User]{}, domainErrors.UserNotFoundByLogin{Login: userLogin}
		}

		return utils.SharedRef[domain.User]{}, err
	}

	user, err := mapper.ConvertPgUserToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.User]{}, err
	}

	return user, nil
}

func (u *UserRepository) FindUserByLoginOrTgLogin(login, tgLogin string) (utils.SharedRef[domain.User], error) {
	ctx := context.Background()
	queries := db.New(u.coonPool)

	pgResult, err := queries.FetchUserByLoginOrTgLogin(ctx, db.FetchUserByLoginOrTgLoginParams{
		Column1: login,
		Column2: tgLogin,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.User]{}, domainErrors.UserNotFoundByLoginOrTgLogin{Login: login, TgLogin: tgLogin}
		}

		return utils.SharedRef[domain.User]{}, err
	}

	user, err := mapper.ConvertPgUserToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.User]{}, err
	}

	return user, nil
}
