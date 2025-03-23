package router

import (
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/internal/modules/comment/handler"
	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(api *gin.RouterGroup, instrument func(gin.HandlerFunc) gin.HandlerFunc) {
	commentGroup := api.Group("/comment-service/comments")
	{
		commentGroup.POST("", middleware.AuthenticationMiddleware(), instrument(handler.CreateComment))
	}
}
