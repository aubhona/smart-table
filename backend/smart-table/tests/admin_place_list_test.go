package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

const (
	testPlaceAddress1 = "testPlaceAddress1"
	testPlaceAddress2 = "testPlaceAddress2"

	testOpeningTime = "12:00"
	testClosingTime = "13:00"

	testTableCount = 10
)

func TestAdminPlaceListHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, userToken, err := CreateDefaultUser()
	assert.NoError(t, err)

	employeeUserUUID, employeeToken, err := CreateUser(
		userDefaultFirstName,
		userDefaultLastName,
		employeeDefaultLogin,
		userDefaultPassword,
		employeeDefaultTgLogin,
	)
	assert.Nil(t, err)

	restaurantUUID, err := CreateRestaurant(
		testRestaurantName1,
		userToken,
		userUUID,
	)
	assert.Nil(t, err)

	placeUUID1, err := CreatePlace(
		userUUID,
		restaurantUUID,
		userToken,
		testPlaceAddress1,
		testOpeningTime,
		testClosingTime,
		testTableCount,
	)
	assert.NoError(t, err)

	placeUUID2, err := CreatePlace(
		userUUID,
		restaurantUUID,
		userToken,
		testPlaceAddress2,
		testOpeningTime,
		testClosingTime,
		testTableCount,
	)
	assert.NoError(t, err)

	err = AddEmployee(
		employeeDefaultLogin,
		"admin",
		userToken,
		userUUID,
		placeUUID1,
	)
	assert.Nil(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceListWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceListParams{
			UserUUID: userUUID,
			JWTToken: userToken,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceListJSONRequestBody{
			RestaurantUUID: restaurantUUID,
		},
	)

	expectedPlaceList := []viewsCodegenAdminPlace.PlaceInfo{
		{
			Address:     testPlaceAddress1,
			OpeningTime: testOpeningTime,
			ClosingTime: testClosingTime,
			TableCount:  testTableCount,
			UUID:        placeUUID1,
		},
		{
			Address:     testPlaceAddress2,
			OpeningTime: testOpeningTime,
			ClosingTime: testClosingTime,
			TableCount:  testTableCount,
			UUID:        placeUUID2,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.Equal(t, expectedPlaceList, response.JSON200.PlaceList)

	response, err = viewsCodegenAdminPlaceClient.PostAdminV1PlaceListWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceListParams{
			UserUUID: employeeUserUUID,
			JWTToken: employeeToken,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceListJSONRequestBody{
			RestaurantUUID: restaurantUUID,
		},
	)

	expectedPlaceList = []viewsCodegenAdminPlace.PlaceInfo{
		{
			Address:     testPlaceAddress1,
			OpeningTime: testOpeningTime,
			ClosingTime: testClosingTime,
			TableCount:  testTableCount,
			UUID:        placeUUID1,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.Equal(t, expectedPlaceList, response.JSON200.PlaceList)
}
