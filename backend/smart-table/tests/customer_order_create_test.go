package smarttable_test

import (
	"fmt"
	"net/http"
	"testing"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"
	"github.com/stretchr/testify/assert"
)

var viewsCodegenCustomerOrderClient, _ = viewsCodegenCustomer.NewClientWithResponses(GetBasePath())

func TestCustomerOrderCreateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.Nil(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.Nil(t, err)

	placeUUID, err := CreateDefaultPlace(token, userUUID, restaurantUUID)
	assert.Nil(t, err)

	tableID := fmt.Sprintf("%s_%d", placeUUID, defaultTableCount)

	hostUUID, err := CreateDefaultCustomer()
	assert.Nil(t, err)

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderCreateWithResponse(
		GetCtx(),
		viewsCodegenCustomer.PostCustomerV1OrderCreateJSONRequestBody{
			CustomerUUID: hostUUID,
			TableID:      tableID,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
}
