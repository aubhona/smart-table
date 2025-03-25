package views

import (
	"context"

	viewsAdminUser "github.com/smart-table/src/views/codegen/admin_user"
)

func (h *AdminV1UserHandler) PostAdminV1UserSignIn(
	ctx context.Context,
	request viewsAdminUser.PostAdminV1UserSignInRequestObject,
) (viewsAdminUser.PostAdminV1UserSignInResponseObject, error) {
	return nil, nil
}
