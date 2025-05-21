package app

import (
	"github.com/google/uuid"
	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	"github.com/smart-table/src/utils"
)

type SmartTableCustomerQueryService interface {
	GetPlaceOrder(placeUUID, orderUUID uuid.UUID) (defsInternalCustomerDTO.OrderInfoDTO, error)
	GetPlaceOrders(placeUUID uuid.UUID, isActive bool) ([]defsInternalCustomerDTO.OrderMainInfoDTO, error)
	EditPlaceOrder(
		orderUUID uuid.UUID,
		tableID string,
		orderStatus utils.Optional[string],
		itemEditGpoup utils.Optional[defsInternalCustomerDTO.ItemEditGroupDTO],
	) error
}
