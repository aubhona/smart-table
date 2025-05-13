package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceOrderListOpened(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceOrderListOpenedRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceOrderListOpenedResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceOrderListOpened200JSONResponse{}, nil
}
