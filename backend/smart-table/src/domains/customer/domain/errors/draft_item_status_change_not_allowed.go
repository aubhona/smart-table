package domainerrors

import (
	"fmt"

	"github.com/google/uuid"
	defsInternalItem "github.com/smart-table/src/codegen/intern/item"
)

type DraftItemStatusChangeNotAllowed struct {
	ItemUUID   uuid.UUID
	ItemStatus defsInternalItem.ItemStatus
}

func (e DraftItemStatusChangeNotAllowed) Error() string {
	return fmt.Sprintf("cannot change status for draft item with uuid='%s', to istatus='%s'", e.ItemUUID, e.ItemStatus)
}
