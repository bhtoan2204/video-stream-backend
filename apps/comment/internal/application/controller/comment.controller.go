package controller

import (
	"reflect"

	"github.com/bhtoan2204/comment/global"
	"github.com/bhtoan2204/comment/internal/application/command_bus"
	"github.com/bhtoan2204/comment/internal/application/command_bus/command"
	"github.com/bhtoan2204/comment/internal/application/middleware"
	"github.com/bhtoan2204/comment/internal/infrastructure/grpc/proto/user"
	"github.com/bhtoan2204/comment/pkg/response"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

type CommentController struct {
	commandBus *command_bus.CommandBus
}

func NewCommentController(
	commandBus *command_bus.CommandBus,
	r *gin.RouterGroup,
) *CommentController {
	ctrl := &CommentController{
		commandBus: commandBus,
	}

	controllerName := reflect.TypeOf(ctrl).Elem().Name()
	meter := otel.Meter("comment-server-meter")
	serverAttribute := attribute.String("controller", controllerName)
	commonLabels := []attribute.KeyValue{serverAttribute}
	requestCount, _ := meter.Int64Counter(
		"comment_server/request_counts",
		metric.WithDescription("The number of requests received"),
	)
	instrument := middleware.NewInstrumentedHandler(requestCount, commonLabels)

	r.POST("", middleware.AuthenticationMiddleware(), instrument(ctrl.CreateComment))

	return ctrl
}

func (controller *CommentController) CreateComment(c *gin.Context) {
	var command command.CreateCommentCommand
	userVal, exists := c.Get("user")
	if !exists {
		global.Logger.Error("User data not found in context")
		response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
		return
	}

	userResp, ok := userVal.(*user.UserResponse)
	if !ok || userResp.Id == "" {
		global.Logger.Error("Invalid user data in context")
		response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
		return
	}
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&command); err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	command.UserId = userResp.Id
	result, err := controller.commandBus.Dispatch(ctx, &command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err.Error())
		return
	}
	response.CreatedResponse(c, 2010, result)
}
