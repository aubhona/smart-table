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

func TestCustomerOrderItemStateHappyPath(t *testing.T) {
	GetTestMutex().Lock()
	defer GetTestMutex().Unlock()
	defer CleanTest()

	_, _, _, hostCustomerUUID, orderUUID, _, menuDishUUID, _, err := CreateDefaultItems() //nolint
	assert.NoError(t, err)

	response, err := viewsCodegenCustomerOrderClient.PostCustomerV1OrderItemStateWithResponse(
		GetCtx(),
		&viewsCodegenCustomer.PostCustomerV1OrderItemStateParams{
			CustomerUUID: hostCustomerUUID,
			JWTToken:     "tipa_token_zhiest",
			OrderUUID:    orderUUID,
		},
		viewsCodegenCustomer.PostCustomerV1OrderItemStateJSONRequestBody{
			DishUUID: menuDishUUID,
			Comment:  &defaultComment,
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

	var metadata viewsCodegenCustomer.ItemStateInfo

	if err := json.Unmarshal(jsonBytes, &metadata); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	assert.Equal(t, menuDishUUID, metadata.ID)
	assert.Equal(t, "369.39", metadata.ResultPrice)
	assert.Equal(t, defaultItemsCount, metadata.Count)

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

	if imageCount != 1 {
		t.Errorf("expected 1 images, got %d", imageCount)
	}
}
