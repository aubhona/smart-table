package domain

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"

	"github.com/smart-table/src/utils"
)

type CustomerRepository interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context, tx pgx.Tx) error

	SaveAndUpdate(ctx context.Context, tx pgx.Tx, customer utils.SharedRef[Customer]) error

	FindCustomerByTgIDForUpdate(ctx context.Context, tx pgx.Tx, customerTgID string) (utils.SharedRef[Customer], error)

	FindCustomerByTgID(ctx context.Context, customerTgID string) (utils.SharedRef[Customer], error)
	FindCustomer(ctx context.Context, customerUUID uuid.UUID) (utils.SharedRef[Customer], error)
}
