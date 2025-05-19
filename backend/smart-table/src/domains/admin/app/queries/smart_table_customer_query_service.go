package app

import (
	"github.com/google/uuid"
	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
)

type SmartTableCustomerQueryService interface {
	GetPlaceOrder(placeUUID uuid.UUID, orderUUID uuid.UUID) (defsInternalCustomerDTO.OrderInfoDTO, error)
	GetPlaceOrders(placeUUID uuid.UUID, isActive bool) ([]defsInternalCustomerDTO.OrderMainInfoDTO, error)
}
