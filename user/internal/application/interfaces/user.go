package interfaces

import "github.com/bhtoan2204/user/internal/application/command"

type UserServiceInterface interface {
	CreateUser(createUserCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
}
