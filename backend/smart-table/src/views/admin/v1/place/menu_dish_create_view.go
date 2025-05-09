package views

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"
	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceMenuDishCreate(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceMenuDishCreateRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceMenuDishCreateResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.MenuDishCreateCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	price, err := decimal.NewFromString(request.Body.Price)
	if err != nil {
		return viewsAdminPlace.PostAdminV1PlaceMenuDishCreate400JSONResponse{
			Code:    viewsAdminPlace.IncorrectPrice,
			Message: err.Error(),
		}, nil
	}

	result, err := handler.Handle(&app.MenuDishCreateCommand{
		UserUUID:  request.Params.UserUUID,
		PlaceUUID: request.Body.PlaceUUID,
		DishUUID:  request.Body.DishUUID,
		Price:     price,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishCreate404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.DishNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishCreate404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishCreate403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminPlace.PostAdminV1PlaceMenuDishCreate200JSONResponse{
		MenuDishUUID: result.MenuDishUUID,
	}, nil
}
