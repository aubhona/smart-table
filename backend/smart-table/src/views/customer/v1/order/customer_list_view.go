package viewscustomerorder

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	appQueriesErrors "github.com/smart-table/src/domains/customer/app/queries/errors"
	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appUseCasesErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	domainerrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
	"go.uber.org/zap"
)

type CustomerInfoImpl struct {
	UUID        uuid.UUID
	TgLogin     string
	TgID        string
	IsActive    bool
	TotalPrice  decimal.Decimal
	ItemInfoMap map[string]viewsCustomerOrder.ItemInfo
	IsHost      bool
}

func convertItemToItemInfo(
	item utils.SharedRef[domain.Item],
) viewsCustomerOrder.ItemInfo {
	itemInfo := viewsCustomerOrder.ItemInfo{
		DishUUID:    item.Get().GetDishUUID(),
		Name:        item.Get().GetName(),
		Status:      viewsCustomerOrder.ItemStatus(item.Get().GetStatus()),
		Description: item.Get().GetDescription(),
		Weight:      item.Get().GetWeight(),
		Calories:    item.Get().GetCalories(),
		Category:    item.Get().GetCategory(),
		Price:       item.Get().GetPrice().String(),
		Count:       0,
		ResultPrice: item.Get().GetPrice().String(),
	}

	if item.Get().GetResolution().HasValue() {
		resolution := viewsCustomerOrder.ItemResolution(item.Get().GetResolution().Value())
		itemInfo.Resolution = &resolution
	}

	if item.Get().GetComment().HasValue() {
		comment := item.Get().GetComment().Value()
		itemInfo.Comment = &comment
	}

	return itemInfo
}

func convertCustomerToCustomerInfoImpl(
	customer utils.SharedRef[domain.Customer],
	hostUUID uuid.UUID,
) CustomerInfoImpl {
	return CustomerInfoImpl{
		UUID:        customer.Get().GetUUID(),
		TgLogin:     customer.Get().GetTgLogin(),
		TgID:        customer.Get().GetTgID(),
		IsActive:    false,
		TotalPrice:  decimal.Zero,
		ItemInfoMap: make(map[string]viewsCustomerOrder.ItemInfo),
		IsHost:      customer.Get().GetUUID() == hostUUID,
	}
}

func convertCustomerInfoImplToCustomerInfo(
	customerInfoImpl *CustomerInfoImpl,
) viewsCustomerOrder.CustomerInfo {
	itemList := make([]viewsCustomerOrder.ItemInfo, 0, len(customerInfoImpl.ItemInfoMap))
	for i := range customerInfoImpl.ItemInfoMap {
		itemList = append(itemList, customerInfoImpl.ItemInfoMap[i])
	}

	return viewsCustomerOrder.CustomerInfo{
		UUID:       customerInfoImpl.UUID,
		TgLogin:    customerInfoImpl.TgLogin,
		TgID:       customerInfoImpl.TgID,
		IsActive:   customerInfoImpl.IsActive,
		TotalPrice: customerInfoImpl.TotalPrice.String(),
		ItemList:   itemList,
		IsHost:     customerInfoImpl.IsHost,
	}
}

func getCustomerInfoList(
	order utils.SharedRef[domain.Order],
) []viewsCustomerOrder.CustomerInfo {
	customerInfoImplMap := make(map[uuid.UUID]CustomerInfoImpl)

	for _, customer := range order.Get().GetCustomers() {
		customerInfoImplMap[customer.Get().GetUUID()] = convertCustomerToCustomerInfoImpl(
			customer,
			order.Get().GetHostUserUUID(),
		)
	}

	for _, item := range order.Get().GetItems() {
		if item.Get().GetIsDraft() {
			continue
		}

		key := fmt.Sprintf("{%s}_{%s}_{%s}", item.Get().GetDishUUID(), item.Get().GetStatus(), item.Get().GetComment().ValueOr(""))

		customerUUID := item.Get().GetCustomer().Get().GetUUID()
		customerInfoImpl := customerInfoImplMap[customerUUID]

		itemInfo, isExists := customerInfoImpl.ItemInfoMap[key]
		if !isExists {
			customerInfoImpl.IsActive = true
			customerInfoImpl.ItemInfoMap[key] = convertItemToItemInfo(item)
			customerInfoImpl.TotalPrice = item.Get().GetPrice()
		} else {
			itemInfoPrice, err := decimal.NewFromString(itemInfo.Price)
			if err != nil {
				panic(err)
			}

			itemInfoResultPrice, err := decimal.NewFromString(itemInfo.ResultPrice)
			if err != nil {
				panic(err)
			}

			itemInfoResultPrice = itemInfoResultPrice.Add(itemInfoPrice)
			itemInfo.ResultPrice = itemInfoResultPrice.String()

			customerInfoImpl.TotalPrice = customerInfoImpl.TotalPrice.Add(itemInfoPrice)

			itemInfo.Count++

			customerInfoImpl.ItemInfoMap[key] = itemInfo
		}

		customerInfoImplMap[customerUUID] = customerInfoImpl
	}

	customerInfoList := make([]viewsCustomerOrder.CustomerInfo, 0, len(customerInfoImplMap))

	for _, customerInfoImpl := range customerInfoImplMap {
		customerInfoList = append(customerInfoList, convertCustomerInfoImplToCustomerInfo(&customerInfoImpl))
	}

	return customerInfoList
}

func (h *CustomerV1OrderHandler) GetCustomerV1OrderCustomerList(
	ctx context.Context,
	request viewsCustomerOrder.GetCustomerV1OrderCustomerListRequestObject,
) (viewsCustomerOrder.GetCustomerV1OrderCustomerListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CustomerListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.CustomerListCommand{
		OrderUUID:    request.Params.OrderUUID,
		CustomerUUID: request.Params.CustomerUUID,
	})

	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appQueriesErrors.UnsuccessMenuDishFetch](err):
		case utils.IsTheSameErrorType[appUseCasesErrors.OrderAccessDenied](err):
			return viewsCustomerOrder.GetCustomerV1OrderCustomerList403JSONResponse{
				Code:    viewsCustomerOrder.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainerrors.OrderNotFound](err):
			return viewsCustomerOrder.GetCustomerV1OrderCustomerList404JSONResponse{
				Code:    viewsCustomerOrder.OrderNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error("Get unexpected error", zap.Any("error", err))

		return nil, err
	}

	return viewsCustomerOrder.GetCustomerV1OrderCustomerList200JSONResponse{
		RoomCode:     result.Order.Get().GetRoomCode(),
		CustomerList: getCustomerInfoList(result.Order),
	}, nil
}
