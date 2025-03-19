package initialize

import (
	"net"
	"os"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command_bus"
	"github.com/bhtoan2204/user/internal/application/controller"
	"github.com/bhtoan2204/user/internal/application/query_bus"
	"github.com/bhtoan2204/user/internal/application/shared"
	"github.com/bhtoan2204/user/internal/dependency"
	"github.com/bhtoan2204/user/internal/domain/service"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/repository"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

func SetupUserRoutes(api *gin.RouterGroup, commandBus *command_bus.CommandBus, queryBus *query_bus.QueryBus) {
	userGroup := api.Group("/users")
	controller.NewUserController(commandBus, queryBus, userGroup)
}

func SetupAuthRoutes(api *gin.RouterGroup, commandBus *command_bus.CommandBus, queryBus *query_bus.QueryBus) {
	authGroup := api.Group("/auth")
	controller.NewAuthController(commandBus, authGroup)
}

func SetupHealthRoutes(api *gin.RouterGroup) {
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
			"port":    global.Listener.Addr().(*net.TCPAddr).Port,
		})
	})
}

func InitRouter() *gin.Engine {
	ginMode := global.Config.Server.GinMode
	if ginMode == "" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(ginMode)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	secMiddleware := secure.New(secure.Config{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
	})
	r.Use(secMiddleware)
	r.Use(otelgin.Middleware("user-service"))

	userContainer, err := dependency.BuildUserContainer()
	if err != nil {
		global.Logger.Fatal("Failed to build user container", zap.Error(err))
		os.Exit(1)
	}
	refreshTokenRepository := repository.NewRefreshTokenRepository(global.MDB)
	refreshTokenService := service.NewRefreshTokenService(refreshTokenRepository)

	userService := service.NewUserService(userContainer.UserRepository, userContainer.ESUserRepository, refreshTokenService)

	userSettingRepository := repository.NewUserSettingRepository(global.MDB)
	userSettingService := service.NewUserSettingService(userSettingRepository)

	commandBus := command_bus.SetUpCommandBus(&shared.ServiceDependencies{
		UserService:        userService,
		UserSettingService: userSettingService,
	})
	queryBus := query_bus.SetUpQueryBus(&shared.ServiceDependencies{
		UserService: userService,
	})

	apiV1 := r.Group("/api/v1/user-service")
	{
		SetupHealthRoutes(apiV1)
		SetupUserRoutes(apiV1, commandBus, queryBus)
		SetupAuthRoutes(apiV1, commandBus, queryBus)
	}
	return r
}
