package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceOrderInfo(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceOrderInfoRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceOrderInfoResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceOrderInfo200JSONResponse{}, nil
}
