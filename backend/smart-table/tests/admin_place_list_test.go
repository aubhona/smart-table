package smarttable_test

import (
	"testing"

	viewsAdmin "github.com/smart-table/src/views/admin/v1/place"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_place"
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

	userUUID, err := CreateDefaultUser()
	assert.Nil(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(
		userUUID,
	)
	assert.Nil(t, err)

	placeUUID1, err := CreatePlace(
		userUUID,
		restaurantUUID,
		testPlaceAddress1,
		testOpeningTime,
		testClosingTime,
		testTableCount,
	)
	assert.Nil(t, err)

	placeUUID2, err := CreatePlace(
		userUUID,
		restaurantUUID,
		testPlaceAddress2,
		testOpeningTime,
		testClosingTime,
		testTableCount,
	)
	assert.Nil(t, err)

	handler := viewsAdmin.AdminV1PlaceHandler{}

	response, err := handler.PostAdminV1PlaceList(GetCtx(), viewsCodegenAdmin.PostAdminV1PlaceListRequestObject{
		Params: viewsCodegenAdmin.PostAdminV1PlaceListParams{
			UserUUID: userUUID,
		},
		Body: &viewsCodegenAdmin.AdminV1PlaceListRequest{
			RestaurantUUID: restaurantUUID,
		},
	})

	assert.NoError(t, err)

	expectedResponse := viewsCodegenAdmin.PostAdminV1PlaceList200JSONResponse{
		PlaceList: []viewsCodegenAdmin.PlaceInfo{
			{
				Address: testPlaceAddress1,
				OpeningTime: testOpeningTime, 
				ClosingTime: testClosingTime,
				TableCount: testTableCount,
				UUID: placeUUID1,
			},
			{
				Address: testPlaceAddress2,
				OpeningTime: testOpeningTime, 
				ClosingTime: testClosingTime,
				TableCount: testTableCount,
				UUID: placeUUID2,
			},
		},
	}

	assert.Equal(t, response, expectedResponse)

	_, ok := response.(viewsCodegenAdmin.PostAdminV1PlaceList200JSONResponse)
	assert.True(t, ok)
}
