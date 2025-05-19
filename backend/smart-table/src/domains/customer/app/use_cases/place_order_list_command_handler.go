package app

import (
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type PlaceOrderListCommandHandlerResult struct {
	OrderList []utils.SharedRef[domain.Order]
}

type PlaceOrderListCommandHandler struct {
	orderRepository domain.OrderRepository
}

func NewPlaceOrderListCommandHandler(
	orderRepository domain.OrderRepository,
) *PlaceOrderListCommandHandler {
	return &PlaceOrderListCommandHandler{
		orderRepository,
	}
}

func (handler *PlaceOrderListCommandHandler) Handle(command *PlaceOrderListCommand) (PlaceOrderListCommandHandlerResult, error) {
	orderList, err := handler.orderRepository.FindOrdersByPlaceUUID(command.PlaceUUID, command.IsActive)
	if err != nil {
		return PlaceOrderListCommandHandlerResult{}, err
	}

	return PlaceOrderListCommandHandlerResult{orderList}, nil
}
