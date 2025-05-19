package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type OrderNotBelongToPLace struct {
	OrderUUID uuid.UUID
	PlaceUUID uuid.UUID
}

func (o OrderNotBelongToPLace) Error() string {
	return fmt.Sprintf("orde with uuid=%s not belong to place with uuid=%s", o.OrderUUID, o.PlaceUUID)
}
