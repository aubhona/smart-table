package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_user"
	"github.com/stretchr/testify/assert"
)

func TestAdminUserSignInHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, _, err := CreateDefaultUser()
	assert.NoError(t, err)

	resp, err := viewsCodegenAdminClient.PostAdminV1UserSignInWithResponse(
		GetCtx(),
		viewsCodegenAdmin.PostAdminV1UserSignInJSONRequestBody{
			Login:    "testLogin",
			Password: "testPassword",
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.NotNil(t, resp.JSON200)
	assert.Equal(t, resp.JSON200.UserUUID, userUUID)
}
