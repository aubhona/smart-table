package smarttable_test

import (
	"encoding/json"
	"net/http"
	"testing"

	viewsCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/google/uuid"
	"github.com/smart-table/src/domains/customer/infra/pg/mapper"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderFinishHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	_, _, _, hostCustomerUUID, orderUUID, _, _, _, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	err = CommitItems(hostCustomerUUID, orderUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderFinishWithResponse(
		GetCtx(),
		&viewsCustomerOrder.PostCustomerV1OrderFinishParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
			JWTToken:     "tipa_token_zhiest",
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode())

	pgOrders, err := GetCustomerQueries().FetchOrders(GetCtx(), []uuid.UUID{orderUUID})
	assert.NoError(t, err)

	order := mapper.PgOrderAggregate{}
	err = json.Unmarshal(pgOrders[0], &order)
	assert.NoError(t, err)

	assert.Equal(t, "payment_waiting", order.PgOrder.Status)
}
