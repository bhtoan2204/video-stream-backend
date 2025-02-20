package initialize

import (
	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/controller"
	"github.com/bhtoan2204/user/internal/application/service"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/repository"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	userRepository := repository.NewUserRepository(global.MDB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(r, userService)
	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
		UserGroup := MainGroup.Group("/users")
		{
			UserGroup.POST("", userController.CreateUser)
		}
	}
	return r
}
