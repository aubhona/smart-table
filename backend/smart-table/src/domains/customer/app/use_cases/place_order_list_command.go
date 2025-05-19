package app

import (
	"github.com/google/uuid"
)

type PlaceOrderListCommand struct {
	PlaceUUID uuid.UUID
	IsActive  bool
}
