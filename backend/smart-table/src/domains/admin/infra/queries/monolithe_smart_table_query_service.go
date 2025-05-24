package queries

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	appQueriesErrors "github.com/smart-table/src/domains/admin/app/queries/errors"
	appServices "github.com/smart-table/src/domains/admin/app/services"
	customerApp "github.com/smart-table/src/domains/customer/app/use_cases"
	customerAppErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	customerDomainErrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
)

type SmartTableCustomerQueryServiceImpl struct {
	placeTableService            *appServices.PlaceTableService
	orderInfoCommandHandler      *customerApp.OrderInfoCommandHandler
	placeOrderListCommandHandler *customerApp.PlaceOrderListCommandHandler
	placeOrderEditCommandHandler *customerApp.PlaceOrderEditCommandHandler
}

func NewSmartTableQueryServiceImpl(
	placeTableService *appServices.PlaceTableService,
	orderInfoCommandHandler *customerApp.OrderInfoCommandHandler,
	placeOrderListCommandHandler *customerApp.PlaceOrderListCommandHandler,
	placeOrderEditCommandHandler *customerApp.PlaceOrderEditCommandHandler,
) *SmartTableCustomerQueryServiceImpl {
	return &SmartTableCustomerQueryServiceImpl{
		placeTableService:            placeTableService,
		orderInfoCommandHandler:      orderInfoCommandHandler,
		placeOrderListCommandHandler: placeOrderListCommandHandler,
		placeOrderEditCommandHandler: placeOrderEditCommandHandler,
	}
}

func (s *SmartTableCustomerQueryServiceImpl) GetPlaceOrder(
	placeUUID, orderUUID uuid.UUID,
) (defsInternalCustomerDTO.OrderInfoDTO, error) {
	response, err := s.orderInfoCommandHandler.Handle(&customerApp.OrderInfoCommand{
		OrderUUID: orderUUID,
	})
	if err != nil {
		return defsInternalCustomerDTO.OrderInfoDTO{}, appQueriesErrors.UnsuccessOrderFetch{InnerError: err}
	}

	orderPlaceUUID, err := s.placeTableService.GetPlaceUUIDFromTableID(response.Order.Get().GetTableID())
	if err != nil {
		return defsInternalCustomerDTO.OrderInfoDTO{}, err
	}

	if orderPlaceUUID != placeUUID {
		return defsInternalCustomerDTO.OrderInfoDTO{}, appQueriesErrors.OrderNotBelongToPLace{
			OrderUUID: orderUUID,
			PlaceUUID: placeUUID,
		}
	}

	return convertOrderToOrderInfoDTO(response.Order, s.placeTableService), nil
}

func (s *SmartTableCustomerQueryServiceImpl) GetPlaceOrders(
	placeUUID uuid.UUID,
	isActive bool,
) ([]defsInternalCustomerDTO.OrderMainInfoDTO, error) {
	response, err := s.placeOrderListCommandHandler.Handle(&customerApp.PlaceOrderListCommand{
		PlaceUUID: placeUUID,
		IsActive:  isActive,
	})
	if err != nil {
		return nil, appQueriesErrors.UnsuccessOrderFetch{InnerError: err}
	}

	result := make([]defsInternalCustomerDTO.OrderMainInfoDTO, 0, len(response.OrderList))

	for i := range response.OrderList {
		orderTotalPrice := decimal.Zero

		for _, item := range response.OrderList[i].Get().GetItems() {
			if !item.Get().GetIsDraft() {
				orderTotalPrice = orderTotalPrice.Add(item.Get().GetPrice())
			}
		}

		result = append(result, convertOrderToOrderMainInfoDTO(response.OrderList[i], orderTotalPrice, s.placeTableService))
	}

	return result, nil
}

