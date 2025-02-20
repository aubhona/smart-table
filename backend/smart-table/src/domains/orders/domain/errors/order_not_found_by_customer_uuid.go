package domain_errors

import (
	"fmt"
	"github.com/google/uuid"
)

type OrderNotFoundByCustomerUuid struct {
	CustomerUuid uuid.UUID
}

func (o OrderNotFoundByCustomerUuid) Error() string {
	return fmt.Sprintf("Order not found by customer uuid: %s", o.CustomerUuid.String())
}
