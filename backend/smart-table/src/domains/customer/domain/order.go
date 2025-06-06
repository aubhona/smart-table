package domain

import (
	"time"

	defsInternalItem "github.com/smart-table/src/codegen/intern/item"

	domainErrors "github.com/smart-table/src/domains/customer/domain/errors"

	"github.com/samber/lo"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	defsInternalOrder "github.com/smart-table/src/codegen/intern/order"
	domain "github.com/smart-table/src/domains/customer/domain/services"
	"github.com/smart-table/src/utils"
	"golang.org/x/exp/slices"
)

var validOrderResolutions = map[defsInternalOrder.OrderResolution]interface{}{
	defsInternalOrder.OrderResolutionCanceledByClient:  nil,
	defsInternalOrder.OrderResolutionCanceledByService: nil,
	defsInternalOrder.OrderResolutionPaid:              nil,
}

var validOrderStatuses = map[defsInternalOrder.OrderStatus]interface{}{
	defsInternalOrder.OrderStatusCanceledByClient:  nil,
	defsInternalOrder.OrderStatusCanceledByService: nil,
	defsInternalOrder.OrderStatusNew:               nil,
	defsInternalOrder.OrderStatusPaid:              nil,
	defsInternalOrder.OrderStatusPaymentWaiting:    nil,
	defsInternalOrder.OrderStatusServing:           nil,
}

var orderStatusesWhichCanResolutions = map[defsInternalOrder.OrderStatus]interface{}{
	defsInternalOrder.OrderStatusCanceledByClient:  nil,
	defsInternalOrder.OrderStatusCanceledByService: nil,
	defsInternalOrder.OrderStatusPaid:              nil,
}

var orderToItemStatusMap = map[defsInternalOrder.OrderStatus]defsInternalItem.ItemStatus{
	defsInternalOrder.OrderStatusCanceledByClient:  defsInternalItem.ItemStatusCanceledByClient,
	defsInternalOrder.OrderStatusCanceledByService: defsInternalItem.ItemStatusCanceledByService,
	defsInternalOrder.OrderStatusNew:               defsInternalItem.ItemStatusNew,
	defsInternalOrder.OrderStatusPaid:              defsInternalItem.ItemStatusPaid,
	defsInternalOrder.OrderStatusPaymentWaiting:    defsInternalItem.ItemStatusPaymentWaiting,
	defsInternalOrder.OrderStatusServing:           defsInternalItem.ItemStatusAccepted,
}

var ItemStatusesCanChangeOnlyWithOrderStatus = map[defsInternalItem.ItemStatus]interface{}{
	defsInternalItem.ItemStatusPaymentWaiting: nil,
	defsInternalItem.ItemStatusPaid:           nil,
}

type Order struct {
	uuid         uuid.UUID
	roomCode     string
	tableID      string
	hostUserUUID uuid.UUID
	status       defsInternalOrder.OrderStatus
	resolution   utils.Optional[defsInternalOrder.OrderResolution]
	customers    []utils.SharedRef[Customer]
	items        []utils.SharedRef[Item]
	deletedItems []utils.SharedRef[Item]
	createdAt    time.Time
	updatedAt    time.Time
}

func NewOrder(
	roomCode, tableID string,
	hostUser utils.SharedRef[Customer],
	uuidGenerator *domain.UUIDGenerator,
) utils.SharedRef[Order] {
	shardID := uuidGenerator.GetShardID()
	orderUUID := uuidGenerator.GenerateShardedUUID(shardID)

	order := Order{
		uuid:         orderUUID,
		roomCode:     roomCode,
		tableID:      tableID,
		hostUserUUID: hostUser.Get().uuid,
		status:       defsInternalOrder.OrderStatusNew,
		resolution:   utils.EmptyOptional[defsInternalOrder.OrderResolution](),
		customers:    []utils.SharedRef[Customer]{hostUser},
		items:        make([]utils.SharedRef[Item], 0),
		deletedItems: make([]utils.SharedRef[Item], 0),
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}

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
		deletedItems: make([]utils.SharedRef[Item], 0),
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
func (o *Order) GetItemsByCustomerUUID(customerUUID uuid.UUID) []utils.SharedRef[Item] {
	return lo.Filter(o.items, func(item utils.SharedRef[Item], _ int) bool {
		return item.Get().customer.Get().uuid == customerUUID
	})
}
func (o *Order) GetDraftedItemsByCustomerUUID(customerUUID uuid.UUID) []utils.SharedRef[Item] {
	return lo.Filter(o.items, func(item utils.SharedRef[Item], _ int) bool {
		return item.Get().customer.Get().uuid == customerUUID && item.Get().isDraft
	})
}
func (o *Order) GetCreatedAt() time.Time { return o.createdAt }
func (o *Order) GetUpdatedAt() time.Time { return o.updatedAt }

func (o *Order) DraftItem(
	dishUUID uuid.UUID,
	customer utils.SharedRef[Customer],
	comment utils.Optional[string],
	name,
	description,
	pictureLink string,
	weight,
	calories int,
	category string,
	price decimal.Decimal,
	uuidGenerator *domain.UUIDGenerator,
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
		calories,
		category,
		price,
		true,
	)

	o.items = append(o.items, itemRef)

	return itemRef
}

