package smarttable_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"

	"github.com/stretchr/testify/assert"
)

func TestAdminOrderListHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, _, placeUUID, hostCustomerUUID, orderUUID, _, _, token, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	err = CommitItems(hostCustomerUUID, orderUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceOrderListWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceOrderListParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceOrderListJSONRequestBody{
			PlaceUUID: placeUUID,
			IsActive:  true,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)

	response.JSON200.OrderList[0].UUID, err = uuid.Parse("149ebb8c-ce00-4492-b468-12d16b89ed36")
	assert.NoError(t, err)

	response.JSON200.OrderList[0].CreatedAt, err = time.Parse(time.RFC3339Nano, "2025-05-19T20:29:41.002238Z")
	assert.NoError(t, err)

	actualJSON, err := json.Marshal(response.JSON200.OrderList)
	assert.NoError(t, err)

	expectedJSON := `
[
  {
		"created_at": "2025-05-19T20:29:41.002238Z",
		"guests_count": 1,
		"status": "new",
		"table_number": 10,
		"total_price": "369.39",
		"uuid": "149ebb8c-ce00-4492-b468-12d16b89ed36"
	}
]
`

	assert.JSONEq(t, expectedJSON, string(actualJSON))
}
