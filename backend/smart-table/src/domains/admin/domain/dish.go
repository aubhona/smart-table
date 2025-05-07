package domain

import (
	"time"

	domainServices "github.com/smart-table/src/domains/admin/domain/services"

	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type Dish struct {
	uuid           uuid.UUID
	restaurantUUID uuid.UUID
	name           string
	description    string
	calories       int
	weight         int
	pictureKey     string
	category       string
	createdAt      time.Time
	updatedAt      time.Time
}

func NewDish(
	restaurantUUID uuid.UUID,
	name,
	description,
	category,
	pictureKey string,
	calories,
	weight int,
	uuidGenerator *domainServices.UUIDGenerator,
) utils.SharedRef[Dish] {
	dish := Dish{
		restaurantUUID: restaurantUUID,
		name:           name,
		description:    description,
		calories:       calories,
		weight:         weight,
		pictureKey:     pictureKey,
		category:       category,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
	}

	shardID := uuidGenerator.GetShardID()
	dish.uuid = uuidGenerator.GenerateShardedUUID(shardID)

	dishRef, _ := utils.NewSharedRef(&dish)

	return dishRef
}

func RestoreDish(
	id,
	restaurantUUID uuid.UUID,
	name,
	description,
	pictureKey,
	category string,
	calories,
	weight int,
	createdAt,
	updatedAt time.Time,
) utils.SharedRef[Dish] {
	dish := Dish{
		uuid:           id,
		restaurantUUID: restaurantUUID,
		name:           name,
		description:    description,
		calories:       calories,
		weight:         weight,
		pictureKey:     pictureKey,
		category:       category,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}

	dishRef, _ := utils.NewSharedRef(&dish)

	return dishRef
}

func (d *Dish) GetUUID() uuid.UUID           { return d.uuid }
func (d *Dish) GetRestaurantUUID() uuid.UUID { return d.restaurantUUID }
func (d *Dish) GetName() string              { return d.name }
func (d *Dish) GetDescription() string       { return d.description }
func (d *Dish) GetCalories() int             { return d.calories }
func (d *Dish) GetWeight() int               { return d.weight }
func (d *Dish) GetPictureKey() string        { return d.pictureKey }
func (d *Dish) GetCategory() string          { return d.category }
func (d *Dish) GetCreatedAt() time.Time      { return d.createdAt }
func (d *Dish) GetUpdatedAt() time.Time      { return d.updatedAt }
