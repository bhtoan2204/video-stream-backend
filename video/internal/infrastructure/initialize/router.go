package initialize

import (
	"net"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/internal/application/command_bus"
	"github.com/bhtoan2204/video/internal/application/controller"
	"github.com/bhtoan2204/video/internal/application/shared"
	"github.com/bhtoan2204/video/internal/domain/services"
	"github.com/bhtoan2204/video/internal/infrastructure/db/mysql/repository"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupVideoRoutes(api *gin.RouterGroup) {
	videoRepository := repository.NewVideoRepository(global.MDB)
	videoService := services.NewVideoService(videoRepository)

	commandBus := command_bus.SetUpCommandBus(&shared.ServiceDependencies{
		VideoService: videoService,
	})

	videoGroup := api.Group("/videos")
	controller.NewVideoController(commandBus, videoGroup)
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
	r.Use(otelgin.Middleware("video-service"))

	apiV1 := r.Group("/api/v1/video-service")
	{
		SetupHealthRoutes(apiV1)
		SetupVideoRoutes(apiV1)
	}

	return r
}
