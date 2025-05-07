package views

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
	"go.uber.org/zap"
)

func (h *AdminV1RestaurantHandler) parseAdminV1RestaurantDishCreateRequest(
	request *viewsAdminRestaurant.PostAdminV1RestaurantDishCreateRequestObject,
) (viewsAdminRestaurant.AdminV1RestaurantDishCreateRequest, error) {
	form, err := request.Body.ReadForm(h.MaxInputFileSizeMB * 1024 * 1024)
	if err != nil {
		return viewsAdminRestaurant.AdminV1RestaurantDishCreateRequest{}, err
	}

	dishName := form.Value["dish_name"][0]
	description := form.Value["description"][0]
	category := form.Value["category"][0]

	calories, err := strconv.Atoi(form.Value["calories"][0])
	if err != nil {
		logging.GetLogger().Error("error occurred while parsing calories", zap.Error(err))
		return viewsAdminRestaurant.AdminV1RestaurantDishCreateRequest{}, err
	}

	weight, err := strconv.Atoi(form.Value["weight"][0])
	if err != nil {
		logging.GetLogger().Error("error occurred while parsing weight", zap.Error(err))
		return viewsAdminRestaurant.AdminV1RestaurantDishCreateRequest{}, err
	}

	restaurantUUID, err := uuid.Parse(form.Value["restaurant_uuid"][0])
	if err != nil {
		logging.GetLogger().Error("error occurred while parsing restaurant uuid", zap.Error(err))
		return viewsAdminRestaurant.AdminV1RestaurantDishCreateRequest{}, err
	}

	fileHeader := form.File["dish_picture_file"][0]

	parsedRequest := viewsAdminRestaurant.AdminV1RestaurantDishCreateRequest{
		DishName:       dishName,
		Description:    description,
		Category:       category,
		Calories:       calories,
		Weight:         weight,
		RestaurantUUID: restaurantUUID,
	}

	parsedRequest.DishPictureFile.InitFromMultipart(fileHeader)

	return parsedRequest, nil
}

func (h *AdminV1RestaurantHandler) PostAdminV1RestaurantDishCreate(
	ctx context.Context,
	request viewsAdminRestaurant.PostAdminV1RestaurantDishCreateRequestObject,
) (viewsAdminRestaurant.PostAdminV1RestaurantDishCreateResponseObject, error) {
	parsedRequest, err := h.parseAdminV1RestaurantDishCreateRequest(&request)
	if err != nil {
		return &viewsAdminRestaurant.PostAdminV1RestaurantDishCreate400JSONResponse{
			Code:    viewsAdminRestaurant.InvalidRequest,
			Message: err.Error(),
		}, nil
	}

	logging.GetLogger().Debug("successfully parsed request", zap.Any("request", parsedRequest))

	handler, err := utils.GetFromContainer[*app.DishCreateCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.DishCreateCommand{
		RestaurantUUID: parsedRequest.RestaurantUUID,
		OwnerUUID:      request.Params.UserUUID,
		DishName:       parsedRequest.DishName,
		Description:    parsedRequest.Description,
		Calories:       parsedRequest.Calories,
		Weight:         parsedRequest.Weight,
		Category:       parsedRequest.Category,
		Image:          parsedRequest.DishPictureFile,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainErrors.RestaurantNotFound](err):
			return viewsAdminRestaurant.PostAdminV1RestaurantDishCreate404JSONResponse{
				Code:    viewsAdminRestaurant.RestaurantNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.RestaurantAccessDenied](err):
			return viewsAdminRestaurant.PostAdminV1RestaurantDishCreate403JSONResponse{
				Code:    viewsAdminRestaurant.AccessDenied,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminRestaurant.PostAdminV1RestaurantDishCreate200JSONResponse{
		DishUUID: result.DishUUID,
	}, nil
}
