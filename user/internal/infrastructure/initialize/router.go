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
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

func SetupUserRoutes(api *gin.RouterGroup) {
	refreshTokenRepository := repository.NewRefreshTokenRepository(global.MDB)
	refreshTokenService := service.NewRefreshTokenService(refreshTokenRepository)
	userContainer, err := dependency.BuildUserContainer()
	if err != nil {
		global.Logger.Fatal("Failed to build user container", zap.Error(err))
		os.Exit(1)
	}
	userService := service.NewUserService(userContainer.UserRepository, userContainer.ESUserRepository, refreshTokenService)

	commandBus := command.SetUpCommandBus(&shared.ServiceDependencies{
		UserService: userService,
	})
	queryBus := query.SetUpQueryBus(&shared.ServiceDependencies{
		UserService: userService,
	})

	userGroup := api.Group("/users")
	controller.NewUserController(commandBus, queryBus, userGroup)
}

func SetupAuthRoutes(api *gin.RouterGroup) {
	refreshTokenRepository := repository.NewRefreshTokenRepository(global.MDB)
	refreshTokenService := service.NewRefreshTokenService(refreshTokenRepository)
	userContainer, err := dependency.BuildUserContainer()
	if err != nil {
		global.Logger.Fatal("Failed to build user container", zap.Error(err))
		os.Exit(1)
	}
	userService := service.NewUserService(userContainer.UserRepository, userContainer.ESUserRepository, refreshTokenService)

	commandBus := command.SetUpCommandBus(&shared.ServiceDependencies{
		UserService: userService,
	})

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
	ginMode := os.Getenv("GIN_MODE")
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

	apiV1 := r.Group("/api/v1/user-service")
	{
		SetupHealthRoutes(apiV1)
		SetupUserRoutes(apiV1)
		SetupAuthRoutes(apiV1)
	}
	return r
}
