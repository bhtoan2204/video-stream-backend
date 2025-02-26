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
	// Command
	r.POST("/create", ctrl.CreateUser)
	r.POST("/login", ctrl.Login)
	r.POST("/refresh", ctrl.RefreshNewToken)

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
