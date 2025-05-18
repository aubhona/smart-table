package smarttable_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/smart-table/src/domains/customer/infra/pg/mapper"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"
	"github.com/stretchr/testify/assert"
)

var viewsCodegenCustomerOrderClient, _ = viewsCodegenCustomer.NewClientWithResponses(GetBasePath())

func CreateDefaultOrder() (userUUID, restaurantUUID, placeUUID, hostUUID, orderUUID uuid.UUID, token string, error error) {
	userUUID, token, err := CreateDefaultUser()
	if err != nil {
		return userUUID, restaurantUUID, placeUUID, hostUUID, orderUUID, token, err
	}

	restaurantUUID, err = CreateDefaultRestaurant(token, userUUID)
	if err != nil {
		return userUUID, restaurantUUID, placeUUID, hostUUID, orderUUID, token, err
	}

	placeUUID, err = CreateDefaultPlace(token, userUUID, restaurantUUID)
	if err != nil {
		return userUUID, restaurantUUID, placeUUID, hostUUID, orderUUID, token, err
	}

	tableID := fmt.Sprintf("%s_%d", placeUUID, defaultTableCount)

	hostUUID, err = CreateDefaultCustomer()
	if err != nil {
		return userUUID, restaurantUUID, placeUUID, hostUUID, orderUUID, token, err
	}

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderCreateWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderCreateParams{
			CustomerUUID: hostUUID,
			JWTToken:     "tipa_token_zhiest",
		},
		viewsCodegenCustomer.PostCustomerV1OrderCreateJSONRequestBody{
			TableID: tableID,
		},
	)
	if err != nil || response.JSON200 == nil || response.StatusCode() != http.StatusOK {
		return userUUID, restaurantUUID, placeUUID, hostUUID, orderUUID, token, err
	}

	orderUUID = response.JSON200.OrderUUID

	return userUUID, restaurantUUID, placeUUID, hostUUID, orderUUID, token, nil
}

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
		&viewsCodegenCustomer.PostCustomerV1OrderCreateParams{
			CustomerUUID: hostUUID,
			JWTToken:     "tipa_token_zhiest",
		},
		viewsCodegenCustomer.PostCustomerV1OrderCreateJSONRequestBody{
			TableID: tableID,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
}

func TestCustomerOrderCreateIdempotency(t *testing.T) {
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

	response1, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderCreateWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderCreateParams{
			CustomerUUID: hostUUID,
			JWTToken:     "tipa_token_zhiest",
		},
		viewsCodegenCustomer.PostCustomerV1OrderCreateJSONRequestBody{
			TableID: tableID,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response1.StatusCode())
	assert.NotNil(t, response1.JSON200)

	response2, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderCreateWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderCreateParams{
			CustomerUUID: hostUUID,
			JWTToken:     "tipa_token_zhiest",
		},
		viewsCodegenCustomer.PostCustomerV1OrderCreateJSONRequestBody{
			TableID: tableID,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response2.StatusCode())
	assert.NotNil(t, response2.JSON200)

	assert.Equal(t, response1.JSON200.OrderUUID, response2.JSON200.OrderUUID)
}

func TestCustomerOrderCreateConnectingToExistingSession(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.NoError(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.NoError(t, err)

	placeUUID, err := CreateDefaultPlace(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	tableID := fmt.Sprintf("%s_%d", placeUUID, defaultTableCount)

	hostUUID, err := CreateDefaultCustomer()
	assert.NoError(t, err)

	customerUUID, err := CreateCustomer("some_login", 123515, 124)
	assert.NoError(t, err)

	response1, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderCreateWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderCreateParams{
			CustomerUUID: hostUUID,
			JWTToken:     "tipa_token_zhiest",
		},
		viewsCodegenCustomer.PostCustomerV1OrderCreateJSONRequestBody{
			TableID: tableID,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response1.StatusCode())
	assert.NotNil(t, response1.JSON200)

	rawOrders, err := GetCustomerQueries().FetchOrders(GetCtx(), []uuid.UUID{response1.JSON200.OrderUUID})
	assert.NoError(t, err)

	pgOrder := mapper.PgOrderAggregate{}

	err = json.Unmarshal(rawOrders[0], &pgOrder)
	assert.NoError(t, err)

	response2, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderCreateWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderCreateParams{
			CustomerUUID: customerUUID,
			JWTToken:     "tipa_token_zhiest",
		},
		viewsCodegenCustomer.PostCustomerV1OrderCreateJSONRequestBody{
			TableID:  tableID,
			RoomCode: &pgOrder.PgOrder.RoomCode,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response2.StatusCode())
	assert.NotNil(t, response2.JSON200)

	assert.Equal(t, response1.JSON200, response2.JSON200)

	rawOrders, err = GetCustomerQueries().FetchOrders(GetCtx(), []uuid.UUID{response1.JSON200.OrderUUID})
	assert.NoError(t, err)

	err = json.Unmarshal(rawOrders[0], &pgOrder)
	assert.NoError(t, err)

	assert.Equal(t, pgOrder.PgOrder.CustomersUUID, []uuid.UUID{hostUUID, customerUUID})
}
