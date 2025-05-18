package viewscustomerorder

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"

	appQueriesErrors "github.com/smart-table/src/domains/customer/app/queries/errors"
	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appUseCasesErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	domainerrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
	"go.uber.org/zap"
)

func (h *CustomerV1OrderHandler) GetCustomerV1OrderCatalog( //nolint
	ctx context.Context,
	request viewsCustomerOrder.GetCustomerV1OrderCatalogRequestObject,
) (viewsCustomerOrder.GetCustomerV1OrderCatalogResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CatalogCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.CatalogCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.UnsuccessMenuDishFetch](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.GetCustomerV1OrderCatalog403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.GetCustomerV1OrderCatalog404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	return viewsCustomerOrder.GetCustomerV1OrderCatalog200MultipartResponse(
		func(writer *multipart.Writer) error {
			catalog := viewsCustomerOrder.Catalog{
				TotalPrice:  result.TotalPrice.String(),
				RoomCode:    result.RoomCode,
				Categories:  result.Categories,
				Menu:        make([]viewsCustomerOrder.MenuDishItem, 0, len(result.Items)),
				GoTipScreen: result.GoTipScreen,
			}

			for i := range result.Items {
				menuDish := &result.Items[i]

				catalog.Menu = append(catalog.Menu, viewsCustomerOrder.MenuDishItem{
					ID:       menuDish.ID,
					Name:     menuDish.Name,
					Count:    menuDish.Count,
					Calories: menuDish.Calories,
					Price:    menuDish.Price,
					Category: menuDish.Category,
				})
			}

			jsonData, err := json.Marshal(catalog)
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

			for i := range result.Items {
				menuDish := &result.Items[i]
				if menuDish.Picture == nil {
					continue
				}

				imageHeaders := textproto.MIMEHeader{}
				imageHeaders.Set("Content-Type", "image/png")
				imageHeaders.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.png"`, menuDish.PictureKey))

				imagePart, err := writer.CreatePart(imageHeaders)
				if err != nil {
					return err
				}

				if _, err := io.Copy(imagePart, menuDish.Picture); err != nil {
					return err
				}
			}

			return nil
		},
	), nil
}
