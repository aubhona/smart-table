package domain_errors

import (
	"fmt"
)

type OrderNotFoundByTableId struct {
	TableId string
}

func (o OrderNotFoundByTableId) Error() string {
	return fmt.Sprintf("Order not found by table id: %s", o.TableId)
}
