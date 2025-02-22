package initialize

import (
	"net"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command"
	"github.com/bhtoan2204/user/internal/application/controller"
	"github.com/bhtoan2204/user/internal/application/service"
	"github.com/bhtoan2204/user/internal/application/shared"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/repository"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	userRepository := repository.NewUserRepository(global.MDB)
	userService := service.NewUserService(userRepository)

	commandBus := command.SetUpCommandBus(&shared.ServiceDependencies{
		UserService: userService,
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
				"port":    global.Listener.Addr().(*net.TCPAddr).Port,
			})
		})

		userGroup := apiV1.Group("/users")
		controller.NewUserController(commandBus, userGroup)
	}
	return r
}
