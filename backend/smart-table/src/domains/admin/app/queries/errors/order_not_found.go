package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type OrderNotFound struct {
	OrderUUID uuid.UUID
}

func (o OrderNotFound) Error() string {
	return fmt.Sprintf("Order not found uuid: %s", o.OrderUUID.String())
}
