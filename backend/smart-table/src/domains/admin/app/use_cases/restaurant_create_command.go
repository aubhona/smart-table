package app

import (
	"github.com/google/uuid"
)

type RestaurantCreateCommand struct {
	OwnerUUID uuid.UUID
	Name      string
}
