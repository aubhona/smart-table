package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceDelete(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceDeleteRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceDeleteResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceDelete204Response{}, nil
}
