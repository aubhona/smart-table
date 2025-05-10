package apperrors

import (
	"fmt"

	"github.com/google/uuid"
)

type OrderAccessDenied struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
}

func (e OrderAccessDenied) Error() string {
	return fmt.Sprintf(
		"access to order '%s' denied for customer '%s'",
		e.OrderUUID.String(),
		e.CustomerUUID.String(),
	)
}
