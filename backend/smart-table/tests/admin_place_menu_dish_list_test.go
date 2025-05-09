package smarttable_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"

	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"

	"github.com/stretchr/testify/assert"
)

func TestAdminPlaceMenuDishListHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, token, err := CreateDefaultUser()
	assert.NoError(t, err)

	restaurantUUID, err := CreateDefaultRestaurant(token, userUUID)
	assert.NoError(t, err)

	placeUUID, err := CreateDefaultPlace(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	dishUUID, err := CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	menuDishUUID, err := CreateDefaultPlaceMenuDish(token, userUUID, placeUUID, dishUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenAdminPlaceClient.PostAdminV1PlaceMenuDishListWithResponse(
		GetCtx(),
		&viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishListParams{
			JWTToken: token,
			UserUUID: userUUID,
		},
		viewsCodegenAdminPlace.PostAdminV1PlaceMenuDishListJSONRequestBody{
			PlaceUUID: placeUUID,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode())

	mediaType, params, err := mime.ParseMediaType(response.HTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		t.Fatalf("cannot parse media type: %v", err)
	}

	if mediaType != "multipart/mixed" {
		t.Fatalf("unexpected media type: %s", mediaType)
	}

	mr := multipart.NewReader(bytes.NewReader(response.Body), params["boundary"])

	part, err := mr.NextPart()

	if err != nil {
		t.Fatalf("failed to read first part: %v", err)
	}

	if part.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("expected first part to be JSON, got: %s", part.Header.Get("Content-Type"))
	}

	jsonBytes, err := io.ReadAll(part)

	if err != nil {
		t.Fatalf("failed to read JSON part: %v", err)
	}

	var metadata []viewsCodegenAdminPlace.MenuDishInfo

	if err := json.Unmarshal(jsonBytes, &metadata); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	t.Logf("parsed %d menu dishes", len(metadata))

	assert.Equal(t, 1, len(metadata))

	assert.Equal(t, metadata[0].ID, menuDishUUID)

	imageCount := 0

	for {
		part, err := mr.NextPart()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			t.Fatalf("error reading part: %v", err)
		}

		if !strings.HasPrefix(part.Header.Get("Content-Type"), "image/") {
			t.Fatalf("expected image part, got: %s", part.Header.Get("Content-Type"))
		}

		imageData, err := io.ReadAll(part)
		if err != nil {
			t.Fatalf("failed to read image part: %v", err)
		}

		t.Logf("read image of %d bytes", len(imageData))

		imageCount++
	}

	if imageCount != len(metadata) {
		t.Errorf("expected %d images, got %d", len(metadata), imageCount)
	}
}
