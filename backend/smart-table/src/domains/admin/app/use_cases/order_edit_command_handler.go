package app

import (
	appQueries "github.com/smart-table/src/domains/admin/app/queries"
	appServices "github.com/smart-table/src/domains/admin/app/services"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
	"go.uber.org/zap"
)

type OrderEditCommandHandler struct {
	placeRepository    domain.PlaceRepository
	userRepository     domain.UserRepository
	appCustomerQueries appQueries.SmartTableCustomerQueryService
	placeTableService  *appServices.PlaceTableService
}

func NewOrderEditCommandHandler(
	placeRepository domain.PlaceRepository,
	userRepository domain.UserRepository,
	appCustomerQueries appQueries.SmartTableCustomerQueryService,
	placeTableService *appServices.PlaceTableService,
) *OrderEditCommandHandler {
	return &OrderEditCommandHandler{
		placeRepository,
		userRepository,
		appCustomerQueries,
		placeTableService,
	}
}

func (handler *OrderEditCommandHandler) Handle(
	orderEditCommand *OrderEditCommand,
) error {
	place, err := handler.placeRepository.FindPlace(orderEditCommand.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding place by uuid", zap.Error(err))
		return err
	}

	if !domain.IsHasAccess(orderEditCommand.UserUUID, place, domain.All) {
		return appErrors.PlaceAccessDenied{
			UserUUID:  orderEditCommand.UserUUID,
			PlaceUUID: orderEditCommand.PlaceUUID,
		}
	}

	tableID := handler.placeTableService.BuildTableID(orderEditCommand.PlaceUUID, orderEditCommand.TableNumber)

	return handler.appCustomerQueries.EditPlaceOrder(
		orderEditCommand.OrderUUID,
		tableID,
		orderEditCommand.OrderStatus,
		orderEditCommand.ItemEditGroup,
	)
}
