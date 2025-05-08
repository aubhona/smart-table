package views

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
	"go.uber.org/zap"
)

func (h *AdminV1RestaurantHandler) PostAdminV1RestaurantDishList( //nolint
	ctx context.Context,
	request viewsAdminRestaurant.PostAdminV1RestaurantDishListRequestObject,
) (viewsAdminRestaurant.PostAdminV1RestaurantDishListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.DishListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.DishListCommand{
		RestaurantUUID: request.Body.RestaurantUUID,
		OwnerUUID:      request.Params.UserUUID,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainErrors.RestaurantNotFound](err):
			return viewsAdminRestaurant.PostAdminV1RestaurantDishList404JSONResponse{
				Code:    viewsAdminRestaurant.RestaurantNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.RestaurantAccessDenied](err):
			return viewsAdminRestaurant.PostAdminV1RestaurantDishList403JSONResponse{
				Code:    viewsAdminRestaurant.AccessDenied,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminRestaurant.PostAdminV1RestaurantDishList200MultipartResponse(func(writer *multipart.Writer) error {
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

		jsonData, err := json.Marshal(dishList)
		if err != nil {
			return err
		}

		partHeaders := textproto.MIMEHeader{}
		partHeaders.Set("Content-Type", "application/json")
		jsonPart, err := writer.CreatePart(partHeaders)

		if err != nil {
			return err
		}

		if _, err := jsonPart.Write(jsonData); err != nil {
			return err
		}

		for _, dish := range result.DishList {
			if dish.Picture == nil {
				continue
			}

			imageHeaders := textproto.MIMEHeader{}
			imageHeaders.Set("Content-Type", "image/png")
			imageHeaders.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.png"`, dish.ID.String()))

			imagePart, err := writer.CreatePart(imageHeaders)
			if err != nil {
				return err
			}

			if _, err := io.Copy(imagePart, dish.Picture); err != nil {
				return err
			}
		}

		return nil
	}), nil
}
