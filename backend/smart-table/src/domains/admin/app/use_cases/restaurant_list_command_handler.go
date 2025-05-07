package app

import (
	"go.uber.org/zap"

	"github.com/smart-table/src/domains/admin/domain"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
)

type RestaurantListCommandHandlerResult struct {
	DomainRestaurantList []utils.SharedRef[domain.Restaurant]
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

func (handler *RestaurantListCommandHandler) Handle(
	restaurantListCommand *RestaurantListCommand,
) (RestaurantListCommandHandlerResult, error) {
	_, err := handler.userRepository.FindUserByUUID(restaurantListCommand.OwnerUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding owner by uuid",
			zap.String("owner_uuid", restaurantListCommand.OwnerUUID.String()),
			zap.Error(err))

		return RestaurantListCommandHandlerResult{}, err
	}

	domainRestaurantList, err := handler.restaurantRepository.FindRestaurantsByOwnerUUID(restaurantListCommand.OwnerUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding restaurants by owner uuid",
			zap.String("owner_uuid", restaurantListCommand.OwnerUUID.String()),
			zap.Error(err))

		return RestaurantListCommandHandlerResult{}, err
	}

	return RestaurantListCommandHandlerResult{domainRestaurantList}, nil
}
