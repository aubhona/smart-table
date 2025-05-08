package smarttable_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	defsInternalAdminDb "github.com/smart-table/src/codegen/intern/admin_user_db"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_user"
	"github.com/stretchr/testify/assert"
)

const (
	userDefaultFirstName = "testFisrtName"
	userDefaultLastName  = "testLastName"
	userDefaultLogin     = "testLogin"
	userDefaultPassword  = "testPassword"
	userDefaultTgLogin   = "testTgLogin"
)

var viewsCodegenAdminClient, _ = viewsCodegenAdmin.NewClientWithResponses(GetBasePath())

func FindUserByLogin(login string) (defsInternalAdminDb.PgUser, error) {
	user := defsInternalAdminDb.PgUser{}

	userJSON, err := GetAdminQueries().FetchUserByLogin(GetCtx(), login)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(userJSON, &user)

	return user, err
}

func CreateDefaultUser() (uuid.UUID, string, error) {
	return CreateUser(
		userDefaultFirstName,
		userDefaultLastName,
		userDefaultLogin,
		userDefaultPassword,
		userDefaultTgLogin,
	)
}

func CreateUser(firstName, lastName, login, password, tgLogin string) (uuid.UUID, string, error) {
	resp, err := viewsCodegenAdminClient.PostAdminV1UserSignUpWithResponse(
		GetCtx(),
		viewsCodegenAdmin.PostAdminV1UserSignUpJSONRequestBody{
			FirstName: firstName,
			LastName:  lastName,
			Login:     login,
			Password:  password,
			TgLogin:   tgLogin,
		},
	)

	if err != nil {
		return uuid.Nil, "", err
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return uuid.Nil, "", fmt.Errorf("unexpected response status: %d", resp.StatusCode())
	}

	return resp.JSON200.UserUUID, resp.JSON200.JwtToken, nil
}

func TestAdminUserSignUpHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, _, err := CreateDefaultUser()
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, userUUID)

	userPg, err := FindUserByLogin("testLogin")
	assert.NoError(t, err)
	assert.Equal(t, "testFisrtName", userPg.FirstName)
	assert.Equal(t, "testLastName", userPg.LastName)
	assert.Equal(t, "testTgLogin", userPg.TgLogin)
}
