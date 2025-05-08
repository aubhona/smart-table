package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
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

	userUUID, userToken, err := CreateDefaultUser()
	assert.Nil(t, err)

	employeeUserUUID, employeeToken, err := CreateUser(
		userDefaultFirstName,
		userDefaultLastName,
		employeeDefaultLogin,
		userDefaultPassword,
		employeeDefaultTgLogin,
	)
	assert.Nil(t, err)

	restaurantUUID1, err := CreateRestaurant(
		testRestaurantName1,
		userToken,
		userUUID,
	)
	assert.Nil(t, err)

	placeUUID, err := CreateDefaultPlace(userToken, userUUID, restaurantUUID1)
	assert.Nil(t, err)

	err = AddEmployee(
		employeeDefaultLogin,
		"admin",
		userToken,
		userUUID,
		placeUUID,
	)
	assert.Nil(t, err)

	restaurantUUID2, err := CreateRestaurant(
		testRestaurantName2,
		employeeToken,
		employeeUserUUID,
	)
	assert.Nil(t, err)

	response, err := viewsCodegenAdminRestaurantClient.GetAdminV1RestaurantListWithResponse(
		GetCtx(), 
		&viewsCodegenAdminRestaurant.GetAdminV1RestaurantListParams{
			UserUUID: employeeUserUUID,
			JWTToken: employeeToken,
		},
	)

	expectedRestaurantList := []viewsCodegenAdminRestaurant.RestaurantInfo{
		{
			Name: testRestaurantName2,
			UUID: restaurantUUID2,
		},
		{
			Name: testRestaurantName1,
			UUID: restaurantUUID1,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, response.JSON200.RestaurantList, expectedRestaurantList)
	assert.Equal(t, response.StatusCode(), http.StatusOK)
}
