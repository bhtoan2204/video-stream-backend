package controller

import (
	"reflect"
	"time"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/internal/application/command_bus"
	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/application/middleware"
	"github.com/bhtoan2204/video/internal/infrastructure/grpc/proto/user"
	"github.com/bhtoan2204/video/pkg/response"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

type VideoController struct {
	commandBus *command_bus.CommandBus
	// queryBus   *query.QueryBus
}

func NewVideoController(commandBus *command_bus.CommandBus, r *gin.RouterGroup) *VideoController {
	ctrl := &VideoController{
		commandBus: commandBus,
		// queryBus:   queryBus,
	}

	controllerName := reflect.TypeOf(ctrl).Elem().Name()
	meter := otel.Meter("video-server-meter")
	serverAttribute := attribute.String("controller", controllerName)
	commonLabels := []attribute.KeyValue{serverAttribute}
	requestCount, _ := meter.Int64Counter(
		"video_server/request_counts",
		metric.WithDescription("The number of requests received"),
	)
	instrument := middleware.NewInstrumentedHandler(requestCount, commonLabels)

	r.GET("/:url", instrument(ctrl.GetVideoByURL))
	r.POST("", middleware.AuthenticationMiddleware(), instrument(ctrl.UploadVideo))

	r.GET("/presigned_url/download", middleware.AuthenticationMiddleware(), instrument(ctrl.GetPresignedURLDownload))
	r.GET("/presigned_url/upload", middleware.AuthenticationMiddleware(), instrument(ctrl.GetPresignedURLUpload))

	return ctrl
}

func (controller *VideoController) UploadVideo(c *gin.Context) {
	userVal := c.Request.Context().Value("user")
	if userVal == nil {
		global.Logger.Error("User data not found in request context")
		response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
		return
	}

	userResp, ok := userVal.(*user.UserResponse)
	if !ok || userResp.Id == "" {
		global.Logger.Error("Invalid user data in context")
		response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
		return
	}

	var command command.UploadVideoCommand
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&command); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	fileMetadata, err := global.S3Client.GetFileMetadata(c.Request.Context(), command.FileKey)
	if err != nil {
		global.Logger.Error("Failed to get file metadata", zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	command.UploadedUser = userResp.Id
	command.FileSize = *fileMetadata.ContentLength
	command.ContentType = *fileMetadata.ContentType

	result, err := controller.commandBus.Dispatch(ctx, &command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, 2010, result)
}

func (controller *VideoController) GetPresignedURLDownload(c *gin.Context) {
	userVal := c.Request.Context().Value("user")
	if userVal == nil {
		global.Logger.Error("User data not found in request context")
		response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
		return
	}

	userResp, ok := userVal.(*user.UserResponse)
	if !ok || userResp.Id == "" {
		global.Logger.Error("Invalid user data in context")
		response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
		return
	}
	key := "dbe2d01d-0610-455d-a542-914643a205fb/videos/20250325151518.mp4"

	presignedDownloadUrl, err := global.S3Client.GeneratePresignedDownloadURL(c.Request.Context(), key, 15*time.Minute)
	if err != nil {
		global.Logger.Error("Failed to generate presigned URL", zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	response.SuccessResponse(c, 2000, presignedDownloadUrl)
}

func (controller *VideoController) GetPresignedURLUpload(c *gin.Context) {
	userVal := c.Request.Context().Value("user")
	if userVal == nil {
		global.Logger.Error("User data not found in request context")
		response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
		return
	}

	userResp, ok := userVal.(*user.UserResponse)
	if !ok || userResp.Id == "" {
		global.Logger.Error("Invalid user data in context")
		response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
		return
	}
	key := userResp.Id + "/" + "videos" + "/" + time.Now().Format("20060102150405") + ".mp4"
	presignedUploadUrl, err := global.S3Client.GeneratePresignedUploadURL(c.Request.Context(), key, 15*time.Minute)
	if err != nil {
		global.Logger.Error("Failed to generate presigned URL", zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	response.SuccessResponse(c, 2000, presignedUploadUrl)
}

func (controller *VideoController) GetVideoByURL(c *gin.Context) {
	url := c.Param("url")
	var command command.GetVideoByURLCommand
	command.URL = url
	ctx := c.Request.Context()
	result, err := controller.commandBus.Dispatch(ctx, &command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, 2000, result)

}
