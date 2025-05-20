package app

import (
	"io"

	"github.com/samber/lo"

	"github.com/shopspring/decimal"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/utils"

	"github.com/google/uuid"
	appQueries "github.com/smart-table/src/domains/customer/app/queries"
	"github.com/smart-table/src/domains/customer/domain"
)

type ItemStateDTO struct {
	Calories    int
	Category    string
	Count       int
	Description string
	ID          uuid.UUID
	Name        string
	Price       decimal.Decimal
	ResultPrice decimal.Decimal
	Weight      int
	PictureKey  string
	Picture     io.Reader
}

type ItemStateCommandHandlerResult struct {
	ItemsState ItemStateDTO
}

type ItemStateCommandHandler struct {
	orderRepository domain.OrderRepository
	appAdminQueries appQueries.SmartTableAdminQueryService
}

func NewItemStateCommandHandler(
	orderRepository domain.OrderRepository,
	appAdminQueries appQueries.SmartTableAdminQueryService,
) *ItemStateCommandHandler {
	return &ItemStateCommandHandler{
		orderRepository,
		appAdminQueries,
	}
}

func (handler *ItemStateCommandHandler) Handle(command *ItemStateCommand) (ItemStateCommandHandlerResult, error) { //nolint
	order, err := handler.orderRepository.FindOrder(command.OrderUUID)
	if err != nil {
		return ItemStateCommandHandlerResult{}, err
	}

	if !order.Get().ContainsCustomer(command.CustomerUUID) {
		return ItemStateCommandHandlerResult{},
			appErrors.OrderAccessDenied{OrderUUID: command.OrderUUID, CustomerUUID: command.CustomerUUID}
	}

	items := lo.Filter(order.Get().GetDraftedItemsByCustomerUUID(command.CustomerUUID), func(item utils.SharedRef[domain.Item], _ int) bool {
		return item.Get().GetDishUUID() == command.DishUUD && item.Get().GetComment() == command.Comment
	})

	menuDish, err := handler.appAdminQueries.GetMenuDish(order.Get().GetTableID(), command.DishUUD, true)
	if err != nil {
		return ItemStateCommandHandlerResult{}, err
	}

	if len(items) > 0 {
		item := items[0]

		menuDish.Calories = item.Get().GetCalories()
		menuDish.Category = item.Get().GetCategory()
		menuDish.Description = item.Get().GetDescription()
		menuDish.Name = item.Get().GetName()
		menuDish.Price = item.Get().GetPrice().String()
		menuDish.Weight = item.Get().GetWeight()
	}

	pictureReader, err := menuDish.Picture.Reader()
	if err != nil {
		return ItemStateCommandHandlerResult{}, err
	}

	price := decimal.RequireFromString(menuDish.Price)

	return ItemStateCommandHandlerResult{
		ItemsState: ItemStateDTO{
			Calories:    menuDish.Calories,
			Category:    menuDish.Category,
			Count:       len(items),
			Description: menuDish.Description,
			ID:          menuDish.ID,
			Name:        menuDish.Name,
			Price:       price,
			ResultPrice: price.Mul(decimal.NewFromInt(int64(len(items)))),
			Weight:      menuDish.Weight,
			PictureKey:  menuDish.PictureKey,
			Picture:     pictureReader,
		},
	}, nil
}
