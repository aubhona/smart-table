package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/domain"
	db "github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/infra/pg/codegen"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/infra/pg/mapper"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/domain_errors"
	// "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	coonPool *pgxpool.Pool
	trx      *pgx.Tx
}

func NewOrderRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool, nil}
}

func (o *UserRepository) Save(ctx context.Context, user utils.SharedRef[domain.User]) error {
	queries := db.New(o.coonPool)

	pgUser, err := mapper.ConvertToPgUser(user)
	if err != nil {
		return err
	}

	_, err = queries.InsertUser(ctx, pgUser)

	return err
}

func (o *UserRepository) Begin(ctx context.Context) error {
	if o.trx != nil {
		return fmt.Errorf("transaction already started")
	}

	trx, err := o.coonPool.Begin(ctx)
	if err != nil {
		return err
	}

	o.trx = &trx

	return nil
}

func (o *UserRepository) Commit(ctx context.Context) error {
	if o.trx == nil {
		return fmt.Errorf("transaction doesn't exist")
	}

	err := (*o.trx).Commit(ctx)
	if err != nil {
		return err
	}

	*o.trx = nil

	return nil
}

func (o *UserRepository) FindUser(ctx context.Context, userLogin string) (utils.SharedRef[domain.User], error) {
	queries := db.New(o.coonPool)

	pgResult, err := queries.FetchUserByLogin(ctx, userLogin)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.User]{}, domain_errors.OrderNotFoundByTableId{TableId: userLogin}
		}
	
		return utils.SharedRef[domain.User]{}, err
	}

	user, err := mapper.ConvertPgUserToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.User]{}, err
	}

	return user, nil
}
