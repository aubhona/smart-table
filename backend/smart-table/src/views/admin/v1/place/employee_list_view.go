package views

import (
	"context"
	"fmt"

	"github.com/smart-table/src/logging"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func convertEmployeeToEmployeeInfo(
	employee utils.SharedRef[domain.Employee],
) viewsAdminPlace.EmployeeInfo {
	return viewsAdminPlace.EmployeeInfo{
		UUID:         employee.Get().GetUser().Get().GetUUID(),
		Login:        employee.Get().GetUser().Get().GetLogin(),
		TgLogin:      employee.Get().GetUser().Get().GetTgLogin(),
		FirstName:    employee.Get().GetUser().Get().GetFirstName(),
		LastName:     employee.Get().GetUser().Get().GetLastName(),
		EmployeeRole: viewsAdminPlace.Role(employee.Get().GetRole()),
	}
}

func (h *AdminV1PlaceHandler) PostAdminV1PlaceEmployeeList( //nolint
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceEmployeeListRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceEmployeeListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.EmployeeListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.EmployeeListCommand{
		UserUUID:  request.Params.UserUUID,
		PlaceUUID: request.Body.PlaceUUID,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceEmployeeList403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceEmployeeList404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	employeeInfoList := make([]viewsAdminPlace.EmployeeInfo, 0, len(result.EmployeeList))

	for _, employee := range result.EmployeeList {
		employeeInfoList = append(employeeInfoList, convertEmployeeToEmployeeInfo(employee))
	}

	return viewsAdminPlace.PostAdminV1PlaceEmployeeList200JSONResponse{
		EmployeeList: employeeInfoList,
	}, nil
}
