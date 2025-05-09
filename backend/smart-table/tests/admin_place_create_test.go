package smarttable_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

const (
	defaultAddress     = "testAddress"
	defaultOpeningTime = "12:00"
	defaultClosingTime = "13:00"
	defaultTableCount  = 10
)

var viewsCodegenAdminPlaceClient, _ = viewsCodegenAdminPlace.NewClientWithResponses(GetBasePath())

func CreateDefaultPlace(token string, userUUID, restaurantUUID uuid.UUID) (uuid.UUID, error) {
	return CreatePlace(
		userUUID,
		restaurantUUID,
		token,
		defaultAddress,
		defaultOpeningTime,
		defaultClosingTime,
		defaultTableCount,
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
			Address:        defaultAddress,
			TableCount:     defaultTableCount,
			OpeningTime:    defaultOpeningTime,
			ClosingTime:    defaultClosingTime,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
}

func TestAdminPlaceCreateRestaurantNotFound(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.Nil(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceCreateWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceCreateParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceCreateJSONRequestBody{
			RestaurantUUID: uuid.New(),
			Address:        defaultAddress,
			TableCount:     defaultTableCount,
			OpeningTime:    defaultOpeningTime,
			ClosingTime:    defaultClosingTime,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode())
	assert.NotNil(t, response.JSON404)
}
