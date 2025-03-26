package viewscustomer

import (
	"context"

	app "github.com/smart-table/src/domains/customer/app/use_cases"
	appErrors "github.com/smart-table/src/domains/customer/app/use_cases/errors"
	"github.com/smart-table/src/utils"
	viewsCustomer "github.com/smart-table/src/views/codegen/customer"
)

func (h *CustomerV1Handler) PostCustomerV1SignUp(
	ctx context.Context,
	request viewsCustomer.PostCustomerV1SignUpRequestObject,
) (viewsCustomer.PostCustomerV1SignUpResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.CustomerRegisterCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.CustomerRegisterCommand{
		TgID:    request.Body.TgID,
		TgLogin: request.Body.TgLogin,
		ChatID:  request.Body.ChatID,
	})
	if err != nil {
		if utils.IsTheSameErrorType(err, appErrors.CustomerAlreadyExist{}) {
			return viewsCustomer.PostCustomerV1SignUp409JSONResponse{
				Code:    viewsCustomer.AlreadyExist,
				Message: err.Error(),
			}, nil
		}

		return nil, err
	}

	return viewsCustomer.PostCustomerV1SignUp200JSONResponse{
		CustomerUUID: result.CustomerUUID,
	}, nil
}
