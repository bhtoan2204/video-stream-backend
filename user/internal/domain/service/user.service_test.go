package service

import (
	"reflect"
	"testing"

	"github.com/bhtoan2204/user/internal/application/command/command"
	commonCommand "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	eSRepository "github.com/bhtoan2204/user/internal/domain/repository/query"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		userRepository      repository.UserRepository
		esUserRepository    eSRepository.ESUserRepository
		refreshTokenService interfaces.RefreshTokenServiceInterface
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepository, tt.args.esUserRepository, tt.args.refreshTokenService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	type fields struct {
		userRepository      repository.UserRepository
		esUserRepository    eSRepository.ESUserRepository
		refreshTokenService interfaces.RefreshTokenServiceInterface
	}
	type args struct {
		createUserCommand *command.CreateUserCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *command.CreateUserCommandResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      tt.fields.userRepository,
				esUserRepository:    tt.fields.esUserRepository,
				refreshTokenService: tt.fields.refreshTokenService,
			}
			got, err := s.CreateUser(tt.args.createUserCommand)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	type fields struct {
		userRepository      repository.UserRepository
		esUserRepository    eSRepository.ESUserRepository
		refreshTokenService interfaces.RefreshTokenServiceInterface
	}
	type args struct {
		loginCommand *command.LoginCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *command.LoginCommandResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      tt.fields.userRepository,
				esUserRepository:    tt.fields.esUserRepository,
				refreshTokenService: tt.fields.refreshTokenService,
			}
			got, err := s.Login(tt.args.loginCommand)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Refresh(t *testing.T) {
	type fields struct {
		userRepository      repository.UserRepository
		esUserRepository    eSRepository.ESUserRepository
		refreshTokenService interfaces.RefreshTokenServiceInterface
	}
	type args struct {
		refreshTokenCommand *command.RefreshTokenCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *commonCommand.RefreshTokenCommandResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      tt.fields.userRepository,
				esUserRepository:    tt.fields.esUserRepository,
				refreshTokenService: tt.fields.refreshTokenService,
			}
			got, err := s.Refresh(tt.args.refreshTokenCommand)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Refresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Refresh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserById(t *testing.T) {
	type fields struct {
		userRepository      repository.UserRepository
		esUserRepository    eSRepository.ESUserRepository
		refreshTokenService interfaces.RefreshTokenServiceInterface
	}
	type args struct {
		getUserByIdCommand *command.GetUserByIdCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *command.GetUserByIdCommandResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      tt.fields.userRepository,
				esUserRepository:    tt.fields.esUserRepository,
				refreshTokenService: tt.fields.refreshTokenService,
			}
			got, err := s.GetUserById(tt.args.getUserByIdCommand)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_SearchUser(t *testing.T) {
	type fields struct {
		userRepository      repository.UserRepository
		esUserRepository    eSRepository.ESUserRepository
		refreshTokenService interfaces.RefreshTokenServiceInterface
	}
	type args struct {
		searchUserQuery *query.SearchUserQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *query.SearchUserQueryResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      tt.fields.userRepository,
				esUserRepository:    tt.fields.esUserRepository,
				refreshTokenService: tt.fields.refreshTokenService,
			}
			got, err := s.SearchUser(tt.args.searchUserQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.SearchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.SearchUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
