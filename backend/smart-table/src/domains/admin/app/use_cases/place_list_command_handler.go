package app

import (
	"go.uber.org/zap"

	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
)

type PlaceListCommandHandlerResult struct {
	PlaceList []utils.SharedRef[domain.Place]
}

type PlaceListCommandHandler struct {
	restaurantRepository domain.RestaurantRepository
	userRepository       domain.UserRepository
	placeRepository      domain.PlaceRepository
	uuidGenerator        *domainServices.UUIDGenerator
}

func NewPlaceListCommandHandler(
	restaurantRepository domain.RestaurantRepository,
	userRepository domain.UserRepository,
	placeRepository domain.PlaceRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *PlaceListCommandHandler {
	return &PlaceListCommandHandler{
		restaurantRepository,
		userRepository,
		placeRepository,
		uuidGenerator,
	}
}

func (handler *PlaceListCommandHandler) Handle(
	placeListCommand *PlaceListCommand,
) (PlaceListCommandHandlerResult, error) {
	restaurant, err := handler.restaurantRepository.FindRestaurant(placeListCommand.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding restaurant", zap.Error(err))

		return PlaceListCommandHandlerResult{}, err
	}

	if restaurant.Get().GetOwner().Get().GetUUID() != placeListCommand.OwnerUUID {
		return PlaceListCommandHandlerResult{}, appErrors.RestaurantAccessDenied{
			UserUUID:       placeListCommand.OwnerUUID,
			RestaurantUUID: placeListCommand.RestaurantUUID,
		}
	}

	placeList, err := handler.placeRepository.FindPlacesByRestaurantUUID(placeListCommand.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding places", zap.Error(err))

		return PlaceListCommandHandlerResult{}, err
	}

	return PlaceListCommandHandlerResult{placeList}, nil
}
