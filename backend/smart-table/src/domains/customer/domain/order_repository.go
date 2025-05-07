package domain

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type OrderRepository interface {
	Begin() (Transaction, error)
	Commit(tx Transaction) error
	Rollback(tx Transaction) error

	Save(tx Transaction, order utils.SharedRef[Order]) error

	FindOrders(orderUUIDs []uuid.UUID) ([]utils.SharedRef[Order], error)
	FindOrder(orderUUID uuid.UUID) (utils.SharedRef[Order], error)
	FindActiveOrderByTableID(tableID string) (utils.SharedRef[Order], error)
}
