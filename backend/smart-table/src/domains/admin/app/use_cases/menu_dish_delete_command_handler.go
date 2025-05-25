package app

import (
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"
)

type MenuDishDeleteCommandHandler struct {
	placeRepository domain.PlaceRepository
	uuidGenerator   *domainServices.UUIDGenerator
}

func NewMenuDishDeleteCommandHandler(
	placeRepository domain.PlaceRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *MenuDishDeleteCommandHandler {
	return &MenuDishDeleteCommandHandler{
		placeRepository,
		uuidGenerator,
	}
}

func (handler *MenuDishDeleteCommandHandler) Handle(
	command *MenuDishDeleteCommand,
) error {
	tx, err := handler.placeRepository.Begin()
	if err != nil {
		return err
	}

	defer utils.Rollback(handler.placeRepository, tx)

	place, err := handler.placeRepository.FindPlaceForUpdate(tx, command.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding place by uuid", zap.Error(err))
		return err
	}

	if !domain.IsHasAccess(command.UserUUID, place, domain.OwnerAndAdmin) {
		return appErrors.PlaceAccessDenied{
			UserUUID:  command.UserUUID,
			PlaceUUID: command.PlaceUUID,
		}
	}

	err = place.Get().DeleteMenuDish(command.MenuDishUUID)
	if err != nil {
		return err
	}

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
