package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceEmployeeEdit(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceEmployeeEditRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceEmployeeEditResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceEmployeeEdit204Response{}, nil
}
