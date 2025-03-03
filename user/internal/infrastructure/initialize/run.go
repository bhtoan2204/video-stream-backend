package initialize

import (
	"os"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/event"
	"github.com/bhtoan2204/user/internal/application/shared"
	"github.com/bhtoan2204/user/internal/domain/listener"
	eSRepository "github.com/bhtoan2204/user/internal/infrastructure/db/elasticsearch/repository"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/repository"
)

func Run() {
	InitLoadConfig()
	InitLogger()
	InitListener()
	InitConsul()
	InitDB()
	InitKafka()
	InitElasticsearch()
	eSUserRepository := eSRepository.NewESUserRepository(global.ESClient)
	userRepository := repository.NewUserRepository(global.MDB)
	userListener := listener.NewUserListener(userRepository, eSUserRepository)
	InitGrpcServer(userRepository)
	eventBus := *event.SetUpEventBus(&shared.ListenerDependencies{
		UserListener: userListener,
	})

	debezium := NewDebezium(eventBus)
	go debezium.Start()

	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
