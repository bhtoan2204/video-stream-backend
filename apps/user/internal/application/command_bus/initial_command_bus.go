package command_bus

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	"github.com/bhtoan2204/user/internal/application/command_bus/handler"
	"github.com/bhtoan2204/user/internal/application/shared"
)

func SetUpCommandBus(deps *shared.ServiceDependencies) *CommandBus {
	bus := NewCommandBus()

	// Register User command handlers
	createUserHandler := handler.NewCreateUserCommandHandler(deps.UserService)
	loginHandler := handler.NewLoginCommandHandler(deps.UserService)
	refreshTokenHandler := handler.NewRefreshTokenCommandHandler(deps.UserService)
	logoutHandler := handler.NewLogoutCommandHandler(deps.UserService)
	setup2faHandler := handler.NewSetup2FACommandHandler(deps.UserService, deps.UserSettingService)

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
	bus.RegisterHandler("Setup2FACommand", func(ctx context.Context, cmd Command) (interface{}, error) {
		return setup2faHandler.Handle(ctx, cmd.(*command.Setup2FACommand))
	})

	return bus
}
