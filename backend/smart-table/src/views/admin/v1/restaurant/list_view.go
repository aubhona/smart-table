package views

import (
	"context"
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
)

func (h *AdminV1RestaurantHandler) GetAdminV1RestaurantList(
	ctx context.Context,
	request viewsAdminRestaurant.GetAdminV1RestaurantListRequestObject,
) (viewsAdminRestaurant.GetAdminV1RestaurantListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.RestaurantListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.RestaurantListCommand{
		OwnerUUID: request.Params.UserUUID,
	})
	if err != nil {
		if utils.IsTheSameErrorType[domainErrors.UserNotFoundByUUID](err) {
			return viewsAdminRestaurant.GetAdminV1RestaurantList404JSONResponse{
				Code:    viewsAdminRestaurant.UserNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminRestaurant.GetAdminV1RestaurantList200JSONResponse{
		RestaurantList: result.GenRestaurantList,
	}, nil
}
