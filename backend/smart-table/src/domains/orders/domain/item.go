package domain

import (
	"time"

	defs_internal_item "github.com/es-debug/backend-academy-2024-go-template/src/codegen/intern/item"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Item struct {
	uuid        uuid.UUID
	orderUuid   uuid.UUID
	dishUuid    uuid.UUID
	customer    utils.SharedRef[Customer]
	comment     utils.Optional[string]
	status      defs_internal_item.ItemStatus
	resolution  utils.Optional[defs_internal_item.ItemResolution]
	name        string
	description string
	pictureLink string
	weight      int
	category    string
	price       decimal.Decimal
	createdAt   time.Time
	updatedAt   time.Time
	isDraft     bool
}

func NewItem(
	uuid uuid.UUID,
	orderUuid uuid.UUID,
	dishUuid uuid.UUID,
	customer utils.SharedRef[Customer],
	comment utils.Optional[string],
	name string,
	description string,
	pictureLink string,
	weight int,
	category string,
	price decimal.Decimal,
	isDraft bool,
) utils.SharedRef[Item] {
	item := Item{
		uuid:        uuid,
		orderUuid:   orderUuid,
		dishUuid:    dishUuid,
		customer:    customer,
		comment:     comment,
		status:      defs_internal_item.ItemStatusNew,
		resolution:  utils.EmptyOptional[defs_internal_item.ItemResolution](),
		name:        name,
		description: description,
		pictureLink: pictureLink,
		weight:      weight,
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
	uuid uuid.UUID,
	orderUuid uuid.UUID,
	dishUuid uuid.UUID,
	customer utils.SharedRef[Customer],
	comment utils.Optional[string],
	status defs_internal_item.ItemStatus,
	resolution utils.Optional[defs_internal_item.ItemResolution],
	name string,
	description string,
	pictureLink string,
	weight int,
	category string,
	price decimal.Decimal,
	isDraft bool,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Item] {
	item := Item{
		uuid:        uuid,
		orderUuid:   orderUuid,
		dishUuid:    dishUuid,
		customer:    customer,
		comment:     comment,
		status:      status,
		resolution:  resolution,
		name:        name,
		description: description,
		pictureLink: pictureLink,
		weight:      weight,
		category:    category,
		price:       price,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		isDraft:     isDraft,
	}

	itemRef, _ := utils.NewSharedRef(&item)
	return itemRef
}

func (i *Item) GetUUID() uuid.UUID                                               { return i.uuid }
func (i *Item) GetOrderUUID() uuid.UUID                                          { return i.orderUuid }
func (i *Item) GetDishUUID() uuid.UUID                                           { return i.dishUuid }
func (i *Item) GetCustomer() utils.SharedRef[Customer]                           { return i.customer }
func (i *Item) GetComment() utils.Optional[string]                               { return i.comment }
func (i *Item) GetStatus() defs_internal_item.ItemStatus                         { return i.status }
func (i *Item) GetResolution() utils.Optional[defs_internal_item.ItemResolution] { return i.resolution }
func (i *Item) GetName() string                                                  { return i.name }
func (i *Item) GetDescription() string                                           { return i.description }
func (i *Item) GetPictureLink() string                                           { return i.pictureLink }
func (i *Item) GetWeight() int                                                   { return i.weight }
func (i *Item) GetCategory() string                                              { return i.category }
func (i *Item) GetPrice() decimal.Decimal                                        { return i.price }
func (i *Item) GetCreatedAt() time.Time                                          { return i.createdAt }
func (i *Item) GetUpdatedAt() time.Time                                          { return i.updatedAt }
func (i *Item) GetIsDraft() bool                                                 { return i.isDraft }

func (i *Item) Commit() {
	i.isDraft = false
}
