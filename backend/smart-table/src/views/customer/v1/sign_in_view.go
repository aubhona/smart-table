package viewscustomer

import (
	"context"
	"fmt"

	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appUseCasesErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsCustomer "github.com/smart-table/src/views/codegen/customer"
)

func (h *CustomerV1Handler) PostCustomerV1SignIn(
	ctx context.Context,
	request viewsCustomer.PostCustomerV1SignInRequestObject,
) (viewsCustomer.PostCustomerV1SignInResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CustomerAuthorizeCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.CustomerAuthorizeCommand{
		InitData: request.Body.InitData,
		TgID:     request.Body.TgID,
		TgLogin:  request.Body.TgLogin,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainErrors.CustomerNotFoundByTgID](err):
			return viewsCustomer.PostCustomerV1SignIn404JSONResponse{
				Code:    viewsCustomer.NotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appUseCasesErrors.IncorrectInitDataError](err):
			return viewsCustomer.PostCustomerV1SignIn403JSONResponse{
				Code:    viewsCustomer.IncorrectInitData,
				Message: err.Error(),
			}, nil
		}

		return nil, err
	}

	return viewsCustomer.PostCustomerV1SignIn200JSONResponse{
		CustomerUUID: result.CustomerUUID,
		JwtToken:     result.JwtToken,
	}, nil
}
