package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

const testEmployeeLogin = "testEmployeeLogin"

func TestAdminEmployeeAddHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.Nil(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.Nil(t, err)

	placeUUID, err := CreateDefaultPlace(token, userUUID, restaurantUUID)
	assert.Nil(t, err)

	_, _, err = CreateUser(
		"testFisrtName",
		"testLastName",
		testEmployeeLogin,
		"testPassword",
		"testEmployeeTgLogin",
	)
	assert.Nil(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceEmployeeAddWithResponse(
		GetCtx(),
		&viewsCodegenAdmin.PostAdminV1PlaceEmployeeAddParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdmin.PostAdminV1PlaceEmployeeAddJSONRequestBody{
			EmployeeLogin: testEmployeeLogin,
			EmployeeRole:  "admin",
			PlaceUUID:     placeUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode())
}
