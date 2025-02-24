package initialize

import (
	"os"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/infrastructure/debezium"
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
	go debezium.ConsumeCDC()
	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
