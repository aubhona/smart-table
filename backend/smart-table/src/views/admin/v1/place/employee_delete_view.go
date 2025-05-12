package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceEmployeeDelete(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceEmployeeDeleteRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceEmployeeDeleteResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceEmployeeDelete204Response{}, nil
}
