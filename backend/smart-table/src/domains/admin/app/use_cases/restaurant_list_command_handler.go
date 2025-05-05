package app

import (
	"context"
	"fmt"

	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	gen "github.com/smart-table/src/views/codegen/admin_restaurant"
)

type RestaurantListCommandHandlerResult struct {
	GenRestaurantList []gen.Restaurant
}

type RestaurantListCommandHandler struct {
	userRepository       domain.UserRepository
	restaurantRepository domain.RestaurantRepository
	uuidGenerator        *domainServices.UUIDGenerator
}

func NewRestaurantListCommandHandler(
	userRepository domain.UserRepository,
	restaurantRepository domain.RestaurantRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *RestaurantListCommandHandler {
	return &RestaurantListCommandHandler{
		userRepository,
		restaurantRepository,
		uuidGenerator,
	}
}

func convertDomainRestauranToGenRestaurant(
	domainRestaurant utils.SharedRef[domain.Restaurant],
) gen.Restaurant {
	return gen.Restaurant{
		Name: domainRestaurant.Get().GetName(),
		UUID: domainRestaurant.Get().GetUUID(),
	}
}

func (handler *RestaurantListCommandHandler) Handle(
	restaurantListCommand *RestaurantListCommand,
) (RestaurantListCommandHandlerResult, error) {
	ctx := context.Background()
	_, err := handler.userRepository.FindUserByUUID(ctx, restaurantListCommand.OwnerUUID)

	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while checking owner_uuid existence: %v", err))
		return RestaurantListCommandHandlerResult{}, err
	}

	domainRestaurantList, err := handler.restaurantRepository.FindRestaurantListByOwnerUUID(ctx, restaurantListCommand.OwnerUUID)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while finding restaurant_list by owner_uuid: %v", err))
		return RestaurantListCommandHandlerResult{}, err
	}

	genRestaurantList := make([]gen.Restaurant, 0, len(domainRestaurantList))

	for _, domainRestaurant := range domainRestaurantList {
		genRestaurantList = append(genRestaurantList, convertDomainRestauranToGenRestaurant(domainRestaurant))
	}

	return RestaurantListCommandHandlerResult{genRestaurantList}, nil
}
