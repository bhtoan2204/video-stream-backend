package initialize

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bhtoan2204/user/global"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

func getDockerInternalIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				// Bỏ qua nếu không phải IPv4
				continue
			}
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("Could not find the internal IP address")
}

func InitConsul() {
	globalConsulConfig := global.Config.ConsulConfig

	consulConfig := api.DefaultConfig()
	consulConfig.Address = globalConsulConfig.Address
	consulConfig.Scheme = globalConsulConfig.Scheme
	consulConfig.Datacenter = globalConsulConfig.DataCenter
	consulConfig.Token = globalConsulConfig.Token

	serviceID := uuid.New().String()
	servicePort := global.Listener.Addr().(*net.TCPAddr).Port
	serviceAddress, err := getDockerInternalIP()
	if err != nil {
		panic(fmt.Sprintf("Failed to get Docker internal IP: ", err))
	}

	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "user-service",
		Address: consulConfig.Address,
		Port:    global.Listener.Addr().(*net.TCPAddr).Port,
		Tags:    []string{"api", "user"},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/api/v1/health", serviceAddress, servicePort),
			Method:   "GET",
			Interval: "10s",
			Timeout:  "5s",
			Notes:    "Basic health check in user service " + fmt.Sprintf("http://%s:%d/api/v1/health", serviceAddress, servicePort),
		},
	}
	fmt.Printf("http://%s:%d/api/v1/health", serviceAddress, servicePort)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		// global.Logger.Error("Failed to connect to Consul:", zap.Error(err))
		panic(err)
	}
	global.ConsulClient = consulClient

	err = global.ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		// global.Logger.Error("Failed to register service:", zap.Error(err))
		panic(err)
	}
	go handleShutdown(serviceID)
}

func handleShutdown(serviceID string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	// global.Logger.Info("Shutting down service, unregistering from Consul...", zap.String("serviceID", serviceID))

	err := global.ConsulClient.Agent().ServiceDeregister(serviceID)
	if err != nil {
		// global.Logger.Error("Failed to unregister service from Consul", zap.Error(err))
		fmt.Print("Failed to unregister service from Consul")
	} else {
		// global.Logger.Info("Service unregistered successfully from Consul")
		fmt.Print("Service unregistered successfully from Consul")
	}

	os.Exit(0)
}
