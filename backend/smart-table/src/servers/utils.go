package servers

import (
	"bytes"
	"io"
	"strings"

	"go.uber.org/zap"
	"gopkg.in/telebot.v4"

	"github.com/gin-gonic/gin"
	"github.com/smart-table/src/config"
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

func botLogger(logger *zap.Logger) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			info := make(map[string]interface{})
			if c.Update().Message != nil {
				info["message_text"] = c.Update().Message.Text
				info["message_chat"] = c.Update().Message.Chat
				info["message_sender"] = c.Update().Message.Sender
			}

			logger.Info("Get update from telegram bot", zap.Any("update_info", info))

			return next(c)
		}
	}
}
