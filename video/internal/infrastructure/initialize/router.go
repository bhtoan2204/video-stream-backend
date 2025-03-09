package initialize

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})
		MainGroup.POST("/upload", func(c *gin.Context) {
			fileHeader, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
				return
			}

			file, err := fileHeader.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not open file"})
				return
			}
			defer file.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err != nil {
				log.Printf("Error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Can not connect to S3"})
				return
			}

			fileKey := utils.GenFileKey(fileHeader.Filename, "video")

			if err = global.S3Client.UploadFile(ctx, fileKey, file); err != nil {
				log.Printf("Error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload file failed"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":  "Upload file successfully",
				"file_key": fileKey,
			})
		})
	}
	return r
}
