package app

import (
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
)

type RestaurantListCommandHandlerResult struct {
	RestaurantList []utils.SharedRef[domain.Restaurant]
}

type RestaurantListCommandHandler struct {
	userRepository       domain.UserRepository
	restaurantRepository domain.RestaurantRepository
	placeRepository      domain.PlaceRepository
	uuidGenerator        *domainServices.UUIDGenerator
}

func NewRestaurantListCommandHandler(
	userRepository domain.UserRepository,
	restaurantRepository domain.RestaurantRepository,
	placeRepository domain.PlaceRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *RestaurantListCommandHandler {
	return &RestaurantListCommandHandler{
		userRepository,
		restaurantRepository,
		placeRepository,
		uuidGenerator,
	}
}

func (handler *RestaurantListCommandHandler) Handle(
	restaurantListCommand *RestaurantListCommand,
) (RestaurantListCommandHandlerResult, error) {
	_, err := handler.userRepository.FindUserByUUID(restaurantListCommand.UserUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding user by uuid",
			zap.String("user_uuid", restaurantListCommand.UserUUID.String()),
			zap.Error(err))

		return RestaurantListCommandHandlerResult{}, err
	}

	restaurantList, err := handler.restaurantRepository.FindRestaurantsByOwnerUUID(restaurantListCommand.UserUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding restaurants by owner uuid",
			zap.String("owner_uuid", restaurantListCommand.UserUUID.String()),
			zap.Error(err))

		return RestaurantListCommandHandlerResult{}, err
	}

	placeList, err := handler.placeRepository.FindPlacesByEmployeeUserUUID(restaurantListCommand.UserUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding places by employee uuid",
			zap.String("user_uuid", restaurantListCommand.UserUUID.String()),
			zap.Error(err))

		return RestaurantListCommandHandlerResult{}, err
	}

	existingRestaurantUUIDs := make(map[uuid.UUID]bool)
	for _, restaurant := range restaurantList {
		existingRestaurantUUIDs[restaurant.Get().GetUUID()] = true
	}

	for _, place := range placeList {
		restaurant := place.Get().GetRestaurant()
		if !existingRestaurantUUIDs[restaurant.Get().GetUUID()] {
			restaurantList = append(restaurantList, restaurant)
			existingRestaurantUUIDs[restaurant.Get().GetUUID()] = true
		}
	}

	return RestaurantListCommandHandlerResult{restaurantList}, nil
}
