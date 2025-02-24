package initialize

import (
	"os"

	"github.com/bhtoan2204/user/global"
)

func Run() {
	InitLoadConfig()
	InitLogger()
	InitListener()
	InitConsul()
	InitDB()
	InitKafka()
	InitElasticsearch()
	// InitGrpcClient()
	InitDebeziumConsumer()
	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
