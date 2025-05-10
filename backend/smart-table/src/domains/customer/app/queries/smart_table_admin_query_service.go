package app

import (
	defsInternalAdminDTO "github.com/smart-table/src/codegen/intern/admin_dto"
)

type SmartTableAdminQueryService interface {
	GetCatalog(tableID string) ([]defsInternalAdminDTO.MenuDishDTO, error)
	TableIDValidate(tableID string) (bool, error)
}
