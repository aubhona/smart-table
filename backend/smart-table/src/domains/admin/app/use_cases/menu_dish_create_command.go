package app

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type MenuDishCreateCommand struct {
	UserUUID  uuid.UUID
	PlaceUUID uuid.UUID
	DishUUID  uuid.UUID
	Price     decimal.Decimal
}
