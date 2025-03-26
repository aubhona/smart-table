package smarttable_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/uuid"
	defsinternalcustomerdb "github.com/smart-table/src/codegen/intern/customer_db"
	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer"
	viewsCustomer "github.com/smart-table/src/views/customer/v1"
	"github.com/stretchr/testify/assert"
)

func FindCustomerByTgID(tgID string) (defsinternalcustomerdb.PgCustomer, error) {
	customer := defsinternalcustomerdb.PgCustomer{}

	customerJSON, err := GetCustomerQueries().FetchCustomerByTgId(context.Background(), tgID)
	if err != nil {
		return customer, err
	}

	err = json.Unmarshal(customerJSON, &customer)

	return customer, err
}

func CreateCustomer(tgLogin, tgID, chatID string) (uuid.UUID, error) {
	handler := viewsCustomer.CustomerV1Handler{}
	response, err := handler.PostCustomerV1SignUp(GetCtx(), viewsCodegenCustomer.PostCustomerV1SignUpRequestObject{
		Body: &viewsCodegenCustomer.CustomerV1OrderCustomerSignUpRequest{
			TgLogin: tgLogin,
			TgID:    tgID,
			ChatID:  chatID,
		},
	})

	if err != nil {
		return uuid.Nil, err
	}

	responseObj, ok := response.(viewsCodegenCustomer.PostCustomerV1SignUp200JSONResponse)
	if !ok {
		return uuid.Nil, errors.New("response is not a PostCustomerV1SignUp200JSONResponse")
	}

	return responseObj.CustomerUUID, nil
}

func TestCustomerRegisterHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	handler := viewsCustomer.CustomerV1Handler{}

	response, err := handler.PostCustomerV1SignUp(GetCtx(), viewsCodegenCustomer.PostCustomerV1SignUpRequestObject{
		Body: &viewsCodegenCustomer.CustomerV1OrderCustomerSignUpRequest{
			TgLogin: "test_login",
			TgID:    "test_id",
			ChatID:  "test_chat_id",
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)

	responseObj, ok := response.(viewsCodegenCustomer.PostCustomerV1SignUp200JSONResponse)
	assert.True(t, ok)
	assert.NotEqual(t, responseObj.CustomerUUID, uuid.Nil)

	customerPg, err := FindCustomerByTgID("test_id")
	assert.NoError(t, err)

	assert.Equal(t, "test_login", customerPg.TgLogin)
	assert.Equal(t, "test_id", customerPg.TgID)
	assert.Equal(t, "test_chat_id", customerPg.ChatID)
}
