package app

import (
	"github.com/google/uuid"

	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	"github.com/smart-table/src/utils"
)

type PlaceOrderEditCommand struct {
	OrderUUID     uuid.UUID
	TableID       string
	OrderStatus   utils.Optional[string]
	ItemEditGroup utils.Optional[defsInternalCustomerDTO.ItemEditGroupDTO]
}
