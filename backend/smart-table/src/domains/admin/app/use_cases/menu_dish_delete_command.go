package app

import (
	"github.com/google/uuid"
)

type MenuDishDeleteCommand struct {
	UserUUID     uuid.UUID
	PlaceUUID    uuid.UUID
	MenuDishUUID uuid.UUID
}
