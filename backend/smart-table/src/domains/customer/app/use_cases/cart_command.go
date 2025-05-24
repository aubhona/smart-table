package app

import "github.com/google/uuid"

type CartCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
	NeedPicture  bool
}
