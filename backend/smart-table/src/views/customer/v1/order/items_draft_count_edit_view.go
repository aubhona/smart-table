package viewscustomerorder

import (
	"context"
	"fmt"

	appQueriesErrors "github.com/smart-table/src/domains/customer/app/queries/errors"
	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appUseCasesErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	domainerrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"

	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
)

func (h *CustomerV1OrderHandler) PostCustomerV1OrderItemsDraftCountEdit(
	ctx context.Context,
	request viewsCustomerOrder.PostCustomerV1OrderItemsDraftCountEditRequestObject,
) (viewsCustomerOrder.PostCustomerV1OrderItemsDraftCountEditResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CartItemsCountEditCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	err = handler.Handle(&app.CartItemsCountEditCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
		Comment:      utils.NewOptionalFromPointer(request.Body.Comment),
		DishUUID:     request.Body.MenuDishUUID,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainerrors.IncorrectDeleteItemsCount](err):
			return viewsCustomerOrder.PostCustomerV1OrderItemsDraftCountEdit400JSONResponse{
				Code:    viewsCustomerOrder.InvalidItemCount,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appQueriesErrors.UnsuccessMenuDishFetch](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.PostCustomerV1OrderItemsDraftCountEdit403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.PostCustomerV1OrderItemsDraftCountEdit404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	return viewsCustomerOrder.PostCustomerV1OrderItemsDraftCountEdit204Response{}, nil
}
