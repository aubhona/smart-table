package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type PlaceNotFoundByAddress struct {
	Address        string
	RestaurantUUID uuid.UUID
}

func (p PlaceNotFoundByAddress) Error() string {
	return fmt.Sprintf("Place not found by address: %s, restaurant_uuid: %s", p.Address, p.RestaurantUUID)
}
