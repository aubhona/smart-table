package smarttable_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	viewsAdmin "github.com/smart-table/src/views/admin/v1/restaurant"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_restaurant"
	"github.com/stretchr/testify/assert"
)

func CreateRestaurant(name string, ownerUUID uuid.UUID) (uuid.UUID, error) {
	handler := viewsAdmin.AdminV1RestaurantHandler{}
	response, err := handler.PostAdminV1RestaurantCreate(GetCtx(), viewsCodegenAdmin.PostAdminV1RestaurantCreateRequestObject{
		Params: viewsCodegenAdmin.PostAdminV1RestaurantCreateParams{
			UserUUID: ownerUUID,
		},
		Body: &viewsCodegenAdmin.AdminV1RestaurantCreateRequest{
			Name: name,
		},
	})

	if err != nil {
		return uuid.Nil, err
	}

	responseObj, ok := response.(viewsCodegenAdmin.PostAdminV1RestaurantCreate200JSONResponse)
	if !ok {
		return uuid.Nil, errors.New("response is not a PostAdminV1UserSignUp200JSONResponse")
	}

	return responseObj.RestaurantUUID, nil
}

func TestAdminRestaurantCreateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	id, err := CreateUser(
		"testFisrtName",
		"testLastName",
		"testLogin",
		"testPassword",
		"testTgLogin",
	)
	assert.Nil(t, err)

	handler := viewsAdmin.AdminV1RestaurantHandler{}

	response, err := handler.PostAdminV1RestaurantCreate(GetCtx(), viewsCodegenAdmin.PostAdminV1RestaurantCreateRequestObject{
		Params: viewsCodegenAdmin.PostAdminV1RestaurantCreateParams{
			UserUUID: id,
		},
		Body: &viewsCodegenAdmin.AdminV1RestaurantCreateRequest{
			Name: "testName",
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)

	_, ok := response.(viewsCodegenAdmin.PostAdminV1RestaurantCreate200JSONResponse)
	assert.True(t, ok)
}
