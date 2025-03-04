package command

import (
	"context"

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
	logoutHandler := handler.NewLogoutCommandHandler(deps.UserService)

	bus.RegisterHandler("CreateUserCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return createUserHandler.Handle(ctx, cmd.(*command.CreateUserCommand))
	})
	bus.RegisterHandler("LoginCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return loginHandler.Handle(ctx, cmd.(*command.LoginCommand))
	})
	bus.RegisterHandler("RefreshTokenCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return refreshTokenHandler.Handle(ctx, cmd.(*command.RefreshTokenCommand))
	})
	bus.RegisterHandler("LogoutCommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return logoutHandler.Handle(ctx, cmd.(*command.LogoutCommand))
	})

	return bus
}
