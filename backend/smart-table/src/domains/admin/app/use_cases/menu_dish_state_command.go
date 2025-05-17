package app

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type MenuDishStateCommandAdminCall struct {
	PlaceUUID    uuid.UUID
	UserUUID     uuid.UUID
	MenuDishUUID uuid.UUID
}

type MenuDishStateCommandInternalCall struct {
	MenuDishUUID uuid.UUID
	TabledID     string
	NeedPicture  bool
}

type MenuDishStateCommand struct {
	AdminCall    utils.Optional[MenuDishStateCommandAdminCall]
	InternalCall utils.Optional[MenuDishStateCommandInternalCall]
}
