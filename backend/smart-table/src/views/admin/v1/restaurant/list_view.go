package views

import (
	"context"
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	"github.com/smart-table/src/domains/admin/domain"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
)

func convertDomainRestaurantToGenRestaurant(
	domainRestaurant utils.SharedRef[domain.Restaurant],
) viewsAdminRestaurant.RestaurantInfo {
	return viewsAdminRestaurant.RestaurantInfo{
		Name: domainRestaurant.Get().GetName(),
		UUID: domainRestaurant.Get().GetUUID(),
	}
}

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
		UserUUID: request.Params.UserUUID,
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

	restaurantList := make([]viewsAdminRestaurant.RestaurantInfo, 0, len(result.RestaurantList))

	for _, domainRestaurant := range result.RestaurantList {
		restaurantList = append(restaurantList, convertDomainRestaurantToGenRestaurant(domainRestaurant))
	}

	return viewsAdminRestaurant.GetAdminV1RestaurantList200JSONResponse{
		RestaurantList: restaurantList,
	}, nil
}
