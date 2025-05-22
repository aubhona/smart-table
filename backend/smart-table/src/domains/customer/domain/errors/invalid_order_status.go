package domainerrors

import (
	"fmt"

	defsInternalOrder "github.com/smart-table/src/codegen/intern/order"
)

type InvalidOrderStatus struct {
	OrderStatus defsInternalOrder.OrderStatus
}

func (e InvalidOrderStatus) Error() string {
	return fmt.Sprintf("invalid order_status='%s'", e.OrderStatus)
}
