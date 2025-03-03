package controller

import (
	"errors"
	"strings"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command"
	realCommand "github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/internal/application/query"
	realQuery "github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/pkg/response"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type UserController struct {
	commandBus *command.CommandBus
	queryBus   *query.QueryBus
}

func NewInstrumentedHandler(counter metric.Int64Counter, commonLabels []attribute.KeyValue) func(gin.HandlerFunc) gin.HandlerFunc {
	return func(handler gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			ctx := c.Request.Context()
			counter.Add(ctx, 1, metric.WithAttributes(commonLabels...))

			span := trace.SpanFromContext(ctx)
			bag := baggage.FromContext(ctx)
			var baggageAttributes []attribute.KeyValue
			baggageAttributes = append(baggageAttributes, commonLabels...)
			for _, member := range bag.Members() {
				baggageAttributes = append(baggageAttributes, attribute.String("baggage key:"+member.Key(), member.Value()))
			}
			span.SetAttributes(baggageAttributes...)

			handler(c)
		}
	}
}

func NewUserController(commandBus *command.CommandBus, queryBus *query.QueryBus, r *gin.RouterGroup) *UserController {
	ctrl := &UserController{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
	meter := otel.Meter("user-server-meter")
	serverAttribute := attribute.String("server-attribute", "foo")
	commonLabels := []attribute.KeyValue{serverAttribute}
	requestCount, _ := meter.Int64Counter(
		"user_server/request_counts",
		metric.WithDescription("The number of requests received"),
	)
	instrument := NewInstrumentedHandler(requestCount, commonLabels)
	// Command
	r.GET("/profile", instrument(ctrl.GetUserProfile))
	r.POST("/create", instrument(ctrl.CreateUser))
	r.POST("/login", instrument(ctrl.Login))
	r.POST("/refresh", instrument(ctrl.RefreshNewToken))
	r.POST("/logout", instrument(ctrl.Logout))

	// Query
	r.GET("", ctrl.SearchUser)
	return ctrl
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var command realCommand.CreateUserCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	result, err := controller.commandBus.Dispatch(&command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}

func (controller *UserController) Login(c *gin.Context) {
	var command realCommand.LoginCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	result, err := controller.commandBus.Dispatch(&command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}

func (controller *UserController) RefreshNewToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		global.Logger.Error("Authorization header is missing")
		response.ErrorBadRequestResponse(c, 4001, errors.New("authorization header is required"))
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		global.Logger.Error("Invalid Authorization header format")
		response.ErrorBadRequestResponse(c, 4001, errors.New("invalid authorization header format"))
		return
	}

	refreshToken := strings.TrimPrefix(authHeader, bearerPrefix)
	if refreshToken == "" {
		global.Logger.Error("Refresh token is empty")
		response.ErrorBadRequestResponse(c, 4001, errors.New("refresh token is empty"))
		return
	}
	var command realCommand.RefreshTokenCommand
	command.RefreshToken = refreshToken

	if err := c.ShouldBindJSON(&command); err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4001, err)
		return
	}
	result, err := controller.commandBus.Dispatch(&command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4001, err.Error())
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
	query.ID = userId
	result, err := controller.queryBus.Dispatch(&query)
	if err != nil {
		global.Logger.Error(query.QueryName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}

func (controller *UserController) SearchUser(c *gin.Context) {
	var query realQuery.SearchUserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		global.Logger.Error("Failed to bind query: ", zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	result, err := controller.queryBus.Dispatch(&query)
	if err != nil {
		global.Logger.Error(query.QueryName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}

func (controller *UserController) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		global.Logger.Error("Authorization header is missing")
		response.ErrorBadRequestResponse(c, 4001, errors.New("authorization header is required"))
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		global.Logger.Error("Invalid Authorization header format")
		response.ErrorBadRequestResponse(c, 4001, errors.New("invalid authorization header format"))
		return
	}

	refreshToken := strings.TrimPrefix(authHeader, bearerPrefix)
	if refreshToken == "" {
		global.Logger.Error("Refresh token is empty")
		response.ErrorBadRequestResponse(c, 4001, errors.New("refresh token is empty"))
		return
	}
	var command realCommand.LogoutCommand
	command.RefreshToken = refreshToken
	if err := c.ShouldBindJSON(&command); err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4001, err)
		return
	}
	result, err := controller.commandBus.Dispatch(&command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4001, err.Error())
		return
	}
	c.JSON(200, result)
}
