package errors

import (
	"fmt"

	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	"github.com/smart-table/src/utils"
)

type DraftItemStatusChangeNotAllowed struct {
	ItemEditGpoup utils.Optional[defsInternalCustomerDTO.ItemEditGroupDTO]
}

func (e DraftItemStatusChangeNotAllowed) Error() string {
	if e.ItemEditGpoup.HasValue() {
		return "cannot change status for draft item, to  item_status='nil'"
	}

	return fmt.Sprintf("cannot change status for draft item, to item_status='%s'", e.ItemEditGpoup.Value().ItemStatus)
}
