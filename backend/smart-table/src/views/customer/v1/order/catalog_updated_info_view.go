package viewscustomerorder

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

func (h *CustomerV1OrderHandler) GetCustomerV1OrderCatalogUpdatedInfo(
	ctx context.Context,
	request viewsCustomerOrder.GetCustomerV1OrderCatalogUpdatedInfoRequestObject,
) (viewsCustomerOrder.GetCustomerV1OrderCatalogUpdatedInfoResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CatalogUpdateInfoCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.CatalogUpdateInfoCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.GetCustomerV1OrderCatalogUpdatedInfo403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.GetCustomerV1OrderCatalogUpdatedInfo404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	response := viewsCustomerOrder.GetCustomerV1OrderCatalogUpdatedInfo200JSONResponse{
		TotalPrice:      result.TotalPrice.String(),
		MenuUpdatedInfo: make([]viewsCustomerOrder.MenuDishItemUpdatedInfo, 0, len(result.MenuUpdatedInfo)),
	}

	for _, item := range result.MenuUpdatedInfo {
		response.MenuUpdatedInfo = append(response.MenuUpdatedInfo, viewsCustomerOrder.MenuDishItemUpdatedInfo{
			ID:    item.ID,
			Count: item.Count,
		})
	}

	return response, nil
}
