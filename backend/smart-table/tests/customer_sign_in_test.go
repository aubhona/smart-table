package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer"
	"github.com/stretchr/testify/assert"
)

const defaultInitData = `user=%7B%22id%22%3A726434232%2C%22first_name%22%3A%22deryun%22%2C%22last_name%22%3A%22%22%2C%22username%22%3A%22deryuuun%22%2C%22language_code%22%3A%22ru%22%2C%22is_premium%22%3Atrue%2C%22allows_write_to_pm%22%3Atrue%2C%22photo_url%22%3A%22https%3A%5C%2F%5C%2Ft.me%5C%2Fi%5C%2Fuserpic%5C%2F320%5C%2FNDS-jJaXxxR3nC55GO3PL6rSC0Tk30spaei_c3W9-Ww.svg%22%7D&chat_instance=-7502140473260042804&chat_type=private&start_param=1836dc3a-ff77-484f-9b07-48c2b564184a_1&auth_date=1747557002&signature=HugKRM_kRPlpUa-myVGnP1GWF_tx37MMkX2SQBc18m-x8d28iVj2_QlD2rxNU3fDwuoj4znsOSC22QzeMv7_Cg&hash=065696260ed4cf5b46a5cd10b528e2d31e252533662abe7199d9f018e64b85f1` //nolint

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
			InitData: defaultInitData,
			TgLogin:  "test_login",
			TgID:     "123",
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
