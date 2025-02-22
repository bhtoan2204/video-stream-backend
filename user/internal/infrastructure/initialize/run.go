package initialize

import (
	"os"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/domain/repository/shared"
	"github.com/bhtoan2204/user/internal/infrastructure/consumer"
)

func Run() {
	InitLoadConfig()
	InitLogger()
	InitListener()
	InitConsul()
	InitDB()
	InitKafka()
	InitElasticsearch()
	eSRepository := shared.NewRepositories(global.ESClient)
	go consumer.ConsumeMessage(eSRepository)
	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
