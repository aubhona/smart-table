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
	"go.uber.org/zap"
)

func (h *AdminV1RestaurantHandler) PostAdminV1RestaurantDishInfoList( //nolint
	ctx context.Context,
	request viewsAdminRestaurant.PostAdminV1RestaurantDishInfoListRequestObject,
) (viewsAdminRestaurant.PostAdminV1RestaurantDishInfoListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.DishListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.DishListCommand{
		RestaurantUUID: request.Body.RestaurantUUID,
		OwnerUUID:      request.Params.UserUUID,
		NeedPicture:    false,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainErrors.RestaurantNotFound](err):
			return viewsAdminRestaurant.PostAdminV1RestaurantDishInfoList404JSONResponse{
				Code:    viewsAdminRestaurant.RestaurantNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.RestaurantAccessDenied](err):
			return viewsAdminRestaurant.PostAdminV1RestaurantDishInfoList403JSONResponse{
				Code:    viewsAdminRestaurant.AccessDenied,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	dishList := make([]viewsAdminRestaurant.DishInfo, 0, len(result.DishList))

	for _, dish := range result.DishList {
		logging.GetLogger().Debug("Adding dish info", zap.String("dish_uuid", dish.ID.String()))

		dishList = append(dishList, viewsAdminRestaurant.DishInfo{
			ID:          dish.ID,
			Name:        dish.Name,
			Description: dish.Description,
			Calories:    dish.Calories,
			Weight:      dish.Weight,
			Category:    dish.Category,
			PictureKey:  dish.ID.String(),
		})
	}

	return viewsAdminRestaurant.PostAdminV1RestaurantDishInfoList200JSONResponse{
		DishList: dishList,
	}, nil
}
