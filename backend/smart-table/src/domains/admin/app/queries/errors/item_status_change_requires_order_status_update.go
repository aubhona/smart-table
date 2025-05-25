package errors

import (
	"fmt"

	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	"github.com/smart-table/src/utils"
)

type ItemStatusChangeRequiresOrderStatusUpdate struct {
	ItemEditGpoup utils.Optional[defsInternalCustomerDTO.ItemEditGroupDTO]
}

func (e ItemStatusChangeRequiresOrderStatusUpdate) Error() string {
	if !e.ItemEditGpoup.HasValue() {
		return "item_status='nil'  can only be changed with a full order status update"
	}

	return fmt.Sprintf("item_status='%s'  can only be changed with a full order status update", e.ItemEditGpoup.Value().ItemStatus)
}
