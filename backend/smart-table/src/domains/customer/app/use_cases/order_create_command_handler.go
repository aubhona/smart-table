package app

import (
	"github.com/smart-table/src/utils"

	"github.com/google/uuid"
	appQueries "github.com/smart-table/src/domains/customer/app/queries"
	appServices "github.com/smart-table/src/domains/customer/app/services"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	domainErrors "github.com/smart-table/src/domains/customer/domain/errors"
	domainServices "github.com/smart-table/src/domains/customer/domain/services"
)

type OrderCreateCommandHandlerResult struct {
	OrderUUID uuid.UUID
}

type OrderCreateCommandHandler struct {
	orderRepository    domain.OrderRepository
	customerRepository domain.CustomerRepository
	uuidGenerator      *domainServices.UUIDGenerator
	roomCodeService    *appServices.RoomCodeService
	appAdminQueries    appQueries.SmartTableAdminQueryService
}

func NewOrderCreateCommandHandler(
	orderRepository domain.OrderRepository,
	customerRepository domain.CustomerRepository,
	uuidGenerator *domainServices.UUIDGenerator,
	roomCodeService *appServices.RoomCodeService,
	appAdminQueries appQueries.SmartTableAdminQueryService,
) *OrderCreateCommandHandler {
	return &OrderCreateCommandHandler{
		orderRepository,
		customerRepository,
		uuidGenerator,
		roomCodeService,
		appAdminQueries,
	}
}

func (handler *OrderCreateCommandHandler) Handle(createCommand *OrderCreateCommand) (OrderCreateCommandHandlerResult, error) {
	user, err := handler.customerRepository.FindCustomer(createCommand.CustomerUUID)
	if err != nil {
		return OrderCreateCommandHandlerResult{}, err
	}

	//nolint
	// isValid, err := handler.appAdminQueries.TableIDValidate(createCommand.TableID)
	// if err != nil {
	// 	return OrderCreateCommandHandlerResult{}, err
	// }

	//nolint
	// if !isValid {
	// 	return OrderCreateCommandHandlerResult{}, appErrors.InvalidTableID{
	// 		TableID: createCommand.TableID,
	// 	}
	// }

	order, err := handler.orderRepository.FindActiveOrderByTableID(createCommand.TableID)
	isNewOrder := false

	if err != nil {
		if utils.IsTheSameErrorType[domainErrors.OrderNotFoundByTableID](err) {
			isNewOrder = true
		} else {
			return OrderCreateCommandHandlerResult{}, err
		}
	}

	if isNewOrder {
		//nolint: godox, gocritic
		// TODO: Check table id existence, check active order id on user
		roomCode, err := handler.roomCodeService.CreateRoomCode(createCommand.TableID, createCommand.CustomerUUID)

		if err != nil {
			return OrderCreateCommandHandlerResult{}, err
		}

		order = domain.NewOrder(roomCode, createCommand.TableID, user, handler.uuidGenerator)
		tx, err := handler.orderRepository.Begin()

		if err != nil {
			return OrderCreateCommandHandlerResult{}, err
		}

		defer utils.Rollback(handler.orderRepository, tx)

		err = handler.orderRepository.Save(tx, order)

		if err != nil {
			return OrderCreateCommandHandlerResult{}, err
		}

		err = handler.orderRepository.Commit(tx)
		if err != nil {
			return OrderCreateCommandHandlerResult{}, err
		}

		return OrderCreateCommandHandlerResult{order.Get().GetUUID()}, nil
	}

	if !createCommand.RoomCode.HasValue() {
		return OrderCreateCommandHandlerResult{}, appErrors.IncorrectRoomCodeError{RoomCode: createCommand.RoomCode}
	}

	if !handler.roomCodeService.VerifyRoomCode(order, createCommand.RoomCode.Value()) {
		return OrderCreateCommandHandlerResult{}, appErrors.IncorrectRoomCodeError{RoomCode: createCommand.RoomCode}
	}

	if !order.Get().ContainsCustomer(user.Get().GetUUID()) {
		order.Get().AddCustomer(user)
	}

	return OrderCreateCommandHandlerResult{order.Get().GetUUID()}, nil
}
