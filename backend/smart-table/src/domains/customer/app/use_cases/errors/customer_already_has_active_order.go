package apperrors

import (
	"fmt"

	"github.com/google/uuid"
)

type CustomerAlreadyHasActiveOrder struct {
	UserUUID  uuid.UUID
	OrderUUID uuid.UUID
}

func (e CustomerAlreadyHasActiveOrder) Error() string {
	return fmt.Sprintf("Customer with '%s' already has active order '%s'", e.UserUUID, e.OrderUUID)
}
