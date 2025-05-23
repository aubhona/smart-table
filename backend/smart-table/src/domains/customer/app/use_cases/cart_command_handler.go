package app

import (
	"io"
	"sync"

	defsInternalAdminDTO "github.com/smart-table/src/codegen/intern/admin_dto"
	appServices "github.com/smart-table/src/domains/customer/app/services"
	"github.com/smart-table/src/logging"
	"go.uber.org/zap"

	"github.com/samber/lo"

	"github.com/shopspring/decimal"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/utils"

	"github.com/google/uuid"
	appQueries "github.com/smart-table/src/domains/customer/app/queries"
	"github.com/smart-table/src/domains/customer/domain"
)

type CartItemDTO struct {
	ID          uuid.UUID
	Count       int
	Name        string
	Price       decimal.Decimal
	ResultPrice decimal.Decimal
	Comment     utils.Optional[string]
	PictureKey  string
	Picture     io.Reader
}

type CartCommandHandlerResult struct {
	Items      []CartItemDTO
	TotalPrice decimal.Decimal
}

type CartCommandHandler struct {
	orderRepository domain.OrderRepository
	appAdminQueries appQueries.SmartTableAdminQueryService
	tipService      *appServices.TipService
}

func NewCartCommandHandler(
	orderRepository domain.OrderRepository,
	appAdminQueries appQueries.SmartTableAdminQueryService,
	tipService *appServices.TipService,
) *CartCommandHandler {
	return &CartCommandHandler{
		orderRepository,
		appAdminQueries,
		tipService,
	}
}

func (handler *CartCommandHandler) Handle(command *CartCommand) (CartCommandHandlerResult, error) { //nolint
	order, err := handler.orderRepository.FindOrder(command.OrderUUID)
	if err != nil {
		return CartCommandHandlerResult{}, err
	}

	if !order.Get().ContainsCustomer(command.CustomerUUID) {
		return CartCommandHandlerResult{},
			appErrors.OrderAccessDenied{OrderUUID: command.OrderUUID, CustomerUUID: command.CustomerUUID}
	}

	items := order.Get().GetDraftedItemsByCustomerUUID(command.CustomerUUID)
	uniqueItems := lo.SliceToMap(items, func(item utils.SharedRef[domain.Item]) (uuid.UUID, interface{}) {
		return item.Get().GetDishUUID(), nil
	})
	menuDishMap := make(map[uuid.UUID]defsInternalAdminDTO.MenuDishDTO)
	groupedItemsMap := make(map[string][]utils.SharedRef[domain.Item])

	mutex := &sync.Mutex{}
	waitGroup := &sync.WaitGroup{}

	for dishUUID := range uniqueItems {
		waitGroup.Add(1)

		go func(dishUUID uuid.UUID) {
			defer waitGroup.Done()

			menuDishDTO, err := handler.appAdminQueries.GetMenuDish(order.Get().GetTableID(), dishUUID, true)
			if err != nil {
				logging.GetLogger().Error("failed to get menu dish", zap.Error(err))
			}

			mutex.Lock()
			defer mutex.Unlock()

			menuDishMap[menuDishDTO.ID] = menuDishDTO
		}(dishUUID)
	}

	for _, item := range items {
		_, exist := groupedItemsMap[handler.tipService.GetItemsGroupKey(item)]
		if !exist {
			groupedItemsMap[handler.tipService.GetItemsGroupKey(item)] = make([]utils.SharedRef[domain.Item], 0)
		}

		groupedItemsMap[handler.tipService.GetItemsGroupKey(item)] = append(groupedItemsMap[handler.tipService.GetItemsGroupKey(item)], item)
	}

	result := CartCommandHandlerResult{
		TotalPrice: order.Get().GetDraftItemsTotalPriceByCustomerUUID(command.CustomerUUID),
		Items:      make([]CartItemDTO, 0, len(groupedItemsMap)),
	}

	waitGroup.Wait()

	for key := range groupedItemsMap {
		item := groupedItemsMap[key][0]

		menuDish := menuDishMap[item.Get().GetDishUUID()]

		menuDish.Calories = item.Get().GetCalories()
		menuDish.Category = item.Get().GetCategory()
		menuDish.Description = item.Get().GetDescription()
		menuDish.Name = item.Get().GetName()
		menuDish.Price = item.Get().GetPrice().String()
		menuDish.Weight = item.Get().GetWeight()

		groupedItems := groupedItemsMap[key]

		pictureReader, err := menuDish.Picture.Reader()
		if err != nil {
			return CartCommandHandlerResult{}, err
		}

		price := decimal.RequireFromString(menuDish.Price)

		result.Items = append(result.Items, CartItemDTO{
			ID:          menuDish.ID,
			Count:       len(groupedItems),
			Name:        menuDish.Name,
			Price:       price,
			Picture:     pictureReader,
			ResultPrice: price.Mul(decimal.NewFromInt(int64(len(groupedItems)))),
			Comment:     item.Get().GetComment(),
			PictureKey:  menuDish.PictureKey,
		})
	}

	return result, nil
}
