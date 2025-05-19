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
	Update(tx Transaction, order utils.SharedRef[Order]) error

	FindOrder(orderUUID uuid.UUID) (utils.SharedRef[Order], error)
	FindOrders(orderUUIDs []uuid.UUID) ([]utils.SharedRef[Order], error)

	FindOrderForUpdate(tx Transaction, orderUUID uuid.UUID) (utils.SharedRef[Order], error)
	FindOrdersForUpdate(tx Transaction, orderUUIDs []uuid.UUID) ([]utils.SharedRef[Order], error)

	FindActiveOrderByTableIDForUpdate(tx Transaction, tableID string) (utils.SharedRef[Order], error)
	FindActiveOrderByCustomerUUID(customerUUID uuid.UUID) (utils.SharedRef[Order], error)
	FindOrdersByPlaceUUID(placeUUID uuid.UUID, isActive bool) ([]utils.SharedRef[Order], error)
}
