package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer"
	"github.com/stretchr/testify/assert"
)

var viewsCodegenCustomerClient, _ = viewsCodegenCustomer.NewClientWithResponses(GetBasePath())

func TestCustomerAuthHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	id, err := CreateCustomer("test_login", 123, 1234)
	assert.Nil(t, err)

	response, err := viewsCodegenCustomerClient.PostCustomerV1SignInWithResponse(
		GetCtx(),
		viewsCodegenCustomer.PostCustomerV1SignInJSONRequestBody{
			TgLogin: "test_login",
			TgID:    "123",
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
	assert.Equal(t, response.JSON200.CustomerUUID, id)

	customerPg, err := FindCustomerByTgID("123")
	assert.NoError(t, err)

	assert.Equal(t, "test_login", customerPg.TgLogin)
	assert.Equal(t, "123", customerPg.TgID)
	assert.Equal(t, "1234", customerPg.ChatID)
	assert.Equal(t, customerPg.UUID, id)
}
