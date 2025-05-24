package views

import (
	"context"
	"fmt"
	"slices"

	"github.com/smart-table/src/logging"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceOrderList(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceOrderListRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceOrderListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.OrderListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.OrderListCommand{
		UserUUID:  request.Params.UserUUID,
		PlaceUUID: request.Body.PlaceUUID,
		IsActive:  request.Body.IsActive,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderList403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderList404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	orderMainInfoList := make([]viewsAdminPlace.OrderMainInfo, 0, len(result.OrderList))

	for _, order := range result.OrderList {
		orderMainInfoList = append(orderMainInfoList, convertMainOrderInfoDTOToOrderMainInfo(&order))
	}

	slices.SortFunc(orderMainInfoList, func(a, b viewsAdminPlace.OrderMainInfo) int {
		return a.CreatedAt.Compare(b.CreatedAt)
	})

	return viewsAdminPlace.PostAdminV1PlaceOrderList200JSONResponse{
		OrderList: orderMainInfoList,
	}, nil
}
