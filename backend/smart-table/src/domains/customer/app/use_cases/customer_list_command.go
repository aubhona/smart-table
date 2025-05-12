package app

import (
	"github.com/google/uuid"
)

type CustomerListCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
}
