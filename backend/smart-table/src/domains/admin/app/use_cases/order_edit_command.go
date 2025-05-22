package app

import (
	"github.com/google/uuid"

	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	"github.com/smart-table/src/utils"
)

type OrderEditCommand struct {
	UserUUID      uuid.UUID
	OrderUUID     uuid.UUID
	PlaceUUID     uuid.UUID
	TableNumber   int
	OrderStatus   utils.Optional[string]
	ItemEditGpoup utils.Optional[defsInternalCustomerDTO.ItemEditGroupDTO]
}
