package viewscustomerorder

import (
	"context"

	appQueriesErrors "github.com/smart-table/src/domains/customer/app/queries/errors"
	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appUseCasesErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	domainerrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
	"go.uber.org/zap"
)

func (h *CustomerV1OrderHandler) PostCustomerV1OrderItemInfoState(
	ctx context.Context,
	request viewsCustomerOrder.PostCustomerV1OrderItemInfoStateRequestObject,
) (viewsCustomerOrder.PostCustomerV1OrderItemInfoStateResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.ItemStateCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.ItemStateCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
		DishUUD:      request.Body.DishUUID,
		Comment:      utils.NewOptionalFromPointer(request.Body.Comment),
		NeedPicture:  false,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.UnsuccessMenuDishFetch](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.PostCustomerV1OrderItemInfoState403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.PostCustomerV1OrderItemInfoState404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	return viewsCustomerOrder.PostCustomerV1OrderItemInfoState200JSONResponse{
		Calories:    result.ItemsState.Calories,
		Category:    result.ItemsState.Category,
		Count:       result.ItemsState.Count,
		Description: result.ItemsState.Description,
		ID:          result.ItemsState.ID,
		Name:        result.ItemsState.Name,
		Price:       result.ItemsState.Price.String(),
		ResultPrice: result.ItemsState.ResultPrice.String(),
		Weight:      result.ItemsState.Weight,
	}, nil
}
