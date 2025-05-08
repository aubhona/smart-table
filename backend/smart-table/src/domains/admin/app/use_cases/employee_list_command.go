package app

import (
	"github.com/google/uuid"
)

type EmployeeListCommand struct {
	UserUUID  uuid.UUID
	PlaceUUID uuid.UUID
}
