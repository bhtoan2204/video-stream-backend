package initialize

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	serviceCounters = make(map[string]int)
	mu              sync.Mutex
)

func getServiceAddress(client *api.Client, serviceName string) (string, error) {
	services, err := client.Agent().Services()
	if err != nil {
		return "", err
	}

	var availableServices []string
	for _, service := range services {
		if service.Service == serviceName {
			host := service.Address
			if strings.Contains(host, ":") {
				var err error
				host, _, err = net.SplitHostPort(host)
				if err != nil {
					return "", err
				}
			}
			address := host + ":" + strconv.Itoa(service.Port)
			availableServices = append(availableServices, address)
		}
	}

	if len(availableServices) == 0 {
		return "", errors.New("service not found")
	}

	mu.Lock()
	index := serviceCounters[serviceName] % len(availableServices)
	serviceCounters[serviceName]++
	mu.Unlock()

	return availableServices[index], nil
}

func userServiceProxy(c *gin.Context) {
	serviceAddress, err := getServiceAddress(global.ConsulClient, "user-service")
	if err != nil {
		global.Logger.Error("User-service not found", zap.Error(err))
		response.ErrorNotFoundResponse(c, 404)
		return
	}
	targetURL, err := url.Parse("http://" + serviceAddress)
	fmt.Println("targetURL", targetURL)
	if err != nil {
		global.Logger.Error("Failed to parse URL", zap.Error(err))
		response.ErrorInternalServerResponse(c, 500)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = strings.Replace(req.URL.Path, "/user-service", "", 1)
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func videoServiceProxy(c *gin.Context) {
	serviceAddress, err := getServiceAddress(global.ConsulClient, "video-service")
	if err != nil {
		global.Logger.Error("Video-service not found", zap.Error(err))
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
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = strings.Replace(req.URL.Path, "/video-service", "", 1)
	}
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

	V1ApiGroup := r.Group("/api/v1")
	{
		V1ApiGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
		V1ApiGroup.GET("/test-kafka", func(c *gin.Context) {
			go ProduceMessage("test-key", "test-message")
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})

		V1ApiGroup.Any("/user-service/*any", userServiceProxy)
		V1ApiGroup.Any("/video-service/*any", videoServiceProxy)
	}
	global.Logger.Info("Router initialized successfully")
	return r
}
