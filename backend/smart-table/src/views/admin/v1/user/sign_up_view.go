package views

import (
	"context"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	"github.com/smart-table/src/utils"
	viewsAdminUser "github.com/smart-table/src/views/codegen/admin_user"
)

func (h *V1AdminUserHandler) PostV1AdminUserSignUp(
	ctx context.Context,
	request viewsAdminUser.PostAdminV1UserSignUpRequestObject,
) (viewsAdminUser.PostAdminV1UserSignUpResponseObject, error) {
	handler, err := utils.GetFromContainer[app.UserSingUpCommandHandler](ctx)

	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.UserSingUpCommand{
		Login:     request.Body.Login,
		TgLogin:   request.Body.TgLogin,
		FirstName: request.Body.FirstName,
		LastName:  request.Body.LastName,
		Password:  request.Body.Password,
	})
	if err != nil {
		return nil, err
	}

	return viewsAdminUser.PostAdminV1UserSignUp200JSONResponse{
		Body: viewsAdminUser.V1AdminUserSignUpResponse{
			UserUUID: result.UserUUID,
		},
		Headers: viewsAdminUser.PostAdminV1UserSignUp200ResponseHeaders{
			SetCookie: result.JwtToken,
		},
	}, nil
}
