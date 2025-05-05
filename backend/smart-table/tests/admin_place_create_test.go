package smarttable_test

import (
	"testing"

	viewsAdmin "github.com/smart-table/src/views/admin/v1/place"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

func TestAdminPlaceCreateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, err := CreateUser(
		"testFisrtName",
		"testLastName",
		"testLogin",
		"testPassword",
		"testTgLogin",
	)
	assert.Nil(t, err)

	restaurantUUID, err := CreateRestaurant(
		"testRestaurantName",
		userUUID,
	)
	assert.Nil(t, err)

	handler := viewsAdmin.AdminV1PlaceHandler{}

	response, err := handler.PostAdminV1PlaceCreate(GetCtx(), viewsCodegenAdmin.PostAdminV1PlaceCreateRequestObject{
		Params: viewsCodegenAdmin.PostAdminV1PlaceCreateParams{
			UserUUID: userUUID,
		},
		Body: &viewsCodegenAdmin.AdminV1PlaceCreateRequest{
			RestaurantUUID: restaurantUUID,
			Address:        "testAddress",
			TableCount:     1,
			OpeningTime:    "13:00",
			ClosingTime:    "14:00",
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)

	_, ok := response.(viewsCodegenAdmin.PostAdminV1PlaceCreate200JSONResponse)
	assert.True(t, ok)
}
