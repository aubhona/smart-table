package pg

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smart-table/src/domains/customer/domain"
	domainErrors "github.com/smart-table/src/domains/customer/domain/errors"
	db "github.com/smart-table/src/domains/customer/infra/pg/codegen"
	"github.com/smart-table/src/domains/customer/infra/pg/mapper"
	"github.com/smart-table/src/utils"
)

type CustomerRepository struct {
	coonPool *pgxpool.Pool
}

func NewCustomerRepository(pool *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{pool}
}

func (c *CustomerRepository) Begin(ctx context.Context) (pgx.Tx, error) {
	return c.coonPool.Begin(ctx)
}
func (c *CustomerRepository) Commit(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}

func (c *CustomerRepository) FindCustomerByTgIDForUpdate(ctx context.Context, tx pgx.Tx, customerTgID string) (utils.SharedRef[domain.Customer], error) {
	queries := db.New(c.coonPool).WithTx(tx)

	pgResult, err := queries.FetchCustomerByTgIdForUpdate(ctx, customerTgID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Customer]{}, domainErrors.CustomerNotFoundByTgID{TgID: customerTgID}
		}

		return utils.SharedRef[domain.Customer]{}, err
	}

	model, err := mapper.ConvertPgCustomerToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.Customer]{}, err
	}

	return model, nil
}

func (c *CustomerRepository) SaveAndUpdate(ctx context.Context, tx pgx.Tx, customer utils.SharedRef[domain.Customer]) error {
	queries := db.New(c.coonPool).WithTx(tx)

	pgCustomer, err := mapper.ConvertToPgCustomer(customer)
	if err != nil {
		return err
	}

	_, err = queries.UpsertCustomer(ctx, pgCustomer)

	return err
}

func (c *CustomerRepository) FindCustomerByTgID(ctx context.Context, customerTgID string) (utils.SharedRef[domain.Customer], error) {
	queries := db.New(c.coonPool)

	pgResult, err := queries.FetchCustomerByTgId(ctx, customerTgID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Customer]{}, domainErrors.CustomerNotFoundByTgID{TgID: customerTgID}
		}

		return utils.SharedRef[domain.Customer]{}, err
	}

	model, err := mapper.ConvertPgCustomerToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.Customer]{}, err
	}

	return model, nil
}

func (c *CustomerRepository) FindCustomer(ctx context.Context, customerUUID uuid.UUID) (utils.SharedRef[domain.Customer], error) {
	queries := db.New(c.coonPool)

	pgResult, err := queries.FetchCustomer(ctx, customerUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Customer]{}, domainErrors.CustomerNotFound{UUID: customerUUID}
		}

		return utils.SharedRef[domain.Customer]{}, err
	}

	model, err := mapper.ConvertPgCustomerToModel(pgResult)
	if err != nil {
		return utils.SharedRef[domain.Customer]{}, err
	}

	return model, nil
}
