package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type MenuDishNotFound struct {
	UUID uuid.UUID
}

func (p MenuDishNotFound) Error() string {
	return fmt.Sprintf("Menu dish not found by uuid: %s", p.UUID)
}
