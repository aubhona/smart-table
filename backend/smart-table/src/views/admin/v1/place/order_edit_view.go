package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceOrderEdit(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceOrderEditRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceOrderEditResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceOrderEdit204Response{}, nil
}