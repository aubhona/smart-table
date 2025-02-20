package domain_errors

import (
	"fmt"
	"github.com/google/uuid"
)

type OrderNotFound struct {
	Uuid uuid.UUID
}

func (o OrderNotFound) Error() string {
	return fmt.Sprintf("Order not found uuid: %s", o.Uuid.String())
}
