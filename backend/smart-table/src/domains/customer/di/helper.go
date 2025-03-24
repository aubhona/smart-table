package di

import (
	"bytes"
	"io"
	"strings"

	"github.com/smart-table/src/config"

	"github.com/gin-gonic/gin"
)

func getRequestBody(c *gin.Context, cfg *config.Config) string {
	if c.Request.Body == nil {
		return ""
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return "ERROR_READING_BODY"
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	bodyStr := string(bodyBytes)
	if len(bodyStr) > cfg.Logging.Server.RequestSymLimit {
		return bodyStr[:cfg.Logging.Server.RequestSymLimit] + "..."
	}

	return bodyStr
}

func getRequestHeaders(c *gin.Context) string {
	headers := c.Request.Header
	headerStrings := make([]string, 0, len(headers))

	for key, values := range headers {
		if key == "Authorization" || key == "Cookie" {
			headerStrings = append(headerStrings, key+": [HIDDEN]")
		} else {
			headerStrings = append(headerStrings, key+": "+strings.Join(values, ", "))
		}
	}

	return strings.Join(headerStrings, "; ")
}

type responseRecorder struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
