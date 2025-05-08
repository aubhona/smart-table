package mapper

import (
	"encoding/json"
	"time"

	"github.com/samber/lo"
	defsinternalAdminDishDb "github.com/smart-table/src/codegen/intern/admin_dish_db"
	defsinternalAdminEmployeeDb "github.com/smart-table/src/codegen/intern/admin_employee_db"
	defsInternalAdminPlaceDb "github.com/smart-table/src/codegen/intern/admin_place_db"
	defsInternalAdminRestaurantDb "github.com/smart-table/src/codegen/intern/admin_restaurant_db"
	defsInternalAdminUserDb "github.com/smart-table/src/codegen/intern/admin_user_db"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/utils"
)

type PgPlaceAggregate struct {
	RestaurantAggregate PgRestaurantAggregate            `json:"restaurant"`
	Place               defsInternalAdminPlaceDb.PgPlace `json:"place"`
	Employees           []PgEmployeeAggregate            `json:"employees"`
}

type PgEmployeeAggregate struct {
	Employee defsinternalAdminEmployeeDb.PgEmployee `json:"employee"`
	User     defsInternalAdminUserDb.PgUser         `json:"user"`
}

type PgRestaurantAggregate struct {
	Restaurant defsInternalAdminRestaurantDb.PgRestaurant `json:"restaurant"`
	Dishes     []defsinternalAdminDishDb.PgDish           `json:"dishes"`
	Owner      defsInternalAdminUserDb.PgUser             `json:"owner"`
}

func restoreUser(user *defsInternalAdminUserDb.PgUser) utils.SharedRef[domain.User] {
	return domain.RestoreUser(
		user.UUID,
		user.Login,
		user.TgID,
		user.TgLogin,
		user.ChatID,
		user.FirstName,
		user.LastName,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	)
}

func restoreDish(dish *defsinternalAdminDishDb.PgDish) utils.SharedRef[domain.Dish] {
	return domain.RestoreDish(
		dish.UUID,
		dish.RestaurantUUID,
		dish.Name,
		dish.Description,
		dish.PictureKey,
		dish.Category,
		dish.Calories,
		dish.Weight,
		dish.CreatedAt,
		dish.UpdatedAt,
	)
}

func restoreEmployee(
	employee *defsinternalAdminEmployeeDb.PgEmployee,
	user utils.SharedRef[domain.User],
) utils.SharedRef[domain.Employee] {
	return domain.RestoreEmployee(
		user,
		employee.PlaceUUID,
		string(employee.Role),
		employee.Active,
		employee.CreatedAt,
		employee.UpdatedAt,
	)
}

func restoreRestaurant(
	restaurant *defsInternalAdminRestaurantDb.PgRestaurant,
	owner utils.SharedRef[domain.User],
	dishes []utils.SharedRef[domain.Dish]) utils.SharedRef[domain.Restaurant] {
	return domain.RestoreRestaurant(
		restaurant.UUID,
		owner,
		dishes,
		restaurant.Name,
		restaurant.CreatedAt,
		restaurant.UpdatedAt,
	)
}

func restorePlace(
	place *defsInternalAdminPlaceDb.PgPlace,
	restaurant utils.SharedRef[domain.Restaurant],
	employees []utils.SharedRef[domain.Employee],
	openingTime,
	closingTime time.Time,
) utils.SharedRef[domain.Place] {
	return domain.RestorePlace(
		place.UUID,
		restaurant,
		employees,
		place.Address,
		place.TableCount,
		openingTime,
		closingTime,
		place.CreatedAt,
		place.UpdatedAt,
	)
}

func restoreFromPgEmployeeAggregate(
	pgEmployeeAggregate *PgEmployeeAggregate,
) utils.SharedRef[domain.Employee] {
	user := restoreUser(&pgEmployeeAggregate.User)

	return restoreEmployee(&pgEmployeeAggregate.Employee, user)
}

func restoreFromPgRestaurantAggregate(
	pgRestaurantAggregate *PgRestaurantAggregate,
) utils.SharedRef[domain.Restaurant] {
	dishes := lo.Map(pgRestaurantAggregate.Dishes, func(dish defsinternalAdminDishDb.PgDish, _ int) utils.SharedRef[domain.Dish] {
		return restoreDish(&dish)
	})
	owner := restoreUser(&pgRestaurantAggregate.Owner)

	return restoreRestaurant(&pgRestaurantAggregate.Restaurant, owner, dishes)
}

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

	return restoreUser(&pgUser), nil
}

