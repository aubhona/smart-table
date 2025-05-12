package viewscustomerorder

import (
	"context"

	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
)

func (h *CustomerV1OrderHandler) GetCustomerV1OrderCart(
	ctx context.Context,
	request viewsCustomerOrder.GetCustomerV1OrderCartRequestObject,
) (viewsCustomerOrder.GetCustomerV1OrderCartResponseObject, error) {
	return nil, nil
}
