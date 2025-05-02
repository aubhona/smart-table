package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type UserNotFoundByUUID struct {
	UUID uuid.UUID
}

func (e UserNotFoundByUUID) Error() string {
	return fmt.Sprintf("User not found by uuid: %s", e.UUID)
}
