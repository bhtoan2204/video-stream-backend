package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/internal/application/common"
	"github.com/bhtoan2204/user/internal/application/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/domain/repository"
	"github.com/bhtoan2204/user/pkg/encrypt_password"
	"github.com/bhtoan2204/user/pkg/jwt_utils"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// CreateUser is a function that creates a new user.
func (s *UserService) CreateUser(createUserCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	birthDate, err := time.Parse(time.DateOnly, createUserCommand.BirthDate)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := encrypt_password.HashPassword(createUserCommand.Password)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	user, err := s.userRepository.Create(&entities.User{
		Username:     createUserCommand.Username,
		PasswordHash: hashedPassword,
		Email:        createUserCommand.Email,
		Phone:        createUserCommand.Phone,
		FirstName:    createUserCommand.FirstName,
		LastName:     createUserCommand.LastName,
		BirthDate:    &birthDate,
	})

	if err != nil {
		return nil, err
	}

	return &command.CreateUserCommandResult{
		Result: &common.UserResult{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: user.BirthDate.Format(time.RFC3339),
		},
	}, nil
}

// Login is a function that logs in a user.
func (s *UserService) Login(loginCommand *command.LoginCommand) (*command.LoginCommandResult, error) {
	user, err := s.userRepository.FindOneByQuery(
		query.QueryOptions{
			Filters: map[string]interface{}{
				"email": loginCommand.Email,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	isVerified, err := encrypt_password.VerifyPassword(user.PasswordHash, loginCommand.Password)

	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, fmt.Errorf("invalid password")
	}

	accessToken, refreshToken, accessExpiration, refreshExpiration, err := jwt_utils.GenerateToken(user)

	if err != nil {
		return nil, err
	}

	return &command.LoginCommandResult{
		Result: &common.LoginResult{
			AccessToken:           accessToken,
			RefreshToken:          refreshToken,
			AccessTokenExpiresAt:  int64(accessExpiration),
			RefreshTokenExpiresAt: int64(refreshExpiration),
		},
	}, nil
}

func (s *UserService) Refresh(refreshTokenCommand *command.RefreshTokenCommand) (*command.RefreshTokenCommandResult, error) {
	claims, err := jwt_utils.ExtractTokenClaims(refreshTokenCommand.RefreshToken)
	if err != nil {
		return nil, err
	}

	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("invalid token claims, missing user id")
	}
	userId := uint(userIdFloat)

	user, err := s.userRepository.FindOneByQuery(
		query.QueryOptions{
			Filters: map[string]interface{}{
				"id": userId,
			},
		},
	)

	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	newAccessToken, newRefreshToken, accessExp, refreshExp, err := jwt_utils.RefreshNewToken(user, refreshTokenCommand.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &command.RefreshTokenCommandResult{
		AccessToken:           newAccessToken,
		RefreshToken:          newRefreshToken,
		AccessTokenExpiresAt:  accessExp,
		RefreshTokenExpiresAt: refreshExp,
	}, nil
}

func (s *UserService) GetUserById(getUserByIdCommand *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error) {
	user, err := s.userRepository.FindOneByQuery(
		query.QueryOptions{
			Filters: map[string]interface{}{
				"id": getUserByIdCommand.ID,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &command.GetUserByIdCommandResult{
		Result: &common.UserResult{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: user.BirthDate.Format(time.RFC3339),
		},
	}, nil
}
