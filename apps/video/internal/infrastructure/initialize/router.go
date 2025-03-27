package initialize

import (
	"net"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/internal/application/command_bus"
	"github.com/bhtoan2204/video/internal/application/controller"
	"github.com/bhtoan2204/video/internal/application/shared"
	"github.com/bhtoan2204/video/internal/domain/services"
	"github.com/bhtoan2204/video/internal/infrastructure/db/mysql/repository"
	"github.com/bhtoan2204/video/internal/infrastructure/socketio"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

func SetupVideoRoutes(api *gin.RouterGroup, socketService *socketio.Service) {
	videoRepository := repository.NewVideoRepository(global.MDB)
	videoService := services.NewVideoService(videoRepository)

	commandBus := command_bus.SetUpCommandBus(&shared.ServiceDependencies{
		VideoService:  videoService,
		SocketService: socketService,
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

	// Initialize Socket.IO server
	socketServer, err := socketio.NewServer()
	redisAdapter := socketio.NewRedisAdapter(global.Redis)
	socketService := socketio.NewService(socketServer, redisAdapter)

	if err != nil {
		global.Logger.Fatal("Failed to create Socket.IO server", zap.Error(err))
	}

	// Initialize Socket.IO handler
	socketHandler := socketio.NewHandler(socketServer)

	// API routes
	apiV1 := r.Group("/api/v1/video-service")
	{
		SetupHealthRoutes(apiV1)
		SetupVideoRoutes(apiV1, socketService)
	}

	// Root Socket.IO endpoints
	r.GET("/socket.io/*any", gin.WrapH(socketHandler))
	r.POST("/socket.io/*any", gin.WrapH(socketHandler))
	r.Handle("OPTIONS", "/socket.io/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Status(204)
	})

	// Start Socket.IO server
	go socketServer.Serve()

	return r
}
