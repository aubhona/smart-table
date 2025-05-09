package app

import defsInternalAdminDTO "github.com/smart-table/src/codegen/intern/admin_dto"

type SmartTableAdminQueryService interface {
	GetMenuDishListByTableID(tableID string) []defsInternalAdminDTO.MenuDishDTO
	IsExistTableID(tableID string) bool
}
