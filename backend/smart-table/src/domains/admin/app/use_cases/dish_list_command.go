package app

import (
	"github.com/google/uuid"
)

type DishListCommand struct {
	RestaurantUUID uuid.UUID
	OwnerUUID      uuid.UUID
	NeedPicture    bool
}
