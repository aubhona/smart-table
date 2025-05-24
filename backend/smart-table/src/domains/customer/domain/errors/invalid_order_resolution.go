package domainerrors

import (
	"fmt"

	defsInternalOrder "github.com/smart-table/src/codegen/intern/order"
)

type InvalidOrderResolution struct {
	OrderResolution defsInternalOrder.OrderResolution
}

func (e InvalidOrderResolution) Error() string {
	return fmt.Sprintf("invalid order_resolution='%s'", e.OrderResolution)
}
