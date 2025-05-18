package app

import "github.com/google/uuid"

type FinishOrderCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
}
