package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
)

func InitHystrix() {
	hystrix.ConfigureCommand("test", hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  100,
		RequestVolumeThreshold: 10,
		ErrorPercentThreshold:  25,
		SleepWindow:            5000,
	})
}
