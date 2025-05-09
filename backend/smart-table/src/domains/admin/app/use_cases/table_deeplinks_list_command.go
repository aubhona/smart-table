package app

import (
	"github.com/google/uuid"
)

type TableDeepLinksListCommand struct {
	UserUUID  uuid.UUID
	PlaceUUID uuid.UUID
}
