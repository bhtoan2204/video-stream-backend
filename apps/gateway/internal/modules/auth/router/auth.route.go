package router

import (
	"github.com/bhtoan2204/gateway/internal/modules/auth/handler"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(api *gin.RouterGroup, instrument func(gin.HandlerFunc) gin.HandlerFunc) {
	authGroup := api.Group("/user-service/auth")
	{
		authGroup.POST("/login", instrument(handler.Login))
		authGroup.POST("/refresh", instrument(handler.RefreshToken))
		authGroup.POST("/logout", instrument(handler.Logout))
		authGroup.GET("/2fa/setup", instrument(handler.Setup2FA))
		authGroup.POST("/2fa/verify", instrument(handler.Verify2FA))
	}
}
