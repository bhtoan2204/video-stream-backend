package router

import (
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/internal/modules/user/handler"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(api *gin.RouterGroup, instrument func(gin.HandlerFunc) gin.HandlerFunc) {
	userGroup := api.Group("/user-service/users")
	{
		userGroup.GET("", middleware.AuthenticationMiddleware(), instrument(handler.GetUserProfile))
		userGroup.PUT("", middleware.AuthenticationMiddleware(), instrument(handler.UpdateUser))
		userGroup.POST("", instrument(handler.CreateUser))
		userGroup.GET("/search", instrument(handler.SearchUser))
	}
}
