package smarttable_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

const (
	employeeDefaultLogin   = "testEmployeeLogin"
	employeeDefaultTgLogin = "testEmployeeTgLogin"
)

func AddEmployee(employeeLogin, employeeRole, token string, userUUID, placeUUID uuid.UUID) error {
	resp, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceEmployeeAddWithResponse(
		GetCtx(),
		&viewsCodegenAdmin.PostAdminV1PlaceEmployeeAddParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdmin.PostAdminV1PlaceEmployeeAddJSONRequestBody{
			EmployeeLogin: employeeLogin,
			EmployeeRole:  viewsCodegenAdmin.Role(employeeRole),
			PlaceUUID:     placeUUID,
		},
	)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected response status: %d", resp.StatusCode())
	}

	return nil
}

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
		userDefaultFirstName,
		userDefaultLastName,
		employeeDefaultLogin,
		userDefaultPassword,
		employeeDefaultTgLogin,
	)
	assert.Nil(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceEmployeeAddWithResponse(
		GetCtx(),
		&viewsCodegenAdmin.PostAdminV1PlaceEmployeeAddParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdmin.PostAdminV1PlaceEmployeeAddJSONRequestBody{
			EmployeeLogin: employeeDefaultLogin,
			EmployeeRole:  "admin",
			PlaceUUID:     placeUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode())
}
