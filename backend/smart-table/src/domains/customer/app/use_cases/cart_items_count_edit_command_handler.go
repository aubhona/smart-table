package app

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	appQueries "github.com/smart-table/src/domains/customer/app/queries"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/domains/customer/domain"
	domainServices "github.com/smart-table/src/domains/customer/domain/services"
	"github.com/smart-table/src/utils"
)

type CartItemsCountEditCommandHandler struct {
	orderRepository domain.OrderRepository
	appAdminQueries appQueries.SmartTableAdminQueryService
	uuidGenerator   *domainServices.UUIDGenerator
}

func NewCartItemsCountEditCommandHandler(
	orderRepository domain.OrderRepository,
	appAdminQueries appQueries.SmartTableAdminQueryService,
	uuidGenerator *domainServices.UUIDGenerator,
) *CartItemsCountEditCommandHandler {
	return &CartItemsCountEditCommandHandler{orderRepository, appAdminQueries, uuidGenerator}
}

func (handler *CartItemsCountEditCommandHandler) Handle(command *CartItemsCountEditCommand) error { //nolint
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

	dataToCreateItem := struct {
		dishUUID    uuid.UUID
		customer    utils.SharedRef[domain.Customer]
		comment     utils.Optional[string]
		name        string
		description string
		pictureLink string
		weight      int
		calories    int
		category    string
		price       decimal.Decimal
	}{
		dishUUID: command.DishUUID,
		customer: order.Get().GetCustomerByUUID(command.CustomerUUID).Value(),
		comment:  command.Comment,
	}

	if command.EditCount >= 0 {
		item := order.Get().GetItemByDishUUID(command.DishUUID)

		if item.HasValue() {
			dataToCreateItem.name = item.Value().Get().GetName()
			dataToCreateItem.description = item.Value().Get().GetDescription()
			dataToCreateItem.pictureLink = item.Value().Get().GetPictureLink()
			dataToCreateItem.weight = item.Value().Get().GetWeight()
			dataToCreateItem.calories = item.Value().Get().GetCalories()
			dataToCreateItem.category = item.Value().Get().GetCategory()
			dataToCreateItem.price = item.Value().Get().GetPrice()
		} else {
			menDish, err := handler.appAdminQueries.GetMenuDish(order.Get().GetTableID(), command.DishUUID, false)
			if err != nil {
				return err
			}

			dataToCreateItem.name = menDish.Name
			dataToCreateItem.description = menDish.Description
			dataToCreateItem.pictureLink = menDish.PictureKey
			dataToCreateItem.weight = menDish.Weight
			dataToCreateItem.calories = menDish.Calories
			dataToCreateItem.category = menDish.Category

			dataToCreateItem.price, err = decimal.NewFromString(menDish.Price)
			if err != nil {
				return err
			}
		}

		for range command.EditCount {
			order.Get().DraftItem(
				dataToCreateItem.dishUUID,
				dataToCreateItem.customer,
				dataToCreateItem.comment,
				dataToCreateItem.name,
				dataToCreateItem.description,
				dataToCreateItem.pictureLink,
				dataToCreateItem.weight,
				dataToCreateItem.calories,
				dataToCreateItem.category,
				dataToCreateItem.price,
				handler.uuidGenerator,
			)
		}
	} else {
		err = order.Get().DeleteItemsByDishUUID(command.DishUUID, -command.EditCount)
		if err != nil {
			return err
		}
	}

	err = handler.orderRepository.Save(tx, order)
	if err != nil {
		return err
	}

	return handler.orderRepository.Commit(tx)
}
