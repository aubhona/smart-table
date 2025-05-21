package app

import (
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type PlaceOrderEditCommandHandler struct {
	orderRepository domain.OrderRepository
}

func NewPlaceOrderEditCommandHandler(
	orderRepository domain.OrderRepository,
) *PlaceOrderEditCommandHandler {
	return &PlaceOrderEditCommandHandler{
		orderRepository,
	}
}

func (handler *PlaceOrderEditCommandHandler) Handle(command *PlaceOrderEditCommand) error {
	tx, err := handler.orderRepository.Begin()
	if err != nil {
		return err
	}

	defer utils.Rollback(handler.orderRepository, tx)

	order, err := handler.orderRepository.FindOrderForUpdate(tx, command.OrderUUID)
	if err != nil {
		return err
	}

	if order.Get().GetTableID() != command.TableID {
		return appErrors.IncorrectTableID{
			RequestTableID: command.TableID,
			ActualTableID:  order.Get().GetTableID(),
		}
	}

	switch {
	case command.OrderStatus.HasValue():
		parsedOrderStatus, err := domain.ParseOrderStatus(command.OrderStatus.Value())
		if err != nil {
			return err
		}

		err = order.Get().SetStatus(parsedOrderStatus)
		if err != nil {
			return err
		}
	case command.ItemEditGpoup.HasValue() && len(command.ItemEditGpoup.Value().ItemUUIDList) != 0:
		err = order.Get().ChangeItemsStatus(command.ItemEditGpoup.Value().ItemUUIDList, command.ItemEditGpoup.Value().ItemStatus)
		if err != nil {
			return err
		}
	default:
		return appErrors.IncorrectEditOrderRequest{}
	}

	err = handler.orderRepository.Update(tx, order)
	if err != nil {
		return err
	}

	return handler.orderRepository.Commit(tx)
}
