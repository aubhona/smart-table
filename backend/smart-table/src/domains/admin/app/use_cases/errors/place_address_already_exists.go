package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type PlaceAddressAlreadyExists struct {
	Address        string
	RestaurantUUID uuid.UUID
}

func (e PlaceAddressAlreadyExists) Error() string {
	return fmt.Sprintf("Place with address '%s' already exists, restaurant_uuid='%s'", e.Address, e.RestaurantUUID)
}
