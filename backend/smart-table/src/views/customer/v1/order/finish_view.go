package viewscustomerorder //nolint

import (
	"context"

	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appUseCasesErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	domainerrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"

	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
)

func (h *CustomerV1OrderHandler) PostCustomerV1OrderFinish(
	ctx context.Context,
	request viewsCustomerOrder.PostCustomerV1OrderFinishRequestObject,
) (viewsCustomerOrder.PostCustomerV1OrderFinishResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.FinishOrderCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	err = handler.Handle(&app.FinishOrderCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.PostCustomerV1OrderFinish403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.PostCustomerV1OrderFinish404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	return viewsCustomerOrder.PostCustomerV1OrderFinish204Response{}, nil
}
