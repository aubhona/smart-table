package smarttable_test

import (
	"net/http"
	"testing"

	"github.com/smart-table/src/domains/admin/domain"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

const (
	employeeLogin1 = "testEmployeeLogin1"
	employeeLogin2 = "testEmployeeLogin2"

	employeeTgLogin1 = "testEmployeeTgLogin1"
	employeeTgLogin2 = "testEmployeeTgLogin2"
)

func TestAdminEmployeeListHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, userToken, err := CreateDefaultUser()
	assert.Nil(t, err)

	employeeUserUUID1, employeeToken1, err := CreateUser(
		userDefaultFirstName,
		userDefaultLastName,
		employeeLogin1,
		userDefaultPassword,
		employeeTgLogin1,
	)
	assert.Nil(t, err)

	employeeUserUUID2, _, err := CreateUser(
		userDefaultFirstName,
		userDefaultLastName,
		employeeLogin2,
		userDefaultPassword,
		employeeTgLogin2,
	)
	assert.Nil(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(userToken, userUUID)
	assert.Nil(t, err)

	placeUUID, err := CreateDefaultPlace(userToken, userUUID, restaurantUUID)
	assert.Nil(t, err)

	err = AddEmployee(
		employeeLogin1,
		domain.AdminRole,
		userToken,
		userUUID,
		placeUUID,
	)
	assert.Nil(t, err)

	err = AddEmployee(
		employeeLogin2,
		domain.WaiterRole,
		userToken,
		userUUID,
		placeUUID,
	)
	assert.Nil(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceEmployeeListWithResponse(
		GetCtx(),
		&viewsCodegenAdmin.PostAdminV1PlaceEmployeeListParams{
			UserUUID: userUUID,
			JWTToken: userToken,
		},
		viewsCodegenAdmin.PostAdminV1PlaceEmployeeListJSONRequestBody{
			PlaceUUID: placeUUID,
		},
	)

	expectedEmployeeList := []viewsCodegenAdmin.EmployeeInfo{
		{
			FirstName:    userDefaultFirstName,
			LastName:     userDefaultLastName,
			Login:        employeeLogin1,
			TgLogin:      employeeTgLogin1,
			EmployeeRole: domain.AdminRole,
			UUID:         employeeUserUUID1,
		},
		{
			FirstName:    userDefaultFirstName,
			LastName:     userDefaultLastName,
			Login:        employeeLogin2,
			TgLogin:      employeeTgLogin2,
			EmployeeRole: domain.WaiterRole,
			UUID:         employeeUserUUID2,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.Equal(t, expectedEmployeeList, response.JSON200.EmployeeList)

	response, err = viewsCodegenAdminPlaceClient.PostAdminV1PlaceEmployeeListWithResponse(
		GetCtx(),
		&viewsCodegenAdmin.PostAdminV1PlaceEmployeeListParams{
			UserUUID: employeeUserUUID1,
			JWTToken: employeeToken1,
		},
		viewsCodegenAdmin.PostAdminV1PlaceEmployeeListJSONRequestBody{
			PlaceUUID: placeUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.Equal(t, expectedEmployeeList, response.JSON200.EmployeeList)
}
