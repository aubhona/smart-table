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
	restaurant, err := handler.restaurantRepository.FindRestaurant(placeCreateCommand.RestaurantUUID)
	if err != nil {
		logging.GetLogger().Error("error while finding restaurant by uuid", zap.Error(err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	if restaurant.Get().GetOwner().Get().GetUUID() != placeCreateCommand.OwnerUUID {
		logging.GetLogger().Error("restaurant access denied",
			zap.String("user_uuid", placeCreateCommand.OwnerUUID.String()),
			zap.String("restaurant_uuid", placeCreateCommand.RestaurantUUID.String()))

		return PlaceCreateCommandHandlerResult{}, appErrors.RestaurantAccessDenied{
			UserUUID:       placeCreateCommand.OwnerUUID,
			RestaurantUUID: placeCreateCommand.RestaurantUUID,
		}
	}

	_, err = handler.placeRepository.FindPlaceByAddress(placeCreateCommand.Address, placeCreateCommand.RestaurantUUID)
	if err == nil {
		logging.GetLogger().Error("place address already exists",
			zap.String("address", placeCreateCommand.Address),
			zap.String("restaurant_uuid", placeCreateCommand.RestaurantUUID.String()))

		return PlaceCreateCommandHandlerResult{}, appErrors.PlaceAddressExists{
			Address:        placeCreateCommand.Address,
			RestaurantUUID: placeCreateCommand.RestaurantUUID,
		}
	}

	if !utils.IsTheSameErrorType[domainErrors.PlaceNotFoundByAddress](err) {
		logging.GetLogger().Error("error while checking place address existence", zap.Error(err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	place, err := domain.NewPlace(
		restaurant,
		placeCreateCommand.Address,
		placeCreateCommand.TableCount,
		placeCreateCommand.OpeningTime,
		placeCreateCommand.ClosingTime,
		handler.uuidGenerator,
	)
	if err != nil {
		logging.GetLogger().Error("error while creating place", zap.Error(err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	tx, err := handler.placeRepository.Begin()
	if err != nil {
		logging.GetLogger().Error("error while beginning transaction", zap.Error(err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	defer utils.Rollback(handler.placeRepository, tx)

	err = handler.placeRepository.Save(tx, place)
	if err != nil {
		logging.GetLogger().Error("error while saving place", zap.Error(err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	err = handler.placeRepository.Commit(tx)
	if err != nil {
		logging.GetLogger().Error("error while committing transaction", zap.Error(err))
		return PlaceCreateCommandHandlerResult{}, err
	}

	return PlaceCreateCommandHandlerResult{place.Get().GetUUID()}, nil
}
