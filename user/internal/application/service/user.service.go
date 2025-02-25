package service

import (
	"errors"
	"fmt"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command/command"
	commonCommand "github.com/bhtoan2204/user/internal/application/common/command"
	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	eSRepository "github.com/bhtoan2204/user/internal/domain/repository/query"
	domain_service "github.com/bhtoan2204/user/internal/domain/service"
	value_object "github.com/bhtoan2204/user/internal/domain/value_object/user"
	"github.com/bhtoan2204/user/pkg/encrypt_password"
	"github.com/bhtoan2204/user/pkg/jwt_utils"
	"github.com/bhtoan2204/user/utils"
	"go.uber.org/zap"
)

type UserService struct {
	userRepository    repository.UserRepository
	esUserRepository  eSRepository.ESUserRepository
	userDomainService domain_service.AuthDomainService
}

func NewUserService(userRepository repository.UserRepository, esUserRepository eSRepository.ESUserRepository) *UserService {
	return &UserService{
		userRepository:    userRepository,
		esUserRepository:  esUserRepository,
		userDomainService: domain_service.NewAuthDomainService(userRepository),
	}
}

// CreateUser is a function that creates a new user.
func (s *UserService) CreateUser(createUserCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	if err := createUserCommand.Validate(); err != nil {
		return nil, err
	}

	// Domain logic
	birthDate, err := value_object.NewBirthDate(createUserCommand.BirthDate)
	if err != nil {
		return nil, err
	}

	email, err := value_object.NewEmail(createUserCommand.Email)
	if err != nil {
		return nil, err
	}

	phone, err := value_object.NewPhone(createUserCommand.Phone)
	if err != nil {
		return nil, err
	}

	password, err := value_object.NewPassword(createUserCommand.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.Create(&entities.User{
		Username:     createUserCommand.Username,
		PasswordHash: password.Hash(),
		Email:        email.Value(),
		Phone:        phone.Value(),
		FirstName:    createUserCommand.FirstName,
		LastName:     createUserCommand.LastName,
		BirthDate:    birthDate.Value(),
	})

	if err != nil {
		global.Logger.Error("Failed to create user: ", zap.Error(err))
		return nil, err
	}

	return &command.CreateUserCommandResult{
		Result: &commonCommand.UserResult{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: user.BirthDate,
		},
	}, nil
}

// Login is a function that logs in a user.
func (s *UserService) Login(loginCommand *command.LoginCommand) (*command.LoginCommandResult, error) {
	if err := loginCommand.Validate(); err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindOneByQuery(
		utils.QueryOptions{
			Filters: map[string]interface{}{
				"email": loginCommand.Email,
			},
		},
	)

	if err != nil {
		global.Logger.Error("Failed to find user: ", zap.Error(err))
		return nil, err
	}

	isVerified, err := encrypt_password.VerifyPassword(user.PasswordHash, loginCommand.Password)

	if err != nil {
		global.Logger.Error("Failed to verify password: ", zap.Error(err))
		return nil, err
	}

	if !isVerified {
		global.Logger.Error("Failed to hash password: ", zap.Error(fmt.Errorf("invalid password")))
		return nil, fmt.Errorf("invalid password")
	}

	accessToken, refreshToken, accessExpiration, refreshExpiration, err := jwt_utils.GenerateToken(user)

	if err != nil {
		global.Logger.Error("Failed to generate token: ", zap.Error(err))
		return nil, err
	}

	return &command.LoginCommandResult{
		Result: &commonCommand.LoginResult{
			AccessToken:           accessToken,
			RefreshToken:          refreshToken,
			AccessTokenExpiresAt:  int64(accessExpiration),
			RefreshTokenExpiresAt: int64(refreshExpiration),
		},
	}, nil
}

func (s *UserService) Refresh(refreshTokenCommand *command.RefreshTokenCommand) (*commonCommand.RefreshTokenCommandResult, error) {
	if err := refreshTokenCommand.Validate(); err != nil {
		return nil, err
	}

	claims, err := jwt_utils.ExtractTokenClaims(refreshTokenCommand.RefreshToken, global.Config.SecurityConfig.JWTRefreshSecret)
	if err != nil {
		global.Logger.Error("Failed to extract token claims: ", zap.Error(err))
		return nil, err
	}

	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		global.Logger.Error("Invalid token claims, missing user id")
		return nil, errors.New("invalid token claims, missing user id")
	}
	userId := uint(userIdFloat)
	fmt.Println("User ID: ", userId)
	user, err := s.userRepository.FindOneByQuery(
		utils.QueryOptions{
			Filters: map[string]interface{}{
				"id": userId,
			},
		},
	)
	fmt.Println("User: ", user)
	if err != nil || user == nil {
		global.Logger.Error("User not found")
		return nil, errors.New("user not found")
	}

	newAccessToken, newRefreshToken, accessExp, refreshExp, err := jwt_utils.RefreshNewToken(user, refreshTokenCommand.RefreshToken)
	if err != nil {
		global.Logger.Error("Failed to refresh new token: ", zap.Error(err))
		return nil, err
	}

	return &commonCommand.RefreshTokenCommandResult{
		AccessToken:           newAccessToken,
		RefreshToken:          newRefreshToken,
		AccessTokenExpiresAt:  accessExp,
		RefreshTokenExpiresAt: refreshExp,
	}, nil
}

func (s *UserService) GetUserById(getUserByIdCommand *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error) {
	user, err := s.userRepository.FindOneByQuery(
		utils.QueryOptions{
			Filters: map[string]interface{}{
				"id": getUserByIdCommand.ID,
			},
		},
	)
	if err != nil {
		global.Logger.Error("Failed to find user: ", zap.Error(err))
		return nil, err
	}

	return &command.GetUserByIdCommandResult{
		Result: &commonCommand.UserResult{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: user.BirthDate,
		},
	}, nil
}

func (s *UserService) SearchUser(searchUserQuery *query.SearchUserQuery) (*query.SearchUserQueryResult, error) {
	searchUserQuery.SetDefaults()
	searchUserQueryResult, err := s.esUserRepository.Search(searchUserQuery)
	if err != nil {
		global.Logger.Error("Failed to search users: ", zap.Error(err))
		return nil, err
	}

	return searchUserQueryResult, nil
}
