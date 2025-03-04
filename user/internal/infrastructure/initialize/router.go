package initialize

import (
	"net"
	"os"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command"
	"github.com/bhtoan2204/user/internal/application/controller"
	"github.com/bhtoan2204/user/internal/application/query"
	"github.com/bhtoan2204/user/internal/application/shared"
	"github.com/bhtoan2204/user/internal/dependency"
	"github.com/bhtoan2204/user/internal/domain/service"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/repository"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Refresh token
	refreshTokenRepository := repository.NewRefreshTokenRepository(global.MDB)
	refreshTokenService := service.NewRefreshTokenService(refreshTokenRepository)

	// User
	userContainer, err := dependency.BuildUserContainer()
	if err != nil {
		global.Logger.Fatal("Failed to build user container", zap.Error(err))
		os.Exit(1)
	}
	userService := service.NewUserService(userContainer.UserRepository, userContainer.ESUserRepository, refreshTokenService)

	// Command and Query
	commandBus := command.SetUpCommandBus(&shared.ServiceDependencies{
		UserService: userService,
	})

	queryBus := query.SetUpQueryBus(&shared.ServiceDependencies{
		UserService: userService,
	})

	r.Use(otelgin.Middleware("user-service"))

	apiV1 := r.Group("/api/v1/user-service")
	{
		apiV1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
				"port":    global.Listener.Addr().(*net.TCPAddr).Port,
			})
		})

		userGroup := apiV1.Group("/users")
		controller.NewUserController(commandBus, queryBus, userGroup)
	}
	return r
}
