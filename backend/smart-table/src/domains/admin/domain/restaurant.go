package domain

import (
	"time"

	"github.com/google/uuid"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/utils"
)

type Restaurant struct {
	uuid      uuid.UUID
	ownerUUID uuid.UUID
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func NewRestaurant(
	ownerUUID uuid.UUID,
	name string,
	uuidGenerator *domainServices.UUIDGenerator,
) utils.SharedRef[Restaurant] {
	restaurant := Restaurant{}

	restaurant.ownerUUID = ownerUUID
	restaurant.name = name
	restaurant.createdAt = time.Now()
	restaurant.updatedAt = time.Now()

	shardID := uuidGenerator.GetShardID()
	restaurant.uuid = uuidGenerator.GenerateShardedUUID(shardID)

	restaurantRef, _ := utils.NewSharedRef(&restaurant)

	return restaurantRef
}

func RestoreRestaurant(
	id uuid.UUID,
	ownerUUID uuid.UUID,
	name string,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Restaurant] {
	restaurant := Restaurant{
		uuid:      id,
		ownerUUID: ownerUUID,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	restaurantRef, _ := utils.NewSharedRef(&restaurant)

	return restaurantRef
}

func (r *Restaurant) GetUUID() uuid.UUID      { return r.uuid }
func (r *Restaurant) GetOwnerUUID() uuid.UUID { return r.ownerUUID }
func (r *Restaurant) GetName() string         { return r.name }
func (r *Restaurant) GetCreatedAt() time.Time { return r.createdAt }
func (r *Restaurant) GetUpdatedAt() time.Time { return r.updatedAt }
