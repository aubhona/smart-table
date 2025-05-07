package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type RestaurantNotFound struct {
	UUID uuid.UUID
}

func (e RestaurantNotFound) Error() string {
	return fmt.Sprintf("Restaurant not found by uuid: %s", e.UUID)
}
