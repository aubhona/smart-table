package smarttable_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/smart-table/src/domains/admin/infra/pg/mapper"
	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

func TestAdminPlaceMenuDishDeleteHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.NoError(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.NoError(t, err)

	placeUUID, err := CreateDefaultPlace(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	dishUUID, err := CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	menuDishUUID, err := CreateDefaultPlaceMenuDish(token, userUUID, placeUUID, dishUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceMenuDishDeleteWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishDeleteParams{
			JWTToken: token,
			UserUUID: userUUID,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishDeleteJSONRequestBody{
			PlaceUUID:    placeUUID,
			MenuDishUUID: menuDishUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode())

	pgPlaces, err := GetAdminQueries().FetchPlacesByUUID(GetCtx(), []uuid.UUID{placeUUID})
	assert.NoError(t, err)

	place := mapper.PgPlaceAggregate{}
	err = json.Unmarshal(pgPlaces[0], &place)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(place.MenuDishes))
}
