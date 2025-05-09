package smarttable_test

import (
	"fmt"
	"net/http"
	"testing"

	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

func TestAdminPlaceTableDeepLinksListHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.NoError(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.NoError(t, err)

	placeUUID, err := CreateDefaultPlace(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceTableDeeplinksListWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceTableDeeplinksListParams{
			JWTToken: token,
			UserUUID: userUUID,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceTableDeeplinksListJSONRequestBody{
			PlaceUUID: placeUUID,
		},
	)

	expectedDeepLinks := make([]string, 0)

	for i := 1; i <= 10; i++ {
		expectedDeepLinks = append(expectedDeepLinks, fmt.Sprintf("%s=%s_%d", GetDeps().Config.Bot.WebAppURL, placeUUID, i))
	}

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())
	assert.NotNil(t, response.JSON200)
	assert.Equal(t, expectedDeepLinks, response.JSON200.Deeplinks)
}
