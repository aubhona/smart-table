package views

import (
	"context"
	"fmt"

	"github.com/smart-table/src/logging"

	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	appQueriesErrors "github.com/smart-table/src/domains/admin/app/queries/errors"
	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceOrderEdit( //nolint
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceOrderEditRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceOrderEditResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.OrderEditCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	OrderStatus := utils.EmptyOptional[string]()
	if request.Body.OrderStatus != nil {
		OrderStatus = utils.NewOptional(string(*request.Body.OrderStatus))
	}

	ItemGroup := utils.EmptyOptional[defsInternalCustomerDTO.ItemEditGroupDTO]()
	if request.Body.ItemGroup != nil {
		ItemGroup = utils.NewOptional(defsInternalCustomerDTO.ItemEditGroupDTO{
			ItemStatus:   string(request.Body.ItemGroup.ItemStatus),
			ItemUUIDList: request.Body.ItemGroup.ItemUUIDList,
		})
	}

	err = handler.Handle(&app.OrderEditCommand{
		UserUUID:      request.Params.UserUUID,
		OrderUUID:     request.Body.OrderUUID,
		PlaceUUID:     request.Body.PlaceUUID,
		TableNumber:   request.Body.TableNumber,
		OrderStatus:   OrderStatus,
		ItemEditGroup: ItemGroup,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.IncorrectEditOrderRequest](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderEdit400JSONResponse{
				Code:    viewsAdminPlace.IncorrectRequestFormat,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appQueriesErrors.OrderNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderEdit404JSONResponse{
				Code:    viewsAdminPlace.OrderNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appQueriesErrors.InvalidOrderStatus](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderEdit403JSONResponse{
				Code:    viewsAdminPlace.InvalidOrderStatus,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appQueriesErrors.ItemStatusChangeRequiresOrderStatusUpdate](err):
		case utils.IsTheSameErrorType[appQueriesErrors.InvalidItemStatus](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderEdit403JSONResponse{
				Code:    viewsAdminPlace.InvalidItemStatus,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appQueriesErrors.DraftItemStatusChangeNotAllowed](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderEdit403JSONResponse{
				Code:    viewsAdminPlace.InvalidItem,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appQueriesErrors.OrderNotBelongToPLace](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderEdit403JSONResponse{
				Code:    viewsAdminPlace.OrderNotBelongToPlace,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderEdit403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderEdit404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminPlace.PostAdminV1PlaceOrderEdit204Response{}, nil
}
