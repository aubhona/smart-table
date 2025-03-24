package app

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type OrderCreateCommand struct {
	TableID      string
	CustomerUUID uuid.UUID
	RoomCode     utils.Optional[string]
}
