package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type DishNotFound struct {
	UUID uuid.UUID
}

func (p DishNotFound) Error() string {
	return fmt.Sprintf("Dish not found by uuid: %s", p.UUID)
}
