package app

import (
	"time"

	"github.com/google/uuid"
)

type PlaceCreateCommand struct {
	OwnerUUID      uuid.UUID
	RestaurantUUID uuid.UUID
	Address        string
	TableCount     int
	OpeningTime    time.Time
	ClosingTime    time.Time
}
