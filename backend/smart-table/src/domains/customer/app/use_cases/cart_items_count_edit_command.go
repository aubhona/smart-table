package app

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type CartItemsCountEditCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
	DishUUID     uuid.UUID
	EditCount    int
	Comment      utils.Optional[string]
}
