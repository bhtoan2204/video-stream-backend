package service

import (
	"fmt"
	"log"
	"time"

	"github.com/bhtoan2204/user/internal/application/command"
	"github.com/bhtoan2204/user/internal/application/common"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/domain/repository"
	"github.com/bhtoan2204/user/pkg/encrypt_password"
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
	user, err := s.userRepository.FindByEmail(loginCommand.Email)
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

	return &command.LoginCommandResult{
		Result: &common.LoginResult{
			AccessToken:           "access_token",
			RefreshToken:          "refresh_token",
			AccessTokenExpiresAt:  time.Now().Add(time.Hour * 24).Unix(),
			RefreshTokenExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}, nil
}
