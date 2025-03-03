package initialize

import (
	"errors"
	"fmt"
	"hash/fnv"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
)

var (
	serviceCounters = make(map[string]int)
	mu              sync.Mutex

	serviceCache = make(map[string]serviceCacheItem)
	cacheMu      sync.Mutex
)

type serviceCacheItem struct {
	addresses []string
	expiry    time.Time
}

// Cache with TTL is 5s
func GetServiceAddresses(client *api.Client, serviceName string) ([]string, error) {
	cacheMu.Lock()
	item, found := serviceCache[serviceName]
	if found && time.Now().Before(item.expiry) {
		cacheMu.Unlock()
		return item.addresses, nil
	}
	cacheMu.Unlock()

	healthServices, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}
	if len(healthServices) == 0 {
		return nil, errors.New("service not found or not healthy")
	}

	var addresses []string
	for _, serviceEntry := range healthServices {
		svc := serviceEntry.Service
		host := svc.Address
		if host == "" {
			host = serviceEntry.Node.Address
		}
		if strings.Contains(host, ":") {
			var err error
			host, _, err = net.SplitHostPort(host)
			if err != nil {
				return nil, err
			}
		}
		address := fmt.Sprintf("%s:%d", host, svc.Port)
		addresses = append(addresses, address)
	}

	cacheMu.Lock()
	serviceCache[serviceName] = serviceCacheItem{
		addresses: addresses,
		expiry:    time.Now().Add(5 * time.Second),
	}
	cacheMu.Unlock()

	return addresses, nil
}

func selectInstance(addresses []string, clientIP, serviceName string) string {
	if clientIP != "" {
		h := fnv.New32a()
		h.Write([]byte(clientIP))
		hashValue := h.Sum32()
		index := int(hashValue) % len(addresses)
		return addresses[index]
	}
	mu.Lock()
	index := serviceCounters[serviceName] % len(addresses)
	serviceCounters[serviceName]++
	mu.Unlock()
	return addresses[index]
}

func newReverseProxy(targetAddress, pathPrefix string) (*httputil.ReverseProxy, error) {
	targetURL, err := url.Parse("http://" + targetAddress)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = strings.TrimPrefix(req.URL.Path, pathPrefix)
	}
	return proxy, nil
}

func serviceProxy(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		addresses, err := GetServiceAddresses(global.ConsulClient, serviceName)
		if err != nil {
			global.Logger.Error("Service not found", zap.Error(err))
			response.ErrorNotFoundResponse(c, response.ErrorServiceUnavailable)
			return
		}
		clientIP := c.ClientIP()
		selected := selectInstance(addresses, clientIP, serviceName)
		global.Logger.Info("Resolved service address", zap.String("address", selected), zap.String("clientIP", clientIP))

		proxy, err := newReverseProxy(selected, "/"+serviceName)
		if err != nil {
			global.Logger.Error("Failed to create reverse proxy", zap.Error(err))
			response.ErrorInternalServerResponse(c, 500)
			return
		}
		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = selected
			req.URL.Path = strings.Replace(req.URL.Path, "/"+serviceName, "", 1)
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
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
	// r.Use(middleware.ApiLogMiddleware())

	V1ApiGroup := r.Group("/api/v1")
	{
		V1ApiGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})
		V1ApiGroup.GET("/test-kafka", func(c *gin.Context) {
			go ProduceMessage("test-key", "test-message")
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})

		// V1ApiGroup.Any("/user-service/*any", serviceProxy("user-service"))
		V1ApiGroup.GET("/user-service/users/profile", middleware.AuthenticationMiddleware(), serviceProxy("user-service"))
		V1ApiGroup.POST("/user-service/users/create", serviceProxy("user-service"))
		V1ApiGroup.POST("/user-service/users/login", serviceProxy("user-service"))
		V1ApiGroup.POST("/user-service/users/refresh", serviceProxy("user-service"))
		V1ApiGroup.GET("/user-service/users", serviceProxy("user-service"))
		V1ApiGroup.POST("/user-service/users/logout", serviceProxy("user-service"))

		// // Query
		// r.GET("", ctrl.SearchUser)
		V1ApiGroup.Any("/video-service/*any", serviceProxy("video-service"))
	}
	global.Logger.Info("Router initialized successfully")
	return r
}
