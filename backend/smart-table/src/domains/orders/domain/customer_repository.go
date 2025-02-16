package domain

import (
	"context"
	"github.com/google/uuid"

	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
)

type CustomerRepository interface {
	SaveAndUpdate(ctx context.Context, customer utils.SharedRef[Customer]) error
	FindCustomerByTgID(ctx context.Context, customerTgId string) (utils.SharedRef[Customer], error)
	FindCustomer(ctx context.Context, customerUuid uuid.UUID) (utils.SharedRef[Customer], error)
}
