package app

import "github.com/google/uuid"

type CatalogUpdateInfoCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
}
