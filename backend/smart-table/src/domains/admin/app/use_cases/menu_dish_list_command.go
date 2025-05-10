package app

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type MenuDishListCommandAdminCall struct {
	PlaceUUID uuid.UUID
	UserUUID  uuid.UUID
}

type MenuDishListCommandInternalCall struct {
	TabledID string
}

type MenuDishListCommand struct {
	AdminCall    utils.Optional[MenuDishListCommandAdminCall]
	InternalCall utils.Optional[MenuDishListCommandInternalCall]
}
