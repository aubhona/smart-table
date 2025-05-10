package queries

import (
	"fmt"
	"io"

	"github.com/oapi-codegen/runtime/types"
	defsInternalAdminDTO "github.com/smart-table/src/codegen/intern/admin_dto"
	adminApp "github.com/smart-table/src/domains/admin/app/use_cases"
	adminAppErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	adminDomainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	appQueriesErrors "github.com/smart-table/src/domains/customer/app/queries/errors"
	"github.com/smart-table/src/utils"
)

type SmartTableAdminQueryServiceImpl struct {
	tableIDValidateHandler     *adminApp.TableIDValidateCommandHandler
	menuDishListCommandHandler *adminApp.MenuDishListCommandHandler
}

func NewSmartTableQueryServiceImpl(
	tableIDValidateHandler *adminApp.TableIDValidateCommandHandler,
	menuDishListCommandHandler *adminApp.MenuDishListCommandHandler,
) *SmartTableAdminQueryServiceImpl {
	return &SmartTableAdminQueryServiceImpl{
		tableIDValidateHandler:     tableIDValidateHandler,
		menuDishListCommandHandler: menuDishListCommandHandler,
	}
}

func (s *SmartTableAdminQueryServiceImpl) GetCatalog(
	tableID string,
) ([]defsInternalAdminDTO.MenuDishDTO, error) {
	response, err := s.menuDishListCommandHandler.Handle(&adminApp.MenuDishListCommand{
		InternalCall: utils.NewOptional(adminApp.MenuDishListCommandInternalCall{
			TabledID: tableID,
		}),
	})
	if err != nil {
		return nil, appQueriesErrors.UnsuccessMenuDishFetch{InnerError: err}
	}

	result := make([]defsInternalAdminDTO.MenuDishDTO, 0, len(response.MenuDishList))

	for i := range response.MenuDishList {
		menuDish := response.MenuDishList[i]
		if !menuDish.Exist {
			continue
		}

		pictureBytes, err := io.ReadAll(menuDish.Picture)
		if err != nil {
			return nil, appQueriesErrors.UnsuccessMenuDishFetch{InnerError: err}
		}

		picture := types.File{}
		picture.InitFromBytes(pictureBytes, fmt.Sprintf("%s.png", menuDish.ID))

		result = append(result, defsInternalAdminDTO.MenuDishDTO{
			ID:          menuDish.ID,
			Name:        menuDish.Name,
			Description: menuDish.Description,
			Calories:    menuDish.Calories,
			Weight:      menuDish.Weight,
			Picture:     picture,
			Category:    menuDish.Category,
			Price:       menuDish.Price.String(),
			PictureKey:  menuDish.ID.String(),
		})
	}

	return result, nil
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
