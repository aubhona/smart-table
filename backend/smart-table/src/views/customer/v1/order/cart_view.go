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
	"go.uber.org/zap"

	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
)

func (h *CustomerV1OrderHandler) GetCustomerV1OrderCart( //nolint
	ctx context.Context,
	request viewsCustomerOrder.GetCustomerV1OrderCartRequestObject,
) (viewsCustomerOrder.GetCustomerV1OrderCartResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CartCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.CartCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.UnsuccessMenuDishFetch](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.GetCustomerV1OrderCart403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.GetCustomerV1OrderCart404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	return viewsCustomerOrder.GetCustomerV1OrderCart200MultipartResponse(
		func(writer *multipart.Writer) error {
			cart := viewsCustomerOrder.CartInfo{
				TotalPrice: result.TotalPrice.String(),
				Items:      make([]viewsCustomerOrder.CartItemInfo, 0, len(result.Items)),
			}

			for i := range result.Items {
				menuDish := &result.Items[i]

				cart.Items = append(cart.Items, viewsCustomerOrder.CartItemInfo{
					DishUUID:    menuDish.ID,
					Name:        menuDish.Name,
					Count:       menuDish.Count,
					Price:       menuDish.Price.String(),
					ResultPrice: menuDish.ResultPrice.String(),
				})
			}

			jsonData, err := json.Marshal(cart)
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
				imageHeaders.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.png"`, menuDish.ID))

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
