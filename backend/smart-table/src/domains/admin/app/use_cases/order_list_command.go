package app

import (
	"github.com/google/uuid"
)

type OrderListCommand struct {
	UserUUID  uuid.UUID
	PlaceUUID uuid.UUID
	IsActive  bool
}
