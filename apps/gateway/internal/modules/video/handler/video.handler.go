package handler

import (
	"net/http"

	"github.com/bhtoan2204/gateway/internal/consul"
	"github.com/bhtoan2204/gateway/internal/modules/video/dto"
	"github.com/gin-gonic/gin"
)

// UploadVideo godoc
// @Summary Upload a video
// @Description Upload a video to the server
// @Tags video
// @Accept json
// @Produce json
// @Param video body dto.UploadVideoRequest true "Video details"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ResponseData
// @Failure 500 {object} response.ResponseData
// @Router /video-service/videos [post]
func UploadVideo(c *gin.Context) {
	var req dto.UploadVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consul.ServiceProxy("video-service")(c)
}

// GetVideoByURL godoc
// @Summary Get a video by URL
// @Description Get a video by URL
// @Tags video
// @Accept json
// @Produce json
// @Param url query string true "URL"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ResponseData
// @Failure 500 {object} response.ResponseData
// @Router /video-service/videos/url [get]
func GetVideoByURL(c *gin.Context) {
	consul.ServiceProxy("video-service")(c)
}

// GetPresignedURL godoc
// @Summary Get a presigned URL
// @Description Get a presigned URL
// @Tags video
// @Accept json
// @Produce json
// @Param url query string true "URL"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ResponseData
// @Failure 500 {object} response.ResponseData
// @Router /video-service/videos/presigned-url [get]
func GetPresignedURL(c *gin.Context) {
	consul.ServiceProxy("video-service")(c)
}
