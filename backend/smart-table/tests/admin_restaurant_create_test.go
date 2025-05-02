package smarttable_test

import (
	"testing"

	viewsAdmin "github.com/smart-table/src/views/admin/v1/restaurant"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_restaurant"
	"github.com/stretchr/testify/assert"
)

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
