package domainerrors

import (
	"fmt"

	defsInternalItem "github.com/smart-table/src/codegen/intern/item"
)

type InvalidItemResolution struct {
	ItemResolution defsInternalItem.ItemResolution
}

func (e InvalidItemResolution) Error() string {
	return fmt.Sprintf("invalid item_resolution='%s'", e.ItemResolution)
}
