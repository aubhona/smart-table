package domain

import (
	"fmt"
	"time"

	"github.com/samber/lo"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	defsInternalOrder "github.com/smart-table/src/codegen/intern/order"
	domain "github.com/smart-table/src/domains/customer/domain/services"
	"github.com/smart-table/src/utils"
	"golang.org/x/exp/slices"
)

type Order struct {
	uuid         uuid.UUID
	roomCode     string
	tableID      string
	hostUserUUID uuid.UUID
	status       defsInternalOrder.OrderStatus
	resolution   utils.Optional[defsInternalOrder.OrderResolution]
	customers    []utils.SharedRef[Customer]
	items        []utils.SharedRef[Item]
	createdAt    time.Time
	updatedAt    time.Time
}

func NewOrder(
	roomCode, tableID string,
	hostUser utils.SharedRef[Customer],
	uuidGenerator *domain.UUIDGenerator,
) utils.SharedRef[Order] {
	order := Order{}

	order.roomCode = roomCode
	order.tableID = tableID
	order.hostUserUUID = hostUser.Get().uuid
	order.status = defsInternalOrder.OrderStatusNew
	order.resolution = utils.EmptyOptional[defsInternalOrder.OrderResolution]()
	order.customers = make([]utils.SharedRef[Customer], 1)
	order.customers[0] = hostUser
	order.items = make([]utils.SharedRef[Item], 0)
	order.createdAt = time.Now()
	order.updatedAt = time.Now()

	shardID := uuidGenerator.GetShardID()
	orderUUID := uuidGenerator.GenerateShardedUUID(shardID)

	order.uuid = orderUUID

	orderRef, _ := utils.NewSharedRef(&order)

	return orderRef
}

func RestoreOrder(
	id uuid.UUID,
	roomCode string,
	tableID string,
	hostUserUUID uuid.UUID,
	status defsInternalOrder.OrderStatus,
	resolution utils.Optional[defsInternalOrder.OrderResolution],
	customers []utils.SharedRef[Customer],
	items []utils.SharedRef[Item],
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Order] {
	order := Order{
		uuid:         id,
		roomCode:     roomCode,
		tableID:      tableID,
		hostUserUUID: hostUserUUID,
		status:       status,
		resolution:   resolution,
		customers:    customers,
		items:        items,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}

	orderRef, _ := utils.NewSharedRef(&order)

	return orderRef
}

func (o *Order) AddCustomer(customer utils.SharedRef[Customer]) {
	o.customers = append(o.customers, customer)
}

func (o *Order) ContainsCustomer(customerUUID uuid.UUID) bool {
	return customerUUID == o.GetHostUserUUID() ||
		slices.ContainsFunc(o.GetCustomers(), func(customer utils.SharedRef[Customer]) bool {
			return customer.Get().GetUUID() == customerUUID
		})
}

func (o *Order) GetUUID() uuid.UUID                       { return o.uuid }
func (o *Order) GetRoomCode() string                      { return o.roomCode }
func (o *Order) GetTableID() string                       { return o.tableID }
func (o *Order) GetHostUserUUID() uuid.UUID               { return o.hostUserUUID }
func (o *Order) GetStatus() defsInternalOrder.OrderStatus { return o.status }
func (o *Order) GetResolution() utils.Optional[defsInternalOrder.OrderResolution] {
	return o.resolution
}
func (o *Order) GetCustomers() []utils.SharedRef[Customer] { return o.customers }
func (o *Order) GetItems() []utils.SharedRef[Item]         { return o.items }
func (o *Order) GetCreatedAt() time.Time                   { return o.createdAt }
func (o *Order) GetUpdatedAt() time.Time                   { return o.updatedAt }

func (o *Order) DraftItem(
	dishUUID uuid.UUID,
	customer utils.SharedRef[Customer],
	comment utils.Optional[string],
	name string,
	description string,
	pictureLink string,
	weight int,
	category string,
	price decimal.Decimal,
	uuidGenerator domain.UUIDGenerator,
) utils.SharedRef[Item] {
	itemRef := NewItem(
		uuidGenerator.GenerateShardedUUID(uuidGenerator.GetShardID()),
		o.uuid,
		dishUUID,
		customer,
		comment,
		name,
		description,
		pictureLink,
		weight,
		category,
		price,
		true,
	)

	o.items = append(o.items, itemRef)

	return itemRef
}

func (o *Order) CommitItem(itemUUID uuid.UUID) (utils.Optional[utils.SharedRef[Item]], error) {
	item, found := lo.Find(o.items, func(item utils.SharedRef[Item]) bool {
		return item.Get().uuid == itemUUID
	})

	if !found {
		return utils.EmptyOptional[utils.SharedRef[Item]](), fmt.Errorf("item not found, item_uuid=%v", itemUUID)
	}

	item.Get().Commit()

	return utils.NewOptional(item), nil
}
