package smarttable_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"

	"github.com/stretchr/testify/assert"
)

func TestAdminOrderInfoHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, _, placeUUID, hostCustomerUUID, orderUUID, _, menuDishUUID, token, err := CreateDefaultItems()
	assert.NoError(t, err)

	err = CommitItems(hostCustomerUUID, orderUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceOrderInfoWithResponse(
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
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)

	response.JSON200.OrderInfo.CustomerList[0].ItemGroupList[0].ItemUUIDList[0], err = uuid.Parse("14f12b05-6406-454d-855f-c25774307c11")
	assert.NoError(t, err)

	response.JSON200.OrderInfo.CustomerList[0].ItemGroupList[0].ItemUUIDList[1], err = uuid.Parse("1156360d-1d71-4a1c-adfb-2405cb9d5e15")
	assert.NoError(t, err)

	response.JSON200.OrderInfo.CustomerList[0].ItemGroupList[0].ItemUUIDList[2], err = uuid.Parse("1a372277-108a-4ea5-a0f7-e16009772644")
	assert.NoError(t, err)

	response.JSON200.OrderInfo.OrderMainInfo.UUID, err = uuid.Parse("149ebb8c-ce00-4492-b468-12d16b89ed36")
	assert.NoError(t, err)

	response.JSON200.OrderInfo.OrderMainInfo.CreatedAt, err = time.Parse(time.RFC3339Nano, "2025-05-19T20:29:41.002238Z")
	assert.NoError(t, err)

	actualJSON, err := json.Marshal(response.JSON200.OrderInfo)
	assert.NoError(t, err)

	expectedJSON := fmt.Sprintf(`
{
  "customer_list": [
    {
      "item_group_list": [
        {
          "comment": "comment",
          "count": 3,
          "item_price": "123.13",
          "item_uuid_list": [
            "14f12b05-6406-454d-855f-c25774307c11",
						"1156360d-1d71-4a1c-adfb-2405cb9d5e15",
						"1a372277-108a-4ea5-a0f7-e16009772644"
          ],
          "menu_dish_uuid": "%s",
          "name": "test_dish",
          "result_price": "369.39",
          "status": "new"
        }
      ],
      "tg_id": "123",
      "tg_login": "testLogin",
      "total_price": "369.39",
      "uuid": "%s"
    }
  ],
  "order_main_info": {
    "created_at": "2025-05-19T20:29:41.002238Z",
    "guests_count": 1,
    "status": "new",
    "table_number": 10,
    "total_price": "369.39",
    "uuid": "149ebb8c-ce00-4492-b468-12d16b89ed36"
  }
}
`, menuDishUUID, hostCustomerUUID)

	assert.JSONEq(t, expectedJSON, string(actualJSON))
}
