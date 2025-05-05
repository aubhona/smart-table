package app

import (
	"context"
	"fmt"

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
	ctx := context.Background()

	restaurant, err := handler.restaurantRepository.FindRestaurantByUUID(ctx, placeListCommand.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while finding restaurant by uuid: %v", err))
		return PlaceListCommandHandlerResult{}, err
	}

	if restaurant.Get().GetOwnerUUID() != placeListCommand.OwnerUUID {
		return PlaceListCommandHandlerResult{}, appErrors.RestaurantAccessDenied{
			UserUUID:       placeListCommand.OwnerUUID,
			RestaurantUUID: placeListCommand.RestaurantUUID,
		}
	}

	placeList, err := handler.placeRepository.FindPlaceListByRestaurantUUID(ctx, placeListCommand.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while finding place_list by owner_uuid: %v", err))
		return PlaceListCommandHandlerResult{}, err
	}

	return PlaceListCommandHandlerResult{placeList}, nil
}
