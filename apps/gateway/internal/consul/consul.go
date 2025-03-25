package consul

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
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

var rrCounter uint64

func InitConsul() {
	globalConsulConfig := global.Config.ConsulConfig

	consulConfig := api.DefaultConfig()
	consulConfig.Address = globalConsulConfig.Address
	consulConfig.Scheme = globalConsulConfig.Scheme
	consulConfig.Datacenter = globalConsulConfig.DataCenter
	consulConfig.Token = globalConsulConfig.Token

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		global.Logger.Error("Failed to connect to Consul:", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("Connected to Consul successfully")
	global.ConsulClient = consulClient
}

func CustomResolverDial(ctx context.Context, network, address string) (net.Conn, error) {
	if global.Config.Server.Mode == "local" {
		return net.Dial("udp", "127.0.0.1:8600")
	}
	return net.Dial("udp", "consul:8600")
}

func ConsulDialContext(serviceName string) func(ctx context.Context, network, addr string) (net.Conn, error) {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial:     CustomResolverDial,
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
		DialContext:         ConsulDialContext(serviceName),
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

func ServiceProxy(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy, err := newReverseProxyWithSRV(serviceName, "/"+serviceName)
		if err != nil {
			global.Logger.Error("Failed to create reverse proxy", zap.Error(err))
			response.ErrorInternalServerResponse(c, http.StatusInternalServerError)
			return
		}

		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			global.Logger.Error("Reverse proxy error", zap.Error(err))
			http.Error(rw, "Bad Gateway", http.StatusBadGateway)
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
