package smarttable_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/samber/lo"
	defsInternalItemDb "github.com/smart-table/src/codegen/intern/item_db"
	"github.com/smart-table/src/domains/customer/infra/pg/mapper"

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderItemsCommitHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	_, _, _, hostCustomerUUID, orderUUID, _, _, _, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderItemsCommitWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderItemsCommitParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
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
		return !item.IsDraft
	})))
}
