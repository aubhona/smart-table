package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type EmployeeAlreadyExists struct {
	EmployeeLogin string
	PlaceUUID     uuid.UUID
}

func (e EmployeeAlreadyExists) Error() string {
	return fmt.Sprintf("employee with login '%s' already exists in place with uuid '%s'", e.EmployeeLogin, e.PlaceUUID)
}
