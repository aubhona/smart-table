package smarttable_test

import (
	"testing"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer"
	viewsCustomer "github.com/smart-table/src/views/customer/v1"
	"github.com/stretchr/testify/assert"
)

func TestCustomerAuthHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	id, err := CreateCustomer("test_login", 123, 1234)
	assert.Nil(t, err)

	handler := viewsCustomer.CustomerV1Handler{}
	response, err := handler.PostCustomerV1SignIn(GetCtx(), viewsCodegenCustomer.PostCustomerV1SignInRequestObject{
		Body: &viewsCodegenCustomer.PostCustomerV1SignInJSONRequestBody{
			TgLogin: "test_login",
			TgID:    "123",
			ChatID:  "1234",
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)

	responseObj, ok := response.(viewsCodegenCustomer.PostCustomerV1SignIn200JSONResponse)
	assert.True(t, ok)
	assert.Equal(t, responseObj.CustomerUUID, id)

	customerPg, err := FindCustomerByTgID("123")
	assert.NoError(t, err)

	assert.Equal(t, "test_login", customerPg.TgLogin)
	assert.Equal(t, "123", customerPg.TgID)
	assert.Equal(t, "1234", customerPg.ChatID)
	assert.Equal(t, customerPg.UUID, id)
}
