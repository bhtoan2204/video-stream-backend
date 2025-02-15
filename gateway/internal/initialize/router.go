package initialize

import (
	"errors"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func getServiceAddress(client *api.Client, serviceName string) (string, error) {
	services, err := client.Agent().Services()
	if err != nil {
		return "", err
	}
	for _, service := range services {
		if service.Service == serviceName {
			return service.Address + ":" + strconv.Itoa(service.Port), nil
		}
	}
	return "", errors.New("service not found")
}

func userServiceProxy(c *gin.Context) {
	serviceAddress, err := getServiceAddress(global.ConsulClient, "user-service")
	if err != nil {
		global.Logger.Error("User-service not found", zap.Error(err))
		response.ErrorNotFoundResponse(c, 404)
		return
	}
	targetURL, err := url.Parse("http://" + serviceAddress)
	if err != nil {
		global.Logger.Error("Failed to parse URL", zap.Error(err))
		response.ErrorInternalServerResponse(c, 500)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	requestsPerSecond := rate.Limit(50)
	burstSize := 10
	rl := middleware.NewRateLimiter(requestsPerSecond, burstSize)

	if global.Config.Server.Mode == "local" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use(gin.Logger())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimitMiddleware(rl))
	r.Use(middleware.ApiLogMiddleware())

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})

		MainGroup.Any("/users", userServiceProxy)
	}
	global.Logger.Info("Router initialized successfully")
	return r
}
