package app

import (
	"github.com/google/uuid"
)

type CatalogCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
}
