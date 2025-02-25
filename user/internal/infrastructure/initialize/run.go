package initialize

import (
	"os"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/infrastructure/event_consumer"
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
	event_consumer.InitDebeziumConsumer()
	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
