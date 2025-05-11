package app

import (
	appQueries "github.com/smart-table/src/domains/customer/app/queries"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type CustomerListCommandHandlerResult struct {
	Order utils.SharedRef[domain.Order]
}

type CustomerListCommandHandler struct {
	orderRepository domain.OrderRepository
	appAdminQueries appQueries.SmartTableAdminQueryService
}

func NewCustomerListCommandHandler(
	orderRepository domain.OrderRepository,
	appAdminQueries appQueries.SmartTableAdminQueryService,
) *CustomerListCommandHandler {
	return &CustomerListCommandHandler{
		orderRepository,
		appAdminQueries,
	}
}

func (handler *CustomerListCommandHandler) Handle(command *CustomerListCommand) (CustomerListCommandHandlerResult, error) {
	order, err := handler.orderRepository.FindOrder(command.OrderUUID)
	if err != nil {
		return CustomerListCommandHandlerResult{}, err
	}

	if !order.Get().ContainsCustomer(command.CustomerUUID) {
		return CustomerListCommandHandlerResult{},
			appErrors.OrderAccessDenied{OrderUUID: command.OrderUUID, CustomerUUID: command.CustomerUUID}
	}

	return CustomerListCommandHandlerResult{order}, nil
}
