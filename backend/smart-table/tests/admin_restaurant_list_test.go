package smarttable_test

import (
	"testing"

	viewsAdmin "github.com/smart-table/src/views/admin/v1/restaurant"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_restaurant"
	"github.com/stretchr/testify/assert"
)

const (
	testRestaurantName1 = "testRestaurantName1"
	testRestaurantName2 = "testRestaurantName2"
)

func TestAdminRestaurantListHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, err := CreateDefaultUser()
	assert.Nil(t, err)

	restaurantUUID1, err := CreateRestaurant(
		testRestaurantName1,
		userUUID,
	)
	assert.Nil(t, err)

	restaurantUUID2, err := CreateRestaurant(
		testRestaurantName2,
		userUUID,
	)
	assert.Nil(t, err)

	handler := viewsAdmin.AdminV1RestaurantHandler{}

	response, err := handler.GetAdminV1RestaurantList(GetCtx(), viewsCodegenAdmin.GetAdminV1RestaurantListRequestObject{
		Params: viewsCodegenAdmin.GetAdminV1RestaurantListParams{
			UserUUID: userUUID,
		},
	})

	assert.NoError(t, err)

	expectedResponse := viewsCodegenAdmin.GetAdminV1RestaurantList200JSONResponse{
		RestaurantList: []viewsCodegenAdmin.RestaurantInfo{
			{
				Name: testRestaurantName1,
				UUID: restaurantUUID1,
			},
			{
				Name: testRestaurantName2,
				UUID: restaurantUUID2,
			},
		},
	}

	assert.Equal(t, response, expectedResponse)

	_, ok := response.(viewsCodegenAdmin.GetAdminV1RestaurantList200JSONResponse)
	assert.True(t, ok)
}
