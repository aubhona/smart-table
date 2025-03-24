package domainerrors

import (
	"fmt"

	"github.com/google/uuid"
)

type OrderNotFoundByCustomerUUID struct {
	CustomerUUID uuid.UUID
}

func (o OrderNotFoundByCustomerUUID) Error() string {
	return fmt.Sprintf("Order not found by customer uuid: %s", o.CustomerUUID.String())
}
