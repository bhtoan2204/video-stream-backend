package service

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
	repository_interface "github.com/bhtoan2204/user/internal/domain/repository/command"
	es_repository_interface "github.com/bhtoan2204/user/internal/domain/repository/query"
	"github.com/bhtoan2204/user/internal/infrastructure/db/in_memory_db"
	repository_test "github.com/bhtoan2204/user/internal/infrastructure/db/in_memory_db/repository"
	"github.com/bhtoan2204/user/pkg/encrypt_password"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandomUsername() string {
	return RandString(8)
}

// RandomEmail generates a random email using a random username and a fixed domain.
func RandomEmail() string {
	domains := []string{"example.com", "test.com", "mail.com"}
	username := RandString(8)
	domain := domains[rand.Intn(len(domains))]
	return fmt.Sprintf("%s@%s", username, domain)
}

var username = RandomUsername()
var email = RandomEmail()

func TestUserService_CreateUser(t *testing.T) {
	type fields struct {
		userRepository      repository_interface.UserRepositoryInterface
		esUserRepository    es_repository_interface.ESUserRepositoryInterface
		refreshTokenService interfaces.RefreshTokenServiceInterface
	}
	type args struct {
		ctx               context.Context
		createUserCommand *command.CreateUserCommand
	}

	gormClient := in_memory_db.CreateTestDb()
	userRepository := repository_test.NewUserRepository(gormClient)
	// esUserRepository := es_repository_test.NewESUserRepository()
	// refreshTokenService := interfaces.NewRefreshTokenService()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *command.CreateUserCommandResult
		wantErr bool
	}{
		{
			name: "Create user successfully",
			fields: fields{
				userRepository:      userRepository,
				esUserRepository:    nil,
				refreshTokenService: nil,
			},
			args: args{
				ctx: context.Background(),
				createUserCommand: &command.CreateUserCommand{
					Username:  username,
					Password:  "Toan@123456",
					Email:     email,
					Phone:     "+84971308623",
					FirstName: "Bray",
					LastName:  "Drose",
					BirthDate: "2002-02-02",
					Address:   "nga 3 hoang mai thanh 3 phu tho",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewUserService(
				tt.fields.userRepository,
				tt.fields.esUserRepository,
				tt.fields.refreshTokenService,
			)
			_, err := s.CreateUser(tt.args.ctx, tt.args.createUserCommand)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	type fields struct {
		userRepository      repository_interface.UserRepositoryInterface
		esUserRepository    es_repository_interface.ESUserRepositoryInterface
		refreshTokenService interfaces.RefreshTokenServiceInterface
	}
	type args struct {
		ctx          context.Context
		loginCommand *command.LoginCommand
	}

	gormClient := in_memory_db.CreateTestDb()
	userRepository := repository_test.NewUserRepository(gormClient)
	refreshTokenRepository := repository_test.NewRefreshTokenRepository(gormClient)
	testNewRefreshTokenService := NewRefreshTokenService(refreshTokenRepository)
	// create user for test
	createdEmail := RandomEmail()
	createdUsername := RandomUsername()
	password := "Toan@12345"
	passwordHash, _ := encrypt_password.HashPassword(password)

	createUserEntities := &entities.User{
		Username:     createdUsername,
		PasswordHash: passwordHash,
		Email:        createdEmail,
		Phone:        "+84971308623",
		FirstName:    "Bray",
		LastName:     "Drose",
		BirthDate:    func() *time.Time { t, _ := time.Parse("2006-01-02", "2002-02-02"); return &t }(),
		Address:      "nga 3 hoang mai thanh 3 phu tho",
	}

	_, err := userRepository.Create(context.Background(), createUserEntities)

	if err != nil {
		t.Errorf("UserService.Login() error = %v", err)
		return
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *command.LoginCommandResult
		wantErr bool
	}{
		{
			name: "Create user successfully",
			fields: fields{
				userRepository:      userRepository,
				esUserRepository:    nil,
				refreshTokenService: testNewRefreshTokenService,
			},
			args: args{
				ctx: context.Background(),
				loginCommand: &command.LoginCommand{
					Email:    createdEmail,
					Password: password,
				},
			},
			want:    &command.LoginCommandResult{},
			wantErr: false,
		},
		{
			name: "Login failed - wrong password",
			fields: fields{
				userRepository:      userRepository,
				esUserRepository:    nil,
				refreshTokenService: testNewRefreshTokenService,
			},
			args: args{
				ctx: context.Background(),
				loginCommand: &command.LoginCommand{
					Email:    createdEmail,
					Password: "WrongPassword123",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      tt.fields.userRepository,
				esUserRepository:    tt.fields.esUserRepository,
				refreshTokenService: tt.fields.refreshTokenService,
			}
			got, err := s.Login(tt.args.ctx, tt.args.loginCommand)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got == nil {
				t.Errorf("UserService.Login() expected a result but got nil")
			}
		})
	}
}

// func TestUserService_Refresh(t *testing.T) {
// 	type fields struct {
// 		userRepository      repository_interface.UserRepositoryInterface
// 		esUserRepository    es_repository_interface.ESUserRepositoryInterface
// 		refreshTokenService interfaces.RefreshTokenServiceInterface
// 	}
// 	type args struct {
// 		ctx                 context.Context
// 		refreshTokenCommand *command.RefreshTokenCommand
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *commonCommand.RefreshTokenCommandResult
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &UserService{
// 				userRepository:      tt.fields.userRepository,
// 				esUserRepository:    tt.fields.esUserRepository,
// 				refreshTokenService: tt.fields.refreshTokenService,
// 			}
// 			got, err := s.Refresh(tt.args.ctx, tt.args.refreshTokenCommand)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UserService.Refresh() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UserService.Refresh() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestUserService_GetUserById(t *testing.T) {
// 	type fields struct {
// 		userRepository      repository_interface.UserRepositoryInterface
// 		esUserRepository    es_repository_interface.ESUserRepositoryInterface
// 		refreshTokenService interfaces.RefreshTokenServiceInterface
// 	}
// 	type args struct {
// 		ctx                context.Context
// 		getUserByIdCommand *command.GetUserByIdCommand
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *command.GetUserByIdCommandResult
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &UserService{
// 				userRepository:      tt.fields.userRepository,
// 				esUserRepository:    tt.fields.esUserRepository,
// 				refreshTokenService: tt.fields.refreshTokenService,
// 			}
// 			got, err := s.GetUserById(tt.args.ctx, tt.args.getUserByIdCommand)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UserService.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UserService.GetUserById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestUserService_SearchUser(t *testing.T) {
// 	type fields struct {
// 		userRepository      repository_interface.UserRepositoryInterface
// 		esUserRepository    es_repository_interface.ESUserRepositoryInterface
// 		refreshTokenService interfaces.RefreshTokenServiceInterface
// 	}
// 	type args struct {
// 		ctx             context.Context
// 		searchUserQuery *query.SearchUserQuery
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *query.SearchUserQueryResult
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &UserService{
// 				userRepository:      tt.fields.userRepository,
// 				esUserRepository:    tt.fields.esUserRepository,
// 				refreshTokenService: tt.fields.refreshTokenService,
// 			}
// 			got, err := s.SearchUser(tt.args.ctx, tt.args.searchUserQuery)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UserService.SearchUser() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UserService.SearchUser() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestUserService_GetUserProfile(t *testing.T) {
// 	type fields struct {
// 		userRepository      repository_interface.UserRepositoryInterface
// 		esUserRepository    es_repository_interface.ESUserRepositoryInterface
// 		refreshTokenService interfaces.RefreshTokenServiceInterface
// 	}
// 	type args struct {
// 		ctx                 context.Context
// 		getUserProfileQuery *query.GetUserProfileQuery
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *query.GetUserProfileQueryResult
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &UserService{
// 				userRepository:      tt.fields.userRepository,
// 				esUserRepository:    tt.fields.esUserRepository,
// 				refreshTokenService: tt.fields.refreshTokenService,
// 			}
// 			got, err := s.GetUserProfile(tt.args.ctx, tt.args.getUserProfileQuery)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UserService.GetUserProfile() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UserService.GetUserProfile() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestUserService_Logout(t *testing.T) {
// 	type fields struct {
// 		userRepository      repository_interface.UserRepositoryInterface
// 		esUserRepository    es_repository_interface.ESUserRepositoryInterface
// 		refreshTokenService interfaces.RefreshTokenServiceInterface
// 	}
// 	type args struct {
// 		ctx           context.Context
// 		logoutCommand *command.LogoutCommand
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *commonCommand.LogoutCommandResult
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &UserService{
// 				userRepository:      tt.fields.userRepository,
// 				esUserRepository:    tt.fields.esUserRepository,
// 				refreshTokenService: tt.fields.refreshTokenService,
// 			}
// 			got, err := s.Logout(tt.args.ctx, tt.args.logoutCommand)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UserService.Logout() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UserService.Logout() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
