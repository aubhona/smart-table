package mapper

import (
	"encoding/json"
	"time"

	defsInternalAdminPlaceDb "github.com/smart-table/src/codegen/intern/admin_place_db"
	defsInternalAdminRestaurantDb "github.com/smart-table/src/codegen/intern/admin_restaurant_db"
	defsInternalAdminUserDb "github.com/smart-table/src/codegen/intern/admin_user_db"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/utils"
)

func ConvertToPgUser(user utils.SharedRef[domain.User]) ([]byte, error) {
	pgUser := defsInternalAdminUserDb.PgUser{
		UUID:         user.Get().GetUUID(),
		Login:        user.Get().GetLogin(),
		TgID:         user.Get().GetTgID(),
		TgLogin:      user.Get().GetTgLogin(),
		ChatID:       user.Get().GetChatID(),
		FirstName:    user.Get().GetFirstName(),
		LastName:     user.Get().GetLastName(),
		PasswordHash: user.Get().GetPasswordHash(),
		CreatedAt:    user.Get().GetCreatedAt(),
		UpdatedAt:    user.Get().GetUpdatedAt(),
	}

	jsonBytes, err := json.Marshal(pgUser)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertPgUserToModel(pgResult []byte) (utils.SharedRef[domain.User], error) {
	pgUser := defsInternalAdminUserDb.PgUser{}
	err := json.Unmarshal(pgResult, &pgUser)

	if err != nil {
		return utils.SharedRef[domain.User]{}, err
	}

	return domain.RestoreUser(
		pgUser.UUID,
		pgUser.Login,
		pgUser.TgID,
		pgUser.TgLogin,
		pgUser.ChatID,
		pgUser.FirstName,
		pgUser.LastName,
		pgUser.PasswordHash,
		pgUser.CreatedAt,
		pgUser.UpdatedAt,
	), nil
}

func ConvertToPgRestaurant(restaurant utils.SharedRef[domain.Restaurant]) ([]byte, error) {
	pgRestaurant := defsInternalAdminRestaurantDb.PgRestaurant{
		UUID:      restaurant.Get().GetUUID(),
		OwnerUUID: restaurant.Get().GetOwnerUUID(),
		Name:      restaurant.Get().GetName(),
		CreatedAt: restaurant.Get().GetCreatedAt(),
		UpdatedAt: restaurant.Get().GetUpdatedAt(),
	}

	jsonBytes, err := json.Marshal(pgRestaurant)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertPgRestaurantToModel(pgResult []byte) (utils.SharedRef[domain.Restaurant], error) {
	pgRestaurant := defsInternalAdminRestaurantDb.PgRestaurant{}
	err := json.Unmarshal(pgResult, &pgRestaurant)

	if err != nil {
		return utils.SharedRef[domain.Restaurant]{}, err
	}

	return domain.RestoreRestaurant(
		pgRestaurant.UUID,
		pgRestaurant.OwnerUUID,
		pgRestaurant.Name,
		pgRestaurant.CreatedAt,
		pgRestaurant.UpdatedAt,
	), nil
}

func ConvertPgRestaurantListToModelList(pgResults []byte) ([]utils.SharedRef[domain.Restaurant], error) {
	var pgRestaurantList []defsInternalAdminRestaurantDb.PgRestaurant
	err := json.Unmarshal(pgResults, &pgRestaurantList)

	if err != nil {
		return []utils.SharedRef[domain.Restaurant]{}, err
	}

	restaurantModelList := make([]utils.SharedRef[domain.Restaurant], 0, len(pgRestaurantList))

	for _, pgRestaurant := range pgRestaurantList {
		restaurant := domain.RestoreRestaurant(
			pgRestaurant.UUID,
			pgRestaurant.OwnerUUID,
			pgRestaurant.Name,
			pgRestaurant.CreatedAt,
			pgRestaurant.UpdatedAt,
		)

		restaurantModelList = append(restaurantModelList, restaurant)
	}

	return restaurantModelList, nil
}

func ConvertToPgPlace(place utils.SharedRef[domain.Place]) ([]byte, error) {
	pgPlace := defsInternalAdminPlaceDb.PgPlace{
		UUID:           place.Get().GetUUID(),
		RestaurantUUID: place.Get().GetRestauranUUID(),
		Address:        place.Get().GetAddress(),
		TableCount:     place.Get().GetTableCount(),
		OpeningTime:    place.Get().GetOpeningTime().Format("15:04"),
		ClosingTime:    place.Get().GetClosingTime().Format("15:04"),
		CreatedAt:      place.Get().GetCreatedAt(),
		UpdatedAt:      place.Get().GetUpdatedAt(),
	}

	jsonBytes, err := json.Marshal(pgPlace)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertPgPlaceToModel(pgResult []byte) (utils.SharedRef[domain.Place], error) {
	pgPlace := defsInternalAdminPlaceDb.PgPlace{}
	err := json.Unmarshal(pgResult, &pgPlace)

	if err != nil {
		return utils.SharedRef[domain.Place]{}, err
	}

	openingTime, err := time.Parse("15:04", pgPlace.OpeningTime)
	if err != nil {
		panic(err)
	}

	closingTime, err := time.Parse("15:04", pgPlace.OpeningTime)
	if err != nil {
		panic(err)
	}

	return domain.RestorePlace(
		pgPlace.UUID,
		pgPlace.RestaurantUUID,
		pgPlace.Address,
		pgPlace.TableCount,
		openingTime,
		closingTime,
		pgPlace.CreatedAt,
		pgPlace.UpdatedAt,
	), nil
}
