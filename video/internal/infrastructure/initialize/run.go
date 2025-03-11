package initialize

import (
	"os"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/internal/infrastructure/db/mysql"
	"github.com/bhtoan2204/video/internal/infrastructure/tracing"
)

func Run() {
	InitLoadConfig()
	InitLogger()
	InitListener()

	deregisterService := InitConsul()
	defer deregisterService()

	mysql.InitDB()
	defer mysql.CloseDB()

	InitStorageService()

	tracerShutdown := tracing.InitProvider()
	defer tracerShutdown()

	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
