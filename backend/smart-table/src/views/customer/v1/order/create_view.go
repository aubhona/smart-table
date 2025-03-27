package viewscustomerorder

import (
	"context"

	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"

	app "github.com/smart-table/src/domains/customer/app/use_cases"
	"github.com/smart-table/src/utils"
	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
)

func (h *CustomerV1OrderHandler) PostCustomerV1OrderCreate(
	ctx context.Context,
	request viewsCustomerOrder.PostCustomerV1OrderCreateRequestObject,
) (viewsCustomerOrder.PostCustomerV1OrderCreateResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.OrderCreateCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.OrderCreateCommand{
		TableID:      request.Body.TableID,
		CustomerUUID: request.Body.CustomerUUID,
		RoomCode:     utils.OptionalFromPointer(request.Body.RoomCode),
	})

	if err != nil {
		if utils.IsTheSameErrorType[appErrors.IncorrectRoomCodeError](err) {
			return viewsCustomerOrder.PostCustomerV1OrderCreate403JSONResponse{
				Code:    viewsCustomerOrder.InvalidRoomCode,
				Message: err.Error(),
			}, nil
		}

		return nil, err
	}

	return viewsCustomerOrder.PostCustomerV1OrderCreate200JSONResponse{
		OrderUUID: result.OrderUUID,
	}, nil
}
