package app

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
)

type CatalogUpdateInfoCommandHandlerResult struct {
	TotalPrice      decimal.Decimal
	MenuUpdatedInfo []struct {
		ID    uuid.UUID
		Count int
	}
}

type CatalogUpdateInfoCommandHandler struct {
	orderRepository domain.OrderRepository
}

func NewCatalogUpdateInfoCommandHandler(orderRepository domain.OrderRepository) *CatalogUpdateInfoCommandHandler {
	return &CatalogUpdateInfoCommandHandler{orderRepository: orderRepository}
}

func (handler *CatalogUpdateInfoCommandHandler) Handle(
	command *CatalogUpdateInfoCommand,
) (CatalogUpdateInfoCommandHandlerResult, error) {
	order, err := handler.orderRepository.FindOrder(command.OrderUUID)
	if err != nil {
		return CatalogUpdateInfoCommandHandlerResult{}, err
	}

	if !order.Get().ContainsCustomer(command.CustomerUUID) {
		return CatalogUpdateInfoCommandHandlerResult{},
			appErrors.OrderAccessDenied{OrderUUID: command.OrderUUID, CustomerUUID: command.CustomerUUID}
	}

	result := CatalogUpdateInfoCommandHandlerResult{
		TotalPrice: order.Get().GetDraftItemsTotalPriceByCustomerUUID(command.CustomerUUID),
		MenuUpdatedInfo: make([]struct {
			ID    uuid.UUID
			Count int
		}, 0),
	}

	countMap := make(map[uuid.UUID]int)

	for _, item := range order.Get().GetItems() {
		if !item.Get().GetIsDraft() {
			continue
		}

		countMap[item.Get().GetDishUUID()]++
	}

	for id, count := range countMap {
		result.MenuUpdatedInfo = append(result.MenuUpdatedInfo, struct {
			ID    uuid.UUID
			Count int
		}{id, count},
		)
	}

	return result, nil
}
