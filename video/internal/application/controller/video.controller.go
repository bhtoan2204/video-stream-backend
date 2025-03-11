package controller

import (
	"reflect"
	"time"

	"github.com/bhtoan2204/video/global"
	"github.com/bhtoan2204/video/internal/application/command_bus"
	"github.com/bhtoan2204/video/internal/application/command_bus/command"
	"github.com/bhtoan2204/video/internal/application/middleware"
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

	r.POST("", middleware.AuthenticationMiddleware(), instrument(ctrl.UploadVideo))
	r.GET("presigned_url", instrument(ctrl.GetPresignedURL))
	return ctrl
}

func (controller *VideoController) UploadVideo(c *gin.Context) {
	var command command.UploadVideoCommand
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&command); err != nil {
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	result, err := controller.commandBus.Dispatch(ctx, &command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	response.SuccessResponse(c, 2010, result)
}

func (controller *VideoController) GetPresignedURL(c *gin.Context) {
	presignedUploadUrl, err := global.S3Client.GeneratePresignedUploadURL(c.Request.Context(), "videos", 15*time.Minute)
	if err != nil {
		global.Logger.Error("Failed to generate presigned URL", zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	presignedDownloadUrl, err := global.S3Client.GeneratePresignedDownloadURL(c.Request.Context(), "videos", 15*time.Minute)
	if err != nil {
		global.Logger.Error("Failed to generate presigned URL", zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	response.SuccessResponse(c, 2000, gin.H{
		"presignedUploadUrl":   presignedUploadUrl,
		"presignedDownloadUrl": presignedDownloadUrl,
	})
}
