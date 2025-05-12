package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceMenuDishDelete(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceMenuDishDeleteRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceMenuDishDeleteResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceMenuDishDelete204Response{}, nil
}
