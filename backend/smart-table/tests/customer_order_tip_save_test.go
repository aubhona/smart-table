package smarttable_test

import (
	"net/http"
	"testing"

	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderTipSaveHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	_, _, _, hostCustomerUUID, orderUUID, _, _, _, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	err = CommitItems(hostCustomerUUID, orderUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderTipSaveWithResponse(
		GetCtx(),
		&viewsCustomerOrder.PostCustomerV1OrderTipSaveParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
			JWTToken:     "tipa_token_zhiest",
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode())
}
