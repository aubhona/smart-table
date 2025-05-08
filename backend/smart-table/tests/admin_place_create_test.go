package smarttable_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

var viewsCodegenAdminPlaceClient, _ = viewsCodegenAdminPlace.NewClientWithResponses(GetBasePath())

func CreateDefaultPlace(token string, userUUID, restaurantUUID uuid.UUID) (uuid.UUID, error) {
	return CreatePlace(
		userUUID,
		restaurantUUID,
		token,
		"testAddress",
		"12:00",
		"13:00",
		10,
	)
}

func CreatePlace(userUUID, restaurantUUID uuid.UUID, token, address, openingTime, closingTime string, tableCount int) (uuid.UUID, error) {
	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceCreateWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceCreateParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceCreateJSONRequestBody{
			RestaurantUUID: restaurantUUID,
			Address:        address,
			TableCount:     tableCount,
			OpeningTime:    openingTime,
			ClosingTime:    closingTime,
		},
	)

	if err != nil {
		return uuid.Nil, err
	}

	if response.JSON200 == nil || response.StatusCode() != http.StatusOK {
		return uuid.Nil, errors.New("response is not a 200")
	}

	return response.JSON200.PlaceUUID, nil
}

func TestAdminPlaceCreateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.Nil(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.Nil(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceCreateWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceCreateParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceCreateJSONRequestBody{
			RestaurantUUID: restaurantUUID,
			Address:        "testAddress",
			TableCount:     1,
			OpeningTime:    "13:00",
			ClosingTime:    "14:00",
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
}
