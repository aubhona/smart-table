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
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceMenuDishList( //nolint
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceMenuDishListRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceMenuDishListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.MenuDishListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.MenuDishListCommand{
		UserUUID:  request.Params.UserUUID,
		PlaceUUID: request.Body.PlaceUUID,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishList403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceMenuDishList404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminPlace.PostAdminV1PlaceMenuDishList200MultipartResponse(func(writer *multipart.Writer) error {
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

		jsonData, err := json.Marshal(menuDishList)
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

		for i := range result.MenuDishList {
			menuDish := result.MenuDishList[i]
			if menuDish.Picture == nil {
				continue
			}

			imageHeaders := textproto.MIMEHeader{}
			imageHeaders.Set("Content-Type", "image/png")
			imageHeaders.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.png"`, menuDish.ID.String()))

			imagePart, err := writer.CreatePart(imageHeaders)
			if err != nil {
				return err
			}

			if _, err := io.Copy(imagePart, menuDish.Picture); err != nil {
				return err
			}
		}

		return nil
	}), nil
}
