package router

import (
	"github.com/bhtoan2204/gateway/internal/middleware"
	"github.com/bhtoan2204/gateway/internal/modules/video/handler"
	"github.com/gin-gonic/gin"
)

func SetupVideoRoutes(api *gin.RouterGroup, instrument func(gin.HandlerFunc) gin.HandlerFunc) {
	videoGroup := api.Group("/video-service/videos")
	{
		videoGroup.GET("/:url", instrument(handler.GetVideoByURL))
		videoGroup.POST("", middleware.AuthenticationMiddleware(), instrument(handler.UploadVideo))
		videoGroup.GET("/presigned_url/download", middleware.AuthenticationMiddleware(), instrument(handler.GetPresignedURLDownload))
		videoGroup.GET("/presigned_url/upload", middleware.AuthenticationMiddleware(), instrument(handler.GetPresignedURLUpload))
	}
}
