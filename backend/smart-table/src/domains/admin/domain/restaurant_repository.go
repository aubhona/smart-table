package domain

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type RestaurantRepository interface {
	Begin() (Transaction, error)
	Commit(x Transaction) error
	Rollback(x Transaction) error

	Save(tx Transaction, restaurant utils.SharedRef[Restaurant]) error
	Update(tx Transaction, restaurant utils.SharedRef[Restaurant]) error

	FindRestaurant(uuid uuid.UUID) (utils.SharedRef[Restaurant], error)
	FindRestaurants(uuids []uuid.UUID) ([]utils.SharedRef[Restaurant], error)

	FindRestaurantForUpdate(tx Transaction, uuid uuid.UUID) (utils.SharedRef[Restaurant], error)
	FindRestaurantsForUpdate(tx Transaction, uuids []uuid.UUID) ([]utils.SharedRef[Restaurant], error)

	FindRestaurantByName(name string) (utils.SharedRef[Restaurant], error)
	FindRestaurantsByOwnerUUID(ownerUUID uuid.UUID) ([]utils.SharedRef[Restaurant], error)
}
