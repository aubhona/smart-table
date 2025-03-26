package servers

import (
	"bytes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/smart-table/src/config"
	"github.com/smart-table/src/dependencies"
	"github.com/smart-table/src/utils"
	viewsUser "github.com/smart-table/src/views/admin/v1/user"
	viewsAdmin "github.com/smart-table/src/views/codegen/admin_user"
	viewsCustomer "github.com/smart-table/src/views/codegen/customer"
	viewsOrder "github.com/smart-table/src/views/customer/v1/order"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func GetRouter(container *dig.Container, deps *dependencies.Dependencies) *gin.Engine {
	router := gin.New()

	//nolint
	// router.SetTrustedProxies() Think about security

	config := cors.DefaultConfig()
	config.AllowOrigins = deps.Config.App.Cors.AllowOrigins
	config.AllowMethods = deps.Config.App.Cors.AllowMethods
	config.AllowHeaders = deps.Config.App.Cors.AllowHeaders
	config.AllowCredentials = deps.Config.App.Cors.AllowCredentials

	router.
		Use(GinZapResponseLogger(deps.Logger, deps.Config)).
		Use(GinZapLogger(deps.Logger, deps.Config)).
		Use(GinZapRecovery(deps.Logger)).
		Use(func(c *gin.Context) {
			c.Set(utils.DiContainerName, container)
			c.Next()
		}).Use(cors.New(config))

	customerStrictHandler := viewsCustomer.NewStrictHandler(&viewsOrder.CustomerV1OrderHandler{}, nil)
	viewsCustomer.RegisterHandlers(router, customerStrictHandler)

	adminStrictHandler := viewsAdmin.NewStrictHandler(&viewsUser.AdminV1UserHandler{}, nil)
	viewsAdmin.RegisterHandlers(router, adminStrictHandler)

	return router
}

func GinZapRecovery(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
				)
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}

func GinZapLogger(logger *zap.Logger, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		requestBody := getRequestBody(c, cfg)
		queryParams := c.Request.URL.Query().Encode()
		requestHeaders := getRequestHeaders(c)

		c.Next()

		logger.Info("HTTP Request",
			zap.String("method", method),
			zap.String("uri", path),
			zap.String("query_params", queryParams),
			zap.String("request_headers", requestHeaders),
			zap.String("request_body", requestBody),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
		)
	}
}

func GinZapResponseLogger(logger *zap.Logger, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		resBody := new(bytes.Buffer)
		rec := &responseRecorder{ResponseWriter: c.Writer, body: resBody}
		c.Writer = rec

		c.Next()

		responseBody := resBody.String()
		if len(responseBody) > cfg.Logging.Server.ResponseSymLimit {
			responseBody = responseBody[:cfg.Logging.Server.ResponseSymLimit] + "..."
		}

		logger.Info("HTTP Response",
			zap.Int("status", c.Writer.Status()),
			zap.String("response_body", responseBody),
		)
	}
}
