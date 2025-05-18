package app

import "github.com/google/uuid"

type ItemsCommitCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
}
