package command

import (
	"github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/internal/application/command/handler"
	"github.com/bhtoan2204/user/internal/application/shared"
)

func SetUpCommandBus(deps *shared.ServiceDependencies) *CommandBus {
	bus := NewCommandBus()

	// Register User command handlers
	createUserHandler := handler.NewCreateUserCommandHandler(deps.UserService)
	loginHandler := handler.NewLoginCommandHandler(deps.UserService)
	refreshTokenHandler := handler.NewRefreshTokenCommandHandler(deps.UserService)

	bus.RegisterHandler("CreateUserCommand", func(cmd Command) (interface{}, error) {
		return createUserHandler.Handle(cmd.(*command.CreateUserCommand))
	})
	bus.RegisterHandler("LoginCommand", func(cmd Command) (interface{}, error) {
		return loginHandler.Handle(cmd.(*command.LoginCommand))
	})
	bus.RegisterHandler("RefreshTokenCommand", func(cmd Command) (interface{}, error) {
		return refreshTokenHandler.Handle(cmd.(*command.RefreshTokenCommand))
	})

	// Register other command handlers here

	return bus
}
