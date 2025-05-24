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

func (h *CustomerV1OrderHandler) GetCustomerV1OrderCatalogInfo( //nolint
	ctx context.Context,
	request viewsCustomerOrder.GetCustomerV1OrderCatalogInfoRequestObject,
) (viewsCustomerOrder.GetCustomerV1OrderCatalogInfoResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CatalogCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.CatalogCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
		NeedPicture:  false,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.UnsuccessMenuDishFetch](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.GetCustomerV1OrderCatalogInfo403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.GetCustomerV1OrderCatalogInfo404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	response := viewsCustomerOrder.GetCustomerV1OrderCatalogInfo200JSONResponse{
		TotalPrice:  result.TotalPrice.String(),
		RoomCode:    result.RoomCode,
		Categories:  result.Categories,
		Menu:        make([]viewsCustomerOrder.MenuDishItem, 0, len(result.Items)),
		GoTipScreen: result.GoTipScreen,
	}

	for i := range result.Items {
		menuDish := &result.Items[i]

		response.Menu = append(response.Menu, viewsCustomerOrder.MenuDishItem{
			ID:       menuDish.ID,
			Name:     menuDish.Name,
			Count:    menuDish.Count,
			Calories: menuDish.Calories,
			Price:    menuDish.Price,
			Category: menuDish.Category,
			Weight:   menuDish.Weight,
		})
	}

	return response, nil
}
