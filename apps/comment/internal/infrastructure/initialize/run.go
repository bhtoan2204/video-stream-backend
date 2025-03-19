package initialize

import (
	"os"

	"github.com/bhtoan2204/comment/global"
	"github.com/bhtoan2204/comment/internal/infrastructure/tracing"
)

func Run() {
	InitLoadConfig()
	InitLogger()
	InitListener()

	// Initialize consul
	deregisterService := InitConsul()
	defer deregisterService()

	// Initialize tracing
	tracerShutdown := tracing.InitProvider()
	defer tracerShutdown()

	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
