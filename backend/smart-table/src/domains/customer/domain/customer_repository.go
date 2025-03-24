package domain

import (
	"context"

	"github.com/google/uuid"

	"github.com/smart-table/src/utils"
)

type CustomerRepository interface {
	SaveAndUpdate(ctx context.Context, customer utils.SharedRef[Customer]) error
	FindCustomerByTgID(ctx context.Context, customerTgID string) (utils.SharedRef[Customer], error)
	FindCustomer(ctx context.Context, customerUUID uuid.UUID) (utils.SharedRef[Customer], error)
}
