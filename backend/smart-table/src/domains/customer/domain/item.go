package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	defsInternalItem "github.com/smart-table/src/codegen/intern/item"
	domainErrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/utils"
)

var validItemStatuses = map[defsInternalItem.ItemStatus]interface{}{
	defsInternalItem.ItemStatusAccepted:          nil,
	defsInternalItem.ItemStatusCanceledByClient:  nil,
	defsInternalItem.ItemStatusCanceledByService: nil,
	defsInternalItem.ItemStatusCooked:            nil,
	defsInternalItem.ItemStatusCooking:           nil,
	defsInternalItem.ItemStatusNew:               nil,
	defsInternalItem.ItemStatusPaid:              nil,
	defsInternalItem.ItemStatusPaymentWaiting:    nil,
	defsInternalItem.ItemStatusServed:            nil,
}

type Item struct {
	uuid        uuid.UUID
	orderUUID   uuid.UUID
	dishUUID    uuid.UUID
	customer    utils.SharedRef[Customer]
	comment     utils.Optional[string]
	status      defsInternalItem.ItemStatus
	resolution  utils.Optional[defsInternalItem.ItemResolution]
	name        string
	description string
	pictureLink string
	weight      int
	calories    int
	category    string
	price       decimal.Decimal
	createdAt   time.Time
	updatedAt   time.Time
	isDraft     bool
}

func NewItem(
	id uuid.UUID,
	orderUUID uuid.UUID,
	dishUUID uuid.UUID,
	customer utils.SharedRef[Customer],
	comment utils.Optional[string],
	name string,
	description string,
	pictureLink string,
	weight int,
	calories int,
	category string,
	price decimal.Decimal,
	isDraft bool,
) utils.SharedRef[Item] {
	item := Item{
		uuid:        id,
		orderUUID:   orderUUID,
		dishUUID:    dishUUID,
		customer:    customer,
		comment:     comment,
		status:      defsInternalItem.ItemStatusNew,
		resolution:  utils.EmptyOptional[defsInternalItem.ItemResolution](),
		name:        name,
		description: description,
		pictureLink: pictureLink,
		weight:      weight,
		calories:    calories,
		category:    category,
		price:       price,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		isDraft:     isDraft,
	}

	itemRef, _ := utils.NewSharedRef(&item)

	return itemRef
}

func RestoreItem(
	id uuid.UUID,
	orderUUID uuid.UUID,
	dishUUID uuid.UUID,
	customer utils.SharedRef[Customer],
	comment utils.Optional[string],
	status defsInternalItem.ItemStatus,
	resolution utils.Optional[defsInternalItem.ItemResolution],
	name string,
	description string,
	pictureLink string,
	weight int,
	calories int,
	category string,
	price decimal.Decimal,
	isDraft bool,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Item] {
	item := Item{
		uuid:        id,
		orderUUID:   orderUUID,
		dishUUID:    dishUUID,
		customer:    customer,
		comment:     comment,
		status:      status,
		resolution:  resolution,
		name:        name,
		description: description,
		pictureLink: pictureLink,
		weight:      weight,
		calories:    calories,
		category:    category,
		price:       price,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		isDraft:     isDraft,
	}

	itemRef, _ := utils.NewSharedRef(&item)

	return itemRef
}

func (i *Item) GetUUID() uuid.UUID                                             { return i.uuid }
func (i *Item) GetOrderUUID() uuid.UUID                                        { return i.orderUUID }
func (i *Item) GetDishUUID() uuid.UUID                                         { return i.dishUUID }
func (i *Item) GetCustomer() utils.SharedRef[Customer]                         { return i.customer }
func (i *Item) GetComment() utils.Optional[string]                             { return i.comment }
func (i *Item) GetStatus() defsInternalItem.ItemStatus                         { return i.status }
func (i *Item) GetResolution() utils.Optional[defsInternalItem.ItemResolution] { return i.resolution }
func (i *Item) GetName() string                                                { return i.name }
func (i *Item) GetDescription() string                                         { return i.description }
func (i *Item) GetPictureLink() string                                         { return i.pictureLink }
func (i *Item) GetWeight() int                                                 { return i.weight }
func (i *Item) GetCategory() string                                            { return i.category }
func (i *Item) GetPrice() decimal.Decimal                                      { return i.price }
func (i *Item) GetCreatedAt() time.Time                                        { return i.createdAt }
func (i *Item) GetUpdatedAt() time.Time                                        { return i.updatedAt }
func (i *Item) GetIsDraft() bool                                               { return i.isDraft }
func (i *Item) GetCalories() int                                               { return i.calories }

func (i *Item) SetStatus(status defsInternalItem.ItemStatus) error {
	if i.GetIsDraft() && status != defsInternalItem.ItemStatusNew {
		return domainErrors.DraftItemStatusChangeNotAllowed{
			ItemUUID:   i.GetUUID(),
			ItemStatus: status,
		}
	}

	i.status = status

	return nil
}

func IsValidItemStatus(status defsInternalItem.ItemStatus) bool {
	_, exists := validItemStatuses[status]
	return exists
}

func ParseItemStatus(raw string) (defsInternalItem.ItemStatus, error) {
	status := defsInternalItem.ItemStatus(raw)
	if !IsValidItemStatus(status) {
		return "", domainErrors.InvalidItemStatus{ItemStatus: status}
	}

	return status, nil
}

func (i *Item) Commit() {
	i.isDraft = false
}
