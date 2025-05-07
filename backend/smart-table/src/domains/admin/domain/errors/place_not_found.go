package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type PlaceNotFound struct {
	UUID uuid.UUID
}

func (p PlaceNotFound) Error() string {
	return fmt.Sprintf("Place not found by uuid: %s", p.UUID)
}
