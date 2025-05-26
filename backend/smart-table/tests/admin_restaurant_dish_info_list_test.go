package smarttable_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	viewsCodegenAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"

	"github.com/stretchr/testify/assert"
)

func TestAdminRestaurantDishInfoListHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.NoError(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.NoError(t, err)

	_, err = CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenAdminRestaurantClient.PostAdminV1RestaurantDishInfoListWithResponse(
		GetCtx(),
		&viewsCodegenAdminRestaurant.PostAdminV1RestaurantDishInfoListParams{
			JWTToken: token,
			UserUUID: userUUID,
		},
		viewsCodegenAdminRestaurant.PostAdminV1RestaurantDishInfoListJSONRequestBody{
			RestaurantUUID: restaurantUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)

	response.JSON200.DishList[0].ID, err = uuid.Parse("149ebb8c-ce00-4492-b468-12d16b89ed36")
	assert.NoError(t, err)

	response.JSON200.DishList[0].PictureKey = "149ebb8c-ce00-4492-b468-12d16b89ed36"

	actualJSON, err := json.Marshal(response.JSON200.DishList)
	assert.NoError(t, err)

	expectedJSON := `
[
	{
		"calories":100, 
		"category":"some_cat", 
		"description":"some_desc", 
		"id":"149ebb8c-ce00-4492-b468-12d16b89ed36", 
		"name":"test_dish", 
		"picture_key":"149ebb8c-ce00-4492-b468-12d16b89ed36", 
		"weight":100
	}
]
`

	assert.JSONEq(t, expectedJSON, string(actualJSON))
}
