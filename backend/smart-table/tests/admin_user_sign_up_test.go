package smarttable_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/uuid"
	defsInternalAdminDb "github.com/smart-table/src/codegen/intern/admin_user_db"
	viewsAdmin "github.com/smart-table/src/views/admin/v1/user"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_user"
	"github.com/stretchr/testify/assert"
)

func FindUserByLogin(login string) (defsInternalAdminDb.PgUser, error) {
	user := defsInternalAdminDb.PgUser{}

	userJSON, err := GetAdminQueries().FetchUserByLogin(context.Background(), login)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(userJSON, &user)

	return user, err
}

func CreateDefaultUser() (uuid.UUID, error) {
	return CreateUser(
		"testFisrtName",
		"testLastName",
		"testLogin",
		"testPassword",
		"testTgLogin",
	)
}

func CreateUser(firstName, lastName, login, password, tgLogin string) (uuid.UUID, error) {
	handler := viewsAdmin.AdminV1UserHandler{}
	response, err := handler.PostAdminV1UserSignUp(GetCtx(), viewsCodegenAdmin.PostAdminV1UserSignUpRequestObject{
		Body: &viewsCodegenAdmin.AdminV1UserSignUpRequest{
			FirstName: firstName,
			LastName:  lastName,
			Login:     login,
			Password:  password,
			TgLogin:   tgLogin,
		},
	})

	if err != nil {
		return uuid.Nil, err
	}

	responseObj, ok := response.(viewsCodegenAdmin.PostAdminV1UserSignUp200JSONResponse)
	if !ok {
		return uuid.Nil, errors.New("response is not a PostAdminV1UserSignUp200JSONResponse")
	}

	return responseObj.Body.UserUUID, nil
}

func TestAdminUserSignUpHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	handler := viewsAdmin.AdminV1UserHandler{}

	response, err := handler.PostAdminV1UserSignUp(GetCtx(), viewsCodegenAdmin.PostAdminV1UserSignUpRequestObject{
		Body: &viewsCodegenAdmin.AdminV1UserSignUpRequest{
			FirstName: "testFisrtName",
			LastName:  "testLastName",
			Login:     "testLogin",
			Password:  "testPassword",
			TgLogin:   "testTgLogin",
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)

	responseObj, ok := response.(viewsCodegenAdmin.PostAdminV1UserSignUp200JSONResponse)
	assert.True(t, ok)
	assert.NotEqual(t, responseObj.Body.UserUUID, uuid.Nil)

	userPg, err := FindUserByLogin("testLogin")
	assert.NoError(t, err)

	assert.Equal(t, "testFisrtName", userPg.FirstName)
	assert.Equal(t, "testLastName", userPg.LastName)
	assert.Equal(t, "testTgLogin", userPg.TgLogin)
	assert.NotNil(t, userPg.PasswordHash)
}
