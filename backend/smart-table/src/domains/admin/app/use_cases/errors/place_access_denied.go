package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type PlaceAccessDenied struct {
	UserUUID  uuid.UUID
	PlaceUUID uuid.UUID
}

func (e PlaceAccessDenied) Error() string {
	return fmt.Sprintf(
		"access to place '%s' denied for user '%s' (not the owner or admin)",
		e.PlaceUUID.String(),
		e.UserUUID.String(),
	)
}
