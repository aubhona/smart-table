package views

import (
	"context"

	viewsCustomer "github.com/smart-table/src/views/codegen/customer"
)

func (h *CustomerV1OrderHandler) PostCustomerV1OrderCustomerSignUp(
	ctx context.Context,
	request viewsCustomer.PostCustomerV1OrderCustomerSignUpRequestObject,
) (viewsCustomer.PostCustomerV1OrderCustomerSignUpResponseObject, error) {
	//nolint: godox, gocritic
	// TODO: impl
	return nil, nil
}
