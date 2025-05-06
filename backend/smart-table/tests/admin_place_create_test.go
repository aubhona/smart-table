package smarttable_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	viewsAdmin "github.com/smart-table/src/views/admin/v1/place"
	viewsCodegenAdmin "github.com/smart-table/src/views/codegen/admin_place"
	"github.com/stretchr/testify/assert"
)

func CreateDefaultPlace(userUUID, restaurantUUID uuid.UUID) (uuid.UUID, error) {
	return CreatePlace(
		userUUID,
		restaurantUUID,
		"testAddress",
		"12:00",
		"13:00",
		10,
	)
}

func CreatePlace(userUUID, restaurantUUID uuid.UUID, address, openingTime, closingTime string, tableCount int) (uuid.UUID, error) {
	handler := viewsAdmin.AdminV1PlaceHandler{}
	response, err := handler.PostAdminV1PlaceCreate(GetCtx(), viewsCodegenAdmin.PostAdminV1PlaceCreateRequestObject{
		Params: viewsCodegenAdmin.PostAdminV1PlaceCreateParams{
			UserUUID: userUUID,
		},
		Body: &viewsCodegenAdmin.AdminV1PlaceCreateRequest{
			RestaurantUUID: restaurantUUID,
			Address:        address,
			TableCount:     tableCount,
			OpeningTime:    openingTime,
			ClosingTime:    closingTime,
		},
	})

	if err != nil {
		return uuid.Nil, err
	}

	responseObj, ok := response.(viewsCodegenAdmin.PostAdminV1PlaceCreate200JSONResponse)
	if !ok {
		return uuid.Nil, errors.New("response is not a PosPostAdminV1PlaceCreate200JSONResponsetAdminV1UserSignUp200JSONResponse")
	}

	return responseObj.PlaceUUID, nil
}

func TestAdminPlaceCreateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, err := CreateDefaultUser()
	assert.Nil(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(userUUID)
	assert.Nil(t, err)

	handler := viewsAdmin.AdminV1PlaceHandler{}

	response, err := handler.PostAdminV1PlaceCreate(GetCtx(), viewsCodegenAdmin.PostAdminV1PlaceCreateRequestObject{
		Params: viewsCodegenAdmin.PostAdminV1PlaceCreateParams{
			UserUUID: userUUID,
		},
		Body: &viewsCodegenAdmin.AdminV1PlaceCreateRequest{
			RestaurantUUID: restaurantUUID,
			Address:        "testAddress",
			TableCount:     1,
			OpeningTime:    "13:00",
			ClosingTime:    "14:00",
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, response)

	_, ok := response.(viewsCodegenAdmin.PostAdminV1PlaceCreate200JSONResponse)
	assert.True(t, ok)
}
