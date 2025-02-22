package interfaces

import (
	"github.com/bhtoan2204/user/internal/application/command/command"
	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/bhtoan2204/user/internal/application/query/query"
)

type UserServiceInterface interface {
	// Command
	CreateUser(createUserCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	Login(loginCommand *command.LoginCommand) (*command.LoginCommandResult, error)
	Refresh(refreshCommand *command.RefreshTokenCommand) (*common.RefreshTokenCommandResult, error)
	GetUserById(getUserByIdCommand *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error)

	// Query
	SearchUser(searchUserQuery *query.SearchUserQuery) (*query.SearchUserQueryResult, error)
}
