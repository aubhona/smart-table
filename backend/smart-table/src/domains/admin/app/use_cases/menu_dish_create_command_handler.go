package app

import (
	"github.com/google/uuid"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"
)

type MenuDishCreateCommandHandlerResult struct {
	MenuDishUUID uuid.UUID
}

type MenuDishCreateCommandHandler struct {
	placeRepository domain.PlaceRepository
	uuidGenerator   *domainServices.UUIDGenerator
}

func NewMenuDishCreateCommandHandler(
	placeRepository domain.PlaceRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *MenuDishCreateCommandHandler {
	return &MenuDishCreateCommandHandler{
		placeRepository,
		uuidGenerator,
	}
}

func (handler *MenuDishCreateCommandHandler) Handle(
	command *MenuDishCreateCommand,
) (MenuDishCreateCommandHandlerResult, error) {
	tx, err := handler.placeRepository.Begin()
	if err != nil {
		return MenuDishCreateCommandHandlerResult{}, err
	}

	defer utils.Rollback(handler.placeRepository, tx)

	place, err := handler.placeRepository.FindPlaceForUpdate(tx, command.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while place by uuid", zap.Error(err))
		return MenuDishCreateCommandHandlerResult{}, err
	}

	if !domain.IsHasAccess(command.UserUUID, place, domain.OwnerAndAdmin) {
		return MenuDishCreateCommandHandlerResult{}, appErrors.PlaceAccessDenied{
			UserUUID:  command.UserUUID,
			PlaceUUID: command.PlaceUUID,
		}
	}

	menuDish, err := place.Get().AddMenuDish(
		command.DishUUID,
		command.Price,
		//nolint
		true, // TODO: Think about it
		handler.uuidGenerator,
	)
	if err != nil {
		return MenuDishCreateCommandHandlerResult{}, err
	}

	err = handler.placeRepository.Update(tx, place)
	if err != nil {
		logging.GetLogger().Error("error while updating place", zap.Error(err))

		return MenuDishCreateCommandHandlerResult{}, err
	}

	err = handler.placeRepository.Commit(tx)
	if err != nil {
		logging.GetLogger().Error("error while committing place", zap.Error(err))

		return MenuDishCreateCommandHandlerResult{}, err
	}

	return MenuDishCreateCommandHandlerResult{MenuDishUUID: menuDish.Get().GetUUID()}, nil
}
