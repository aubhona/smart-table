package app

import (
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type EmployeeAddCommandHandler struct {
	placeRepository domain.PlaceRepository
	userRepository  domain.UserRepository
}

func NewEmployeeAddCommandHandler(
	placeRepository domain.PlaceRepository,
	userRepository domain.UserRepository,
) *EmployeeAddCommandHandler {
	return &EmployeeAddCommandHandler{
		placeRepository,
		userRepository,
	}
}

func (handler *EmployeeAddCommandHandler) Handle(
	employeeAddCommand *EmployeeAddCommand,
) error {
	tx, err := handler.placeRepository.Begin()
	if err != nil {
		return err
	}

	defer utils.Rollback(handler.placeRepository, tx)

	place, err := handler.placeRepository.FindPlace(employeeAddCommand.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while place by uuid", zap.Error(err))
		return err
	}

	if place.Get().GetRestaurant().Get().GetOwner().Get().GetUUID() != employeeAddCommand.UserUUID {
		return appErrors.RestaurantAccessDenied{
			UserUUID:       employeeAddCommand.UserUUID,
			RestaurantUUID: employeeAddCommand.PlaceUUID,
		}
	}

	employeeProfile, err := handler.userRepository.FindUserByLogin(employeeAddCommand.EmployeeLogin)
	if err != nil {
		logging.GetLogger().Error("error while finding employee by login", zap.Error(err))
		return err
	}

	placeEmployees := place.Get().GetEmployees()

	isExist := slices.ContainsFunc(placeEmployees, func(employee utils.SharedRef[domain.Employee]) bool {
		return employee.Get().GetUser().Get().GetUUID() == employeeProfile.Get().GetUUID()
	})
	if isExist {
		logging.GetLogger().Error("employee already exists in this place",
			zap.String("employee_login", employeeAddCommand.EmployeeLogin),
			zap.String("place_uuid", employeeAddCommand.PlaceUUID.String()))

		return appErrors.EmployeeAlreadyExists{
			EmployeeLogin: employeeAddCommand.EmployeeLogin,
			PlaceUUID:     employeeAddCommand.PlaceUUID,
		}
	}

	place.Get().AddEmployee(
		employeeProfile,
		employeeAddCommand.EmployeeRole,
	)

	err = handler.placeRepository.Update(tx, place)
	if err != nil {
		logging.GetLogger().Error("error while updating place", zap.Error(err))

		return err
	}

	err = handler.placeRepository.Commit(tx)
	if err != nil {
		logging.GetLogger().Error("error while committing place", zap.Error(err))

		return err
	}

	return nil
}
