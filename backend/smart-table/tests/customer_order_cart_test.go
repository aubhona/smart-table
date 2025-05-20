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

func TestCustomerOrderCartHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	_, _, _, hostCustomerUUID, orderUUID, _, menuDishUUID, _, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.GetCustomerV1OrderCartWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.GetCustomerV1OrderCartParams{
			CustomerUUID: hostCustomerUUID,
			JWTToken:     "tipa_token_zhiest",
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

	var metadata viewsCodegenCustomer.CartInfo

	if err := json.Unmarshal(jsonBytes, &metadata); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	t.Logf("parsed %d menu dishes", len(metadata.Items))

	assert.Equal(t, 1, len(metadata.Items))
	assert.Equal(t, menuDishUUID, metadata.Items[0].DishUUID)
	assert.Equal(t, "369.39", metadata.Items[0].ResultPrice)
	assert.Equal(t, defaultItemsCount, metadata.Items[0].Count)
	assert.Equal(t, "369.39", metadata.TotalPrice)

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

	if imageCount != len(metadata.Items) {
		t.Errorf("expected %d images, got %d", len(metadata.Items), imageCount)
	}
}
