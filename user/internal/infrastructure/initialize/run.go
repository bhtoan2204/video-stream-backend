package initialize

import (
	"os"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/event"
	"github.com/bhtoan2204/user/internal/application/shared"
	"github.com/bhtoan2204/user/internal/dependency"
	"github.com/bhtoan2204/user/internal/infrastructure/db/elasticsearch"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql"
	"github.com/bhtoan2204/user/internal/infrastructure/grpc"
	"go.uber.org/zap"
)

func Run() {
	// Initialize configurations
	InitLoadConfig()
	InitLogger()
	InitListener()

	// Initialize consul
	deregisterService := InitConsul()
	defer deregisterService()

	// Initialize write database
	mysql.InitDB()
	defer mysql.CloseDB()

	// Initialize kafka
	InitKafka()
	defer global.KafkaProducer.Close()
	defer global.KafkaConsumer.Close()

	// Initialize read database
	elasticsearch.InitElasticsearch()

	// Initialize tracing
	tracerShutdown := InitProvider()
	defer tracerShutdown()

	// Initialize container
	userContainer, err := dependency.BuildUserContainer()

	if err != nil {
		global.Logger.Fatal("Failed to build user container", zap.Error(err))
		os.Exit(1)
	}

	// Initialize gRPC server
	grpc.StartGrpcServer(userContainer.UserRepository)

	// Initialize event bus
	eventBus := *event.SetUpEventBus(&shared.ListenerDependencies{
		UserListener: userContainer.UserListener,
	})

	debezium := NewDebezium(eventBus)
	debeziumConsumer := debezium.Start()
	defer debeziumConsumer.Close()

	r := InitRouter()
	if err := r.RunListener(global.Listener); err != nil {
		os.Exit(1)
	}
}
