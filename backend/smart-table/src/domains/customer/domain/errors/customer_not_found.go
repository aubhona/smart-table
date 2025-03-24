package domainerrors

import (
	"fmt"

	"github.com/google/uuid"
)

type CustomerNotFound struct {
	UUID uuid.UUID
}

func (c CustomerNotFound) Error() string {
	return fmt.Sprintf("Customer not found by uuid: %s", c.UUID.String())
}
