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

type PlaceCreateCommandHandlerResult struct {
	PlaceUUID uuid.UUID
}

type PlaceCreateCommandHandler struct {
	restaurantRepository domain.RestaurantRepository
	placeRepository      domain.PlaceRepository
	uuidGenerator        *domainServices.UUIDGenerator
}

func NewPlaceCreateCommandHandler(
	restaurantRepository domain.RestaurantRepository,
	placeRepository domain.PlaceRepository,
	uuidGenerator *domainServices.UUIDGenerator,
) *PlaceCreateCommandHandler {
	return &PlaceCreateCommandHandler{
		restaurantRepository,
		placeRepository,
		uuidGenerator,
	}
}

func (handler *PlaceCreateCommandHandler) Handle(
	placeCreateCommand *PlaceCreateCommand,
) (PlaceCreateCommandHandlerResult, error) {
	ctx := context.Background()

	restaurant, err := handler.restaurantRepository.FindRestaurantByUUID(ctx, placeCreateCommand.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while finding restaurant by uuid: %v", err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	if restaurant.Get().GetOwnerUUID() != placeCreateCommand.OwnerUUID {
		return PlaceCreateCommandHandlerResult{}, appErrors.RestaurantAccessDenied{
			UserUUID:       placeCreateCommand.OwnerUUID,
			RestaurantUUID: placeCreateCommand.RestaurantUUID,
		}
	}

	isExist, err := handler.placeRepository.CheckAddressExist(ctx, placeCreateCommand.Address, placeCreateCommand.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while checking place address existence: %v", err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	if isExist {
		return PlaceCreateCommandHandlerResult{}, appErrors.PlaceAddressExists{
			Address:        placeCreateCommand.Address,
			RestaurantUUID: placeCreateCommand.RestaurantUUID,
		}
	}

	place, err := domain.NewPlace(
		placeCreateCommand.RestaurantUUID,
		placeCreateCommand.Address,
		placeCreateCommand.TableCount,
		placeCreateCommand.OpeningTime,
		placeCreateCommand.ClosingTime,
		handler.uuidGenerator,
	)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while new place creating: %v", err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	tx, err := handler.placeRepository.Begin(ctx)
	if err != nil {
		return PlaceCreateCommandHandlerResult{}, err
	}

	defer func(placeRepository domain.PlaceRepository, ctx context.Context, tx pgx.Tx) {
		_ = placeRepository.Commit(ctx, tx)
	}(handler.placeRepository, ctx, tx)

	err = handler.placeRepository.Save(ctx, tx, place)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while place saving: %v", err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	return PlaceCreateCommandHandlerResult{place.Get().GetUUID()}, nil
}
