package errors

import (
	"fmt"

	"github.com/smart-table/src/utils"
)

type InvalidOrderResolution struct {
	OrderResolution utils.Optional[string]
}

func (e InvalidOrderResolution) Error() string {
	return fmt.Sprintf("invalid order_resolution='%s'", e.OrderResolution.ValueOr("nil"))
}
