package app

import (
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/utils"
	"go.uber.org/zap"

	"github.com/google/uuid"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
)

type RestaurantCreateCommandHandlerResult struct {
	RestaurantUUID uuid.UUID
}

type RestaurantCreateCommandHandler struct {
	userRepository       domain.UserRepository
	restaurantRepository domain.RestaurantRepository
	uuidGenerator        *domainServices.UUIDGenerator
}

func NewRestaurantCreateCommandHandler(
	userRepository domain.UserRepository,
	restaurantRepository domain.RestaurantRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *RestaurantCreateCommandHandler {
	return &RestaurantCreateCommandHandler{
		userRepository,
		restaurantRepository,
		uuidGenerator,
	}
}

func (handler *RestaurantCreateCommandHandler) Handle(
	restaurantCreateCommand *RestaurantCreateCommand,
) (RestaurantCreateCommandHandlerResult, error) {
	_, err := handler.restaurantRepository.FindRestaurantByName(restaurantCreateCommand.Name)

	if err == nil {
		logging.GetLogger().Error("restaurant name already exists",
			zap.String("name", restaurantCreateCommand.Name))

		return RestaurantCreateCommandHandlerResult{}, appErrors.RestaurantNameAlreadyExists{
			Name: restaurantCreateCommand.Name,
		}
	}

	if !utils.IsTheSameErrorType[domainErrors.RestaurantNotFoundByName](err) {
		logging.GetLogger().Error("error while checking restaurant name existence",
			zap.Error(err))
		return RestaurantCreateCommandHandlerResult{}, err
	}

	owner, err := handler.userRepository.FindUserByUUID(restaurantCreateCommand.OwnerUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding owner by uuid",
			zap.String("owner_uuid", restaurantCreateCommand.OwnerUUID.String()),
			zap.Error(err))

		return RestaurantCreateCommandHandlerResult{}, err
	}

	restaurant := domain.NewRestaurant(
		owner,
		restaurantCreateCommand.Name,
		handler.uuidGenerator,
	)

	tx, err := handler.restaurantRepository.Begin()
	if err != nil {
		logging.GetLogger().Error("error while beginning transaction",
			zap.Error(err))
		return RestaurantCreateCommandHandlerResult{}, err
	}

	defer utils.Rollback(handler.restaurantRepository, tx)

	err = handler.restaurantRepository.Save(tx, restaurant)
	if err != nil {
		logging.GetLogger().Error("error while saving restaurant",
			zap.String("restaurant_name", restaurantCreateCommand.Name),
			zap.Error(err))

		return RestaurantCreateCommandHandlerResult{}, err
	}

	err = handler.restaurantRepository.Commit(tx)
	if err != nil {
		logging.GetLogger().Error("error while committing transaction",
			zap.Error(err))
		return RestaurantCreateCommandHandlerResult{}, err
	}

	return RestaurantCreateCommandHandlerResult{restaurant.Get().GetUUID()}, nil
}
