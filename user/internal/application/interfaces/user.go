package interfaces

import "github.com/bhtoan2204/user/internal/application/command/command"

type UserServiceInterface interface {
	CreateUser(createUserCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	Login(loginCommand *command.LoginCommand) (*command.LoginCommandResult, error)
	Refresh(refreshCommand *command.RefreshTokenCommand) (*command.RefreshTokenCommandResult, error)
	GetUserById(getUserByIdCommand *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error)
}
