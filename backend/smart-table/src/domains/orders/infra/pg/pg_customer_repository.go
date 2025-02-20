package pg

import (
	"context"
	"errors"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain"
	domain_errors "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain/errors"
	db "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/infra/pg/codegen"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/infra/pg/mapper"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository struct {
	coonPool *pgxpool.Pool
}

func NewCustomerRepository(pool *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{pool}
}

func (c *CustomerRepository) SaveAndUpdate(ctx context.Context, customer utils.SharedRef[domain.Customer]) error {
	queries := db.New(c.coonPool)

	pgCustomer, err := mapper.ConvertToPgCustomer(customer)
	if err != nil {
		return err
	}

	_, err = queries.UpsertCustomer(ctx, pgCustomer)

	return err
}

func (c *CustomerRepository) FindCustomerByTgID(ctx context.Context, customerTgId string) (utils.SharedRef[domain.Customer], error) {
	queries := db.New(c.coonPool)

	pgResult, err := queries.FetchCustomerByTgId(ctx, customerTgId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Customer]{}, domain_errors.CustomerNotFoundByTgID{TgId: customerTgId}
		}

		return utils.SharedRef[domain.Customer]{}, err
	}

	model, err := mapper.ConvertPgCustomerToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.Customer]{}, err
	}

	return model, nil
}

func (c *CustomerRepository) FindCustomer(ctx context.Context, customerUuid uuid.UUID) (utils.SharedRef[domain.Customer], error) {
	queries := db.New(c.coonPool)

	pgResult, err := queries.FetchCustomer(ctx, customerUuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Customer]{}, domain_errors.CustomerNotFound{Uuid: customerUuid}
		}

		return utils.SharedRef[domain.Customer]{}, err
	}

	model, err := mapper.ConvertPgCustomerToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.Customer]{}, err
	}

	return model, nil
}
