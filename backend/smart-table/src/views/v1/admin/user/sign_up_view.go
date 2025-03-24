package views

import (
	"context"

	viewsAdminUser "github.com/smart-table/src/views/codegen/admin_user"
)

func (h *V1AdminUserHandler) PostV1AdminUserSignUp(
	ctx context.Context,
	request viewsAdminUser.PostV1AdminUserSignUpRequestObject,
) (viewsAdminUser.PostV1AdminUserSignUpResponseObject, error) {
	return nil, nil
}
