package interfaces

import (
	"context"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	common "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/bhtoan2204/user/internal/application/query_bus/query"
)

type UserServiceInterface interface {
	// Command
	CreateUser(ctx context.Context, createUserCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
	Login(ctx context.Context, loginCommand *command.LoginCommand) (*command.LoginCommandResult, error)
	Refresh(ctx context.Context, refreshCommand *command.RefreshTokenCommand) (*common.RefreshTokenCommandResult, error)
	GetUserById(ctx context.Context, getUserByIdCommand *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error)
	Logout(ctx context.Context, logoutCommand *command.LogoutCommand) (*common.LogoutCommandResult, error)
	UpdateUser(ctx context.Context, updateUserCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error)
	// Query
	SearchUser(ctx context.Context, searchUserQuery *query.SearchUserQuery) (*query.SearchUserQueryResult, error)
	GetUserProfile(ctx context.Context, getUserProfileQuery *query.GetUserProfileQuery) (*query.GetUserProfileQueryResult, error)
}
