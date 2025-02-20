package app

import (
	"context"
	"errors"
	app "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/app/services"
	app_errors "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/app/use_cases/errors"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain"
	domain_errors "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain/errors"
	domain_services "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain/services"
	"github.com/google/uuid"
)

type OrderCreateCommandHandlerResult struct {
	OrderUUID uuid.UUID
}

type OrderCreateCommandHandler struct {
	orderRepository    domain.OrderRepository
	customerRepository domain.CustomerRepository
	uuidGenerator      *domain_services.UUIDGenerator
	roomCodeService    *app.RoomCodeService
}

func NewOrderCreateCommandHandler(
	orderRepository domain.OrderRepository,
	customerRepository domain.CustomerRepository,
	uuidGenerator *domain_services.UUIDGenerator,
	roomCodeService *app.RoomCodeService,
) *OrderCreateCommandHandler {
	return &OrderCreateCommandHandler{
		orderRepository,
		customerRepository,
		uuidGenerator,
		roomCodeService,
	}
}

func (handler *OrderCreateCommandHandler) Handle(createCommand *OrderCreateCommand) (OrderCreateCommandHandlerResult, error) {
	user, err := handler.customerRepository.FindCustomer(context.Background(), createCommand.CustomerUuid)
	if err != nil {
		return OrderCreateCommandHandlerResult{}, err
	}

	order, err := handler.orderRepository.FindActiveOrderByTableId(context.Background(), createCommand.TableId)
	isNewOrder := false

	if err != nil {
		if errors.Is(err, domain_errors.OrderNotFoundByTableId{}) {
			isNewOrder = true
		} else {
			return OrderCreateCommandHandlerResult{}, err
		}
	}

	if isNewOrder {
		// TODO: Check table id existence
		roomCode, err := handler.roomCodeService.CreateRoomCode(createCommand.TableId, createCommand.CustomerUuid)
		if err != nil {
			return OrderCreateCommandHandlerResult{}, err
		}

		order = domain.NewOrder(roomCode, createCommand.TableId, user, handler.uuidGenerator)

		return OrderCreateCommandHandlerResult{order.Get().GetUUID()}, nil
	}

	if !createCommand.RoomCode.HasValue() {
		return OrderCreateCommandHandlerResult{}, app_errors.IncorrectRoomCodeError{RoomCode: createCommand.RoomCode}
	}

	if !handler.roomCodeService.VerifyRoomCode(order, createCommand.RoomCode.Value()) {
		return OrderCreateCommandHandlerResult{}, app_errors.IncorrectRoomCodeError{RoomCode: createCommand.RoomCode}
	}

	return OrderCreateCommandHandlerResult{order.Get().GetUUID()}, nil
}
