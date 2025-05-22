package domainerrors

import (
	"fmt"

	defsInternalItem "github.com/smart-table/src/codegen/intern/item"
)

type ItemStatusChangeRequiresOrderStatusUpdate struct {
	ItemStatus defsInternalItem.ItemStatus
}

func (e ItemStatusChangeRequiresOrderStatusUpdate) Error() string {
	return fmt.Sprintf("item_status='%s'  can only be changed with a full order status update", e.ItemStatus)
}
