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

	userUUID, token, err := CreateDefaultUser()
	assert.NoError(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(
		token,
		userUUID,
	)
	assert.NoError(t, err)

	placeUUID1, err := CreatePlace(
		userUUID,
		restaurantUUID,
		token,
		testPlaceAddress1,
		testOpeningTime,
		testClosingTime,
		testTableCount,
	)
	assert.NoError(t, err)

	placeUUID2, err := CreatePlace(
		userUUID,
		restaurantUUID,
		token,
		testPlaceAddress2,
		testOpeningTime,
		testClosingTime,
		testTableCount,
	)
	assert.NoError(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceListWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceListParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceListJSONRequestBody{
			RestaurantUUID: restaurantUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)

	expectedResponse := viewsCodegenAdminPlace.PostAdminV1PlaceList200JSONResponse{
		PlaceList: []viewsCodegenAdminPlace.PlaceInfo{
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
		},
	}

	assert.Equal(t, expectedResponse.PlaceList, response.JSON200.PlaceList)
}
