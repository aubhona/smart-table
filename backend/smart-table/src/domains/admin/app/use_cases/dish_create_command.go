package app

import (
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type DishCreateCommand struct {
	RestaurantUUID uuid.UUID
	OwnerUUID      uuid.UUID
	DishName       string
	Description    string
	Calories       int
	Weight         int
	Category       string
	Image          openapi_types.File
}
