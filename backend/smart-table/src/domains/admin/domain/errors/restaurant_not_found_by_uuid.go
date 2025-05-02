package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type RestaurantNotFoundByUUID struct {
	UUID uuid.UUID
}

func (e RestaurantNotFoundByUUID) Error() string {
	return fmt.Sprintf("Restaurant not found by uuid: %s", e.UUID)
}
