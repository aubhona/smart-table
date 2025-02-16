package mapper

import (
	"encoding/json"
	defs_internal_item "github.com/es-debug/backend-academy-2024-go-template/src/codegen/intern/item"
	defs_internal_order "github.com/es-debug/backend-academy-2024-go-template/src/codegen/intern/order"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/thoas/go-funk"

	defs_internal_customer_db "github.com/es-debug/backend-academy-2024-go-template/src/codegen/intern/customer_db"
	defs_internal_item_db "github.com/es-debug/backend-academy-2024-go-template/src/codegen/intern/item_db"
	defs_internal_order_db "github.com/es-debug/backend-academy-2024-go-template/src/codegen/intern/order_db"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type PgOrderAggregate struct {
	PgOrder     defs_internal_order_db.PgOrder         `json:"order"`
	PgItems     []defs_internal_item_db.PgItem         `json:"items"`
	PgCustomers []defs_internal_customer_db.PgCustomer `json:"customers"`
}

func ConvertToPgOrder(order utils.SharedRef[domain.Order]) ([]byte, error) {
	customersUUIDs := make([]openapi_types.UUID, len(order.Get().GetCustomers()))
	for i, customerRef := range order.Get().GetCustomers() {
		customersUUIDs[i] = customerRef.Get().GetUUID()
	}

	var resolution *string
	if order.Get().GetResolution().HasValue() {
		res := string(order.Get().GetResolution().Value())
		resolution = &res
	}

	pgOrder := defs_internal_order_db.PgOrder{
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

	pgItem := defs_internal_item_db.PgItem{
		UUID:         item.Get().GetUUID(),
		OrderUUID:    item.Get().GetOrderUUID(),
		DishUUID:     item.Get().GetDishUUID(),
		CustomerUUID: customer.Get().GetUUID(),
		Name:         item.Get().GetName(),
		Description:  item.Get().GetDescription(),
		PictureLink:  item.Get().GetPictureLink(),
		Weight:       item.Get().GetWeight(),
		Category:     item.Get().GetCategory(),
		Price:        float32(item.Get().GetPrice().InexactFloat64()),
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
	pgCustomer := defs_internal_customer_db.PgCustomer{
		UUID:       customer.Get().GetUUID(),
		TgID:       customer.Get().GetTgId(),
		TgLogin:    customer.Get().GetTgLogin(),
		AvatarLink: customer.Get().GetAvatarLink(),
		ChatID:     customer.Get().GetChatId(),
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
	pgItems := make([]defs_internal_item_db.PgItem, 0, len(items))

	for _, item := range items {
		pgItemBytes, err := ConvertToPgItem(item)
		if err != nil {
			return nil, err
		}

		var pgItem defs_internal_item_db.PgItem
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
	pgCustomer := defs_internal_customer_db.PgCustomer{}
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

func ConvertPgOrderAggregateToModel(pgResult []byte) (utils.SharedRef[domain.Order], error) {
	pgOrderAggregate := PgOrderAggregate{}
	err := json.Unmarshal(pgResult, &pgOrderAggregate)
	if err != nil {
		return utils.SharedRef[domain.Order]{}, err
	}

	uuidToCustomer := make(map[uuid.UUID]utils.SharedRef[domain.Customer], len(pgOrderAggregate.PgCustomers))
	for _, pgCustomer := range pgOrderAggregate.PgCustomers {
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
	for i, pgItem := range pgOrderAggregate.PgItems {
		customerRef := uuidToCustomer[pgItem.CustomerUUID]

		comment := utils.EmptyOptional[string]()
		if pgItem.Comment != nil {
			comment = utils.NewOptional(*pgItem.Comment)
		}

		resolution := utils.EmptyOptional[defs_internal_item.ItemResolution]()
		if pgItem.Resolution != nil {
			resolution = utils.NewOptional(defs_internal_item.ItemResolution(*pgItem.Resolution))
		}

		items[i] = domain.RestoreItem(
			pgItem.UUID,
			pgItem.OrderUUID,
			pgItem.DishUUID,
			customerRef,
			comment,
			defs_internal_item.ItemStatus(pgItem.Status),
			resolution,
			pgItem.Name,
			pgItem.Description,
			pgItem.PictureLink,
			pgItem.Weight,
			pgItem.Category,
			decimal.NewFromFloat(float64(pgItem.Price)),
			pgItem.IsDraft,
			pgItem.CreatedAt,
			pgItem.UpdatedAt,
		)
	}

	resolution := utils.EmptyOptional[defs_internal_order.OrderResolution]()
	if pgOrderAggregate.PgOrder.Resolution != nil {
		resolution = utils.NewOptional(defs_internal_order.OrderResolution(*pgOrderAggregate.PgOrder.Resolution))
	}

	order := domain.RestoreOrder(
		pgOrderAggregate.PgOrder.UUID,
		pgOrderAggregate.PgOrder.RoomCode,
		pgOrderAggregate.PgOrder.TableID,
		pgOrderAggregate.PgOrder.HostUserUUID,
		defs_internal_order.OrderStatus(pgOrderAggregate.PgOrder.Status),
		resolution,
		funk.Values(uuidToCustomer).([]utils.SharedRef[domain.Customer]),
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
