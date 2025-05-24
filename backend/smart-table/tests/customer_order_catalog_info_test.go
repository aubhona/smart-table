package smarttable_test

import (
	"net/http"
	"testing"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderCatalogInfoHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, token, err := CreateDefaultOrder()
	assert.NoError(t, err)

	dishUUID, err := CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	menuDishUUID, err := CreateDefaultPlaceMenuDish(token, userUUID, placeUUID, dishUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.GetCustomerV1OrderCatalogInfoWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.GetCustomerV1OrderCatalogInfoParams{
			CustomerUUID: hostCustomerUUID,
			JWTToken:     "tipa_token_zhiest",
			OrderUUID:    orderUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
	assert.Equal(t, 1, len(response.JSON200.Menu))
	assert.Equal(t, menuDishUUID, response.JSON200.Menu[0].ID)
	assert.Equal(t, "some_cat", response.JSON200.Menu[0].Category)
	assert.Equal(t, "0", response.JSON200.TotalPrice)
	assert.Equal(t, []string{"some_cat"}, response.JSON200.Categories)
	assert.Equal(t, false, response.JSON200.GoTipScreen)
	assert.Equal(t, 100, response.JSON200.Menu[0].Calories)
}
