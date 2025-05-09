package app

import (
	"github.com/google/uuid"
)

type MenuDishListCommand struct {
	PlaceUUID uuid.UUID
	UserUUID  uuid.UUID
}
