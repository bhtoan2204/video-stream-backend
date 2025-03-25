package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command_bus/command"
	commonCommand "github.com/bhtoan2204/user/internal/application/common/command"
	common "github.com/bhtoan2204/user/internal/application/common/query"
	"github.com/bhtoan2204/user/internal/application/query_bus/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
	repository_interface "github.com/bhtoan2204/user/internal/domain/repository/command"
	es_repository_interface "github.com/bhtoan2204/user/internal/domain/repository/query"
	value_object "github.com/bhtoan2204/user/internal/domain/value_object/user"
	"github.com/bhtoan2204/user/pkg/encrypt_password"
	"github.com/bhtoan2204/user/pkg/jwt_utils"
	"github.com/bhtoan2204/user/utils"
)

type UserService struct {
	userRepository      repository_interface.UserRepositoryInterface
	esUserRepository    es_repository_interface.ESUserRepositoryInterface
	refreshTokenService interfaces.RefreshTokenServiceInterface
}

func NewUserService(
	userRepository repository_interface.UserRepositoryInterface,
	esUserRepository es_repository_interface.ESUserRepositoryInterface,
	refreshTokenService interfaces.RefreshTokenServiceInterface,
) *UserService {
	return &UserService{
		userRepository:      userRepository,
		esUserRepository:    esUserRepository,
		refreshTokenService: refreshTokenService,
	}
}

func (s *UserService) CreateUser(ctx context.Context, createUserCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	if err := createUserCommand.Validate(); err != nil {
		return nil, err
	}

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

	user, err := s.userRepository.Create(ctx, &entities.User{
		Username:     createUserCommand.Username,
		PasswordHash: password.Hash(),
		Email:        email.Value(),
		Phone:        phone.Value(),
		FirstName:    createUserCommand.FirstName,
		LastName:     createUserCommand.LastName,
		BirthDate:    birthDate.Value(),
		Address:      createUserCommand.Address,
	})

	if err != nil {
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

func (s *UserService) Login(ctx context.Context, loginCommand *command.LoginCommand) (*command.LoginCommandResult, error) {
	if err := loginCommand.Validate(); err != nil {
		return nil, err
	}

	email, err := value_object.NewEmail(loginCommand.Email)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindOneByQuery(ctx,
		&utils.QueryOptions{
			Filters: map[string]interface{}{
				"email": email.Value(),
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

	if err := s.refreshTokenService.CreateRefreshToken(ctx, refreshToken, user.ID, time.UnixMilli(refreshExpiration)); err != nil {
		return nil, err
	}

	if err != nil {
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

func (s *UserService) Refresh(ctx context.Context, refreshTokenCommand *command.RefreshTokenCommand) (*commonCommand.RefreshTokenCommandResult, error) {
	if err := refreshTokenCommand.Validate(); err != nil {
		return nil, err
	}

	claims, err := jwt_utils.ExtractTokenClaims(refreshTokenCommand.RefreshToken, global.Config.SecurityConfig.JWTRefreshSecret)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindOneByQuery(ctx,
		&utils.QueryOptions{
			Filters: map[string]interface{}{
				"id": claims["id"],
			},
		},
	)

	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	isRefreshTokenValid, err := s.refreshTokenService.CheckRefreshToken(ctx, refreshTokenCommand.RefreshToken)
	if err != nil {
		return nil, err
	}
	if !isRefreshTokenValid {
		return nil, errors.New("refresh token is invalid")
	}

	err = s.refreshTokenService.RevokedByQuery(ctx, map[string]interface{}{
		"refresh_token": refreshTokenCommand.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	newAccessToken, newRefreshToken, accessExp, refreshExp, err := jwt_utils.RefreshNewToken(user, refreshTokenCommand.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &commonCommand.RefreshTokenCommandResult{
		AccessToken:           newAccessToken,
		RefreshToken:          newRefreshToken,
		AccessTokenExpiresAt:  accessExp,
		RefreshTokenExpiresAt: refreshExp,
	}, nil
}

func (s *UserService) GetUserById(ctx context.Context, getUserByIdCommand *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error) {
	user, err := s.userRepository.FindOneByQuery(ctx,
		&utils.QueryOptions{
			Filters: map[string]interface{}{
				"id": getUserByIdCommand.ID,
			},
		},
	)
	if err != nil {
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

func (s *UserService) SearchUser(ctx context.Context, searchUserQuery *query.SearchUserQuery) (*query.SearchUserQueryResult, error) {
	searchUserQuery.SetDefaults()
	searchUserQueryResult, pagination, err := s.esUserRepository.Search(ctx, searchUserQuery)
	if err != nil {
		return nil, err
	}

	result := make([]common.UserResult, len(*searchUserQueryResult))
	for i, user := range *searchUserQueryResult {
		birthDate, err := value_object.NewBirthDate(user.BirthDate.Format("2006-01-02"))
		if err != nil {
			return nil, err
		}
		result[i] = common.UserResult{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: birthDate.String(),
			Address:   user.Address,
		}
	}

	return &query.SearchUserQueryResult{
		Result:         &result,
		PaginateResult: pagination,
	}, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, getUserProfileQuery *query.GetUserProfileQuery) (*query.GetUserProfileQueryResult, error) {
	user, err := s.userRepository.FindOneByQuery(ctx,
		&utils.QueryOptions{
			Filters: map[string]interface{}{
				"id": getUserProfileQuery.ID,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	birthDate, err := value_object.NewBirthDate(user.BirthDate.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	return &query.GetUserProfileQueryResult{
		Result: &common.UserResult{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: birthDate.String(),
			Address:   user.Address,
		},
	}, nil
}

func (s *UserService) Logout(ctx context.Context, logoutCommand *command.LogoutCommand) (*commonCommand.LogoutCommandResult, error) {
	if err := logoutCommand.Validate(); err != nil {
		return nil, err
	}

	err := s.refreshTokenService.RevokedByQuery(ctx, map[string]interface{}{
		"refresh_token": logoutCommand.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	return &commonCommand.LogoutCommandResult{
		Success: true,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, updateUserCommand *command.UpdateUserCommand) (*command.UpdateUserCommandResult, error) {
	if err := updateUserCommand.Validate(); err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindOneByQuery(ctx,
		&utils.QueryOptions{
			Filters: map[string]interface{}{
				"id": updateUserCommand.ID,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	user.FirstName = updateUserCommand.FirstName
	user.LastName = updateUserCommand.LastName
	user.Phone = updateUserCommand.Phone
	user.Address = updateUserCommand.Address
	birthDate, err := time.Parse("2006-01-02", updateUserCommand.BirthDate)
	if err != nil {
		return nil, err
	}
	user.BirthDate = &birthDate
	user.Avatar = updateUserCommand.Avatar

	err = s.userRepository.UpdateOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &command.UpdateUserCommandResult{Success: true}, nil
}
