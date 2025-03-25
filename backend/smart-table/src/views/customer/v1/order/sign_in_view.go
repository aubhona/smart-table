package views

import (
	"context"

	viewsCustomer "github.com/smart-table/src/views/codegen/customer"
)

func (h *CustomerV1OrderHandler) PostCustomerV1OrderCustomerSignIn(
	ctx context.Context,
	request viewsCustomer.PostCustomerV1OrderCustomerSignInRequestObject,
) (viewsCustomer.PostCustomerV1OrderCustomerSignInResponseObject, error) {
	//nolint: godox, gocritic
	// TODO: impl
	return nil, nil
}
