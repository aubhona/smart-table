package domain_errors

import (
	"fmt"
	"github.com/google/uuid"
)

type CustomerNotFound struct {
	Uuid uuid.UUID
}

func (c CustomerNotFound) Error() string {
	return fmt.Sprintf("Customer not found by uuid: %s", c.Uuid.String())
}
