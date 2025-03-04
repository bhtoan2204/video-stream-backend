package controller

import (
	"reflect"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command"
	realCommand "github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/internal/application/middleware"
	"github.com/bhtoan2204/user/internal/application/query"
	realQuery "github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/pkg/response"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

type UserController struct {
	commandBus *command.CommandBus
	queryBus   *query.QueryBus
}

func NewUserController(commandBus *command.CommandBus, queryBus *query.QueryBus, r *gin.RouterGroup) *UserController {
	ctrl := &UserController{
		commandBus: commandBus,
		queryBus:   queryBus,
	}

	controllerName := reflect.TypeOf(ctrl).Elem().Name()
	meter := otel.Meter("user-server-meter")
	serverAttribute := attribute.String("controller", controllerName)
	commonLabels := []attribute.KeyValue{serverAttribute}
	requestCount, _ := meter.Int64Counter(
		"user_server/request_counts",
		metric.WithDescription("The number of requests received"),
	)
	instrument := middleware.NewInstrumentedHandler(requestCount, commonLabels)
	// Command
	r.GET("/profile", instrument(ctrl.GetUserProfile))
	r.POST("/create", instrument(ctrl.CreateUser))

	// Query
	r.GET("", ctrl.SearchUser)
	return ctrl
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var command realCommand.CreateUserCommand
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
	c.JSON(200, result)
}

func (controller *UserController) GetUserProfile(c *gin.Context) {
	userId := c.Request.Header.Get("X-User-ID")
	if userId == "" {
		global.Logger.Error("User ID is missing")
		response.ErrorUnauthorizedResponse(c, 401)
		return
	}
	var query realQuery.GetUserProfileQuery
	ctx := c.Request.Context()
	query.ID = userId
	result, err := controller.queryBus.Dispatch(ctx, &query)
	if err != nil {
		global.Logger.Error(query.QueryName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}

func (controller *UserController) SearchUser(c *gin.Context) {
	var query realQuery.SearchUserQuery
	ctx := c.Request.Context()
	if err := c.ShouldBindQuery(&query); err != nil {
		global.Logger.Error("Failed to bind query: ", zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	result, err := controller.queryBus.Dispatch(ctx, &query)
	if err != nil {
		global.Logger.Error(query.QueryName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}
