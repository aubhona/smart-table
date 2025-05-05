package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

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
	ctx := context.Background()
	isExist, err := handler.restaurantRepository.CheckNameExist(ctx, restaurantCreateCommand.Name)

	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while checking restaurant name existence: %v", err))
		return RestaurantCreateCommandHandlerResult{}, err
	}

	if isExist {
		return RestaurantCreateCommandHandlerResult{}, appErrors.RestaurantNameExists{
			Name: restaurantCreateCommand.Name,
		}
	}

	_, err = handler.userRepository.FindUserByUUID(ctx, restaurantCreateCommand.OwnerUUID)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while checking owner_uuid existence: %v", err))
		return RestaurantCreateCommandHandlerResult{}, err
	}

	restaurant := domain.NewRestaurant(
		restaurantCreateCommand.OwnerUUID,
		restaurantCreateCommand.Name,
		handler.uuidGenerator,
	)

	tx, err := handler.restaurantRepository.Begin(ctx)
	if err != nil {
		return RestaurantCreateCommandHandlerResult{}, err
	}

	defer func(restaurantRepository domain.RestaurantRepository, ctx context.Context, tx pgx.Tx) {
		_ = restaurantRepository.Commit(ctx, tx)
	}(handler.restaurantRepository, ctx, tx)

	err = handler.restaurantRepository.Save(ctx, tx, restaurant)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while restaurant saving: %v", err))
		return RestaurantCreateCommandHandlerResult{}, err
	}

	return RestaurantCreateCommandHandlerResult{restaurant.Get().GetUUID()}, nil
}
