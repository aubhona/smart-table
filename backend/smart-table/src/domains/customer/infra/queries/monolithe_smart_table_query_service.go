package queries

import (
	defsInternalAdminDTO "github.com/smart-table/src/codegen/intern/admin_dto"
	adminApp "github.com/smart-table/src/domains/admin/app/use_cases"
	adminAppErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	adminDomainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/utils"
)

type SmartTableAdminQueryServiceImpl struct {
	tableIDValidateHandler *adminApp.TableIDValidateCommandHandler
}

func NewSmartTableQueryServiceImpl(
	tableIDValidateHandler *adminApp.TableIDValidateCommandHandler,
) *SmartTableAdminQueryServiceImpl {
	return &SmartTableAdminQueryServiceImpl{
		tableIDValidateHandler: tableIDValidateHandler,
	}
}

func (s *SmartTableAdminQueryServiceImpl) GetMenuDishListByTableID(tableID string) []defsInternalAdminDTO.MenuDishDTO {
	_ = tableID

	return []defsInternalAdminDTO.MenuDishDTO{}
}

func (s *SmartTableAdminQueryServiceImpl) TableIDValidate(tableID string) (bool, error) {
	result, err := s.tableIDValidateHandler.Handle(&adminApp.TableIDValidateCommand{
		TableID: tableID,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[adminDomainErrors.PlaceNotFound](err):
			return false, nil
		case utils.IsTheSameErrorType[adminAppErrors.InvalidTableNumber](err):
			return false, nil
		}
	}

	return result.IsValid, nil
}