func (s *SmartTableCustomerQueryServiceImpl) EditPlaceOrder(
	orderUUID uuid.UUID,
	tableID string,
	orderStatus utils.Optional[string],
	itemEditGroup utils.Optional[defsInternalCustomerDTO.ItemEditGroupDTO],
) error {
	err := s.placeOrderEditCommandHandler.Handle(&customerApp.PlaceOrderEditCommand{
		OrderUUID:     orderUUID,
		TableID:       tableID,
		OrderStatus:   orderStatus,
		ItemEditGroup: itemEditGroup,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[customerAppErrors.IncorrectTableID](err):
			PlaceUUID, err := s.placeTableService.GetPlaceUUIDFromTableID(tableID)
			if err != nil {
				return err
			}

			return appQueriesErrors.OrderNotBelongToPLace{
				PlaceUUID: PlaceUUID,
				OrderUUID: orderUUID,
			}
		case utils.IsTheSameErrorType[customerAppErrors.IncorrectEditOrderRequest](err):
			return appQueriesErrors.IncorrectEditOrderRequest{}
		case utils.IsTheSameErrorType[customerDomainErrors.OrderNotFound](err):
			return appQueriesErrors.OrderNotFound{OrderUUID: orderUUID}
		case utils.IsTheSameErrorType[customerDomainErrors.InvalidOrderStatus](err):
			return appQueriesErrors.InvalidOrderStatus{OrderStatus: orderStatus}
		case utils.IsTheSameErrorType[customerDomainErrors.InvalidOrderResolution](err):
			return appQueriesErrors.InvalidOrderResolution{OrderResolution: orderStatus}
		case utils.IsTheSameErrorType[customerDomainErrors.InvalidItemStatus](err):
			return appQueriesErrors.InvalidItemStatus{ItemEditGpoup: itemEditGroup}
		case utils.IsTheSameErrorType[customerDomainErrors.DraftItemStatusChangeNotAllowed](err):
			return appQueriesErrors.DraftItemStatusChangeNotAllowed{ItemEditGpoup: itemEditGroup}
		case utils.IsTheSameErrorType[customerDomainErrors.ItemStatusChangeRequiresOrderStatusUpdate](err):
			return appQueriesErrors.ItemStatusChangeRequiresOrderStatusUpdate{ItemEditGpoup: itemEditGroup}
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from customer command handler: %v", err))

		return appQueriesErrors.UnsuccessOrderEdit{InnerError: err}
	}

	return nil
}

// Converters.
type CustomerInfoImpl struct {
	UUID         uuid.UUID
	TgLogin      string
	TgID         string
	ItemGroupMap map[string]*defsInternalCustomerDTO.ItemGroupInfoDTO
	TotalPrice   decimal.Decimal
}

func createItemGroupInfoDTO(
	item utils.SharedRef[domain.Item],
) defsInternalCustomerDTO.ItemGroupInfoDTO {
	itemGroupInfo := defsInternalCustomerDTO.ItemGroupInfoDTO{
		MenuDishUUID: item.Get().GetDishUUID(),
		ItemUUIDList: []uuid.UUID{item.Get().GetUUID()},
		Status:       string(item.Get().GetStatus()),
		Name:         item.Get().GetName(),
		ItemPrice:    item.Get().GetPrice().String(),
		ResultPrice:  item.Get().GetPrice().String(),
		Count:        1,
	}

	if item.Get().GetComment().HasValue() {
		comment := item.Get().GetComment().Value()
		itemGroupInfo.Comment = &comment
	}

	if item.Get().GetResolution().HasValue() {
		resolution := string(item.Get().GetResolution().Value())
		itemGroupInfo.Resolution = &resolution
	}

	return itemGroupInfo
}

func convertCustomerToCustomerInfoImpl(
	customer utils.SharedRef[domain.Customer],
) CustomerInfoImpl {
	return CustomerInfoImpl{
		UUID:         customer.Get().GetUUID(),
		TgLogin:      customer.Get().GetTgLogin(),
		TgID:         customer.Get().GetTgID(),
		ItemGroupMap: make(map[string]*defsInternalCustomerDTO.ItemGroupInfoDTO),
		TotalPrice:   decimal.Zero,
	}
}

func convertCustomerInfoImplToCustomerInfo(
	customerInfoImpl *CustomerInfoImpl,
) defsInternalCustomerDTO.CustomerInfoDTO {
	itemGroupList := make([]defsInternalCustomerDTO.ItemGroupInfoDTO, 0, len(customerInfoImpl.ItemGroupMap))
	for i := range customerInfoImpl.ItemGroupMap {
		itemGroupList = append(itemGroupList, *customerInfoImpl.ItemGroupMap[i])
	}

	return defsInternalCustomerDTO.CustomerInfoDTO{
		UUID:          customerInfoImpl.UUID,
		TgLogin:       customerInfoImpl.TgLogin,
		TgID:          customerInfoImpl.TgID,
		ItemGroupList: itemGroupList,
		TotalPrice:    customerInfoImpl.TotalPrice.String(),
	}
}

func convertOrderToOrderMainInfoDTO(
	order utils.SharedRef[domain.Order],
	totalPrice decimal.Decimal,
	placeTableService *appServices.PlaceTableService,
) defsInternalCustomerDTO.OrderMainInfoDTO {
	orderMainInfo := defsInternalCustomerDTO.OrderMainInfoDTO{
		UUID:        order.Get().GetUUID(),
		Status:      string(order.Get().GetStatus()),
		GuestsCount: len(order.Get().GetCustomers()),
		CreatedAt:   order.Get().GetCreatedAt(),
		TotalPrice:  totalPrice.String(),
	}

	tableNumber, err := placeTableService.GetTableNumberFromTableID(order.Get().GetTableID())
	if err != nil {
		panic(err)
	}

	orderMainInfo.TableNumber = tableNumber

	if order.Get().GetResolution().HasValue() {
		resolution := string(order.Get().GetResolution().Value())
		orderMainInfo.Resolution = &resolution
	}

	return orderMainInfo
}

func convertOrderToOrderInfoDTO(
	order utils.SharedRef[domain.Order],
	placeTableService *appServices.PlaceTableService,
) defsInternalCustomerDTO.OrderInfoDTO {
	customerInfoImplMap := make(map[uuid.UUID]*CustomerInfoImpl)

	for _, customer := range order.Get().GetCustomers() {
		customerInfoImpl := convertCustomerToCustomerInfoImpl(customer)
		customerInfoImplMap[customer.Get().GetUUID()] = &customerInfoImpl
	}

	orderTotalPrice := decimal.Zero

	for _, item := range order.Get().GetItems() {
		if item.Get().GetIsDraft() {
			continue
		}

		key := fmt.Sprintf("{%s}_{%s}_{%s}", item.Get().GetDishUUID(), item.Get().GetStatus(), item.Get().GetComment().ValueOr(""))

		customerUUID := item.Get().GetCustomer().Get().GetUUID()
		customerInfoImpl := customerInfoImplMap[customerUUID]

		itemGroup, isExists := customerInfoImpl.ItemGroupMap[key]
		if !isExists {
			ItemGroup := createItemGroupInfoDTO(item)
			customerInfoImpl.ItemGroupMap[key] = &ItemGroup
			customerInfoImpl.TotalPrice = customerInfoImpl.TotalPrice.Add(item.Get().GetPrice())
			orderTotalPrice = orderTotalPrice.Add(item.Get().GetPrice())
		} else {
			itemGroup.ItemUUIDList = append(itemGroup.ItemUUIDList, item.Get().GetUUID())

			itemPrice, err := decimal.NewFromString(itemGroup.ItemPrice)
			if err != nil {
				panic(err)
			}

			resultPrice, err := decimal.NewFromString(itemGroup.ResultPrice)
			if err != nil {
				panic(err)
			}

			resultPrice = resultPrice.Add(itemPrice)
			itemGroup.ResultPrice = resultPrice.String()

			customerInfoImpl.TotalPrice = customerInfoImpl.TotalPrice.Add(itemPrice)

			itemGroup.Count++

			orderTotalPrice = orderTotalPrice.Add(itemPrice)
		}
	}

	customerList := make([]defsInternalCustomerDTO.CustomerInfoDTO, 0, len(customerInfoImplMap))
	for _, customerInfoImpl := range customerInfoImplMap {
		customerList = append(customerList, convertCustomerInfoImplToCustomerInfo(customerInfoImpl))
	}

	orderMainInfo := convertOrderToOrderMainInfoDTO(order, orderTotalPrice, placeTableService)

	return defsInternalCustomerDTO.OrderInfoDTO{
		OrderMainInfo: orderMainInfo,
		CustomerList:  customerList,
	}
}
