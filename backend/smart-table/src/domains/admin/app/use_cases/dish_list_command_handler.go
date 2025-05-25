package app

import (
	"io"
	"slices"
	"strings"
	"sync"

	"github.com/google/uuid"

	app "github.com/smart-table/src/domains/admin/app/queries"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"

	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/logging"
)

type DishListDTO struct {
	ID             uuid.UUID
	RestaurantUUID uuid.UUID
	Name           string
	Description    string
	Calories       int
	Weight         int
	Picture        io.ReadCloser
	Category       string
}

type DishListCommandHandlerResult struct {
	DishList []DishListDTO
}

type DishListCommandHandler struct {
	restaurantRepository domain.RestaurantRepository
	s3QueryService       *app.S3QueryService
}

func NewDishListCommandHandler(
	restaurantRepository domain.RestaurantRepository,
	s3QueryService *app.S3QueryService,
) *DishListCommandHandler {
	return &DishListCommandHandler{
		restaurantRepository,
		s3QueryService,
	}
}

func (handler *DishListCommandHandler) Handle(
	command *DishListCommand,
) (DishListCommandHandlerResult, error) {
	restaurant, err := handler.restaurantRepository.FindRestaurant(command.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding restaurant by uuid", zap.Error(err))

		return DishListCommandHandlerResult{}, err
	}

	//nolint
	// TODO: Add role checking
	if restaurant.Get().GetOwner().Get().GetUUID() != command.OwnerUUID {
		logging.GetLogger().Error("restaurant access denied",
			zap.String("user_uuid", command.OwnerUUID.String()),
			zap.String("restaurant_uuid", command.RestaurantUUID.String()))

		return DishListCommandHandlerResult{}, appErrors.RestaurantAccessDenied{
			UserUUID:       command.OwnerUUID,
			RestaurantUUID: command.RestaurantUUID,
		}
	}

	waitGroup := sync.WaitGroup{}
	dishDTOList := make([]DishListDTO, 0, len(restaurant.Get().GetDishes()))
	mut := sync.Mutex{}

	for _, dish := range restaurant.Get().GetDishes() {
		logging.GetLogger().Debug("Get dish", zap.String("dish_uuid", dish.Get().GetUUID().String()))

		waitGroup.Add(1)

		go func(dish utils.SharedRef[domain.Dish]) {
			defer waitGroup.Done()

			var image io.ReadCloser

			if command.NeedPicture {
				image, err = handler.s3QueryService.GetImage(dish.Get().GetPictureKey())
				if err != nil {
					logging.GetLogger().Error(
						"error while getting image from S3",
						zap.String("picture_key", dish.Get().GetPictureKey()),
						zap.Error(err),
					)

					return
				}
			}

			mut.Lock()
			defer mut.Unlock()

			dishDTOList = append(dishDTOList, DishListDTO{
				ID:             dish.Get().GetUUID(),
				RestaurantUUID: dish.Get().GetUUID(),
				Name:           dish.Get().GetName(),
				Description:    dish.Get().GetDescription(),
				Calories:       dish.Get().GetCalories(),
				Weight:         dish.Get().GetWeight(),
				Picture:        image,
				Category:       dish.Get().GetCategory(),
			})

			logging.GetLogger().Debug("Add dish list dto", zap.String("dish_uuid", dish.Get().GetUUID().String()))
		}(dish)
	}

	waitGroup.Wait()

	slices.SortFunc(dishDTOList, func(a, b DishListDTO) int {
		return strings.Compare(a.Name, b.Name)
	})

	return DishListCommandHandlerResult{DishList: dishDTOList}, nil
}
