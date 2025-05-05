package app

import (
	"github.com/google/uuid"
)

type RestaurantListCommand struct {
	OwnerUUID uuid.UUID
}
