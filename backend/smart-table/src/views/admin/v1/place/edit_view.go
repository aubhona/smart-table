package views

import (
	"context"

	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceEdit(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceEditRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceEditResponseObject, error) {
	return viewsAdminPlace.PostAdminV1PlaceEdit204Response{}, nil
}
