package smarttable_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/smart-table/src/domains/customer/infra/pg/mapper"
	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

func TestAdminOrderEditHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, _, placeUUID, hostCustomerUUID, orderUUID, _, _, token, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	err = CommitItems(hostCustomerUUID, orderUUID)
	assert.NoError(t, err)

	responseOrderInfo, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceOrderInfoWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceOrderInfoParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceOrderInfoJSONRequestBody{
			PlaceUUID: placeUUID,
			OrderUUID: orderUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, responseOrderInfo.StatusCode())
	assert.NotNil(t, responseOrderInfo.JSON200)

	itemGroup := viewsCodegenAdminPlace.ItemEditGroup{
		ItemStatus:   viewsCodegenAdminPlace.ItemStatus("cooked"),
		ItemUUIDList: []uuid.UUID{responseOrderInfo.JSON200.OrderInfo.CustomerList[0].ItemGroupList[0].ItemUUIDList[0]},
	}

	responseOrderEdit, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceOrderEditWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceOrderEditParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceOrderEditJSONRequestBody{
			PlaceUUID:   placeUUID,
			OrderUUID:   orderUUID,
			TableNumber: responseOrderInfo.JSON200.OrderInfo.OrderMainInfo.TableNumber,
			OrderStatus: nil,
			ItemGroup:   &itemGroup,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, responseOrderEdit.StatusCode())

	pgOrders, err := GetCustomerQueries().FetchOrders(GetCtx(), []uuid.UUID{orderUUID})
	assert.NoError(t, err)

	order := mapper.PgOrderAggregate{}
	err = json.Unmarshal(pgOrders[0], &order)
	assert.NoError(t, err)

	assert.Equal(t, "cooked", order.PgItems[0].Status)

	orderStatus := viewsCodegenAdminPlace.OrderStatus("canceled_by_service")

	responseOrderEdit, err = viewsCodegenAdminPlaceClient.PostAdminV1PlaceOrderEditWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceOrderEditParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceOrderEditJSONRequestBody{
			PlaceUUID:   placeUUID,
			OrderUUID:   orderUUID,
			TableNumber: responseOrderInfo.JSON200.OrderInfo.OrderMainInfo.TableNumber,
			OrderStatus: &orderStatus,
			ItemGroup:   nil,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, responseOrderEdit.StatusCode())

	pgOrders, err = GetCustomerQueries().FetchOrders(GetCtx(), []uuid.UUID{orderUUID})
	assert.NoError(t, err)

	order = mapper.PgOrderAggregate{}
	err = json.Unmarshal(pgOrders[0], &order)
	assert.NoError(t, err)

	assert.Equal(t, "canceled_by_service", order.PgOrder.Status)
	assert.Equal(t, "canceled_by_service", order.PgItems[0].Status)
}
