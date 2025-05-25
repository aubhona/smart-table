package errors

import (
	"fmt"

	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	"github.com/smart-table/src/utils"
)

type InvalidItemResolution struct {
	ItemEditGpoup utils.Optional[defsInternalCustomerDTO.ItemEditGroupDTO]
}

func (e InvalidItemResolution) Error() string {
	if !e.ItemEditGpoup.HasValue() {
		return "invalid item_resolution='nil'"
	}

	return fmt.Sprintf("invalid item_resolution='%s'", e.ItemEditGpoup.Value().ItemStatus)
}
