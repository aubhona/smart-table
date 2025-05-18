package app //nolint

import (
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type FinishOrderCommandHandler struct {
	orderRepository domain.OrderRepository
}

func NewFinishOrderCommandHandler(orderRepository domain.OrderRepository) *FinishOrderCommandHandler {
	return &FinishOrderCommandHandler{orderRepository: orderRepository}
}

func (handler *FinishOrderCommandHandler) Handle(
	command *FinishOrderCommand,
) error {
	tx, err := handler.orderRepository.Begin()
	if err != nil {
		return err
	}

	defer utils.Rollback(handler.orderRepository, tx)

	order, err := handler.orderRepository.FindOrderForUpdate(tx, command.OrderUUID)
	if err != nil {
		return err
	}

	if !order.Get().ContainsCustomer(command.CustomerUUID) {
		return appErrors.OrderAccessDenied{OrderUUID: command.OrderUUID, CustomerUUID: command.CustomerUUID}
	}

	order.Get().MarkWaitingPayment()

	err = handler.orderRepository.Update(tx, order)
	if err != nil {
		return err
	}

	return handler.orderRepository.Commit(tx)
}
