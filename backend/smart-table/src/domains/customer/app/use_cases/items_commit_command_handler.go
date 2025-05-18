package app

import (
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type ItemsCommitCommandHandler struct {
	orderRepository domain.OrderRepository
}

func NewItemsCommitCommandHandler(orderRepository domain.OrderRepository) *ItemsCommitCommandHandler {
	return &ItemsCommitCommandHandler{orderRepository: orderRepository}
}

func (handler *ItemsCommitCommandHandler) Handle(
	command *ItemsCommitCommand,
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

	order.Get().CommitItems(command.CustomerUUID)

	err = handler.orderRepository.Update(tx, order)
	if err != nil {
		return err
	}

	return handler.orderRepository.Commit(tx)
}
