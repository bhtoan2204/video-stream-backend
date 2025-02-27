package initialize

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/pkg/response"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var (
	serviceCounters = make(map[string]int)
	mu              sync.Mutex
)

// func getServiceAddress(client *api.Client, serviceName string) (string, error) {
// 	healthServices, _, err := client.Health().Service(serviceName, "", true, nil)
// 	if err != nil {
// 		return "", err
// 	}

// 	var availableServices []string
// 	for _, serviceEntry := range healthServices {
// 		svc := serviceEntry.Service
// 		host := svc.Address
// 		if strings.Contains(host, ":") {
// 			var err error
// 			host, _, err = net.SplitHostPort(host)
// 			if err != nil {
// 				return "", err
// 			}
// 		}
// 		address := host + ":" + strconv.Itoa(svc.Port)
// 		availableServices = append(availableServices, address)
// 	}

// 	if len(availableServices) == 0 {
// 		return "", errors.New("service not found or not healthy")
// 	}

// 	mu.Lock()
// 	index := serviceCounters[serviceName] % len(availableServices)
// 	serviceCounters[serviceName]++
// 	mu.Unlock()

// 	return availableServices[index], nil
// }

func userServiceProxy(c *gin.Context) {
	targetURL, err := url.Parse("http://user-service.service.consul")
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
	targetURL, err := url.Parse("http://video-service.service.consul")
	if err != nil {
		global.Logger.Error("Failed to parse URL", zap.Error(err))
		response.ErrorInternalServerResponse(c, 500)
		return
	}
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
	// requestsPerSecond := rate.Limit(50)
	// burstSize := 10
	// rl := middleware.NewRateLimiter(requestsPerSecond, burstSize)

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
	// r.Use(middleware.RateLimitMiddleware(rl))
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
