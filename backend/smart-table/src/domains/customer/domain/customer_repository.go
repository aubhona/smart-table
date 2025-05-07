package domain

import (
	"github.com/google/uuid"

	"github.com/smart-table/src/utils"
)

type CustomerRepository interface {
	Begin() (Transaction, error)
	Commit(tx Transaction) error
	Rollback(tx Transaction) error

	SaveAndUpdate(tx Transaction, customer utils.SharedRef[Customer]) error

	FindCustomerByTgIDForUpdate(tx Transaction, customerTgID string) (utils.SharedRef[Customer], error)

	FindCustomerByTgID(customerTgID string) (utils.SharedRef[Customer], error)
	FindCustomer(customerUUID uuid.UUID) (utils.SharedRef[Customer], error)
}
