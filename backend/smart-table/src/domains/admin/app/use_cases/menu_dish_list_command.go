package app

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type MenuDishListCommandAdminCall struct {
	PlaceUUID   uuid.UUID
	UserUUID    uuid.UUID
	NeedPicture bool
}

type MenuDishListCommandInternalCall struct {
	TabledID    string
	NeedPicture bool
}

type MenuDishListCommand struct {
	AdminCall    utils.Optional[MenuDishListCommandAdminCall]
	InternalCall utils.Optional[MenuDishListCommandInternalCall]
}
