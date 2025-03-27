package smarttable_test

import (
	"testing"

	viewsAdmin "github.com/smart-table/src/views/admin/v1/user"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_user"
	"github.com/stretchr/testify/assert"
)

func TestAdminSignInHappyPath(t *testing.T) {
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

	handler := viewsAdmin.AdminV1UserHandler{}

	response, err := handler.PostAdminV1UserSignIn(GetCtx(), viewsCodegenAdmin.PostAdminV1UserSignInRequestObject{
		Body: &viewsCodegenAdmin.AdminV1UserSignInRequest{
			Login:    "testLogin",
			Password: "testPassword",
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)

	responseObj, ok := response.(viewsCodegenAdmin.PostAdminV1UserSignIn200JSONResponse)
	assert.True(t, ok)
	assert.Equal(t, responseObj.Body.UserUUID, id)
}
