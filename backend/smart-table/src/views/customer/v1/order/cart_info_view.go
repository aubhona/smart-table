package viewscustomerorder

import (
	"context"

	appQueriesErrors "github.com/smart-table/src/domains/customer/app/queries/errors"
	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appUseCasesErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	domainerrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"

	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
)

func (h *CustomerV1OrderHandler) GetCustomerV1OrderCartInfo( //nolint
	ctx context.Context,
	request viewsCustomerOrder.GetCustomerV1OrderCartInfoRequestObject,
) (viewsCustomerOrder.GetCustomerV1OrderCartInfoResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CartCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.CartCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
		NeedPicture:  false,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.UnsuccessMenuDishFetch](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.GetCustomerV1OrderCartInfo403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.GetCustomerV1OrderCartInfo404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	response := viewsCustomerOrder.GetCustomerV1OrderCartInfo200JSONResponse{
		TotalPrice: result.TotalPrice.String(),
		Items:      make([]viewsCustomerOrder.CartItemInfo, 0, len(result.Items)),
	}

	for i := range result.Items {
		menuDish := &result.Items[i]

		response.Items = append(response.Items, viewsCustomerOrder.CartItemInfo{
			ID:          menuDish.ID,
			Name:        menuDish.Name,
			Count:       menuDish.Count,
			Price:       menuDish.Price.String(),
			ResultPrice: menuDish.ResultPrice.String(),
			Comment:     menuDish.Comment.ToPointer(),
		})
	}

	return response, nil
}
