package smarttable_test

import (
	"net/http"
	"testing"

	"github.com/google/uuid"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func CreateDefaultItems() (
	userUUID,
	restaurantUUID,
	placeUUID,
	hostCustomerUUID,
	orderUUID,
	dishUUID,
	menuDishUUID uuid.UUID,
	token string,
	error error,
) {
	userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, token, error = CreateDefaultOrder()
	if error != nil {
		return userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, dishUUID, menuDishUUID, token, error
	}

	dishUUID, error = CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	if error != nil {
		return userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, dishUUID, menuDishUUID, token, error
	}

	menuDishUUID, error = CreateDefaultPlaceMenuDish(token, userUUID, placeUUID, dishUUID)
	if error != nil {
		return userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, dishUUID, menuDishUUID, token, error
	}

	comment := "comment"

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderItemsDraftCountEditWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderItemsDraftCountEditParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
		},
		viewsCodegenCustomer.CustomerV1OrderItemsDraftCountEditRequest{
			Count:        3,
			MenuDishUUID: menuDishUUID,
			Comment:      &comment,
		},
	)

	if err != nil || response.StatusCode() != http.StatusNoContent {
		return userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, dishUUID, menuDishUUID, token, err
	}

	return userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, dishUUID, menuDishUUID, token, nil
}

func TestCustomerOrderItemsCountDraftEditHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, token, err := CreateDefaultOrder()
	assert.NoError(t, err)

	dishUUID, err := CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	menuDishUUID, err := CreateDefaultPlaceMenuDish(token, userUUID, placeUUID, dishUUID)
	assert.NoError(t, err)

	comment := "comment"

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderItemsDraftCountEditWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderItemsDraftCountEditParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
		},
		viewsCodegenCustomer.CustomerV1OrderItemsDraftCountEditRequest{
			Count:        3,
			MenuDishUUID: menuDishUUID,
			Comment:      &comment,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode())

	response, err = viewsCodegenCustomerOrderClient.PostCustomerV1OrderItemsDraftCountEditWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderItemsDraftCountEditParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
		},
		viewsCodegenCustomer.CustomerV1OrderItemsDraftCountEditRequest{
			Count:        -2,
			MenuDishUUID: menuDishUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode())
}
