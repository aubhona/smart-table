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

	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer_order"

	"github.com/stretchr/testify/assert"
)

func TestCustomerOrderCatalogHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	userUUID, restaurantUUID, placeUUID, hostCustomerUUID, orderUUID, token, err := CreateDefaultOrder()
	assert.NoError(t, err)

	dishUUID, err := CreateDefaultRestaurantDish(token, userUUID, restaurantUUID)
	assert.NoError(t, err)

	menuDishUUID, err := CreateDefaultPlaceMenuDish(token, userUUID, placeUUID, dishUUID)
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.GetCustomerV1OrderCatalogWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.GetCustomerV1OrderCatalogParams{
			CustomerUUID: hostCustomerUUID,
			OrderUUID:    orderUUID,
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

	var metadata viewsCodegenCustomer.Catalog

	if err := json.Unmarshal(jsonBytes, &metadata); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	t.Logf("parsed %d menu dishes", len(metadata.Menu))

	assert.Equal(t, 1, len(metadata.Menu))
	assert.Equal(t, metadata.Menu[0].ID, menuDishUUID)
	assert.Equal(t, metadata.TotalPrice, "0")

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

	if imageCount != len(metadata.Menu) {
		t.Errorf("expected %d images, got %d", len(metadata.Menu), imageCount)
	}
}
