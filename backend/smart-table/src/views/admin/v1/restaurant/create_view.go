package views

import (
	"context"
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
)

func (h *AdminV1RestaurantHandler) PostAdminV1RestaurantCreate(
	ctx context.Context,
	request viewsAdminRestaurant.PostAdminV1RestaurantCreateRequestObject,
) (viewsAdminRestaurant.PostAdminV1RestaurantCreateResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.RestaurantCreateCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.RestaurantCreateCommand{
		OwnerUUID: request.Params.UserUUID,
		Name:      request.Body.RestaurantName,
	})
	if err != nil {
		if utils.IsTheSameErrorType[domainErrors.UserNotFoundByUUID](err) {
			return viewsAdminRestaurant.PostAdminV1RestaurantCreate404JSONResponse{
				Code:    viewsAdminRestaurant.UserNotFound,
				Message: err.Error(),
			}, nil
		} else if utils.IsTheSameErrorType[appErrors.RestaurantNameExists](err) {
			return viewsAdminRestaurant.PostAdminV1RestaurantCreate403JSONResponse{
				Code:    viewsAdminRestaurant.AlreadyExist,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminRestaurant.PostAdminV1RestaurantCreate200JSONResponse{
		RestaurantUUID: result.RestaurantUUID,
	}, nil
}
