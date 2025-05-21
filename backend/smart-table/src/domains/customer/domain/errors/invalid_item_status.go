package domainerrors

import (
	"fmt"

	defsInternalItem "github.com/smart-table/src/codegen/intern/item"
)

type InvalidItemStatus struct {
	ItemStatus defsInternalItem.ItemStatus
}

func (e InvalidItemStatus) Error() string {
	return fmt.Sprintf("invalid item_status='%s'", e.ItemStatus)
}
