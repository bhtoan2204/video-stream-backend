package router

import (
	"net/http"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/consul"
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/internal/websocket"
	"github.com/gin-contrib/secure"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
)

func NewInstrumentedHandler(counter metric.Int64Counter, commonLabels []attribute.KeyValue) func(gin.HandlerFunc) gin.HandlerFunc {
	return func(handler gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			ctx := c.Request.Context()
			counter.Add(ctx, 1, metric.WithAttributes(commonLabels...))
			span := trace.SpanFromContext(ctx)
			bag := baggage.FromContext(ctx)
			var baggageAttributes []attribute.KeyValue
			baggageAttributes = append(baggageAttributes, commonLabels...)
			for _, member := range bag.Members() {
				baggageAttributes = append(baggageAttributes, attribute.String("baggage key:"+member.Key(), member.Value()))
			}
			span.SetAttributes(baggageAttributes...)

			handler(c)
		}
	}
}

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "local" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	}

	r.RedirectTrailingSlash = false

	secMiddleware := secure.New(secure.Config{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
	})

	r.Use(secMiddleware)
	r.Use(otelgin.Middleware("gateway-server"))
	r.Use(gin.Logger())
	r.Use(middleware.CORSMiddleware())

	// Rate limiter
	requestsPerSecond := rate.Limit(50)
	burstSize := 10
	rl := middleware.NewRateLimiter(requestsPerSecond, burstSize)
	r.Use(middleware.RateLimitMiddleware(rl))
	r.Use(middleware.ApiLogMiddleware())
	r.Use(middleware.HMACMiddleware())

	meter := otel.Meter("gateway-server-meter")
	serverAttribute := attribute.String("controller", "gateway")
	commonLabels := []attribute.KeyValue{serverAttribute}
	requestCount, _ := meter.Int64Counter(
		"gateway_server/request_counts",
		metric.WithDescription("The number of requests received"),
	)
	instrument := NewInstrumentedHandler(requestCount, commonLabels)

	apiV1 := r.Group("/api/v1")
	{
		SetupHealthRoutes(apiV1)
		SetupUserRoutes(apiV1, instrument)
		SetupAuthRoutes(apiV1, instrument)
		SetupVideoRoutes(apiV1)
	}

	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/user/*any", func(c *gin.Context) {
			websocket.ProxyWebsocketWithConsul(c, "user-service", "/ws/user")
		})
		wsGroup.GET("/video/*any", func(c *gin.Context) {
			websocket.ProxyWebsocketWithConsul(c, "video-service", "/ws/video")
		})
	}

	global.Logger.Info("Router initialized successfully")
	return r
}

func SetupHealthRoutes(api *gin.RouterGroup) {
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}

func SetupUserRoutes(api *gin.RouterGroup, instrument func(gin.HandlerFunc) gin.HandlerFunc) {
	userGroup := api.Group("/user-service/users")
	{
		userGroup.GET("", middleware.AuthenticationMiddleware(), instrument(consul.ServiceProxy("user-service")))
		userGroup.PUT("", middleware.AuthenticationMiddleware(), instrument(consul.ServiceProxy("user-service")))
		userGroup.POST("", instrument(consul.ServiceProxy("user-service")))
		userGroup.GET("/search", instrument(consul.ServiceProxy("user-service")))
	}
}

func SetupAuthRoutes(api *gin.RouterGroup, instrument func(gin.HandlerFunc) gin.HandlerFunc) {
	authGroup := api.Group("/user-service/auth")
	{
		authGroup.POST("/login", instrument(consul.ServiceProxy("user-service")))
		authGroup.POST("/refresh", instrument(consul.ServiceProxy("user-service")))
		authGroup.POST("/logout", instrument(consul.ServiceProxy("user-service")))
		authGroup.GET("/2fa/setup", instrument(consul.ServiceProxy("user-service")))
		authGroup.POST("/2fa/verify", instrument(consul.ServiceProxy("user-service")))
	}
}

func SetupVideoRoutes(api *gin.RouterGroup) {
	videoGroup := api.Group("/video-service/videos")
	{
		videoGroup.GET("/:url", consul.ServiceProxy("video-service"))
		videoGroup.POST("", middleware.AuthenticationMiddleware(), consul.ServiceProxy("video-service"))
		videoGroup.GET("/presigned_url", middleware.AuthenticationMiddleware(), consul.ServiceProxy("video-service"))
	}
}
