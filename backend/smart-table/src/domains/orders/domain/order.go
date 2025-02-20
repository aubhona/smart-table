package domain

import (
	"errors"
	"fmt"
	"time"

	defs_internal_order "github.com/es-debug/backend-academy-2024-go-template/src/codegen/intern/order"
	domain "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain/services"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/thoas/go-funk"
)

type Order struct {
	uuid         uuid.UUID
	roomCode     string
	tableId      string
	hostUserUuid uuid.UUID
	status       defs_internal_order.OrderStatus
	resolution   utils.Optional[defs_internal_order.OrderResolution]
	customers    []utils.SharedRef[Customer]
	items        []utils.SharedRef[Item]
	createdAt    time.Time
	updatedAt    time.Time
}

func NewOrder(roomCode string, tableId string, hostUser utils.SharedRef[Customer], uuidGenerator *domain.UUIDGenerator) utils.SharedRef[Order] {
	order := Order{}

	order.roomCode = roomCode
	order.tableId = tableId
	order.hostUserUuid = hostUser.Get().uuid
	order.status = defs_internal_order.OrderStatusNew
	order.resolution = utils.EmptyOptional[defs_internal_order.OrderResolution]()
	order.customers = make([]utils.SharedRef[Customer], 1)
	order.customers[0] = hostUser
	order.items = make([]utils.SharedRef[Item], 0)
	order.createdAt = time.Now()
	order.updatedAt = time.Now()

	shardId := uuidGenerator.GetShardID()
	orderUuid := uuidGenerator.GenerateShardedUUID(shardId)

	order.uuid = orderUuid

	orderRef, _ := utils.NewSharedRef(&order)

	return orderRef
}

func RestoreOrder(
	uuid uuid.UUID,
	roomCode string,
	tableId string,
	hostUserUuid uuid.UUID,
	status defs_internal_order.OrderStatus,
	resolution utils.Optional[defs_internal_order.OrderResolution],
	customers []utils.SharedRef[Customer],
	items []utils.SharedRef[Item],
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Order] {
	order := Order{
		uuid:         uuid,
		roomCode:     roomCode,
		tableId:      tableId,
		hostUserUuid: hostUserUuid,
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

func (o *Order) GetUUID() uuid.UUID                         { return o.uuid }
func (o *Order) GetRoomCode() string                        { return o.roomCode }
func (o *Order) GetTableID() string                         { return o.tableId }
func (o *Order) GetHostUserUUID() uuid.UUID                 { return o.hostUserUuid }
func (o *Order) GetStatus() defs_internal_order.OrderStatus { return o.status }
func (o *Order) GetResolution() utils.Optional[defs_internal_order.OrderResolution] {
	return o.resolution
}
func (o *Order) GetCustomers() []utils.SharedRef[Customer] { return o.customers }
func (o *Order) GetItems() []utils.SharedRef[Item]         { return o.items }
func (o *Order) GetCreatedAt() time.Time                   { return o.createdAt }
func (o *Order) GetUpdatedAt() time.Time                   { return o.updatedAt }

func (o *Order) DraftItem(
	dishUuid uuid.UUID,
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
		dishUuid,
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

func (o *Order) CommitItem(itemUuid uuid.UUID) (utils.Optional[utils.SharedRef[Item]], error) {
	item, found := funk.Find(o.items, func(item utils.SharedRef[Item]) bool {
		return item.Get().uuid == itemUuid
	}).(utils.SharedRef[Item])

	if !found {
		return utils.EmptyOptional[utils.SharedRef[Item]](), errors.New(fmt.Sprintf("item not found, item_uuid=%v", itemUuid))
	}

	item.Get().Commit()

	return utils.NewOptional(item), nil
}