func ConvertToPgDishes(dishes []utils.SharedRef[domain.Dish]) ([]byte, error) {
	pgDishes := lo.Map(dishes, func(dish utils.SharedRef[domain.Dish], _ int) defsinternalAdminDishDb.PgDish {
		return defsinternalAdminDishDb.PgDish{
			UUID:           dish.Get().GetUUID(),
			RestaurantUUID: dish.Get().GetRestaurantUUID(),
			Name:           dish.Get().GetName(),
			Description:    dish.Get().GetDescription(),
			PictureKey:     dish.Get().GetPictureKey(),
			Category:       dish.Get().GetCategory(),
			Calories:       dish.Get().GetCalories(),
			Weight:         dish.Get().GetWeight(),
			CreatedAt:      dish.Get().GetCreatedAt(),
			UpdatedAt:      dish.Get().GetUpdatedAt(),
		}
	})

	return json.Marshal(pgDishes)
}

func ConvertToPgEmployees(employees []utils.SharedRef[domain.Employee]) ([]byte, error) {
	pgEmployees := lo.Map(employees, func(employee utils.SharedRef[domain.Employee], _ int) defsinternalAdminEmployeeDb.PgEmployee {
		return defsinternalAdminEmployeeDb.PgEmployee{
			UserUUID:  employee.Get().GetUser().Get().GetUUID(),
			PlaceUUID: employee.Get().GetPlaceUUID(),
			Role:      defsinternalAdminEmployeeDb.PgEmployeeRole(employee.Get().GetRole()),
			Active:    employee.Get().GetActive(),
			CreatedAt: employee.Get().GetCreatedAt(),
			UpdatedAt: employee.Get().GetUpdatedAt(),
		}
	})

	return json.Marshal(pgEmployees)
}

func ConvertToPgRestaurant(restaurant utils.SharedRef[domain.Restaurant]) ([]byte, error) {
	pgRestaurant := defsInternalAdminRestaurantDb.PgRestaurant{
		UUID:      restaurant.Get().GetUUID(),
		OwnerUUID: restaurant.Get().GetOwner().Get().GetUUID(),
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
	pgRestaurantAggregate := PgRestaurantAggregate{}
	err := json.Unmarshal(pgResult, &pgRestaurantAggregate)

	if err != nil {
		return utils.SharedRef[domain.Restaurant]{}, err
	}

	return restoreFromPgRestaurantAggregate(&pgRestaurantAggregate), nil
}

func ConvertPgRestaurantsToModel(pgResults [][]byte) ([]utils.SharedRef[domain.Restaurant], error) {
	restaurants := make([]utils.SharedRef[domain.Restaurant], 0, len(pgResults))

	for _, pgResult := range pgResults {
		restaurant, err := ConvertPgRestaurantToModel(pgResult)
		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}

func ConvertToPgPlace(place utils.SharedRef[domain.Place]) ([]byte, error) {
	pgPlace := defsInternalAdminPlaceDb.PgPlace{
		UUID:           place.Get().GetUUID(),
		RestaurantUUID: place.Get().GetRestaurant().Get().GetUUID(),
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

func ConvertPgPlaceToModel(pgPlaceAggregateResult []byte) (utils.SharedRef[domain.Place], error) {
	pgPlaceAggregate := PgPlaceAggregate{}
	err := json.Unmarshal(pgPlaceAggregateResult, &pgPlaceAggregate)

	if err != nil {
		return utils.SharedRef[domain.Place]{}, err
	}

	openingTime, err := time.Parse("15:04:05", pgPlaceAggregate.Place.OpeningTime)
	if err != nil {
		return utils.SharedRef[domain.Place]{}, err
	}

	closingTime, err := time.Parse("15:04:05", pgPlaceAggregate.Place.ClosingTime)
	if err != nil {
		return utils.SharedRef[domain.Place]{}, err
	}

	restaurant := restoreFromPgRestaurantAggregate(&pgPlaceAggregate.RestaurantAggregate)
	employees := lo.Map(pgPlaceAggregate.Employees, func(employee PgEmployeeAggregate, _ int) utils.SharedRef[domain.Employee] {
		return restoreFromPgEmployeeAggregate(&employee)
	})

	return restorePlace(&pgPlaceAggregate.Place, restaurant, employees, openingTime, closingTime), nil
}

func ConvertPgPlacesToModel(pgPlacesAggregateResult [][]byte) ([]utils.SharedRef[domain.Place], error) {
	places := make([]utils.SharedRef[domain.Place], 0, len(pgPlacesAggregateResult))

	for _, pgPlace := range pgPlacesAggregateResult {
		place, err := ConvertPgPlaceToModel(pgPlace)
		if err != nil {
			return nil, err
		}

		places = append(places, place)
	}

	return places, nil
}
