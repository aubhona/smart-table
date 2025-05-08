package app

import (
	"github.com/google/uuid"
)

type PlaceListCommand struct {
	UserUUID       uuid.UUID
	RestaurantUUID uuid.UUID
}
