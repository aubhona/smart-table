package smarttable_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderCustomerListHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	_, _, _, hostCustomerUUID, orderUUID, _, menuDishUUID, _, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	err = CommitItems(hostCustomerUUID, orderUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.GetCustomerV1OrderCustomerListWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.GetCustomerV1OrderCustomerListParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
			JWTToken:     "tipa_token_zhiest",
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)

	actualJSON, err := json.Marshal(response.JSON200.CustomerList)
	assert.NoError(t, err)

	expectedJSON := fmt.Sprintf(`
[
  {
    "is_active": true,
    "is_host": true,
    "item_list": [
      {
        "calories": 100,
        "category": "some_cat",
        "comment": "comment",
        "count": 3,
        "description": "some_desc",
        "name": "test_dish",
        "price": "123.13",
        "result_price": "369.39",
        "status": "new",
        "dish_uuid": "%s",
        "weight": 100
      }
    ],
    "tg_id": "123",
    "tg_login": "testLogin",
    "total_price": "369.39",
    "uuid": "%s"
  }
]
`, menuDishUUID, hostCustomerUUID)

	assert.JSONEq(t, expectedJSON, string(actualJSON))
}
