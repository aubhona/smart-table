package app

import (
	"io"
	"slices"

	"github.com/samber/lo"
	defsInternalOrder "github.com/smart-table/src/codegen/intern/order"

	"github.com/shopspring/decimal"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/utils"

	"github.com/google/uuid"
	appQueries "github.com/smart-table/src/domains/customer/app/queries"
	"github.com/smart-table/src/domains/customer/domain"
)

type CatalogItemDTO struct {
	ID         uuid.UUID
	Calories   int
	Weight     int
	Count      int
	Name       string
	Price      string
	PictureKey string
	Category   string
	Picture    io.Reader
}

type CatalogCommandHandlerResult struct {
	GoTipScreen bool
	Items       []CatalogItemDTO
	Categories  []string
	TotalPrice  decimal.Decimal
	RoomCode    string
}

type CatalogCommandHandler struct {
	orderRepository domain.OrderRepository
	appAdminQueries appQueries.SmartTableAdminQueryService
}

func NewCatalogCommandHandler(
	orderRepository domain.OrderRepository,
	appAdminQueries appQueries.SmartTableAdminQueryService,
) *CatalogCommandHandler {
	return &CatalogCommandHandler{
		orderRepository,
		appAdminQueries,
	}
}

func (handler *CatalogCommandHandler) Handle(command *CatalogCommand) (CatalogCommandHandlerResult, error) {
	order, err := handler.orderRepository.FindOrder(command.OrderUUID)
	if err != nil {
		return CatalogCommandHandlerResult{}, err
	}

	if !order.Get().ContainsCustomer(command.CustomerUUID) {
		return CatalogCommandHandlerResult{},
			appErrors.OrderAccessDenied{OrderUUID: command.OrderUUID, CustomerUUID: command.CustomerUUID}
	}

	menuDishList, err := handler.appAdminQueries.GetCatalog(order.Get().GetTableID(), command.NeedPicture)
	if err != nil {
		return CatalogCommandHandlerResult{}, err
	}

	menuDishUUIDToItemsMap := make(map[uuid.UUID][]utils.SharedRef[domain.Item])

	for _, item := range order.Get().GetDraftedItemsByCustomerUUID(command.CustomerUUID) {
		_, exist := menuDishUUIDToItemsMap[item.Get().GetDishUUID()]

		if !exist {
			menuDishUUIDToItemsMap[item.Get().GetDishUUID()] = make([]utils.SharedRef[domain.Item], 0)
		}

		menuDishUUIDToItemsMap[item.Get().GetDishUUID()] = append(menuDishUUIDToItemsMap[item.Get().GetDishUUID()], item)
	}

	result := CatalogCommandHandlerResult{
		RoomCode:    order.Get().GetRoomCode(),
		TotalPrice:  order.Get().GetDraftItemsTotalPriceByCustomerUUID(command.CustomerUUID),
		Items:       make([]CatalogItemDTO, 0, len(menuDishList)),
		GoTipScreen: order.Get().GetStatus() == defsInternalOrder.OrderStatusPaymentWaiting,
	}

	uniqueCategories := make(map[string]interface{})

	for i := range menuDishList {
		menuDish := &menuDishList[i]

		items, exist := menuDishUUIDToItemsMap[menuDish.ID]
		if exist {
			menuDish.Calories = items[0].Get().GetCalories()
			menuDish.Category = items[0].Get().GetCategory()
			menuDish.Description = items[0].Get().GetDescription()
			menuDish.Name = items[0].Get().GetName()
			menuDish.Price = items[0].Get().GetPrice().String()
			menuDish.Weight = items[0].Get().GetWeight()
		}

		pictureReader, err := menuDish.Picture.Reader()
		if err != nil {
			return CatalogCommandHandlerResult{}, err
		}

		result.Items = append(result.Items, CatalogItemDTO{
			ID:         menuDish.ID,
			Calories:   menuDish.Calories,
			Weight:     menuDish.Weight,
			Count:      len(items),
			Name:       menuDish.Name,
			Price:      menuDish.Price,
			PictureKey: menuDish.PictureKey,
			Category:   menuDish.Category,
			Picture:    pictureReader,
		})

		uniqueCategories[menuDish.Category] = nil
	}

	result.Categories = lo.Keys(uniqueCategories)

	slices.Sort(result.Categories)

	return result, nil
}
