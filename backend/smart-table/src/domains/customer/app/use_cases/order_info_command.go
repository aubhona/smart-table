package app

import (
	"github.com/google/uuid"
)

type OrderInfoCommand struct {
	OrderUUID uuid.UUID
}
