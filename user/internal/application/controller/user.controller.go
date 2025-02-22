package controller

import (
	"github.com/bhtoan2204/user/internal/application/command"
	realCommand "github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	commandBus *command.CommandBus
}

func NewUserController(commandBus *command.CommandBus, r *gin.RouterGroup) *UserController {
	ctrl := &UserController{
		commandBus: commandBus,
	}

	r.POST("", ctrl.CreateUser)
	r.POST("/login", ctrl.Login)
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
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}

func (controller *UserController) Login(c *gin.Context) {
	var command realCommand.LoginCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	result, err := controller.commandBus.Dispatch(&command)
	if err != nil {
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}

func (controller *UserController) RefreshNewToken(c *gin.Context) {
	var command realCommand.RefreshTokenCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		response.ErrorBadRequestResponse(c, 4001, err)
		return
	}
	result, err := controller.commandBus.Dispatch(&command)
	if err != nil {
		response.ErrorBadRequestResponse(c, 4001, err.Error())
		return
	}

	c.JSON(200, result)
}
