package main

import (
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/src/custom"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/di"
	views_order "github.com/es-debug/backend-academy-2024-go-template/src/views/codegen/order"
	views "github.com/es-debug/backend-academy-2024-go-template/src/views/mobile/v1/order/create"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	deps := custom.InitDependencies()
	logger := deps.Logger

	container, err := di.BuildContainer(deps)
	if err != nil {
		logger.Fatal(err.Error())
	}

	router := gin.New()

	gin.Default()

	router.
		Use(GinZapLogger(logger)).
		Use(GinZapRecovery(logger)).
		Use(func(c *gin.Context) {
			c.Set(di.KDiContainerName, container)
			c.Next()
		})

	strictHandler := views_order.NewStrictHandler(&views.MobileV1OrderCreateHandler{}, nil)
	views_order.RegisterHandlers(router, strictHandler)

	err = router.Run(deps.Config.App.Port)
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func GinZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		logger.Info("Request",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
		)
	}
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
