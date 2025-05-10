package mapper

import (
	"encoding/json"

	"github.com/samber/lo"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	defsInternalItem "github.com/smart-table/src/codegen/intern/item"
	defsInternalOrder "github.com/smart-table/src/codegen/intern/order"

	openapiTypes "github.com/oapi-codegen/runtime/types"
	defsInternalCustomerDb "github.com/smart-table/src/codegen/intern/customer_db"
	defsInternalItemDb "github.com/smart-table/src/codegen/intern/item_db"
	defsInternalOrderDb "github.com/smart-table/src/codegen/intern/order_db"
	"github.com/smart-table/src/domains/customer/domain"
	"github.com/smart-table/src/utils"
)

type PgOrderAggregate struct {
	PgOrder     defsInternalOrderDb.PgOrder         `json:"order"`
	PgItems     []defsInternalItemDb.PgItem         `json:"items"`
	PgCustomers []defsInternalCustomerDb.PgCustomer `json:"customers"`
}

func ConvertToPgOrder(order utils.SharedRef[domain.Order]) ([]byte, error) {
	customersUUIDs := make([]openapiTypes.UUID, len(order.Get().GetCustomers()))
	for i, customerRef := range order.Get().GetCustomers() {
		customersUUIDs[i] = customerRef.Get().GetUUID()
	}

	var resolution *string

	if order.Get().GetResolution().HasValue() {
		res := string(order.Get().GetResolution().Value())
		resolution = &res
	}

	pgOrder := defsInternalOrderDb.PgOrder{
		UUID:          order.Get().GetUUID(),
		RoomCode:      order.Get().GetRoomCode(),
		TableID:       order.Get().GetTableID(),
		HostUserUUID:  order.Get().GetHostUserUUID(),
		CustomersUUID: customersUUIDs,
		Status:        string(order.Get().GetStatus()),
		Resolution:    resolution,
		CreatedAt:     order.Get().GetCreatedAt(),
		UpdatedAt:     order.Get().GetUpdatedAt(),
	}

	jsonBytes, err := json.Marshal(pgOrder)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertToPgItem(item utils.SharedRef[domain.Item]) ([]byte, error) {
	var comment *string

	if item.Get().GetComment().HasValue() {
		cmt := item.Get().GetComment().Value()
		comment = &cmt
	}

	var resolution *string

	if item.Get().GetResolution().HasValue() {
		res := string(item.Get().GetResolution().Value())
		resolution = &res
	}

	customer := item.Get().GetCustomer()

	pgItem := defsInternalItemDb.PgItem{
		UUID:         item.Get().GetUUID(),
		OrderUUID:    item.Get().GetOrderUUID(),
		DishUUID:     item.Get().GetDishUUID(),
		CustomerUUID: customer.Get().GetUUID(),
		Name:         item.Get().GetName(),
		Description:  item.Get().GetDescription(),
		PictureLink:  item.Get().GetPictureLink(),
		Weight:       item.Get().GetWeight(),
		Category:     item.Get().GetCategory(),
		Price:        item.Get().GetPrice().String(),
		Comment:      comment,
		Status:       string(item.Get().GetStatus()),
		Resolution:   resolution,
		IsDraft:      item.Get().GetIsDraft(),
		CreatedAt:    item.Get().GetCreatedAt(),
		UpdatedAt:    item.Get().GetUpdatedAt(),
	}

	jsonBytes, err := json.Marshal(pgItem)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertToPgCustomer(customer utils.SharedRef[domain.Customer]) ([]byte, error) {
	pgCustomer := defsInternalCustomerDb.PgCustomer{
		UUID:       customer.Get().GetUUID(),
		TgID:       customer.Get().GetTgID(),
		TgLogin:    customer.Get().GetTgLogin(),
		AvatarLink: customer.Get().GetAvatarLink(),
		ChatID:     customer.Get().GetChatID(),
		CreatedAt:  customer.Get().GetCreatedAt(),
		UpdatedAt:  customer.Get().GetUpdatedAt(),
	}

	jsonBytes, err := json.Marshal(pgCustomer)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertToPgItems(items []utils.SharedRef[domain.Item]) ([]byte, error) {
	pgItems := make([]defsInternalItemDb.PgItem, 0, len(items))

	for _, item := range items {
		pgItemBytes, err := ConvertToPgItem(item)
		if err != nil {
			return nil, err
		}

		var pgItem defsInternalItemDb.PgItem
		if err := json.Unmarshal(pgItemBytes, &pgItem); err != nil {
			return nil, err
		}

		pgItems = append(pgItems, pgItem)
	}

	jsonBytes, err := json.Marshal(pgItems)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertPgCustomerToModel(pgResult []byte) (utils.SharedRef[domain.Customer], error) {
	pgCustomer := defsInternalCustomerDb.PgCustomer{}
	err := json.Unmarshal(pgResult, &pgCustomer)

	if err != nil {
		return utils.SharedRef[domain.Customer]{}, err
	}

	return domain.RestoreCustomer(
		pgCustomer.UUID,
		pgCustomer.TgID,
		pgCustomer.TgLogin,
		pgCustomer.AvatarLink,
		pgCustomer.ChatID,
		pgCustomer.CreatedAt,
		pgCustomer.UpdatedAt,
	), nil
}

func ConvertPgOrderAggregateToModel(pgResult []byte) (utils.SharedRef[domain.Order], error) { //nolint
	pgOrderAggregate := PgOrderAggregate{}
	err := json.Unmarshal(pgResult, &pgOrderAggregate)

	if err != nil {
		return utils.SharedRef[domain.Order]{}, err
	}

	uuidToCustomer := make(map[uuid.UUID]utils.SharedRef[domain.Customer], len(pgOrderAggregate.PgCustomers))

	for i := range pgOrderAggregate.PgCustomers {
		pgCustomer := &pgOrderAggregate.PgCustomers[i]

		uuidToCustomer[pgCustomer.UUID] = domain.RestoreCustomer(
			pgCustomer.UUID,
			pgCustomer.TgID,
			pgCustomer.TgLogin,
			pgCustomer.AvatarLink,
			pgCustomer.ChatID,
			pgCustomer.CreatedAt,
			pgCustomer.UpdatedAt,
		)
	}

	items := make([]utils.SharedRef[domain.Item], len(pgOrderAggregate.PgItems))

	for i := range pgOrderAggregate.PgItems {
		pgItem := &pgOrderAggregate.PgItems[i]
		customerRef := uuidToCustomer[pgItem.CustomerUUID]

		comment := utils.EmptyOptional[string]()
		if pgItem.Comment != nil {
			comment = utils.NewOptional(*pgItem.Comment)
		}

		resolution := utils.EmptyOptional[defsInternalItem.ItemResolution]()
		if pgItem.Resolution != nil {
			resolution = utils.NewOptional(defsInternalItem.ItemResolution(*pgItem.Resolution))
		}

		price, err := decimal.NewFromString(pgItem.Price)
		if err != nil {
			return utils.SharedRef[domain.Order]{}, err
		}

		items[i] = domain.RestoreItem(
			pgItem.UUID,
			pgItem.OrderUUID,
			pgItem.DishUUID,
			customerRef,
			comment,
			defsInternalItem.ItemStatus(pgItem.Status),
			resolution,
			pgItem.Name,
			pgItem.Description,
			pgItem.PictureLink,
			pgItem.Weight,
			pgItem.Calories,
			pgItem.Category,
			price,
			pgItem.IsDraft,
			pgItem.CreatedAt,
			pgItem.UpdatedAt,
		)
	}

	resolution := utils.EmptyOptional[defsInternalOrder.OrderResolution]()
	if pgOrderAggregate.PgOrder.Resolution != nil {
		resolution = utils.NewOptional(defsInternalOrder.OrderResolution(*pgOrderAggregate.PgOrder.Resolution))
	}

	order := domain.RestoreOrder(
		pgOrderAggregate.PgOrder.UUID,
		pgOrderAggregate.PgOrder.RoomCode,
		pgOrderAggregate.PgOrder.TableID,
		pgOrderAggregate.PgOrder.HostUserUUID,
		defsInternalOrder.OrderStatus(pgOrderAggregate.PgOrder.Status),
		resolution,
		lo.Values(uuidToCustomer),
		items,
		pgOrderAggregate.PgOrder.CreatedAt,
		pgOrderAggregate.PgOrder.UpdatedAt,
	)

	return order, nil
}

func ConvertPgOrderAggregatesToModels(pgResults [][]byte) ([]utils.SharedRef[domain.Order], error) {
	orders := make([]utils.SharedRef[domain.Order], 0, len(pgResults))

	for _, pgResult := range pgResults {
		orderRef, err := ConvertPgOrderAggregateToModel(pgResult)
		if err != nil {
			return nil, err
		}

		orders = append(orders, orderRef)
	}

	return orders, nil
}
