package smarttable_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	viewsCodegenAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
	"github.com/stretchr/testify/assert"
)

var viewsCodegenAdminRestaurantClient, _ = viewsCodegenAdminRestaurant.NewClientWithResponses(GetBasePath())

func CreateDefaultRestaurant(token string, ownerUUID uuid.UUID) (uuid.UUID, error) {
	return CreateRestaurant(
		"testRestaurantName",
		token,
		ownerUUID,
	)
}

func CreateRestaurant(restaurantName, token string, ownerUUID uuid.UUID) (uuid.UUID, error) {
	response, err := viewsCodegenAdminRestaurantClient.PostAdminV1RestaurantCreateWithResponse(
		GetCtx(),
		&viewsCodegenAdminRestaurant.PostAdminV1RestaurantCreateParams{
			UserUUID: ownerUUID,
			JWTToken: token,
		},
		viewsCodegenAdminRestaurant.PostAdminV1RestaurantCreateJSONRequestBody{
			RestaurantName: restaurantName,
		},
	)

	if err != nil {
		return uuid.Nil, err
	}

	if response.JSON200 == nil || response.StatusCode() != http.StatusOK {
		return uuid.Nil, errors.New("response is not a PostAdminV1RestaurantCreate200JSONResponse")
	}

	return response.JSON200.RestaurantUUID, nil
}

func TestAdminRestaurantCreateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.Nil(t, err)

	response, err := viewsCodegenAdminRestaurantClient.PostAdminV1RestaurantCreateWithResponse(
		GetCtx(),
		&viewsCodegenAdminRestaurant.PostAdminV1RestaurantCreateParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		viewsCodegenAdminRestaurant.PostAdminV1RestaurantCreateJSONRequestBody{
			RestaurantName: "testRestaurantName",
		},
	)

	assert.NoError(t, err)
	assert.NotNil(t, response.JSON200)
	assert.Equal(t, response.StatusCode(), http.StatusOK)
}
