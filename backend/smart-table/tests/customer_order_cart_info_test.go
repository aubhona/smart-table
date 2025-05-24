package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderCartInfoHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	_, _, _, hostCustomerUUID, orderUUID, _, menuDishUUID, _, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.GetCustomerV1OrderCartInfoWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.GetCustomerV1OrderCartInfoParams{
			CustomerUUID: hostCustomerUUID,
			JWTToken:     "tipa_token_zhiest",
			OrderUUID:    orderUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
	assert.Equal(t, 1, len(response.JSON200.Items))
	assert.Equal(t, menuDishUUID, response.JSON200.Items[0].ID)
	assert.Equal(t, "369.39", response.JSON200.Items[0].ResultPrice)
	assert.Equal(t, defaultItemsCount, response.JSON200.Items[0].Count)
	assert.Equal(t, "369.39", response.JSON200.TotalPrice)
}
