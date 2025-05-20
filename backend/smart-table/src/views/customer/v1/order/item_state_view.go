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

func (h *CustomerV1OrderHandler) PostCustomerV1OrderItemState(
	ctx context.Context,
	request viewsCustomerOrder.PostCustomerV1OrderItemStateRequestObject,
) (viewsCustomerOrder.PostCustomerV1OrderItemStateResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.ItemStateCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.ItemStateCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
		DishUUD:      request.Body.DishUUID,
		Comment:      utils.NewOptionalFromPointer(request.Body.Comment),
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.UnsuccessMenuDishFetch](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.PostCustomerV1OrderItemState403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.PostCustomerV1OrderItemState404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	return viewsCustomerOrder.PostCustomerV1OrderItemState200MultipartResponse(func(writer *multipart.Writer) error {
		menuDish := viewsCustomerOrder.ItemStateInfo{
			Calories:    result.ItemsState.Calories,
			Category:    result.ItemsState.Category,
			Count:       result.ItemsState.Count,
			Description: result.ItemsState.Description,
			ID:          result.ItemsState.ID,
			Name:        result.ItemsState.Name,
			Price:       result.ItemsState.Price.String(),
			ResultPrice: result.ItemsState.ResultPrice.String(),
			Weight:      result.ItemsState.Weight,
		}

		jsonData, err := json.Marshal(menuDish)
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

		imageHeaders := textproto.MIMEHeader{}
		imageHeaders.Set("Content-Type", "image/png")
		imageHeaders.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.png"`, result.ItemsState.PictureKey))

		imagePart, err := writer.CreatePart(imageHeaders)
		if err != nil {
			return err
		}

		if _, err := io.Copy(imagePart, result.ItemsState.Picture); err != nil {
			return err
		}

		return nil
	}), nil
}
