package app

import (
	appServices "github.com/smart-table/src/domains/admin/app/services"
	"go.uber.org/zap"

	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
)

type TableDeepLinksListCommandHandlerResult struct {
	TableDeepLinksList []string
}

type TableDeepLinksListCommandHandler struct {
	placeRepository   domain.PlaceRepository
	placeTableService *appServices.PlaceTableService
}

func NewTableDeepLinksListCommandHandler(
	placeRepository domain.PlaceRepository,
	placeTableService *appServices.PlaceTableService,
) *TableDeepLinksListCommandHandler {
	return &TableDeepLinksListCommandHandler{
		placeRepository,
		placeTableService,
	}
}

func (handler *TableDeepLinksListCommandHandler) Handle(
	command *TableDeepLinksListCommand,
) (TableDeepLinksListCommandHandlerResult, error) {
	place, err := handler.placeRepository.FindPlace(command.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding restaurant by uuid", zap.Error(err))

		return TableDeepLinksListCommandHandlerResult{}, err
	}

	if !domain.IsHasAccess(command.UserUUID, place, domain.All) {
		return TableDeepLinksListCommandHandlerResult{}, appErrors.PlaceAccessDenied{
			UserUUID:  command.UserUUID,
			PlaceUUID: command.PlaceUUID,
		}
	}

	return TableDeepLinksListCommandHandlerResult{
		handler.placeTableService.GetTableDeepLinkForQR(place),
	}, nil
}
