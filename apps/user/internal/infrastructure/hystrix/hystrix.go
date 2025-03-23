package hystrix

import (
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/bhtoan2204/user/global"
	"go.uber.org/zap"
)

func InitHystrix() {
	hystrix.ConfigureCommand("user-service", hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  100,
		ErrorPercentThreshold:  25,
		SleepWindow:            5000,
		RequestVolumeThreshold: 10,
	})

	err := hystrix.Do("user-service", func() error {
		global.Logger.Info("Attempting service call...")
		time.Sleep(2 * time.Second)
		return fmt.Errorf("service unavailable")
	}, func(err error) error {
		global.Logger.Error("Service call failed: ", zap.Error(err))
		return nil
	})

	if err != nil {
		global.Logger.Error("Final error: ", zap.Error(err))
	} else {
		global.Logger.Info("Service call succeeded")
	}
}
