package initialize

import (
	"net"

	"github.com/bhtoan2204/comment/global"
	"github.com/bhtoan2204/comment/internal/application/command_bus"
	"github.com/bhtoan2204/comment/internal/application/controller"
	"github.com/bhtoan2204/comment/internal/application/shared"
	"github.com/bhtoan2204/comment/internal/domain/service"
	"github.com/bhtoan2204/comment/internal/infrastructure/db/mysql/repository"
	"github.com/bhtoan2204/comment/internal/infrastructure/grpc"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupHealthRoutes(api *gin.RouterGroup) {
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
			"port":    global.Listener.Addr().(*net.TCPAddr).Port,
		})
	})
}

func SetupCommentRoutes(api *gin.RouterGroup, commandBus *command_bus.CommandBus) {
	commentGroup := api.Group("/comments")
	controller.NewCommentController(commandBus, commentGroup)
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
	r.Use(otelgin.Middleware("comment-service"))

	commentRepository := repository.NewCommentRepository(global.MDB)
	videoPort := grpc.NewVideoAdapter(global.VideoGRPCClient)
	commentService := service.NewCommentService(commentRepository, videoPort)
	depsService := shared.ServiceDependencies{
		CommentService: commentService,
	}
	commandBus := command_bus.SetUpCommandBus(&depsService)
	apiV1 := r.Group("/api/v1/comment-service")
	{
		SetupHealthRoutes(apiV1)
		SetupCommentRoutes(apiV1, commandBus)
	}
	return r
}
