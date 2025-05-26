package app

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type ItemStateCommand struct {
	CustomerUUID uuid.UUID
	OrderUUID    uuid.UUID
	DishUUD      uuid.UUID
	Comment      utils.Optional[string]
	NeedPicture  bool
}
