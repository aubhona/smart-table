package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type RestaurantAccessDenied struct {
	UserUUID       uuid.UUID
	RestaurantUUID uuid.UUID
}

func (e RestaurantAccessDenied) Error() string {
	return fmt.Sprintf(
		"access to restaurant '%s' denied for user '%s' (not the owner)",
		e.RestaurantUUID.String(),
		e.UserUUID.String(),
	)
}
