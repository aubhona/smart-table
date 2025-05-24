package views

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/smart-table/src/logging"

	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	appQueriesErrors "github.com/smart-table/src/domains/admin/app/queries/errors"
	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceOrderInfo(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceOrderInfoRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceOrderInfoResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.OrderInfoCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.OrderInfoCommand{
		UserUUID:  request.Params.UserUUID,
		PlaceUUID: request.Body.PlaceUUID,
		OrderUUID: request.Body.OrderUUID,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.OrderNotBelongToPLace](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderInfo403JSONResponse{
				Code:    viewsAdminPlace.OrderNotBelongToPlace,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderInfo403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceOrderInfo404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminPlace.PostAdminV1PlaceOrderInfo200JSONResponse{
		OrderInfo: convertOrderInfoDTOToOrderInfo(&result.OrderInfo),
	}, nil
}

// Converters.
func convertItemGroupInfoDTOToItemGroupInfo(
	itemGroupInfoDTO *defsInternalCustomerDTO.ItemGroupInfoDTO,
) viewsAdminPlace.ItemGroupInfo {
	itemGroupInfo := viewsAdminPlace.ItemGroupInfo{
		MenuDishUUID: itemGroupInfoDTO.MenuDishUUID,
		ItemUUIDList: itemGroupInfoDTO.ItemUUIDList,
		Status:       viewsAdminPlace.ItemStatus(itemGroupInfoDTO.Status),
		Name:         itemGroupInfoDTO.Name,
		ItemPrice:    itemGroupInfoDTO.ItemPrice,
		ResultPrice:  itemGroupInfoDTO.ResultPrice,
		Count:        itemGroupInfoDTO.Count,
		Comment:      itemGroupInfoDTO.Comment,
	}

	if itemGroupInfoDTO.Resolution != nil {
		resolution := viewsAdminPlace.ItemResolution(*itemGroupInfoDTO.Resolution)
		itemGroupInfo.Resolution = &resolution
	}

	return itemGroupInfo
}

func convertCustomerInfoDTOToCustomerInfo(
	customerInfoDTO *defsInternalCustomerDTO.CustomerInfoDTO,
) viewsAdminPlace.CustomerInfo {
	itemGroupList := make([]viewsAdminPlace.ItemGroupInfo, 0, len(customerInfoDTO.ItemGroupList))
	for i := range customerInfoDTO.ItemGroupList {
		itemGroupList = append(itemGroupList, convertItemGroupInfoDTOToItemGroupInfo(&customerInfoDTO.ItemGroupList[i]))
	}

	slices.SortFunc(itemGroupList, func(a, b viewsAdminPlace.ItemGroupInfo) int {
		return strings.Compare(a.Name, b.Name)
	})

	return viewsAdminPlace.CustomerInfo{
		UUID:          customerInfoDTO.UUID,
		TgLogin:       customerInfoDTO.TgLogin,
		TgID:          customerInfoDTO.TgID,
		ItemGroupList: itemGroupList,
		TotalPrice:    customerInfoDTO.TotalPrice,
	}
}

func convertMainOrderInfoDTOToOrderMainInfo(
	orderMainInfoDTO *defsInternalCustomerDTO.OrderMainInfoDTO,
) viewsAdminPlace.OrderMainInfo {
	orderMainInfo := viewsAdminPlace.OrderMainInfo{
		UUID:        orderMainInfoDTO.UUID,
		Status:      viewsAdminPlace.OrderStatus(orderMainInfoDTO.Status),
		GuestsCount: orderMainInfoDTO.GuestsCount,
		CreatedAt:   orderMainInfoDTO.CreatedAt,
		TotalPrice:  orderMainInfoDTO.TotalPrice,
		TableNumber: orderMainInfoDTO.TableNumber,
	}

	if orderMainInfoDTO.Resolution != nil {
		resolution := viewsAdminPlace.OrderResolution(*orderMainInfoDTO.Resolution)
		orderMainInfo.Resolution = &resolution
	}

	return orderMainInfo
}

func convertOrderInfoDTOToOrderInfo(
	orderInfoDTO *defsInternalCustomerDTO.OrderInfoDTO,
) viewsAdminPlace.OrderInfo {
	customerList := make([]viewsAdminPlace.CustomerInfo, 0, len(orderInfoDTO.CustomerList))
	for _, customerInfoDTO := range orderInfoDTO.CustomerList {
		customerList = append(customerList, convertCustomerInfoDTOToCustomerInfo(&customerInfoDTO))
	}

	slices.SortFunc(customerList, func(a, b viewsAdminPlace.CustomerInfo) int {
		return strings.Compare(a.TgLogin, b.TgLogin)
	})

	orderMainInfo := convertMainOrderInfoDTOToOrderMainInfo(&orderInfoDTO.OrderMainInfo)

	return viewsAdminPlace.OrderInfo{
		OrderMainInfo: orderMainInfo,
		CustomerList:  customerList,
	}
}
