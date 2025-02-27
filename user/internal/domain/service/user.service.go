package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command/command"
	commonCommand "github.com/bhtoan2204/user/internal/application/common/command"
	common "github.com/bhtoan2204/user/internal/application/common/query"
	"github.com/bhtoan2204/user/internal/application/query/query"
	"github.com/bhtoan2204/user/internal/domain/entities"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	eSRepository "github.com/bhtoan2204/user/internal/domain/repository/query"
	value_object "github.com/bhtoan2204/user/internal/domain/value_object/user"
	"github.com/bhtoan2204/user/pkg/encrypt_password"
	"github.com/bhtoan2204/user/pkg/jwt_utils"
	"github.com/bhtoan2204/user/utils"
)

type UserService struct {
	userRepository      repository.UserRepository
	esUserRepository    eSRepository.ESUserRepository
	refreshTokenService interfaces.RefreshTokenServiceInterface
}

func NewUserService(
	userRepository repository.UserRepository,
	esUserRepository eSRepository.ESUserRepository,
	refreshTokenService interfaces.RefreshTokenServiceInterface,
) *UserService {
	return &UserService{
		userRepository:      userRepository,
		esUserRepository:    esUserRepository,
		refreshTokenService: refreshTokenService,
	}
}

func (s *UserService) CreateUser(createUserCommand *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
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

	user, err := s.userRepository.Create(&entities.User{
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

	if err := s.refreshTokenService.CreateRefreshToken(refreshToken, user.ID, time.UnixMilli(refreshExpiration)); err != nil {
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

func (s *UserService) Refresh(refreshTokenCommand *command.RefreshTokenCommand) (*commonCommand.RefreshTokenCommandResult, error) {
	if err := refreshTokenCommand.Validate(); err != nil {
		return nil, err
	}

	claims, err := jwt_utils.ExtractTokenClaims(refreshTokenCommand.RefreshToken, global.Config.SecurityConfig.JWTRefreshSecret)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindOneByQuery(
		utils.QueryOptions{
			Filters: map[string]interface{}{
				"id": claims["id"],
			},
		},
	)

	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	isRefreshTokenValid, err := s.refreshTokenService.CheckRefreshToken(refreshTokenCommand.RefreshToken)
	if err != nil {
		return nil, err
	}
	if !isRefreshTokenValid {
		return nil, errors.New("refresh token is invalid")
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

func (s *UserService) GetUserById(getUserByIdCommand *command.GetUserByIdCommand) (*command.GetUserByIdCommandResult, error) {
	user, err := s.userRepository.FindOneByQuery(
		utils.QueryOptions{
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

func (s *UserService) SearchUser(searchUserQuery *query.SearchUserQuery) (*query.SearchUserQueryResult, error) {
	searchUserQuery.SetDefaults()
	searchUserQueryResult, pagination, err := s.esUserRepository.Search(searchUserQuery)
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
