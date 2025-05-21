package errors

import (
	"fmt"

	"github.com/smart-table/src/utils"
)

type InvalidOrderStatus struct {
	OrderStatus utils.Optional[string]
}

func (e InvalidOrderStatus) Error() string {
	return fmt.Sprintf("invalid order_status='%s'", e.OrderStatus.ValueOr("nil"))
}
