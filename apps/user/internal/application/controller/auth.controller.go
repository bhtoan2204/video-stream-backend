package controller

import (
	"errors"
	"reflect"
	"strings"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command_bus"
	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	"github.com/bhtoan2204/user/internal/application/middleware"

	"github.com/bhtoan2204/user/pkg/response"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

type AuthController struct {
	commandBus *command_bus.CommandBus
}

func NewAuthController(commandBus *command_bus.CommandBus, r *gin.RouterGroup) *AuthController {
	ctrl := &AuthController{
		commandBus: commandBus,
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
	r.POST("/login", instrument(ctrl.Login))
	r.POST("/refresh", instrument(ctrl.RefreshNewToken))
	r.POST("/logout", instrument(ctrl.Logout))
	r.GET("/2fa/setup", instrument(ctrl.Setup2FA))

	return ctrl
}

func (controller *AuthController) Login(c *gin.Context) {
	var command command.LoginCommand
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&command); err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	result, err := controller.commandBus.Dispatch(ctx, &command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, 2000, result)
}

func (controller *AuthController) RefreshNewToken(c *gin.Context) {
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
	var command command.RefreshTokenCommand
	ctx := c.Request.Context()
	command.RefreshToken = refreshToken

	if err := c.ShouldBindJSON(&command); err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4001, err)
		return
	}
	result, err := controller.commandBus.Dispatch(ctx, &command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4001, err.Error())
		return
	}
	response.SuccessResponse(c, 2000, result)
}

func (controller *AuthController) Logout(c *gin.Context) {
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
	var command command.LogoutCommand
	ctx := c.Request.Context()
	command.RefreshToken = refreshToken
	if err := c.ShouldBindJSON(&command); err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4001, err)
		return
	}
	result, err := controller.commandBus.Dispatch(ctx, &command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, 4001, err.Error())
		return
	}
	response.SuccessResponse(c, 2000, result)
}

func (controller *AuthController) Setup2FA(c *gin.Context) {
	var command command.Setup2FACommand
	ctx := c.Request.Context()
	// if err := c.ShouldBindJSON(&command); err != nil {
	// 	global.Logger.Error(command.CommandName(), zap.Error(err))
	// 	response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
	// 	return
	// }
	command.UserID = "afdec9b1-b8f6-442e-8a94-46ac570c95a8"
	result, err := controller.commandBus.Dispatch(ctx, &command)
	if err != nil {
		global.Logger.Error(command.CommandName(), zap.Error(err))
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, 2000, result)
}
