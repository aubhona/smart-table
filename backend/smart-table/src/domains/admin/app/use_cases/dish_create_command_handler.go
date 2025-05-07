package app

import (
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/queries"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"

	"github.com/google/uuid"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
)

type DishCreateCommandHandlerResult struct {
	DishUUID uuid.UUID
}

type DishCreateCommandHandler struct {
	restaurantRepository domain.RestaurantRepository
	s3QueryService       *app.S3QueryService
	uuidGenerator        *domainServices.UUIDGenerator
}

func NewDishCreateCommandHandler(
	restaurantRepository domain.RestaurantRepository,
	s3QueryService *app.S3QueryService,
	uuidGenerator *domainServices.UUIDGenerator,
) *DishCreateCommandHandler {
	return &DishCreateCommandHandler{
		restaurantRepository,
		s3QueryService,
		uuidGenerator,
	}
}

func (handler *DishCreateCommandHandler) Handle(
	command *DishCreateCommand,
) (DishCreateCommandHandlerResult, error) {
	tx, err := handler.restaurantRepository.Begin()
	if err != nil {
		return DishCreateCommandHandlerResult{}, err
	}

	defer utils.Rollback(handler.restaurantRepository, tx)

	restaurant, err := handler.restaurantRepository.FindRestaurantForUpdate(tx, command.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding restaurant by uuid", zap.Error(err))
		return DishCreateCommandHandlerResult{}, err
	}

	if restaurant.Get().GetOwner().Get().GetUUID() != command.OwnerUUID {
		logging.GetLogger().Error("restaurant access denied",
			zap.String("user_uuid", command.OwnerUUID.String()),
			zap.String("restaurant_uuid", command.RestaurantUUID.String()))

		return DishCreateCommandHandlerResult{}, appErrors.RestaurantAccessDenied{
			UserUUID:       command.OwnerUUID,
			RestaurantUUID: command.RestaurantUUID,
		}
	}

	pictureKey := fmt.Sprintf("%s/%s", restaurant.Get().GetUUID(), command.DishName)

	reader, err := command.Image.Reader()
	if err != nil {
		logging.GetLogger().Error("error while reading image file", zap.Error(err))

		return DishCreateCommandHandlerResult{}, err
	}

	err = handler.s3QueryService.StoreImage(reader, command.Image.FileSize(), pictureKey)
	if err != nil {
		logging.GetLogger().Error("error while storing image file", zap.Error(err))

		return DishCreateCommandHandlerResult{}, err
	}

	dish := restaurant.Get().AddDish(
		command.DishName,
		command.Description,
		command.Category,
		pictureKey,
		command.Calories,
		command.Weight,
		handler.uuidGenerator,
	)

	err = handler.restaurantRepository.Update(tx, restaurant)
	if err != nil {
		logging.GetLogger().Error("error while updating restaurant", zap.Error(err))

		return DishCreateCommandHandlerResult{}, err
	}

	err = handler.restaurantRepository.Commit(tx)
	if err != nil {
		logging.GetLogger().Error("error while committing restaurant", zap.Error(err))

		return DishCreateCommandHandlerResult{}, err
	}

	return DishCreateCommandHandlerResult{dish.Get().GetUUID()}, nil
}
