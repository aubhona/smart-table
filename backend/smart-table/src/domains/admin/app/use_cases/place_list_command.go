package app

import (
	"github.com/google/uuid"
)

type PlaceListCommand struct {
	OwnerUUID      uuid.UUID
	RestaurantUUID uuid.UUID
}
