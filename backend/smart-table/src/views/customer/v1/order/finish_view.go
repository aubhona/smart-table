package viewscustomerorder

import (
	"context"

	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
)

func (h *CustomerV1OrderHandler) PostCustomerV1OrderFinish(
	ctx context.Context,
	request viewsCustomerOrder.PostCustomerV1OrderFinishRequestObject,
) (viewsCustomerOrder.PostCustomerV1OrderFinishResponseObject, error) {
	return nil, nil
}
