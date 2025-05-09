package domain

import (
	"time"

	domainServices "github.com/smart-table/src/domains/admin/domain/services"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/smart-table/src/utils"
)

type MenuDish struct {
	uuid      uuid.UUID
	placeUUID uuid.UUID
	dish      utils.SharedRef[Dish]
	price     decimal.Decimal
	exist     bool
	createdAt time.Time
	updatedAt time.Time
}

func NewMenuDish(
	placeUUID uuid.UUID,
	dish utils.SharedRef[Dish],
	price decimal.Decimal,
	exist bool,
	uuidGenerator *domainServices.UUIDGenerator,
) utils.SharedRef[MenuDish] {
	menuDish := MenuDish{
		placeUUID: placeUUID,
		dish:      dish,
		price:     price,
		exist:     exist,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	shardID := uuidGenerator.GetShardID()
	menuDish.uuid = uuidGenerator.GenerateShardedUUID(shardID)

	menuDishRef, _ := utils.NewSharedRef(&menuDish)

	return menuDishRef
}

func RestoreMenuDish(
	id uuid.UUID,
	placeUUID uuid.UUID,
	dish utils.SharedRef[Dish],
	price decimal.Decimal,
	exist bool,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[MenuDish] {
	menuDish := MenuDish{
		uuid:      id,
		placeUUID: placeUUID,
		dish:      dish,
		price:     price,
		exist:     exist,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	menuDishRef, _ := utils.NewSharedRef(&menuDish)

	return menuDishRef
}

func (m *MenuDish) GetUUID() uuid.UUID             { return m.uuid }
func (m *MenuDish) GetPlaceUUID() uuid.UUID        { return m.placeUUID }
func (m *MenuDish) GetDish() utils.SharedRef[Dish] { return m.dish }
func (m *MenuDish) GetPrice() decimal.Decimal      { return m.price }
func (m *MenuDish) GetExist() bool                 { return m.exist }
func (m *MenuDish) GetCreatedAt() time.Time        { return m.createdAt }
func (m *MenuDish) GetUpdatedAt() time.Time        { return m.updatedAt }
