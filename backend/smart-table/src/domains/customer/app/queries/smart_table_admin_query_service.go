package app

import (
	"github.com/google/uuid"
	defsInternalAdminDTO "github.com/smart-table/src/codegen/intern/admin_dto"
)

type SmartTableAdminQueryService interface {
	GetMenuDish(tableID string, dishUUID uuid.UUID, withPicture bool) (defsInternalAdminDTO.MenuDishDTO, error)
	GetCatalog(tableID string) ([]defsInternalAdminDTO.MenuDishDTO, error)
	TableIDValidate(tableID string) (bool, error)
}
