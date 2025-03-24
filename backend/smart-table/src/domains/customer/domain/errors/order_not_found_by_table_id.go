package domainerrors

import (
	"fmt"
)

type OrderNotFoundByTableID struct {
	TableID string
}

func (o OrderNotFoundByTableID) Error() string {
	return fmt.Sprintf("Order not found by table id: %s", o.TableID)
}
