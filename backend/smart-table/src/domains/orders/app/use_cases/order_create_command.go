package app

import (
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
)

type OrderCreateCommand struct {
	TableId      string
	CustomerUuid uuid.UUID
	RoomCode     utils.Optional[string]
}
