package viewscustomerorder

import (
	"context"
	"errors"

	"github.com/smart-table/src/logging"
	"go.uber.org/zap"

	appQueriesErrors "github.com/smart-table/src/domains/customer/app/queries/errors"
	appUseCasesErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"

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
		CustomerUUID: request.Params.CustomerUUID,
		RoomCode:     utils.OptionalFromPointer(request.Body.RoomCode),
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appUseCasesErrors.IncorrectRoomCodeError](err):
			return viewsCustomerOrder.PostCustomerV1OrderCreate403JSONResponse{
				Code:    viewsCustomerOrder.InvalidRoomCode,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appUseCasesErrors.CustomerAlreadyHasActiveOrder](err):
			var orderErr appUseCasesErrors.CustomerAlreadyHasActiveOrder

			errors.As(err, &orderErr)

			return viewsCustomerOrder.PostCustomerV1OrderCreate200JSONResponse{
				OrderUUID: orderErr.OrderUUID,
			}, nil
		case utils.IsTheSameErrorType[appQueriesErrors.PlaceNotFound](err):
		case utils.IsTheSameErrorType[appQueriesErrors.InvalidTableNumber](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.InvalidTableID](err):
			return viewsCustomerOrder.PostCustomerV1OrderCreate403JSONResponse{
				Code:    viewsCustomerOrder.InvalidTableID,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	return viewsCustomerOrder.PostCustomerV1OrderCreate200JSONResponse{
		OrderUUID: result.OrderUUID,
	}, nil
}
