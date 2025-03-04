package initialize

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/gin-contrib/secure"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
)

var rrCounter uint64

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

func customResolverDial(ctx context.Context, network, address string) (net.Conn, error) {
	if global.Config.Server.Mode == "local" {
		return net.Dial("udp", "127.0.0.1:8600")
	}
	return net.Dial("udp", "consul:8600")
}

func consulDialContext(serviceName string) func(ctx context.Context, network, addr string) (net.Conn, error) {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial:     customResolverDial,
	}
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		_, srvs, err := resolver.LookupSRV(ctx, serviceName, "tcp", "service.consul")
		if err != nil {
			return nil, fmt.Errorf("failed to lookup SRV for %s: %w", serviceName, err)
		}
		if len(srvs) == 0 {
			return nil, fmt.Errorf("no SRV records found for %s", serviceName)
		}

		idx := int(rrCounter % uint64(len(srvs)))
		rrCounter++
		selected := srvs[idx]

		targetHost := strings.TrimSuffix(selected.Target, ".")
		targetPort := selected.Port
		target := fmt.Sprintf("%s:%d", targetHost, targetPort)

		global.Logger.Info("Consul SRV lookup result (round robin)", zap.String("target", target))

		dialer := net.Dialer{
			Timeout:  30 * time.Second,
			Resolver: resolver,
		}
		return dialer.DialContext(ctx, network, target)
	}
}

func newReverseProxyWithSRV(serviceName, pathPrefix string) (*httputil.ReverseProxy, error) {
	targetURL, err := url.Parse("http://" + serviceName)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	transport := &http.Transport{
		DialContext:         consulDialContext(serviceName),
		TLSHandshakeTimeout: 10 * time.Second,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		// TODO: Add more transport options like TLS, etc.
	}
	proxy.Transport = transport

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = strings.TrimPrefix(req.URL.Path, pathPrefix)
		req.Host = serviceName

		ctx := req.Context()
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	}
	return proxy, nil
}

func serviceProxy(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy, err := newReverseProxyWithSRV(serviceName, "/"+serviceName)
		if err != nil {
			global.Logger.Error("Failed to create reverse proxy", zap.Error(err))
			response.ErrorInternalServerResponse(c, http.StatusInternalServerError)
			return
		}
		proxy.ServeHTTP(c.Writer, c.Request)
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
		userGroup.GET("/profile", middleware.AuthenticationMiddleware(), instrument(serviceProxy("user-service")))
		userGroup.POST("/create", instrument(serviceProxy("user-service")))
		userGroup.GET("", instrument(serviceProxy("user-service")))
	}
}

func SetupAuthRoutes(api *gin.RouterGroup, instrument func(gin.HandlerFunc) gin.HandlerFunc) {
	authGroup := api.Group("/user-service/auth")
	{
		authGroup.POST("/login", instrument(serviceProxy("user-service")))
		authGroup.POST("/refresh", instrument(serviceProxy("user-service")))
		authGroup.POST("/logout", instrument(serviceProxy("user-service")))
	}
}

func SetupVideoRoutes(api *gin.RouterGroup) {
	api.Any("/video-service/*any", serviceProxy("video-service"))
}
