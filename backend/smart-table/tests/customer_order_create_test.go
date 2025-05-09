package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"
	"github.com/stretchr/testify/assert"
)

const defaultTableID = "32-232"

var viewsCodegenCustomerOrderClient, _ = viewsCodegenCustomer.NewClientWithResponses(GetBasePath())

func TestCustomerOrderCreateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	hostUUID, err := CreateDefaultCustomer()
	assert.Nil(t, err)

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderCreateWithResponse(
		GetCtx(),
		viewsCodegenCustomer.PostCustomerV1OrderCreateJSONRequestBody{
			CustomerUUID: hostUUID,
			TableID:      defaultTableID,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
}
