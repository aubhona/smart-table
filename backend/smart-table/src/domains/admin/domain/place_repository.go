package domain

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type PlaceRepository interface {
	Begin() (Transaction, error)
	Commit(tx Transaction) error
	Rollback(tx Transaction) error

	Save(tx Transaction, place utils.SharedRef[Place]) error
	Update(tx Transaction, place utils.SharedRef[Place]) error

	FindPlace(uuid uuid.UUID) (utils.SharedRef[Place], error)
	FindPlaces(uuids []uuid.UUID) ([]utils.SharedRef[Place], error)

	FindPlacesByRestaurantUUID(restaurantUUID uuid.UUID) ([]utils.SharedRef[Place], error)
	FindPlacesByEmployeeUserUUID(userUUID uuid.UUID) ([]utils.SharedRef[Place], error)
	FindPlaceByAddress(address string, restaurantUUID uuid.UUID) (utils.SharedRef[Place], error)
}
