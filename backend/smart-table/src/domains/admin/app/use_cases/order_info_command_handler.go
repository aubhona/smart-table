package app

import (
	defsInternalCustomerDTO "github.com/smart-table/src/codegen/intern/customer_dto"
	appQueries "github.com/smart-table/src/domains/admin/app/queries"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
	"go.uber.org/zap"
)

type OrderInfoCommandHandlerResult struct {
	OrderInfo defsInternalCustomerDTO.OrderInfoDTO
}

type OrderInfoCommandHandler struct {
	placeRepository    domain.PlaceRepository
	userRepository     domain.UserRepository
	appCustomerQueries appQueries.SmartTableCustomerQueryService
}

func NewOrderInfoCommandHandler(
	placeRepository domain.PlaceRepository,
	userRepository domain.UserRepository,
	appCustomerQueries appQueries.SmartTableCustomerQueryService,
) *OrderInfoCommandHandler {
	return &OrderInfoCommandHandler{
		placeRepository,
		userRepository,
		appCustomerQueries,
	}
}

func (handler *OrderInfoCommandHandler) Handle(
	orderInfoCommand *OrderInfoCommand,
) (OrderInfoCommandHandlerResult, error) {
	place, err := handler.placeRepository.FindPlace(orderInfoCommand.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding place by uuid", zap.Error(err))
		return OrderInfoCommandHandlerResult{}, err
	}

	if !domain.IsHasAccess(orderInfoCommand.UserUUID, place, domain.All) {
		return OrderInfoCommandHandlerResult{}, appErrors.PlaceAccessDenied{
			UserUUID:  orderInfoCommand.UserUUID,
			PlaceUUID: orderInfoCommand.PlaceUUID,
		}
	}

	order, err := handler.appCustomerQueries.GetPlaceOrder(orderInfoCommand.PlaceUUID, orderInfoCommand.OrderUUID)
	if err != nil {
		return OrderInfoCommandHandlerResult{}, err
	}

	return OrderInfoCommandHandlerResult{order}, nil
}
