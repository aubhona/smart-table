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

func (h *AdminV1PlaceHandler) PostAdminV1PlaceMenuDishInfoList( //nolint
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceMenuDishInfoListRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceMenuDishInfoListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.MenuDishListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.MenuDishListCommand{
		AdminCall: utils.NewOptional(app.MenuDishListCommandAdminCall{
			UserUUID:    request.Params.UserUUID,
			PlaceUUID:   request.Body.PlaceUUID,
			NeedPicture: false,
		}),
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishInfoList403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishInfoList404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	menuDishList := make([]viewsAdminPlace.MenuDishInfo, 0, len(result.MenuDishList))

	for i := range result.MenuDishList {
		menuDish := result.MenuDishList[i]

		menuDishList = append(menuDishList, viewsAdminPlace.MenuDishInfo{
			ID:          menuDish.ID,
			Name:        menuDish.Name,
			Description: menuDish.Description,
			Calories:    menuDish.Calories,
			Weight:      menuDish.Weight,
			Category:    menuDish.Category,
			PictureKey:  menuDish.ID.String(),
			Price:       menuDish.Price.String(),
			Exist:       menuDish.Exist,
		})
	}

	return viewsAdminPlace.PostAdminV1PlaceMenuDishInfoList200JSONResponse{
		MenuDishList: menuDishList,
	}, nil
}
