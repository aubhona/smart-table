package app

import (
	"github.com/google/uuid"
)

type OrderInfoCommand struct {
	UserUUID  uuid.UUID
	OrderUUID uuid.UUID
	PlaceUUID uuid.UUID
}
