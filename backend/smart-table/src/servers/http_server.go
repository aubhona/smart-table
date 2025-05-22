package servers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	viewsCodegenAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
	viewsCodegenAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
	viewsCodegenAdminUser "github.com/smart-table/src/views/codegen/admin_user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smart-table/src/config"
	"github.com/smart-table/src/dependencies"
	appAdmin "github.com/smart-table/src/domains/admin/app/services"
	appCustomer "github.com/smart-table/src/domains/customer/app/services"
	"github.com/smart-table/src/utils"
	viewsPlace "github.com/smart-table/src/views/admin/v1/place"
	viewsRestaurant "github.com/smart-table/src/views/admin/v1/restaurant"
	viewsUser "github.com/smart-table/src/views/admin/v1/user"
	viewsCodegenCustomer "github.com/smart-table/src/views/codegen/customer"
	viewsCodegenCustomerOrder "github.com/smart-table/src/views/codegen/customer_order"
	viewsCustomer "github.com/smart-table/src/views/customer/v1"
	viewsCustomerOrder "github.com/smart-table/src/views/customer/v1/order"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func NewGinRouter(container *dig.Container, deps *dependencies.Dependencies) *gin.Engine {
	router := gin.New()

	//nolint
	// router.SetTrustedProxies() Think about security

	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = deps.Config.App.Cors.AllowOrigins
	cfg.AllowMethods = deps.Config.App.Cors.AllowMethods
	cfg.AllowHeaders = deps.Config.App.Cors.AllowHeaders
	cfg.AllowCredentials = deps.Config.App.Cors.AllowCredentials

	router.
		Use(GinZapResponseLogger(deps.Logger, deps.Config)).
		Use(GinZapLogger(deps.Logger, deps.Config)).
		Use(GinZapRecovery(deps.Logger)).
		Use(func(c *gin.Context) {
			c.Set(utils.DiContainerName, container)
			c.Set(utils.DependenciesName, deps)
			c.Next()
		}).Use(cors.New(cfg))

	privateAdmin := router.Group("/")
	privateCustomer := router.Group("/")

	if deps.Config.App.Admin.Jwt.Enable {
		jwtService, err := utils.GetFromDiContainer[*appAdmin.JwtService](container)
		if err != nil {
			deps.Logger.Error(fmt.Sprintf("Error while getting JWT service: %v", err))
			panic(err)
		}

		privateAdmin.Use(JWTAuthMiddleware(deps.Logger, jwtService, "User-UUID"))
	}

	if deps.Config.App.Customer.Jwt.Enable {
		jwtService, err := utils.GetFromDiContainer[*appCustomer.JwtService](container)
		if err != nil {
			deps.Logger.Error(fmt.Sprintf("Error while getting JWT service: %v", err))
			panic(err)
		}

		privateCustomer.Use(JWTAuthMiddleware(deps.Logger, jwtService, "Customer-UUID"))
	}

	customerStrictHandler := viewsCodegenCustomer.NewStrictHandler(&viewsCustomer.CustomerV1Handler{}, nil)
	viewsCodegenCustomer.RegisterHandlers(router, customerStrictHandler)

	customerOrderStrictHandler := viewsCodegenCustomerOrder.NewStrictHandler(&viewsCustomerOrder.CustomerV1OrderHandler{}, nil)
	viewsCodegenCustomerOrder.RegisterHandlers(privateCustomer, customerOrderStrictHandler)

	adminUserStrictHandler := viewsCodegenAdminUser.NewStrictHandler(&viewsUser.AdminV1UserHandler{}, nil)
	viewsCodegenAdminUser.RegisterHandlers(router, adminUserStrictHandler)

	adminRestaurantStrictHandler := viewsCodegenAdminRestaurant.NewStrictHandler(&viewsRestaurant.AdminV1RestaurantHandler{}, nil)
	viewsCodegenAdminRestaurant.RegisterHandlers(privateAdmin, adminRestaurantStrictHandler)

	adminPlaceStrictHandler := viewsCodegenAdminPlace.NewStrictHandler(&viewsPlace.AdminV1PlaceHandler{}, nil)
	viewsCodegenAdminPlace.RegisterHandlers(privateAdmin, adminPlaceStrictHandler)

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

		cookies := c.Request.Cookies()

		var cookieString string
		for _, cookie := range cookies {
			cookieString += cookie.Name + "=" + cookie.Value + "; "
		}

		logger.Info("HTTP Request",
			zap.String("method", method),
			zap.String("uri", path),
			zap.String("query_params", queryParams),
			zap.String("request_headers", requestHeaders),
			zap.String("request_body", requestBody),
			zap.String("cookies", cookieString),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
		)

		c.Next()
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

func JWTAuthMiddleware(logger *zap.Logger, jwtService utils.JwtService, headerName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("JWT-Token")

		userUUID, err := uuid.Parse(c.GetHeader(headerName))
		if err != nil {
			logger.Error(fmt.Sprintf("Error while parsing user_uuid: %v", err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "invalid_headers",
				"message": "Invalid user_uuid header",
			})

			return
		}

		if jwtService.ValidateJWT(tokenString, userUUID) != nil {
			logger.Warn("Invalid JWT token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "invalid_token",
				"message": "Invalid authentication token",
			})

			return
		}

		c.Next()
	}
}
