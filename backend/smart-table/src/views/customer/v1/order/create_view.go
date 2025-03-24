package views

import (
	"context"

	app "github.com/smart-table/src/domains/customer/app/use_cases"
	"github.com/smart-table/src/utils"
	viewsCustomer "github.com/smart-table/src/views/codegen/customer"
)

func (h *CustomerV1OrderHandler) PostCustomerV1OrderCreate(
	ctx context.Context,
	request viewsCustomer.PostCustomerV1OrderCreateRequestObject,
) (viewsCustomer.PostCustomerV1OrderCreateResponseObject, error) {
	handler, err := utils.GetFromContainer[app.OrderCreateCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.OrderCreateCommand{
		TableID:      request.Body.TableID,
		CustomerUUID: request.Body.CustomerUUID,
		RoomCode:     utils.OptionalFromPointer(request.Body.RoomCode),
	})

	if err != nil {
		return nil, err
	}

	return viewsCustomer.PostCustomerV1OrderCreate200JSONResponse{
		OrderUUID: result.OrderUUID,
	}, nil
}
