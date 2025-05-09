package smarttable_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	viewsCodegenAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
	"github.com/stretchr/testify/assert"
)

func CreateDefaultRestaurantDish(
	token string,
	userUUID,
	restaurantUUID uuid.UUID,
) (uuid.UUID, error) {
	dishFile, err := os.Open("data/dish.png")
	if err != nil {
		return uuid.Nil, err
	}

	return CreateRestaurantDish(
		token,
		userUUID,
		restaurantUUID,
		"test_dish",
		"some_desc",
		"some_cat",
		100,
		100,
		"dish_file_name",
		dishFile,
	)
}

func CreateRestaurantDish(
	token string,
	userUUID,
	restaurantUUID uuid.UUID,
	name, description, category string,
	calories, weight int,
	pictureFileName string,
	picture io.Reader,
) (uuid.UUID, error) {
	var buf bytes.Buffer

	writer := multipart.NewWriter(&buf)
	_ = writer.WriteField("restaurant_uuid", restaurantUUID.String())
	_ = writer.WriteField("dish_name", name)
	_ = writer.WriteField("description", description)
	_ = writer.WriteField("category", category)
	_ = writer.WriteField("calories", strconv.Itoa(calories))
	_ = writer.WriteField("weight", strconv.Itoa(weight))

	part, err := writer.CreateFormFile("dish_picture_file", pictureFileName)

	if err != nil {
		return uuid.Nil, err
	}

	if _, err := io.Copy(part, picture); err != nil {
		return uuid.Nil, err
	}

	err = writer.Close()
	if err != nil {
		return uuid.Nil, err
	}

	resp, err := viewsCodegenAdminRestaurantClient.PostAdminV1RestaurantDishCreateWithBodyWithResponse(
		GetCtx(),
		&viewsCodegenAdminRestaurant.PostAdminV1RestaurantDishCreateParams{
			UserUUID: userUUID,
			JWTToken: token,
		},
		writer.FormDataContentType(),
		&buf,
	)

	if err != nil {
		return uuid.Nil, err
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return uuid.Nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode())
	}

	return resp.JSON200.DishUUID, nil
}

func TestAdminRestaurantDishCreateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.NoError(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.NoError(t, err)

	dishFile, err := os.Open("data/dish.png")
	assert.NoError(t, err)

	response, err := CreateRestaurantDish(
		token,
		userUUID,
		restaurantUUID,
		"test_dish",
		"some_desc",
		"some_cat",
		100,
		100,
		"dish_file_name",
		dishFile,
	)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}
