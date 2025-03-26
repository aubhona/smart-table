package views

import (
	"context"
	"errors"
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminUser "github.com/smart-table/src/views/codegen/admin_user"
)

func (h *AdminV1UserHandler) PostAdminV1UserSignIn(
	ctx context.Context,
	request viewsAdminUser.PostAdminV1UserSignInRequestObject,
) (viewsAdminUser.PostAdminV1UserSignInResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.UserSingInCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.UserSingInCommand{
		Login:    request.Body.Login,
		Password: request.Body.Password,
	})
	if err != nil {
		if errors.Is(err, appErrors.IncorrectPassword{}) {
			return viewsAdminUser.PostAdminV1UserSignIn401JSONResponse{
				Code:    "Incorrect password",
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminUser.PostAdminV1UserSignIn200JSONResponse{
		Body: viewsAdminUser.V1AdminUserSignInResponse{
			UserUUID: result.UserUUID,
		},
		Headers: viewsAdminUser.PostAdminV1UserSignIn200ResponseHeaders{
			SetCookie: result.JwtToken,
		},
	}, nil
}
