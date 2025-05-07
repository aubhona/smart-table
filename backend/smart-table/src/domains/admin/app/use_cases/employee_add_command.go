package app

import (
	"github.com/google/uuid"
)

type EmployeeAddCommand struct {
	UserUUID      uuid.UUID
	PlaceUUID     uuid.UUID
	EmployeeLogin string
	EmployeeRole  string
}
