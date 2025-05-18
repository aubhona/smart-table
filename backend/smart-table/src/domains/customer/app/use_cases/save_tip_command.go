package app

import "github.com/google/uuid"

type SaveTipCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
}
