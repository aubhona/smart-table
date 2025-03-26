package viewscustomer

import (
	"context"

	app "github.com/smart-table/src/domains/customer/app/use_cases"
	domainErrors "github.com/smart-table/src/domains/customer/domain/errors"
	"github.com/smart-table/src/utils"
	viewsCustomer "github.com/smart-table/src/views/codegen/customer"
)

func (h *CustomerV1Handler) PostCustomerV1SignIn(
	ctx context.Context,
	request viewsCustomer.PostCustomerV1SignInRequestObject,
) (viewsCustomer.PostCustomerV1SignInResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CustomerAuthorizeCommandHandler](ctx)

	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.CustomerAuthorizeCommand{
		TgID:    request.Body.TgID,
		TgLogin: request.Body.TgLogin,
		ChatID:  request.Body.ChatID,
	})
	if err != nil {
		if utils.IsTheSameErrorType(err, domainErrors.CustomerNotFoundByTgID{}) {
			return viewsCustomer.PostCustomerV1SignIn404JSONResponse{
				Code:    viewsCustomer.NotFound,
				Message: err.Error(),
			}, nil
		}

		return nil, err
	}

	return viewsCustomer.PostCustomerV1SignIn200JSONResponse{
		CustomerUUID: result.CustomerUUID,
	}, nil
}
