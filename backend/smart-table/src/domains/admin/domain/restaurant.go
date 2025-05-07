package domain

import (
	"time"

	"github.com/google/uuid"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/utils"
)

type Restaurant struct {
	uuid      uuid.UUID
	owner     utils.SharedRef[User]
	dishes    []utils.SharedRef[Dish]
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func NewRestaurant(
	owner utils.SharedRef[User],
	name string,
	uuidGenerator *domainServices.UUIDGenerator,
) utils.SharedRef[Restaurant] {
	restaurant := Restaurant{
		owner:     owner,
		name:      name,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		dishes:    make([]utils.SharedRef[Dish], 0),
	}

	shardID := uuidGenerator.GetShardID()
	restaurant.uuid = uuidGenerator.GenerateShardedUUID(shardID)

	restaurantRef, _ := utils.NewSharedRef(&restaurant)

	return restaurantRef
}

func RestoreRestaurant(
	id uuid.UUID,
	owner utils.SharedRef[User],
	dishes []utils.SharedRef[Dish],
	name string,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Restaurant] {
	restaurant := Restaurant{
		uuid:      id,
		owner:     owner,
		dishes:    dishes,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	restaurantRef, _ := utils.NewSharedRef(&restaurant)

	return restaurantRef
}

func (r *Restaurant) AddDish(
	dishName,
	description,
	category,
	pictureKey string,
	calories,
	weight int,
	uuidGenerator *domainServices.UUIDGenerator,
) utils.SharedRef[Dish] {
	dish := NewDish(r.uuid, dishName, description, category, pictureKey, calories, weight, uuidGenerator)
	r.dishes = append(r.dishes, dish)

	return dish
}

func (r *Restaurant) GetUUID() uuid.UUID              { return r.uuid }
func (r *Restaurant) GetOwner() utils.SharedRef[User] { return r.owner }
func (r *Restaurant) GetName() string                 { return r.name }
func (r *Restaurant) GetCreatedAt() time.Time         { return r.createdAt }
func (r *Restaurant) GetUpdatedAt() time.Time         { return r.updatedAt }
func (r *Restaurant) GetDishes() []utils.SharedRef[Dish] {
	return r.dishes
}
