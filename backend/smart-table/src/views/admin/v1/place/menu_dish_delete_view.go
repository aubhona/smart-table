package views

import (
	"context"
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceMenuDishDelete(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceMenuDishDeleteRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceMenuDishDeleteResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.MenuDishDeleteCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	err = handler.Handle(&app.MenuDishDeleteCommand{
		UserUUID:     request.Params.UserUUID,
		PlaceUUID:    request.Body.PlaceUUID,
		MenuDishUUID: request.Body.MenuDishUUID,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishDelete404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.MenuDishNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishDelete404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishDelete403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminPlace.PostAdminV1PlaceMenuDishDelete204Response{}, nil
}
