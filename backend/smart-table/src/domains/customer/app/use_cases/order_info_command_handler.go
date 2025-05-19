package app

import (
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type OrderInfoCommandHandlerResult struct {
	Order utils.SharedRef[domain.Order]
}

type OrderInfoCommandHandler struct {
	orderRepository domain.OrderRepository
}

func NewOrderInfoCommandHandler(
	orderRepository domain.OrderRepository,
) *OrderInfoCommandHandler {
	return &OrderInfoCommandHandler{
		orderRepository,
	}
}

func (handler *OrderInfoCommandHandler) Handle(command *OrderInfoCommand) (OrderInfoCommandHandlerResult, error) {
	order, err := handler.orderRepository.FindOrder(command.OrderUUID)
	if err != nil {
		return OrderInfoCommandHandlerResult{}, err
	}

	return OrderInfoCommandHandlerResult{order}, nil
}
