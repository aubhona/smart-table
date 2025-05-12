package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceMenuDishEdit(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceMenuDishEditRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceMenuDishEditResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceMenuDishEdit204Response{}, nil
}
