package smarttable_test

import (
	"net/http"
	"testing"

	"github.com/smart-table/src/utils"
	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderCatalogUpdateInfoTestHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, token, err := CreateDefaultOrder()
	assert.NoError(t, err)

	dishUUID, err := CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	menuDishUUID, err := CreateDefaultPlaceMenuDish(token, userUUID, placeUUID, dishUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.GetCustomerV1OrderCatalogUpdatedInfoWithResponse(
		GetCtx(),
		&viewsCustomerOrder.GetCustomerV1OrderCatalogUpdatedInfoParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
		},
	)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
	assert.Equal(t, "0", response.JSON200.TotalPrice)
	assert.Equal(t, []viewsCustomerOrder.MenuDishItemUpdatedInfo{}, response.JSON200.MenuUpdatedInfo)

	err = EditItems(hostCustomerUUID, orderUUID, menuDishUUID, 3, utils.EmptyOptional[string]())
	assert.NoError(t, err)

	response, err = viewsCodegenCustomerOrderClient.GetCustomerV1OrderCatalogUpdatedInfoWithResponse(
		GetCtx(),
		&viewsCustomerOrder.GetCustomerV1OrderCatalogUpdatedInfoParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
		},
	)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
	assert.Equal(t, "369.39", response.JSON200.TotalPrice)
	assert.Equal(t, []viewsCustomerOrder.MenuDishItemUpdatedInfo{
		{
			ID:    menuDishUUID,
			Count: 3,
		},
	}, response.JSON200.MenuUpdatedInfo)
}
