package views

import (
	"context"

	viewsAdminUser "github.com/smart-table/src/views/codegen/admin_user"
)

func (h *V1AdminUserHandler) PostV1AdminUserSignUp(
	ctx context.Context,
	request viewsAdminUser.PostAdminV1UserSignUpRequestObject,
) (viewsAdminUser.PostAdminV1UserSignUpResponseObject, error) {
	return nil, nil
}
