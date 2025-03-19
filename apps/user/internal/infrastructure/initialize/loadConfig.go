package initialize

import (
	"fmt"
	"os"

	"github.com/bhtoan2204/user/global"
	"github.com/spf13/viper"
)

func InitLoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")
	configName := os.Getenv("GO_ENV")
	if configName == "" {
		configName = "local"
	}
	viper.SetConfigName(configName)

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode configuration: %s", err))
	}
}
