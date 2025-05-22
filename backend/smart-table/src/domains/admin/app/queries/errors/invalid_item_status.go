package errors

import (
	"fmt"

	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	"github.com/smart-table/src/utils"
)

type InvalidItemStatus struct {
	ItemEditGpoup utils.Optional[defsInternalCustomerDTO.ItemEditGroupDTO]
}

func (e InvalidItemStatus) Error() string {
	if e.ItemEditGpoup.HasValue() {
		return "invalid item_status='nil'"
	}

	return fmt.Sprintf("invalid item_status='%s'", e.ItemEditGpoup.Value().ItemStatus)
}
