package queries

import (
	defsInternalAdminDTO "github.com/smart-table/src/codegen/intern/admin_dto"
	adminApp "github.com/smart-table/src/domains/admin/app/use_cases"
)

type SmartTableAdminQueryServiceImpl struct {
	adminDishListHandler adminApp.MenuDishListCommandHandler
}

func (s *SmartTableAdminQueryServiceImpl) GetMenuDishListByTableID(tableID string) []defsInternalAdminDTO.MenuDishDTO {
	_ = s.adminDishListHandler
	_ = tableID

	return []defsInternalAdminDTO.MenuDishDTO{}
}

func (s *SmartTableAdminQueryServiceImpl) IsExistTableID(tableID string) bool {
	_ = tableID

	return false
}
