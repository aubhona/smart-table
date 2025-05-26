package smarttable_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderItemInfoStateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	_, _, _, hostCustomerUUID, orderUUID, _, menuDishUUID, _, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderItemInfoStateWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderItemInfoStateParams{
			CustomerUUID: hostCustomerUUID,
			JWTToken:     "tipa_token_zhiest",
			OrderUUID:    orderUUID,
		},
		viewsCodegenCustomer.PostCustomerV1OrderItemInfoStateJSONRequestBody{
			DishUUID: menuDishUUID,
			Comment:  &defaultComment,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)

	response.JSON200.ID, err = uuid.Parse("149ebb8c-ce00-4492-b468-12d16b89ed36")
	assert.NoError(t, err)

	actualJSON, err := json.Marshal(response.JSON200)
	assert.NoError(t, err)

	expectedJSON := `
{
	"calories":100, 
	"category":"some_cat", 
	"count":3, 
	"description":"some_desc", 
	"id":"149ebb8c-ce00-4492-b468-12d16b89ed36", 
	"name":"test_dish", 
	"price":"123.13", 
	"result_price":"369.39", 
	"weight":100
}
`

	assert.JSONEq(t, expectedJSON, string(actualJSON))
}
