package smarttable_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

func CreateDefaultPlaceMenuDish(token string, userUUID, placeUUID, dishUUID uuid.UUID) (uuid.UUID, error) {
	return CreatePlaceMenuDish(
		token,
		userUUID,
		placeUUID,
		dishUUID,
		"123.13",
	)
}

func CreatePlaceMenuDish(token string, userUUID, placeUUID, dishUUID uuid.UUID, price string) (uuid.UUID, error) {
	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceMenuDishCreateWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishCreateParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishCreateJSONRequestBody{
			PlaceUUID: placeUUID,
			DishUUID:  dishUUID,
			Price:     price,
		},
	)

	if err != nil {
		return uuid.Nil, err
	}

	if response.JSON200 == nil || response.StatusCode() != http.StatusOK {
		return uuid.Nil, errors.New("response is not a 200")
	}

	return response.JSON200.MenuDishUUID, nil
}

func TestAdminPlaceMenuDishCreateHappyPath(t *testing.T) {
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

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceMenuDishCreateWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishCreateParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishCreateJSONRequestBody{
			PlaceUUID: placeUUID,
			DishUUID:  dishUUID,
			Price:     "123.123",
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
}

func TestAdminPlaceMenuDishCreatePlaceNotFound(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.NoError(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.NoError(t, err)

	dishUUID, err := CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceMenuDishCreateWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishCreateParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishCreateJSONRequestBody{
			PlaceUUID: uuid.New(),
			DishUUID:  dishUUID,
			Price:     "123.123",
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode())
	assert.NotNil(t, response.JSON404)
}
