package app

import (
	"github.com/google/uuid"
)

type RestaurantListCommand struct {
	UserUUID uuid.UUID
}
