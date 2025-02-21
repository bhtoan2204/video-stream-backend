package controller

import (
	"github.com/bhtoan2204/user/internal/application/command"
	"github.com/bhtoan2204/user/internal/application/interfaces"
	"github.com/bhtoan2204/user/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService interfaces.UserServiceInterface
}

func NewUserController(userService interfaces.UserServiceInterface, r *gin.RouterGroup) *UserController {
	ctrl := &UserController{
		userService: userService,
	}

	r.POST("", ctrl.CreateUser)
	r.POST("/login", ctrl.Login)
	return ctrl
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var command command.CreateUserCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	result, err := controller.userService.CreateUser(&command)
	if err != nil {
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}

func (controller *UserController) Login(c *gin.Context) {
	var command command.LoginCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		response.ErrorBadRequestResponse(c, 4000, err)
		return
	}
	result, err := controller.userService.Login(&command)
	if err != nil {
		response.ErrorBadRequestResponse(c, 4000, err.Error())
		return
	}
	c.JSON(200, result)
}
