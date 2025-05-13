package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceOrderListClosed(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceOrderListClosedRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceOrderListClosedResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceOrderListClosed200JSONResponse{}, nil
}
