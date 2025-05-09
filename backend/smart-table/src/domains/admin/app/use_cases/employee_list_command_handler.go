package app

import (
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"
)

type EmployeeListCommandHandlerResult struct {
	EmployeeList []utils.SharedRef[domain.Employee]
}

type EmployeeListCommandHandler struct {
	placeRepository domain.PlaceRepository
	userRepository  domain.UserRepository
}

func NewEmployeeListCommandHandler(
	placeRepository domain.PlaceRepository,
	userRepository domain.UserRepository,
) *EmployeeListCommandHandler {
	return &EmployeeListCommandHandler{
		placeRepository,
		userRepository,
	}
}

func (handler *EmployeeListCommandHandler) Handle(
	employeeListCommand *EmployeeListCommand,
) (EmployeeListCommandHandlerResult, error) {
	place, err := handler.placeRepository.FindPlace(employeeListCommand.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding place by uuid", zap.Error(err))
		return EmployeeListCommandHandlerResult{}, err
	}

	if !domain.IsHasAccess(employeeListCommand.UserUUID, place, domain.All) {
		return EmployeeListCommandHandlerResult{}, appErrors.PlaceAccessDenied{
			UserUUID:  employeeListCommand.UserUUID,
			PlaceUUID: employeeListCommand.PlaceUUID,
		}
	}

	return EmployeeListCommandHandlerResult{place.Get().GetEmployees()}, nil
}
