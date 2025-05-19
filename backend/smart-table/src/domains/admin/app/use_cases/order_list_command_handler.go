package app

import (
	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	appQueries "github.com/smart-table/src/domains/admin/app/queries"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
	"go.uber.org/zap"
)

type OrderListCommandHandlerResult struct {
	OrderList []defsInternalCustomerDTO.OrderMainInfoDTO
}

type OrderListCommandHandler struct {
	placeRepository    domain.PlaceRepository
	userRepository     domain.UserRepository
	appCustomerQueries appQueries.SmartTableCustomerQueryService
}

func NewOrderListCommandHandler(
	placeRepository domain.PlaceRepository,
	userRepository domain.UserRepository,
	appCustomerQueries appQueries.SmartTableCustomerQueryService,
) *OrderListCommandHandler {
	return &OrderListCommandHandler{
		placeRepository,
		userRepository,
		appCustomerQueries,
	}
}

func (handler *OrderListCommandHandler) Handle(
	orderListCommand *OrderListCommand,
) (OrderListCommandHandlerResult, error) {
	place, err := handler.placeRepository.FindPlace(orderListCommand.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding place by uuid", zap.Error(err))
		return OrderListCommandHandlerResult{}, err
	}

	if !domain.IsHasAccess(orderListCommand.UserUUID, place, domain.All) {
		return OrderListCommandHandlerResult{}, appErrors.PlaceAccessDenied{
			UserUUID:  orderListCommand.UserUUID,
			PlaceUUID: orderListCommand.PlaceUUID,
		}
	}

	orderList, err := handler.appCustomerQueries.GetPlaceOrders(orderListCommand.PlaceUUID, orderListCommand.IsActive)
	if err != nil {
		return OrderListCommandHandlerResult{}, err
	}

	return OrderListCommandHandlerResult{orderList}, nil
}
