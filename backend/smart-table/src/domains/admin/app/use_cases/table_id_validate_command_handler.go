package app

import (
	app "github.com/smart-table/src/domains/admin/app/services"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
	"go.uber.org/zap"
)

type TableIDValidateCommandHandlerResult struct {
	IsValid bool
}

type TableIDValidateCommandHandler struct {
	placeRepository   domain.PlaceRepository
	placeTableService *app.PlaceTableService
}

func NewTableIDValidateCommandHandler(
	placeRepository domain.PlaceRepository,
	placeTableService *app.PlaceTableService,
) *TableIDValidateCommandHandler {
	return &TableIDValidateCommandHandler{
		placeRepository,
		placeTableService,
	}
}

func (handler *TableIDValidateCommandHandler) Handle(
	tableIDValidateCommand *TableIDValidateCommand,
) (TableIDValidateCommandHandlerResult, error) {
	placeUUID, err := handler.placeTableService.GetPlaceUUIDFromTableID(tableIDValidateCommand.TableID)
	if err != nil {
		logging.GetLogger().Error("error while getting place_uuid from table_id", zap.Error(err))
		return TableIDValidateCommandHandlerResult{}, err
	}

	place, err := handler.placeRepository.FindPlace(placeUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding place by uuid", zap.Error(err))
		return TableIDValidateCommandHandlerResult{}, err
	}

	tableNumber, err := handler.placeTableService.GetTableNumberFromTableID(tableIDValidateCommand.TableID)
	if err != nil {
		logging.GetLogger().Error("error while getting table_number from table_id", zap.Error(err))
		return TableIDValidateCommandHandlerResult{}, err
	}

	if !place.Get().ValidateTableNumber(tableNumber) {
		return TableIDValidateCommandHandlerResult{}, appErrors.InvalidTableNumber{
			TableNumber: tableNumber,
			PlaceUUID:   placeUUID,
		}
	}

	return TableIDValidateCommandHandlerResult{true}, nil
}
