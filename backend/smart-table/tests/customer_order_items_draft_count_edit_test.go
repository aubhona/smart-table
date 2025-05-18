package smarttable_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/samber/lo"
	defsInternalItemDb "github.com/smart-table/src/codegen/intern/item_db"
	"github.com/smart-table/src/domains/customer/infra/pg/mapper"
	"github.com/smart-table/src/utils"

	"github.com/google/uuid"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

var (
	defaultItemsCount = 3
	defaultComment    = "comment"
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

	err := EditItems(hostCustomerUUID, orderUUID, menuDishUUID, defaultItemsCount, utils.NewOptional(defaultComment))

	return userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, dishUUID, menuDishUUID, token, err
}

func EditItems(customerUUID, orderUUID, menuDishUUID uuid.UUID, count int, comment utils.Optional[string]) error {
	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderItemsDraftCountEditWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderItemsDraftCountEditParams{
			CustomerUUID: customerUUID,
			JWTToken:     "tipa_token_zhiest",
			OrderUUID:    orderUUID,
		},
		viewsCodegenCustomer.CustomerV1OrderItemsDraftCountEditRequest{
			Count:        count,
			MenuDishUUID: menuDishUUID,
			Comment:      comment.ToPointer(),
		},
	)

	if err != nil || response.StatusCode() != http.StatusNoContent {
		return err
	}

	return nil
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
			JWTToken:     "tipa_token_zhiest",
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

	pgOrders, err := GetCustomerQueries().FetchOrders(GetCtx(), []uuid.UUID{orderUUID})
	assert.NoError(t, err)

	order := mapper.PgOrderAggregate{}
	err = json.Unmarshal(pgOrders[0], &order)
	assert.NoError(t, err)

	assert.Equal(t, 3, len(order.PgItems))
	assert.Equal(t, 3, len(lo.Filter(order.PgItems, func(item defsInternalItemDb.PgItem, _ int) bool {
		return item.IsDraft
	})))

	response, err = viewsCodegenCustomerOrderClient.PostCustomerV1OrderItemsDraftCountEditWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderItemsDraftCountEditParams{
			CustomerUUID: hostCustomerUUID,
			JWTToken:     "tipa_token_zhiest",
			OrderUUID:    orderUUID,
		},
		viewsCodegenCustomer.CustomerV1OrderItemsDraftCountEditRequest{
			Count:        -2,
			MenuDishUUID: menuDishUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode())

	pgOrders, err = GetCustomerQueries().FetchOrders(GetCtx(), []uuid.UUID{orderUUID})
	assert.NoError(t, err)

	order = mapper.PgOrderAggregate{}
	err = json.Unmarshal(pgOrders[0], &order)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(order.PgItems))
}
