package app

import (
	"errors"
	"io"

	appServices "github.com/smart-table/src/domains/admin/app/services"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"

	app "github.com/smart-table/src/domains/admin/app/queries"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"

	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
)

type MenuDishStateDTO struct {
	ID          uuid.UUID
	PlaceUUID   uuid.UUID
	Name        string
	Description string
	Calories    int
	Weight      int
	Picture     io.ReadCloser
	Category    string
	Price       decimal.Decimal
	Exist       bool
}

type MenuDishStateCommandHandlerResult struct {
	MenuDish MenuDishStateDTO
}

type MenuDishStateCommandHandler struct {
	placeRepository   domain.PlaceRepository
	s3QueryService    *app.S3QueryService
	placeTableService *appServices.PlaceTableService
}

func NewMenuDishStateCommandHandler(
	placeRepository domain.PlaceRepository,
	s3QueryService *app.S3QueryService,
	placeTableService *appServices.PlaceTableService,
) *MenuDishStateCommandHandler {
	return &MenuDishStateCommandHandler{
		placeRepository,
		s3QueryService,
		placeTableService,
	}
}

func (handler *MenuDishStateCommandHandler) Handle( //nolint
	command *MenuDishStateCommand,
) (MenuDishStateCommandHandlerResult, error) {
	var place utils.SharedRef[domain.Place]

	var menuDish utils.SharedRef[domain.MenuDish]

	var image io.ReadCloser

	needPicture := false

	var err error

	if !command.InternalCall.HasValue() && !command.AdminCall.HasValue() {
		return MenuDishStateCommandHandlerResult{}, errors.New("invalid command")
	}

	if command.AdminCall.HasValue() {
		place, err = handler.placeRepository.FindPlace(command.AdminCall.Value().PlaceUUID)
		if err != nil {
			logging.GetLogger().Error("error while finding restaurant by uuid", zap.Error(err))

			return MenuDishStateCommandHandlerResult{}, err
		}

		if !domain.IsHasAccess(command.AdminCall.Value().UserUUID, place, domain.All) {
			return MenuDishStateCommandHandlerResult{}, appErrors.PlaceAccessDenied{
				UserUUID:  command.AdminCall.Value().UserUUID,
				PlaceUUID: command.AdminCall.Value().PlaceUUID,
			}
		}

		menuDishOpt := place.Get().GetMenuDishByUUID(command.AdminCall.Value().MenuDishUUID)
		if !menuDishOpt.HasValue() {
			return MenuDishStateCommandHandlerResult{}, domainErrors.MenuDishNotFound{UUID: command.AdminCall.Value().MenuDishUUID}
		}

		menuDish = menuDishOpt.Value()
	} else {
		placeUUID, err := handler.placeTableService.GetPlaceUUIDFromTableID(command.InternalCall.Value().TabledID)
		if err != nil {
			return MenuDishStateCommandHandlerResult{}, err
		}

		place, err = handler.placeRepository.FindPlace(placeUUID)
		if err != nil {
			logging.GetLogger().Error("error while finding restaurant by uuid", zap.Error(err))

			return MenuDishStateCommandHandlerResult{}, err
		}

		menuDishOpt := place.Get().GetMenuDishByUUID(command.InternalCall.Value().MenuDishUUID)
		if !menuDishOpt.HasValue() {
			return MenuDishStateCommandHandlerResult{}, domainErrors.MenuDishNotFound{UUID: command.InternalCall.Value().MenuDishUUID}
		}

		needPicture = command.InternalCall.Value().NeedPicture

		menuDish = menuDishOpt.Value()
	}

	if needPicture {
		image, err = handler.s3QueryService.GetImage(menuDish.Get().GetDish().Get().GetPictureKey())
		if err != nil {
			logging.GetLogger().Error(
				"error while getting image from S3",
				zap.String("picture_key", menuDish.Get().GetDish().Get().GetPictureKey()),
				zap.Error(err),
			)

			return MenuDishStateCommandHandlerResult{}, err
		}
	}

	return MenuDishStateCommandHandlerResult{MenuDish: MenuDishStateDTO{
		ID:          menuDish.Get().GetUUID(),
		PlaceUUID:   menuDish.Get().GetPlaceUUID(),
		Name:        menuDish.Get().GetDish().Get().GetName(),
		Description: menuDish.Get().GetDish().Get().GetDescription(),
		Calories:    menuDish.Get().GetDish().Get().GetCalories(),
		Weight:      menuDish.Get().GetDish().Get().GetWeight(),
		Category:    menuDish.Get().GetDish().Get().GetCategory(),
		Price:       menuDish.Get().GetPrice(),
		Picture:     image,
		Exist:       menuDish.Get().GetExist(),
	}}, nil
}