func (o *Order) CommitItems(customerUUID uuid.UUID) {
	for _, item := range o.items {
		if item.Get().GetIsDraft() && item.Get().GetCustomer().Get().GetUUID() == customerUUID {
			item.Get().Commit()
		}
	}
}

func (o *Order) setResolution(resolution defsInternalOrder.OrderResolution) {
	o.resolution = utils.NewOptional(resolution)
}

func IsValidOrderResolution(resolution defsInternalOrder.OrderResolution) bool {
	_, exists := validOrderResolutions[resolution]
	return exists
}

func ParseOrderResolution(raw string) (defsInternalOrder.OrderResolution, error) {
	resolution := defsInternalOrder.OrderResolution(raw)
	if !IsValidOrderResolution(resolution) {
		return "", domainErrors.InvalidOrderResolution{OrderResolution: resolution}
	}

	return resolution, nil
}

func (o *Order) SetStatus(status defsInternalOrder.OrderStatus) error {
	parsedItemStatus, exist := orderToItemStatusMap[status]
	if !exist {
		return domainErrors.InvalidOrderStatus{OrderStatus: status}
	}

	for _, item := range o.items {
		if item.Get().GetIsDraft() {
			continue
		}

		err := item.Get().SetStatus(parsedItemStatus)
		if err != nil {
			return err
		}
	}

	o.status = status

	_, exists := orderStatusesWhichCanResolutions[status]
	if exists {
		parsedOrderResolution, err := ParseOrderResolution(string(status))
		if err != nil {
			return err
		}

		o.setResolution(parsedOrderResolution)
	}

	return nil
}

func IsValidOrderStatus(status defsInternalOrder.OrderStatus) bool {
	_, exists := validOrderStatuses[status]
	return exists
}

func ParseOrderStatus(raw string) (defsInternalOrder.OrderStatus, error) {
	status := defsInternalOrder.OrderStatus(raw)
	if !IsValidOrderStatus(status) {
		return "", domainErrors.InvalidOrderStatus{OrderStatus: status}
	}

	return status, nil
}

func (o *Order) ChangeItemsStatus(itemUUIDList []uuid.UUID, status string) error {
	parsedItemStatus, err := ParseItemStatus(status)
	if err != nil {
		return err
	}

	_, exists := ItemStatusesCanChangeOnlyWithOrderStatus[parsedItemStatus]
	if exists {
		return domainErrors.ItemStatusChangeRequiresOrderStatusUpdate{
			ItemStatus: parsedItemStatus,
		}
	}

	itemUUIDSet := make(map[uuid.UUID]interface{})
	for _, itemUUID := range itemUUIDList {
		itemUUIDSet[itemUUID] = nil
	}

	parsedOrderStatus, err := ParseOrderStatus(status)
	needToChangeOrderStatus := err == nil

	for _, item := range o.items {
		_, exists := itemUUIDSet[item.Get().GetUUID()]
		if exists {
			err = item.Get().SetStatus(parsedItemStatus)
			if err != nil {
				return err
			}
		} else if item.Get().GetStatus() != parsedItemStatus {
			needToChangeOrderStatus = false
		}
	}

	if needToChangeOrderStatus {
		err := o.SetStatus(parsedOrderStatus)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Order) GetCustomerByUUID(uuid uuid.UUID) utils.Optional[utils.SharedRef[Customer]] {
	customer, found := lo.Find(o.customers, func(item utils.SharedRef[Customer]) bool {
		return item.Get().uuid == uuid
	})
	if !found {
		return utils.EmptyOptional[utils.SharedRef[Customer]]()
	}

	return utils.NewOptional(customer)
}

func (o *Order) GetItemByDishUUID(dishUUID uuid.UUID) utils.Optional[utils.SharedRef[Item]] {
	item, found := lo.Find(o.items, func(item utils.SharedRef[Item]) bool {
		return item.Get().dishUUID == dishUUID
	})
	if !found {
		return utils.EmptyOptional[utils.SharedRef[Item]]()
	}

	return utils.NewOptional(item)
}

func (o *Order) DeleteItemsByDishUUID(dishUUID uuid.UUID, count int) error {
	items := make([]utils.SharedRef[Item], 0, len(o.items))
	deletedItems := make([]utils.SharedRef[Item], 0, count)

	for _, item := range o.items {
		if !item.Get().isDraft || item.Get().dishUUID != dishUUID || len(deletedItems) == count {
			items = append(items, item)
			continue
		}

		deletedItems = append(deletedItems, item)
	}

	if len(deletedItems) != count {
		return domainErrors.IncorrectDeleteItemsCount{
			Count: count,
		}
	}

	o.items = items
	o.deletedItems = deletedItems

	return nil
}

func (o *Order) GetDraftItemsTotalPriceByCustomerUUID(customerUUID uuid.UUID) decimal.Decimal {
	result := decimal.Zero

	for _, item := range o.GetDraftedItemsByCustomerUUID(customerUUID) {
		result = result.Add(item.Get().price)
	}

	return result
}

func (o *Order) GetDeletesItems() []utils.SharedRef[Item] {
	return o.deletedItems
}

func (o *Order) MarkWaitingPayment() {
	o.status = defsInternalOrder.OrderStatusPaymentWaiting
	for _, item := range o.items {
		_ = item.Get().SetStatus(defsInternalItem.ItemStatusPaymentWaiting)
	}
}
