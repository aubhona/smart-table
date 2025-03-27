package app

import (
	"context"

	"github.com/smart-table/src/utils"

	"github.com/google/uuid"
	app "github.com/smart-table/src/domains/customer/app/services"
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
	roomCodeService    *app.RoomCodeService
}

func NewOrderCreateCommandHandler(
	orderRepository domain.OrderRepository,
	customerRepository domain.CustomerRepository,
	uuidGenerator *domainServices.UUIDGenerator,
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
	ctx := context.Background()
	user, err := handler.customerRepository.FindCustomer(context.Background(), createCommand.CustomerUUID)

	if err != nil {
		return OrderCreateCommandHandlerResult{}, err
	}

	order, err := handler.orderRepository.FindActiveOrderByTableID(context.Background(), createCommand.TableID)
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
		err = handler.orderRepository.Begin(ctx)

		if err != nil {
			return OrderCreateCommandHandlerResult{}, err
		}

		defer func(orderRepository domain.OrderRepository, ctx context.Context) {
			_ = orderRepository.Commit(ctx)
		}(handler.orderRepository, ctx)

		err = handler.orderRepository.Save(ctx, order)

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

	return OrderCreateCommandHandlerResult{order.Get().GetUUID()}, nil
}
