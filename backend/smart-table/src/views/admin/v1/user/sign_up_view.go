package views

import (
	"context"
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminUser "github.com/smart-table/src/views/codegen/admin_user"
)

func (h *AdminV1UserHandler) PostAdminV1UserSignUp(
	ctx context.Context,
	request viewsAdminUser.PostAdminV1UserSignUpRequestObject,
) (viewsAdminUser.PostAdminV1UserSignUpResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.UserSingUpCommandHandler](ctx)
	if err != nil {
		if utils.IsTheSameErrorType[appErrors.LoginOrTgLoginAlreadyExists](err) {
			return viewsAdminUser.PostAdminV1UserSignUp409JSONResponse{
				Code:    viewsAdminUser.AlreadyExist,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))

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
		if utils.IsTheSameErrorType[appErrors.LoginOrTgLoginAlreadyExists](err) {
			return viewsAdminUser.PostAdminV1UserSignUp409JSONResponse{
				Code:    viewsAdminUser.AlreadyExist,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminUser.PostAdminV1UserSignUp200JSONResponse{
		Body: viewsAdminUser.AdminV1UserSignUpResponse{
			UserUUID: result.UserUUID,
		},
		Headers: viewsAdminUser.PostAdminV1UserSignUp200ResponseHeaders{
			SetCookie: fmt.Sprintf("%s=%s", "jwt", result.JwtToken),
		},
	}, nil
}
