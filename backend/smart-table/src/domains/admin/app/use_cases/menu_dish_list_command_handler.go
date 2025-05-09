package app

import (
	"io"
	"sync"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"

	app "github.com/smart-table/src/domains/admin/app/queries"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"

	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
)

type MenuDishListDTO struct {
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

type MenuDishListCommandHandlerResult struct {
	MenuDishList []MenuDishListDTO
}

type MenuDishListCommandHandler struct {
	placeRepository domain.PlaceRepository
	s3QueryService  *app.S3QueryService
}

func NewMenuDishListCommandHandler(
	placeRepository domain.PlaceRepository,
	s3QueryService *app.S3QueryService,
) *MenuDishListCommandHandler {
	return &MenuDishListCommandHandler{
		placeRepository,
		s3QueryService,
	}
}

func (handler *MenuDishListCommandHandler) Handle(
	command *MenuDishListCommand,
) (MenuDishListCommandHandlerResult, error) {
	place, err := handler.placeRepository.FindPlace(command.PlaceUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding restaurant by uuid", zap.Error(err))

		return MenuDishListCommandHandlerResult{}, err
	}

	if !domain.IsHasAccess(command.UserUUID, place, domain.All) {
		return MenuDishListCommandHandlerResult{}, appErrors.PlaceAccessDenied{
			UserUUID:  command.UserUUID,
			PlaceUUID: command.PlaceUUID,
		}
	}

	waitGroup := sync.WaitGroup{}
	menuDishDTOList := make([]MenuDishListDTO, 0, len(place.Get().GetMenuDishes()))
	mut := sync.Mutex{}

	for _, menuDish := range place.Get().GetMenuDishes() {
		waitGroup.Add(1)

		go func(menuDish utils.SharedRef[domain.MenuDish]) {
			defer waitGroup.Done()

			image, err := handler.s3QueryService.GetImage(menuDish.Get().GetDish().Get().GetPictureKey())
			if err != nil {
				logging.GetLogger().Error(
					"error while getting image from S3",
					zap.String("picture_key", menuDish.Get().GetDish().Get().GetPictureKey()),
					zap.Error(err),
				)

				return
			}

			mut.Lock()
			defer mut.Unlock()

			menuDishDTOList = append(menuDishDTOList, MenuDishListDTO{
				ID:          menuDish.Get().GetUUID(),
				PlaceUUID:   menuDish.Get().GetPlaceUUID(),
				Name:        menuDish.Get().GetDish().Get().GetName(),
				Description: menuDish.Get().GetDish().Get().GetDescription(),
				Calories:    menuDish.Get().GetDish().Get().GetCalories(),
				Weight:      menuDish.Get().GetDish().Get().GetWeight(),
				Picture:     image,
				Category:    menuDish.Get().GetDish().Get().GetCategory(),
				Price:       menuDish.Get().GetPrice(),
				Exist:       menuDish.Get().GetExist(),
			})
		}(menuDish)
	}

	waitGroup.Wait()

	return MenuDishListCommandHandlerResult{MenuDishList: menuDishDTOList}, nil
}
