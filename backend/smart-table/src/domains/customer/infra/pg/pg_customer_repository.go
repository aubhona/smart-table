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

func (c *CustomerRepository) Begin() (domain.Transaction, error) {
	ctx := context.Background()
	tx, err := c.coonPool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	return &pgTx{tx: tx, ctx: ctx}, nil
}

func (c *CustomerRepository) Commit(tx domain.Transaction) error {
	return tx.Commit()
}

func (c *CustomerRepository) Rollback(tx domain.Transaction) error {
	return tx.Rollback()
}

func (c *CustomerRepository) FindCustomerByTgIDForUpdate(
	tx domain.Transaction,
	customerTgID string,
) (utils.SharedRef[domain.Customer], error) {
	ctx := context.Background()
	trx := tx.(*pgTx)

	queries := db.New(c.coonPool).WithTx(trx.tx)

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

func (c *CustomerRepository) SaveAndUpdate(tx domain.Transaction, customer utils.SharedRef[domain.Customer]) error {
	ctx := context.Background()
	trx := tx.(*pgTx)

	queries := db.New(c.coonPool).WithTx(trx.tx)

	pgCustomer, err := mapper.ConvertToPgCustomer(customer)
	if err != nil {
		return err
	}

	_, err = queries.UpsertCustomer(ctx, pgCustomer)

	return err
}

func (c *CustomerRepository) FindCustomerByTgID(customerTgID string) (utils.SharedRef[domain.Customer], error) {
	ctx := context.Background()
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

func (c *CustomerRepository) FindCustomer(customerUUID uuid.UUID) (utils.SharedRef[domain.Customer], error) {
	ctx := context.Background()
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
